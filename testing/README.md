# testing

## Constants

```golang
const InvalidOperationName contract.OperationName = "ARGH"
```

## Variables

```golang
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

    // InvalidOp is purely fictional operation (exists only for test purpose)
    InvalidOp = contract.Operation{Name: InvalidOperationName, SemverType: contract.SemverNoUpdate}
    // InvalidDowngradeOp is not expected to exist (downgrade + semver no update)
    InvalidDowngradeOp = contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverNoUpdate}
    // InvalidUpgradeOp is not expected to exist (upgrade + semver no update)
    InvalidUpgradeOp = contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverNoUpdate}
)
```

## Functions

### func [ValidatePackageMap](./validate.go#L9)

`func ValidatePackageMap(actual, expectedChanges contract.PackageMap) []error`

### func [ValidatePackageVersion](./validate.go#L59)

`func ValidatePackageVersion(actualVersion, expectedVersion contract.PkgVersion) error`

### func [ValidatePackageVersionSemver](./validate.go#L77)

`func ValidatePackageVersionSemver(actualVersion, expectedVersion contract.PkgVersion) error`

### func [ValidateWrapperPackage](./validate.go#L33)

`func ValidateWrapperPackage(actualPackage, expectedPackage contract.PkgWrapper) error`

## Types

### type [TestPkgWrapper](./validate.go#L97)

`type TestPkgWrapper struct { ... }`

#### func [GetDummyPackage](./values.go#L32)

`func GetDummyPackage() *TestPkgWrapper`

#### func (*TestPkgWrapper) [GetLink](./validate.go#L116)

`func (w *TestPkgWrapper) GetLink() string`

#### func (*TestPkgWrapper) [GetName](./validate.go#L107)

`func (w *TestPkgWrapper) GetName() string`

#### func (*TestPkgWrapper) [GetVersion](./validate.go#L113)

`func (w *TestPkgWrapper) GetVersion() contract.PkgVersion`

#### func (*TestPkgWrapper) [IsAbandoned](./validate.go#L110)

`func (w *TestPkgWrapper) IsAbandoned() bool`

#### func (*TestPkgWrapper) [IsDevOnly](./validate.go#L119)

`func (w *TestPkgWrapper) IsDevOnly() bool`

#### func (*TestPkgWrapper) [IsRootDevRequirement](./validate.go#L125)

`func (w *TestPkgWrapper) IsRootDevRequirement() bool`

#### func (*TestPkgWrapper) [IsRootRequirement](./validate.go#L122)

`func (w *TestPkgWrapper) IsRootRequirement() bool`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
