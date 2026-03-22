package diff

import (
	"fmt"
	"regexp"
	"strconv"

	"composer-diff/composer"
)

// Diff compares two composer.lock files and returns the differences
func Diff(lockPrevious, lockCurrent, reqPrevious, reqCurrent []byte) (*Output, error) {
	// Validate mutual dependency: both requirements or both nil
	hasPrev := reqPrevious != nil && len(reqPrevious) > 0
	hasCurr := reqCurrent != nil && len(reqCurrent) > 0

	if hasPrev != hasCurr {
		return nil, fmt.Errorf("composer requirement files must be provided together: both or neither")
	}

	// Parse lock files
	lockPrevData, err := composer.ParseLock(lockPrevious)
	if err != nil {
		return nil, fmt.Errorf("parsing lockPrevious: %w", err)
	}

	lockCurrData, err := composer.ParseLock(lockCurrent)
	if err != nil {
		return nil, fmt.Errorf("parsing lockCurrent: %w", err)
	}

	// Parse requirement files if provided
	var reqPrevData, reqCurrData *composer.ComposerReq
	if hasPrev {
		reqPrevData, err = composer.ParseReq(reqPrevious)
		if err != nil {
			return nil, fmt.Errorf("parsing reqPrevious: %w", err)
		}

		reqCurrData, err = composer.ParseReq(reqCurrent)
		if err != nil {
			return nil, fmt.Errorf("parsing reqCurrent: %w", err)
		}
	}

	// Build maps for O(1) lookup
	pkgsPrevious := buildPackageMap(lockPrevData, reqPrevData, true)     // Previous uses reqPrevious
	packagesCurrent := buildPackageMap(lockCurrData, reqCurrData, false) // Current uses reqCurrent

	// Find differences
	output := &Output{
		Packages: []PackageInfo{},
	}

	// Find removed and updated packages
	for name, pkgPrev := range pkgsPrevious.packages {
		if pkgCurr, exists := packagesCurrent.packages[name]; exists {
			// Package exists in both - check if updated
			if pkgPrev.version != pkgCurr.version {
				info := createPackageInfo(
					name,
					pkgPrev,
					pkgCurr,
					packagesCurrent.rootRequire[name],    // For updated, check current's require
					packagesCurrent.rootRequireDev[name], // For updated, check current's require-dev
				)
				info.Update = detectUpdate(pkgPrev.version, pkgCurr.version)
				output.Packages = append(output.Packages, info)
			}
		} else {
			// Package removed (only in previous)
			info := createPackageInfo(
				name,
				pkgPrev,
				nil,
				pkgsPrevious.rootRequire[name],    // For removed, check previous's require
				pkgsPrevious.rootRequireDev[name], // For removed, check previous's require-dev
			)
			info.Update = UpdateType{Type: "REMOVED", SubType: "NONE", Direction: "NONE"}
			output.Packages = append(output.Packages, info)
		}
	}

	// Find added packages
	for name, pkgCurr := range packagesCurrent.packages {
		if _, exists := pkgsPrevious.packages[name]; !exists {
			// Package added (only in current)
			info := createPackageInfo(
				name,
				nil,
				pkgCurr,
				packagesCurrent.rootRequire[name],    // For added, check current's require
				packagesCurrent.rootRequireDev[name], // For added, check current's require-dev
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
func buildPackageMap(lock *composer.ComposerLock, req *composer.ComposerReq, isPrevious bool) *packageMap {
	pm := &packageMap{
		packages:       make(map[string]*packageData),
		rootRequire:    make(map[string]bool),
		rootRequireDev: make(map[string]bool),
	}

	// If requirement provided, populate root requirement maps
	if req != nil {
		for pkg := range req.Require {
			pm.rootRequire[pkg] = true
		}
		for pkg := range req.RequireDev {
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
func createPackageInfo(name string, pkgPrev, pkgCurr *packageData, isRoot, isRootDev bool) PackageInfo {
	info := PackageInfo{
		Name:                 name,
		IsRootRequirement:    isRoot,
		IsRootDevRequirement: isRootDev,
	}

	// Determine if abandoned (use current version if available, else previous)
	if pkgCurr != nil {
		info.IsAbandoned = composer.IsAbandoned(pkgCurr.pkg)
		info.Link = composer.GetLink(pkgCurr.pkg)
		info.Current = parseVersion(pkgCurr.version, pkgCurr.pkg)
	} else if pkgPrev != nil {
		info.IsAbandoned = composer.IsAbandoned(pkgPrev.pkg)
		info.Link = composer.GetLink(pkgPrev.pkg)
	}

	if pkgPrev != nil {
		info.Previous = parseVersion(pkgPrev.version, pkgPrev.pkg)
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
func detectUpdate(versionPrevious, versionCurrent string) UpdateType {
	result := UpdateType{
		Type:      "UPDATED",
		SubType:   "NONE",
		Direction: "NONE",
	}

	tagPrevious := parseSemver(versionPrevious)
	tagCurrent := parseSemver(versionCurrent)

	// If either version is not semver, direction is UNKNOWN
	if tagPrevious == nil || tagCurrent == nil {
		result.Direction = "UNKNOWN"
		return result
	}

	// Compare MAJOR
	majorPrevious, _ := strconv.Atoi(tagPrevious.Major)
	majorCurrent, _ := strconv.Atoi(tagCurrent.Major)

	if majorPrevious != majorCurrent {
		result.SubType = "MAJOR"
		if majorCurrent > majorPrevious {
			result.Direction = "UP"
		} else {
			result.Direction = "DOWN"
		}
		return result
	}

	// Compare MINOR
	minorPrevious, _ := strconv.Atoi(tagPrevious.Minor)
	minorCurrent, _ := strconv.Atoi(tagCurrent.Minor)

	if minorPrevious != minorCurrent {
		result.SubType = "MINOR"
		if minorCurrent > minorPrevious {
			result.Direction = "UP"
		} else {
			result.Direction = "DOWN"
		}
		return result
	}

	// Compare PATCH
	patchPrevious, _ := strconv.Atoi(tagPrevious.Patch)
	patchCurrent, _ := strconv.Atoi(tagCurrent.Patch)

	if patchPrevious != patchCurrent {
		result.SubType = "PATCH"
		if patchCurrent > patchPrevious {
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
