package diff

import (
	"fmt"
	"regexp"
	"strconv"

	"composer-diff/composer"
)

// Diff compares two composer.lock files and returns the differences
func Diff(composerLockA, composerLockB, composerJsonA, composerJsonB []byte) (*Output, error) {
	// Validate mutual dependency: both json or both nil
	hasA := composerJsonA != nil && len(composerJsonA) > 0
	hasB := composerJsonB != nil && len(composerJsonB) > 0

	if hasA != hasB {
		return nil, fmt.Errorf("composer.json files must be provided together: both or neither")
	}

	// Parse lock files
	lockA, err := composer.ParseLock(composerLockA)
	if err != nil {
		return nil, fmt.Errorf("parsing composerLockA: %w", err)
	}

	lockB, err := composer.ParseLock(composerLockB)
	if err != nil {
		return nil, fmt.Errorf("parsing composerLockB: %w", err)
	}

	// Parse json files if provided
	var jsonA, jsonB *composer.ComposerJson
	if hasA {
		jsonA, err = composer.ParseJson(composerJsonA)
		if err != nil {
			return nil, fmt.Errorf("parsing composerJsonA: %w", err)
		}

		jsonB, err = composer.ParseJson(composerJsonB)
		if err != nil {
			return nil, fmt.Errorf("parsing composerJsonB: %w", err)
		}
	}

	// Build maps for O(1) lookup
	packagesA := buildPackageMap(lockA, jsonA, true)  // A uses jsonA
	packagesB := buildPackageMap(lockB, jsonB, false) // B uses jsonB

	// Find differences
	output := &Output{
		Packages: []PackageInfo{},
	}

	// Find removed and updated packages
	for name, pkgA := range packagesA.packages {
		if pkgB, exists := packagesB.packages[name]; exists {
			// Package exists in both - check if updated
			if pkgA.version != pkgB.version {
				info := createPackageInfo(
					name,
					pkgA,
					pkgB,
					packagesB.rootRequire[name],    // For updated, check B's require
					packagesB.rootRequireDev[name], // For updated, check B's require-dev
				)
				info.Update = detectUpdate(pkgA.version, pkgB.version)
				output.Packages = append(output.Packages, info)
			}
		} else {
			// Package removed (only in A)
			info := createPackageInfo(
				name,
				pkgA,
				nil,
				packagesA.rootRequire[name],    // For removed, check A's require
				packagesA.rootRequireDev[name], // For removed, check A's require-dev
			)
			info.Update = UpdateType{Type: "REMOVED", SubType: "NONE", Direction: "NONE"}
			output.Packages = append(output.Packages, info)
		}
	}

	// Find added packages
	for name, pkgB := range packagesB.packages {
		if _, exists := packagesA.packages[name]; !exists {
			// Package added (only in B)
			info := createPackageInfo(
				name,
				nil,
				pkgB,
				packagesB.rootRequire[name],    // For added, check B's require
				packagesB.rootRequireDev[name], // For added, check B's require-dev
			)
			info.Update = UpdateType{Type: "ADDED", SubType: "NONE", Direction: "NONE"}
			output.Packages = append(output.Packages, info)
		}
	}

	return output, nil
}

// packageMap holds package information for efficient lookup
type packageMap struct {
	packages       map[string]*packageData
	rootRequire    map[string]bool
	rootRequireDev map[string]bool
}

type packageData struct {
	version   string
	pkg       *composer.Package
	isFromDev bool // true if from packages-dev section
}

// buildPackageMap creates an efficient lookup map for packages
func buildPackageMap(lock *composer.ComposerLock, json *composer.ComposerJson, isA bool) *packageMap {
	pm := &packageMap{
		packages:       make(map[string]*packageData),
		rootRequire:    make(map[string]bool),
		rootRequireDev: make(map[string]bool),
	}

	// If json provided, populate root requirement maps
	if json != nil {
		for pkg := range json.Require {
			pm.rootRequire[pkg] = true
		}
		for pkg := range json.RequireDev {
			pm.rootRequireDev[pkg] = true
		}
	}

	// Add regular packages
	if lock.Packages != nil {
		for i := range lock.Packages {
			pkg := &lock.Packages[i]
			if _, exists := pm.packages[pkg.Name]; !exists { // Skip duplicates
				pm.packages[pkg.Name] = &packageData{
					version:   pkg.Version,
					pkg:       pkg,
					isFromDev: false,
				}
			}
		}
	}

	// Add dev packages
	if lock.PackagesDev != nil {
		for i := range lock.PackagesDev {
			pkg := &lock.PackagesDev[i]
			if _, exists := pm.packages[pkg.Name]; !exists { // Skip duplicates
				pm.packages[pkg.Name] = &packageData{
					version:   pkg.Version,
					pkg:       pkg,
					isFromDev: true,
				}
			}
		}
	}

	return pm
}

