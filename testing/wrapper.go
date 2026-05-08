package difftesting

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"

	"github.com/yoanm/go-deps-diff/contract"
)

type TestPkgWrapper struct {
	Name               string
	Abandoned          bool
	Version            contract.PkgVersion
	Link               string
	DevOnly            bool // true if only in lock file "packages-dev" section (dev-only dependency)
	RootRequirement    bool // true if exists in requirement file "require" section
	RootDevRequirement bool // true if exists in requirement file "require-dev" section
}

// Ensure that *TestPkgWrapper implements contract.PkgWrapper.
var _ contract.PkgWrapper = (*TestPkgWrapper)(nil)

func (w *TestPkgWrapper) GetName() string {
	return w.Name
}

func (w *TestPkgWrapper) IsAbandoned() bool {
	return w.Abandoned
}

func (w *TestPkgWrapper) GetVersion() contract.PkgVersion {
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

func GetDummyPackage() *TestPkgWrapper {
	major, minor, patch := rand.Int(), rand.Int(), rand.Int()
	version := strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch)

	return &TestPkgWrapper{
		Name:               "vendor/package-" + strconv.Itoa(rand.Int()),
		Abandoned:          true,
		Version:            contract.PkgVersion{Raw: version, Label: version, Semver: &contract.Semver{Major: major, Minor: minor, Patch: patch, Extra: ""}},
		Link:               "",
		DevOnly:            false,
		RootRequirement:    false,
		RootDevRequirement: false,
	}
}

func ValidatePkgWrapper(actualPackage, expectedPackage contract.PkgWrapper) error {
	if actualPackage.GetName() != expectedPackage.GetName() {
		return fmt.Errorf("unexpected GetName() value. Actual: %s", actualPackage.GetName())
	}

	if err := ValidatePkgVersion(actualPackage.GetVersion(), expectedPackage.GetVersion()); err != nil {
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

func ValidatePkgVersion(actualVersion, expectedVersion contract.PkgVersion) error {
	if actualVersion.Raw != expectedVersion.Raw {
		return fmt.Errorf("unexpected Raw value. Expected: %v Actual: %v", expectedVersion.Raw, actualVersion.Raw)
	}

	if actualVersion.Label != expectedVersion.Label {
		return fmt.Errorf("unexpected Label value. Expected: %v Actual: %v", expectedVersion.Label, actualVersion.Label)
	}

	if err := ValidateSemver(actualVersion.Semver, expectedVersion.Semver); err != nil {
		return err
	}

	if !reflect.DeepEqual(actualVersion, expectedVersion) {
		return fmt.Errorf("unexpected differences. Expected: %+v, Actual: %+v", expectedVersion, expectedVersion)
	}

	return nil
}
