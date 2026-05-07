package shared_test

import (
	"github.com/yoanm/go-deps-diff/shared"
	"math/rand"
	"strconv"
)

const InvalidOperationName shared.OperationName = "ARGH"

var (
	AdditionOp          = shared.Operation{Name: shared.AdditionOperation, SemverType: shared.SemverNoUpdate}
	RemovalOp           = shared.Operation{Name: shared.RemovalOperation, SemverType: shared.SemverNoUpdate}
	SameOp              = shared.Operation{Name: shared.NoChangeOperation, SemverType: shared.SemverNoUpdate}
	UpgradeMajorOp      = shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverMajorUpdate}
	UpgradeMinorOp      = shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverMinorUpdate}
	UpgradePatchOp      = shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverPatchUpdate}
	DowngradeMajorOp    = shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverMajorUpdate}
	DowngradeMinorOp    = shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverMinorUpdate}
	DowngradePatchOp    = shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverPatchUpdate}
	UnknownUpdateOp     = shared.Operation{Name: shared.UnknownUpdateOperation, SemverType: shared.SemverUnknownUpdate}
	SemverExtraUpdateOp = shared.Operation{Name: shared.UnknownUpdateOperation, SemverType: shared.SemverExtraUpdate}

	// InvalidOp is purely fictional operation (exists only for test purpose)
	InvalidOp = shared.Operation{Name: InvalidOperationName, SemverType: shared.SemverNoUpdate}
	// InvalidDowngradeOp is not expected to exist (downgrade + semver no update)
	InvalidDowngradeOp = shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverNoUpdate}
	// InvalidUpgradeOp is not expected to exist (upgrade + semver no update)
	InvalidUpgradeOp = shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverNoUpdate}
)

func GetDummyPackage() *TestPkgWrapper {
	major, minor, patch := rand.Int(), rand.Int(), rand.Int()
	version := strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch)

	return &TestPkgWrapper{
		Name:               "vendor/package-" + strconv.Itoa(rand.Int()),
		Abandoned:          true,
		Version:            shared.PkgVersion{Raw: version, Label: version, Semver: &shared.SemverVersion{Major: major, Minor: minor, Patch: patch, Extra: ""}}, // nolint:lll // Meaningless for tests
		Link:               "",
		DevOnly:            false,
		RootRequirement:    false,
		RootDevRequirement: false,
	}
}
