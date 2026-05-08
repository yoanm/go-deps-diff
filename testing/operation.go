package difftesting

import (
	"fmt"
	"reflect"

	"github.com/yoanm/go-deps-diff/contract"
)

const InvalidOperationName contract.OperationName = "ARGH"

var (
	AdditionOp          = contract.Operation{Name: contract.AdditionOperation, SemverType: contract.SemverNoUpdate}
	RemovalOp           = contract.Operation{Name: contract.RemovalOperation, SemverType: contract.SemverNoUpdate}
	SameOp              = contract.Operation{Name: contract.NoChangeOperation, SemverType: contract.SemverNoUpdate}
	UpgradeMajorOp      = contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverMajorUpdate}
	UpgradeMinorOp      = contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverMinorUpdate}
	UpgradePatchOp      = contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverPatchUpdate}
	DowngradeMajorOp    = contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverMajorUpdate}
	DowngradeMinorOp    = contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverMinorUpdate}
	DowngradePatchOp    = contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverPatchUpdate}
	UnknownUpdateOp     = contract.Operation{Name: contract.UnknownUpdateOperation, SemverType: contract.SemverUnknownUpdate}
	SemverExtraUpdateOp = contract.Operation{Name: contract.UnknownUpdateOperation, SemverType: contract.SemverExtraUpdate}

	// InvalidOp is purely fictional operation (exists only for test purpose).
	InvalidOp = contract.Operation{Name: InvalidOperationName, SemverType: contract.SemverNoUpdate}
	// InvalidDowngradeOp is not expected to exist (downgrade + semver no update).
	InvalidDowngradeOp = contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverNoUpdate}
	// InvalidUpgradeOp is not expected to exist (upgrade + semver no update).
	InvalidUpgradeOp = contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverNoUpdate}
)

func ValidateOperation(actualOperation, expectedOperation contract.Operation) error {
	if actualOperation.Name != expectedOperation.Name {
		return fmt.Errorf("unexpected Name value. Expected: %s Actual: %s", expectedOperation.Name, actualOperation.Name)
	}

	if actualOperation.SemverType != expectedOperation.SemverType {
		return fmt.Errorf(
			"unexpected SemverType value. Expected: %s Actual: %s",
			expectedOperation.SemverType,
			actualOperation.SemverType,
		)
	}

	if !reflect.DeepEqual(actualOperation, expectedOperation) {
		return fmt.Errorf("unexpected differences. Expected: %+v, Actual: %+v", expectedOperation, actualOperation)
	}

	return nil
}
