package difftesting

import (
	"fmt"

	"github.com/yoanm/go-deps-diff/contract"
)

func ValidateSemverVersion(actual, expected *contract.Semver) error {
	switch {
	case actual == nil && expected == nil:
		return nil
	case actual == nil && expected != nil:
		return fmt.Errorf("unexpected Semver value. Expected: %v Actual: NIL", expected)
	case actual != nil && expected == nil:
		return fmt.Errorf("unexpected Semver value. Expected: NIL Actual: %v", actual)
	case actual.Major != expected.Major:
		return fmt.Errorf("unexpected Semver Major value. Expected: %d Actual: %d", expected.Major, actual.Major)
	case actual.Minor != expected.Minor:
		return fmt.Errorf("unexpected Semver Minor value. Expected: %d Actual: %d", expected.Minor, actual.Minor)
	case actual.Patch != expected.Patch:
		return fmt.Errorf("unexpected Semver Patch value. Expected: %d Actual: %d", expected.Patch, actual.Patch)
	case actual.Extra != expected.Extra:
		return fmt.Errorf("unexpected Semver Extra value. Expected: %s Actual: %s", expected.Extra, actual.Extra)
	}

	return nil
}