// createPackageInfo creates a PackageInfo entry for a package
func createPackageInfo(name string, pkgA, pkgB *packageData, isRoot, isRootDev bool) PackageInfo {
	info := PackageInfo{
		Name:                 name,
		IsRootRequirement:    isRoot,
		IsRootDevRequirement: isRootDev,
	}

	// Determine if abandoned (use current version if available, else previous)
	if pkgB != nil {
		info.IsAbandoned = composer.IsAbandoned(pkgB.pkg)
		info.Link = composer.GetLink(pkgB.pkg)
		info.Current = parseVersion(pkgB.version, pkgB.pkg)
	} else if pkgA != nil {
		info.IsAbandoned = composer.IsAbandoned(pkgA.pkg)
		info.Link = composer.GetLink(pkgA.pkg)
	}

	if pkgA != nil {
		info.Previous = parseVersion(pkgA.version, pkgA.pkg)
	}

	return info
}

// parseVersion parses a version string into a PkgVersion
func parseVersion(version string, pkg *composer.Package) PkgVersion {
	// Try to parse as semver first
	if tag := parseSemver(version); tag != nil {
		return tag
	}

	// Not semver - check if it's a commit reference
	if pkg != nil {
		ref := composer.GetCommitReference(pkg)
		if ref != "" {
			return NewPkgVersionCommit(version, ref)
		}
	}

	// Fallback to commit with version as commit string
	return NewPkgVersionCommit(version, version)
}

// parseSemver parses a semantic version string
// Returns nil if parsing fails
func parseSemver(version string) *PkgVersionTag {
	// Pattern: optional 'v', then MAJOR.MINOR.PATCH, then optional extra
	pattern := `^(v)?(\d+)\.(\d+)\.(\d+)(.*)$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(version)
	if matches == nil {
		return nil
	}

	// matches[0] = full match
	// matches[1] = optional v
	// matches[2] = major
	// matches[3] = minor
	// matches[4] = patch
	// matches[5] = extra (may be empty)

	return NewPkgVersionTag(
		version,
		matches[2],
		matches[3],
		matches[4],
		matches[5],
	)
}

// detectUpdate detects the type and direction of a version update
func detectUpdate(versionA, versionB string) UpdateType {
	result := UpdateType{
		Type:      "UPDATED",
		SubType:   "NONE",
		Direction: "NONE",
	}

	tagA := parseSemver(versionA)
	tagB := parseSemver(versionB)

	// If either version is not semver, direction is UNKNOWN
	if tagA == nil || tagB == nil {
		result.Direction = "UNKNOWN"
		return result
	}

	// Compare MAJOR
	majorA, _ := strconv.Atoi(tagA.Major)
	majorB, _ := strconv.Atoi(tagB.Major)

	if majorA != majorB {
		result.SubType = "MAJOR"
		if majorB > majorA {
			result.Direction = "UP"
		} else {
			result.Direction = "DOWN"
		}
		return result
	}

	// Compare MINOR
	minorA, _ := strconv.Atoi(tagA.Minor)
	minorB, _ := strconv.Atoi(tagB.Minor)

	if minorA != minorB {
		result.SubType = "MINOR"
		if minorB > minorA {
			result.Direction = "UP"
		} else {
			result.Direction = "DOWN"
		}
		return result
	}

	// Compare PATCH
	patchA, _ := strconv.Atoi(tagA.Patch)
	patchB, _ := strconv.Atoi(tagB.Patch)

	if patchA != patchB {
		result.SubType = "PATCH"
		if patchB > patchA {
			result.Direction = "UP"
		} else {
			result.Direction = "DOWN"
		}
		return result
	}

	// All numeric components equal
	result.Direction = "NONE"
	return result
}
