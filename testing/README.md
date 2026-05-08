# difftesting

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

    // InvalidOp is purely fictional operation (exists only for test purpose).
    InvalidOp = contract.Operation{Name: InvalidOperationName, SemverType: contract.SemverNoUpdate}
    // InvalidDowngradeOp is not expected to exist (downgrade + semver no update).
    InvalidDowngradeOp = contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverNoUpdate}
    // InvalidUpgradeOp is not expected to exist (upgrade + semver no update).
    InvalidUpgradeOp = contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverNoUpdate}
)
```

## Functions

### func [ValidateChanges](./diff.go#L9)

`func ValidateChanges(actual, expectedChanges contract.DiffMap) []error`

### func [ValidateOperation](./operation.go#L33)

`func ValidateOperation(actualOperation, expectedOperation contract.Operation) error`

### func [ValidatePackageMap](./validate.go#L9)

`func ValidatePackageMap(actual, expectedChanges contract.PackageMap) []error`

### func [ValidatePkgVersion](./wrapper.go#L101)

`func ValidatePkgVersion(actualVersion, expectedVersion contract.PkgVersion) error`

### func [ValidatePkgWrapper](./wrapper.go#L69)

`func ValidatePkgWrapper(actualPackage, expectedPackage contract.PkgWrapper) error`

### func [ValidateSemverVersion](./semver.go#L9)

`func ValidateSemverVersion(actual, expected *semver.Version) error`

## Types

### type [TestPkgWrapper](./wrapper.go#L13)

`type TestPkgWrapper struct { ... }`

#### func [GetDummyPackage](./wrapper.go#L54)

`func GetDummyPackage() *TestPkgWrapper`

#### func (*TestPkgWrapper) [GetLink](./wrapper.go#L38)

`func (w *TestPkgWrapper) GetLink() string`

#### func (*TestPkgWrapper) [GetName](./wrapper.go#L26)

`func (w *TestPkgWrapper) GetName() string`

#### func (*TestPkgWrapper) [GetVersion](./wrapper.go#L34)

`func (w *TestPkgWrapper) GetVersion() contract.PkgVersion`

#### func (*TestPkgWrapper) [IsAbandoned](./wrapper.go#L30)

`func (w *TestPkgWrapper) IsAbandoned() bool`

#### func (*TestPkgWrapper) [IsDevOnly](./wrapper.go#L42)

`func (w *TestPkgWrapper) IsDevOnly() bool`

#### func (*TestPkgWrapper) [IsRootDevRequirement](./wrapper.go#L50)

`func (w *TestPkgWrapper) IsRootDevRequirement() bool`

#### func (*TestPkgWrapper) [IsRootRequirement](./wrapper.go#L46)

`func (w *TestPkgWrapper) IsRootRequirement() bool`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
