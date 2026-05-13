package depsdiff_test

import (
	"fmt"
	"testing"

	depsdiff "github.com/yoanm/go-deps-diff"
	"github.com/yoanm/go-deps-diff/contract"

	difftesting "github.com/yoanm/go-deps-diff/testing"
)

// createPackageMapWithChanges generates two package maps with controlled changes.
// changeDistribution: map with keys "unchanged", "added", "removed", "upgraded", "downgraded".
func createPackageMapWithChanges(size int, changeDistribution map[string]int) (contract.PackageMap, contract.PackageMap) {
	previous := make(contract.PackageMap)
	current := make(contract.PackageMap)

	changeCount := 0
	unchanged := changeDistribution["unchanged"]
	added := changeDistribution["added"]
	removed := changeDistribution["removed"]
	upgraded := changeDistribution["upgraded"]
	downgraded := changeDistribution["downgraded"]

	// Unchanged packages
	for i := 0; i < unchanged && changeCount < size; i++ {
		pkgName := fmt.Sprintf("vendor/package-%d", changeCount)
		version := fmt.Sprintf("1.0.%d", i)
		wrapper := &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Only required fields are set for testing
			Name:    pkgName,
			Version: contract.PkgVersion{Raw: version, Label: version, Semver: &contract.Semver{Major: 1, Minor: 0, Patch: i, Extra: ""}},
		}
		previous[pkgName] = wrapper
		current[pkgName] = wrapper
		changeCount++
	}

	// Added packages
	for i := 0; i < added && changeCount < size; i++ {
		pkgName := fmt.Sprintf("vendor/package-%d", changeCount)
		version := "1.0.0"
		current[pkgName] = &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Only required fields are set for testing
			Name:    pkgName,
			Version: contract.PkgVersion{Raw: version, Label: version, Semver: &contract.Semver{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
		}
		changeCount++
	}

	// Removed packages
	for i := 0; i < removed && changeCount < size; i++ {
		pkgName := fmt.Sprintf("vendor/package-%d", changeCount)
		version := "1.0.0"
		previous[pkgName] = &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Only required fields are set for testing
			Name:    pkgName,
			Version: contract.PkgVersion{Raw: version, Label: version, Semver: &contract.Semver{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
		}
		changeCount++
	}

	// Upgraded packages
	for i := 0; i < upgraded && changeCount < size; i++ {
		pkgName := fmt.Sprintf("vendor/package-%d", changeCount)
		previousVersion := "1.0.0"
		currentVersion := "2.0.0"
		previous[pkgName] = &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Only required fields are set for testing
			Name:    pkgName,
			Version: contract.PkgVersion{Raw: previousVersion, Label: previousVersion, Semver: &contract.Semver{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
		}
		current[pkgName] = &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Only required fields are set for testing
			Name:    pkgName,
			Version: contract.PkgVersion{Raw: currentVersion, Label: currentVersion, Semver: &contract.Semver{Major: 2, Minor: 0, Patch: 0, Extra: ""}},
		}
		changeCount++
	}

	// Downgraded packages
	for i := 0; i < downgraded && changeCount < size; i++ {
		pkgName := fmt.Sprintf("vendor/package-%d", changeCount)
		previousVersion := "2.0.0"
		currentVersion := "1.0.0"
		previous[pkgName] = &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Only required fields are set for testing
			Name:    pkgName,
			Version: contract.PkgVersion{Raw: previousVersion, Label: previousVersion, Semver: &contract.Semver{Major: 2, Minor: 0, Patch: 0, Extra: ""}},
		}
		current[pkgName] = &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Only required fields are set for testing
			Name:    pkgName,
			Version: contract.PkgVersion{Raw: currentVersion, Label: currentVersion, Semver: &contract.Semver{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
		}
		changeCount++
	}

	return previous, current
}

// BenchmarkDiff_SmallPackageMap benchmarks the Diff function with a small package map (~50 packages).
func BenchmarkDiff_SmallPackageMap(b *testing.B) {
	const smallSize = 50

	changeDistribution := map[string]int{
		"unchanged":  25,
		"added":      10,
		"removed":    5,
		"upgraded":   5,
		"downgraded": 5,
	}

	previous, current := createPackageMapWithChanges(smallSize, changeDistribution)

	b.ResetTimer()

	for range b.N {
		changes := depsdiff.Diff(previous, current)
		if nil == changes { // Just to avoid the compiler to optimize the call away !
			b.Fatal("Nil changes !")
		}
	}
}

// BenchmarkDiff_MediumPackageMap benchmarks the Diff function with a medium package map (~500 packages).
func BenchmarkDiff_MediumPackageMap(b *testing.B) {
	const mediumSize = 500

	changeDistribution := map[string]int{
		"unchanged":  250,
		"added":      100,
		"removed":    50,
		"upgraded":   50,
		"downgraded": 50,
	}

	previous, current := createPackageMapWithChanges(mediumSize, changeDistribution)

	b.ResetTimer()

	for range b.N {
		changes := depsdiff.Diff(previous, current)
		if nil == changes { // Just to avoid the compiler to optimize the call away !
			b.Fatal("Nil changes !")
		}
	}
}

// BenchmarkDiff_ManyChanges benchmarks the Diff function with many changes (worst-case scenario).
// Uses a small total size but with mostly changes.
func BenchmarkDiff_ManyChanges(b *testing.B) {
	const smallSize = 50

	changeDistribution := map[string]int{
		"unchanged":  5,
		"added":      15,
		"removed":    10,
		"upgraded":   10,
		"downgraded": 10,
	}

	previous, current := createPackageMapWithChanges(smallSize, changeDistribution)

	b.ResetTimer()

	for range b.N {
		changes := depsdiff.Diff(previous, current)
		if nil == changes { // Just to avoid the compiler to optimize the call away !
			b.Fatal("Nil changes !")
		}
	}
}

// BenchmarkDiff_NoChanges benchmarks the Diff function when packages are unchanged.
func BenchmarkDiff_NoChanges(b *testing.B) {
	const smallSize = 50

	changeDistribution := map[string]int{
		"unchanged":  50,
		"added":      0,
		"removed":    0,
		"upgraded":   0,
		"downgraded": 0,
	}

	previous, current := createPackageMapWithChanges(smallSize, changeDistribution)

	b.ResetTimer()

	for range b.N {
		changes := depsdiff.Diff(previous, current)
		if nil == changes { // Just to avoid the compiler to optimize the call away !
			b.Fatal("Nil changes !")
		}
	}
}
