package difftesting

import (
	"fmt"

	"github.com/yoanm/go-deps-diff/contract"
)

func ValidateChanges(actual, expectedChanges contract.DiffMap) []error {
	var errList []error

	for _, expectedChange := range expectedChanges {
		pkgName := expectedChange.Package.GetName()
		actualChange, exists := actual[pkgName]

		if !exists {
			errList = append(errList, fmt.Errorf("package %s is expected to exist", pkgName))
		} else {
			if err := ValidatePkgWrapper(actualChange.Package, expectedChange.Package); err != nil {
				errList = append(errList, fmt.Errorf("package %s has unexpected Package differences: %w", pkgName, err))
			}

			if err := ValidateOperation(actualChange.Operation, expectedChange.Operation); err != nil {
				errList = append(errList, fmt.Errorf("package %s has unexpected Operation differences: %w", pkgName, err))
			}

			if err := ValidatePkgVersion(actualChange.PreviousVersion, expectedChange.PreviousVersion); err != nil {
				errList = append(errList, fmt.Errorf("package %s has unexpected PreviousVersion differences: %w", pkgName, err))
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
