# shared_test

## Constants

```golang
const InvalidOperationName shared.OperationName = "ARGH"
```

## Variables

```golang
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
```

## Functions

### func [ValidatePackageMap](./validate.go#L9)

`func ValidatePackageMap(actual, expectedChanges shared.PackageMap) []error`

### func [ValidatePackageVersion](./validate.go#L59)

`func ValidatePackageVersion(actualVersion, expectedVersion shared.PkgVersion) error`

### func [ValidatePackageVersionSemver](./validate.go#L78)

`func ValidatePackageVersionSemver(actualVersion, expectedVersion shared.PkgVersion) error`

### func [ValidateWrapperPackage](./validate.go#L33)

`func ValidateWrapperPackage(actualPackage, expectedPackage shared.PkgWrapper) error`

## Types

### type [TestPkgWrapper](./validate.go#L98)

`type TestPkgWrapper struct { ... }`

#### func [GetDummyPackage](./values.go#L32)

`func GetDummyPackage() *TestPkgWrapper`

#### func (*TestPkgWrapper) [GetLink](./validate.go#L117)

`func (w *TestPkgWrapper) GetLink() string`

#### func (*TestPkgWrapper) [GetName](./validate.go#L108)

`func (w *TestPkgWrapper) GetName() string`

#### func (*TestPkgWrapper) [GetVersion](./validate.go#L114)

`func (w *TestPkgWrapper) GetVersion() shared.PkgVersion`

#### func (*TestPkgWrapper) [IsAbandoned](./validate.go#L111)

`func (w *TestPkgWrapper) IsAbandoned() bool`

#### func (*TestPkgWrapper) [IsDevOnly](./validate.go#L120)

`func (w *TestPkgWrapper) IsDevOnly() bool`

#### func (*TestPkgWrapper) [IsRootDevRequirement](./validate.go#L126)

`func (w *TestPkgWrapper) IsRootDevRequirement() bool`

#### func (*TestPkgWrapper) [IsRootRequirement](./validate.go#L123)

`func (w *TestPkgWrapper) IsRootRequirement() bool`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
