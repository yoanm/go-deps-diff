# shared_test

## Functions

### func [ValidatePackageMap](./test.go#L9)

`func ValidatePackageMap(actual, expectedChanges shared.PackageMap) []error`

### func [ValidatePackageVersion](./test.go#L59)

`func ValidatePackageVersion(actualVersion, expectedVersion shared.PkgVersion) error`

### func [ValidateWrapperPackage](./test.go#L33)

`func ValidateWrapperPackage(actualPackage, expectedPackage shared.PkgWrapper) error`

## Types

### type [TestPkgWrapper](./test.go#L74)

`type TestPkgWrapper struct { ... }`

#### func (*TestPkgWrapper) [GetLink](./test.go#L93)

`func (w *TestPkgWrapper) GetLink() string`

#### func (*TestPkgWrapper) [GetName](./test.go#L84)

`func (w *TestPkgWrapper) GetName() string`

#### func (*TestPkgWrapper) [GetVersion](./test.go#L90)

`func (w *TestPkgWrapper) GetVersion() shared.PkgVersion`

#### func (*TestPkgWrapper) [IsAbandoned](./test.go#L87)

`func (w *TestPkgWrapper) IsAbandoned() bool`

#### func (*TestPkgWrapper) [IsDevOnly](./test.go#L96)

`func (w *TestPkgWrapper) IsDevOnly() bool`

#### func (*TestPkgWrapper) [IsRootDevRequirement](./test.go#L102)

`func (w *TestPkgWrapper) IsRootDevRequirement() bool`

#### func (*TestPkgWrapper) [IsRootRequirement](./test.go#L99)

`func (w *TestPkgWrapper) IsRootRequirement() bool`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
