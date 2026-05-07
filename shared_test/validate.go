package shared_test

import (
	"fmt"
	"github.com/yoanm/go-deps-diff/shared"
	"reflect"
)

func ValidatePackageMap(actual, expectedChanges shared.PackageMap) []error {
	var errList []error
	for _, expectedChange := range expectedChanges {
		pkgName := expectedChange.GetName()
		actualChange, exists := actual[pkgName]
		if !exists {
			errList = append(errList, fmt.Errorf("Package %s is expected to exist", pkgName))
		} else {

			if err := ValidateWrapperPackage(actualChange, expectedChange); err != nil {
				errList = append(errList, fmt.Errorf("package %s has unexpected Package differences: %w", pkgName, err))
			}
		}
	}

	for pkgName := range actual {
		if change, exists := expectedChanges[pkgName]; !exists {
			errList = append(errList, fmt.Errorf("package %s is not expected to exist. %+v", pkgName, change))
		}
	}

	return errList
}

func ValidateWrapperPackage(actualPackage, expectedPackage shared.PkgWrapper) error {
	if actualPackage.GetName() != expectedPackage.GetName() {
		return fmt.Errorf("unexpected GetName() value. Actual: %s", actualPackage.GetName())
	}
	if err := ValidatePackageVersion(actualPackage.GetVersion(), expectedPackage.GetVersion()); err != nil {
		return fmt.Errorf("unexpected Version differences: %w", err)
	}
	if actualPackage.IsAbandoned() != expectedPackage.IsAbandoned() {
		return fmt.Errorf("unexpected IsAbandoned() value. Expected: %v Actual: %v", expectedPackage.IsAbandoned(), actualPackage.IsAbandoned())
	}
	if actualPackage.IsDevOnly() != expectedPackage.IsDevOnly() {
		return fmt.Errorf("unexpected IsDevOnly() value. Expected: %v Actual: %v", expectedPackage.IsDevOnly(), actualPackage.IsDevOnly())
	}
	if actualPackage.IsRootRequirement() != expectedPackage.IsRootRequirement() {
		return fmt.Errorf("unexpected IsRootRequirement() value. Expected: %v Actual: %v", expectedPackage.IsRootRequirement(), actualPackage.IsRootRequirement())
	}
	if actualPackage.IsRootDevRequirement() != expectedPackage.IsRootDevRequirement() {
		return fmt.Errorf("unexpected IsRootRequirement() value. Expected: %v Actual: %v", expectedPackage.IsRootDevRequirement(), actualPackage.IsRootDevRequirement())
	}
	if actualPackage.GetLink() != expectedPackage.GetLink() {
		return fmt.Errorf("unexpected GetLink() value. Expected: %v Actual: %v", expectedPackage.GetLink(), actualPackage.GetLink())
	}

	return nil
}

func ValidatePackageVersion(actualVersion, expectedVersion shared.PkgVersion) error {
	if actualVersion.Raw != expectedVersion.Raw {
		return fmt.Errorf("unexpected Raw value. Expected: %v Actual: %v", expectedVersion.Raw, actualVersion.Raw)
	}
	if actualVersion.Label != expectedVersion.Label {
		return fmt.Errorf("unexpected Label value. Expected: %v Actual: %v", expectedVersion.Label, actualVersion.Label)
	}
	err := ValidatePackageVersionSemver(actualVersion, expectedVersion)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(actualVersion, expectedVersion) {
		return fmt.Errorf("unexpected differences. Expected: %+v, Actual: %+v", expectedVersion, expectedVersion)
	}

	return nil
}

func ValidatePackageVersionSemver(actualVersion, expectedVersion shared.PkgVersion) error {
	if actualVersion.Semver == expectedVersion.Semver {
		return nil
	} else if actualVersion.Semver == nil && expectedVersion.Semver != nil {
		return fmt.Errorf("unexpected Semver value. Expected: %v Actual: NIL", expectedVersion.Semver)
	} else if actualVersion.Semver != nil && expectedVersion.Semver == nil {
		return fmt.Errorf("unexpected Semver value. Expected: NIL Actual: %v", actualVersion.Semver)
	} else if actualVersion.Semver.Major != expectedVersion.Semver.Major {
		return fmt.Errorf("unexpected Semver Major value. Expected: %d Actual: %d", expectedVersion.Semver.Major, actualVersion.Semver.Major)
	} else if actualVersion.Semver.Minor != expectedVersion.Semver.Minor {
		return fmt.Errorf("unexpected Semver Minor value. Expected: %d Actual: %d", expectedVersion.Semver.Minor, actualVersion.Semver.Minor)
	} else if actualVersion.Semver.Patch != expectedVersion.Semver.Patch {
		return fmt.Errorf("unexpected Semver Patch value. Expected: %d Actual: %d", expectedVersion.Semver.Patch, actualVersion.Semver.Patch)
	} else if actualVersion.Semver.Extra != expectedVersion.Semver.Extra {
		return fmt.Errorf("unexpected Semver Extra value. Expected: %s Actual: %s", expectedVersion.Semver.Extra, actualVersion.Semver.Extra)
	}

	return nil
}

type TestPkgWrapper struct {
	Name               string
	Abandoned          bool
	Version            shared.PkgVersion
	Link               string
	DevOnly            bool // true if only in lock file "packages-dev" section (dev-only dependency)
	RootRequirement    bool // true if exists in requirement file "require" section
	RootDevRequirement bool // true if exists in requirement file "require-dev" section
}

func (w *TestPkgWrapper) GetName() string {
	return w.Name
}
func (w *TestPkgWrapper) IsAbandoned() bool {
	return w.Abandoned
}
func (w *TestPkgWrapper) GetVersion() shared.PkgVersion {
	return w.Version
}
func (w *TestPkgWrapper) GetLink() string {
	return w.Link
}
func (w *TestPkgWrapper) IsDevOnly() bool {
	return w.DevOnly
}
func (w *TestPkgWrapper) IsRootRequirement() bool {
	return w.RootRequirement
}
func (w *TestPkgWrapper) IsRootDevRequirement() bool {
	return w.RootDevRequirement
}
