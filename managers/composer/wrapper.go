package composer

import (
	"fmt"
	"log/slog"

	"github.com/yoanm/go-deps-diff/shared"
)

func BuildMapFromBytes(reqContent, lockContent []byte) (shared.PackageMap, error) {
	reqData, err := ParseReq(reqContent)
	if err != nil {
		return nil, fmt.Errorf("parsing requirement file content: %w", err)
	}

	lockData, err := ParseLock(lockContent)
	if err != nil {
		return nil, fmt.Errorf("parsing lock file content: %w", err)
	}

	return BuildMap(reqData, lockData)
}

// BuildMap creates an efficient lookup map for composer packages.
func BuildMap(reqData *ComposerReq, lockData *ComposerLock) (shared.PackageMap, error) {
	packageMap := make(map[string]shared.PkgWrapper)

	// Add regular packages
	if lockData.Packages != nil {
		for i := range lockData.Packages {
			pkg := &lockData.Packages[i]
			packageMap[pkg.Name] = createWrapper(pkg, reqData, false)
		}
	}

	// Add dev packages
	if lockData.PackagesDev != nil {
		for i := range lockData.PackagesDev {
			pkg := &lockData.PackagesDev[i]
			// Avoid overwriting regular packages with dev packages (unlikely but just in case)
			if _, exists := packageMap[pkg.Name]; !exists {
				packageMap[pkg.Name] = createWrapper(pkg, reqData, true)
			}
		}
	}

	return packageMap, nil
}

func createWrapper(pkg *Package, reqData *ComposerReq, isDevOnly bool) *ComposerPackageWrapper {
	_, isRootReq := reqData.Require[pkg.Name]
	_, isRootDevReq := reqData.RequireDev[pkg.Name]

	return &ComposerPackageWrapper{
		name:                 pkg.Name,
		isAbandoned:          isAbandonedPkg(pkg),
		version:              parsePkgVersion(pkg),
		link:                 getPkgLink(pkg),
		isDevOnly:            isDevOnly,
		isRootRequirement:    isRootReq,
		isRootDevRequirement: isRootDevReq,
	}
}

// isAbandonedPkg safely extracts the abandoned status from the composer package
// Returns true if the field is explicitly set to true (boolean or string "true").
func isAbandonedPkg(pkg *Package) bool {
	if pkg.Abandoned == nil {
		return false
	}

	switch v := pkg.Abandoned.(type) {
	case bool:
		return v
	case string:
		// Non-empty string is considered abandoned (e.g. an url to the replacement), "true" is explicitly true
		return v == "true" || v != ""
	default:
		return false
	}
}

const shortRefLength = 7

// parsePkgVersion parses a version string into a PkgVersionOld.
func parsePkgVersion(pkg *Package) shared.PkgVersion {
	var (
		semver *shared.SemverVersion
		err    error
	)

	if !shared.IsSemverValid(pkg.Version) { // Not semver - check if there is a commit reference
		if ref := getPkgRef(pkg); ref != "" {
			shortRef := ref
			if len(shortRef) > shortRefLength {
				shortRef = shortRef[:shortRefLength]
			}

			return shared.PkgVersion{
				Raw:    ref,
				Label:  pkg.Version + "#" + shortRef,
				Semver: nil,
			}
		}
	} else if semver, err = shared.ParseSemverVersion(pkg.Version); err != nil {
		slog.Warn(fmt.Errorf("error while parsing semver version %s: %w", pkg.Version, err).Error())
	}

	return shared.PkgVersion{
		Raw:    pkg.Version,
		Label:  pkg.Version,
		Semver: semver,
	}
}

// getPkgRef extracts the commit hash from a package
// Prefers dist.reference over source.reference.
func getPkgRef(pkg *Package) string {
	if pkg.Dist != nil && pkg.Dist.Reference != "" {
		return pkg.Dist.Reference
	}

	if pkg.Source != nil && pkg.Source.Reference != "" {
		return pkg.Source.Reference
	}

	return ""
}

// getPkgLink extracts the best available link from a package
// Priority: wiki -> docs -> source -> homepage.
func getPkgLink(pkg *Package) string {
	if pkg.Support != nil {
		if pkg.Support.Wiki != "" {
			return pkg.Support.Wiki
		}

		if pkg.Support.Docs != "" {
			return pkg.Support.Docs
		}

		if pkg.Support.Source != "" {
			return pkg.Support.Source
		}
	}

	if pkg.Homepage != "" {
		return pkg.Homepage
	}

	return ""
}
