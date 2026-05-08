package difftesting

import (
	"fmt"

	"github.com/yoanm/go-deps-diff/contract"
)

func ValidatePackageMap(actual, expectedChanges contract.PackageMap) []error {
	var errList []error

	for _, expectedChange := range expectedChanges {
		pkgName := expectedChange.GetName()

		actualChange, exists := actual[pkgName]
		if !exists {
			errList = append(errList, fmt.Errorf("package %s is expected to exist", pkgName))
		} else {
			if err := ValidatePkgWrapper(actualChange, expectedChange); err != nil {
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
