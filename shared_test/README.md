# shared_test

## Functions

### func [ValidatePackageMap](./test.go#L10)

`func ValidatePackageMap(actual, expectedChanges shared.PackageMap) []error`

### func [ValidatePackageVersion](./test.go#L60)

`func ValidatePackageVersion(actualVersion, expectedVersion shared.PkgVersion) error`

### func [ValidateWrapperOperation](./test.go#L75)

`func ValidateWrapperOperation(actualOperation, expectedOperation depsdiff.Operation) error`

### func [ValidateWrapperPackage](./test.go#L34)

`func ValidateWrapperPackage(actualPackage, expectedPackage shared.PkgWrapper) error`

## Types

### type [TestPkgWrapper](./test.go#L93)

`type TestPkgWrapper struct { ... }`

#### func (*TestPkgWrapper) [GetLink](./test.go#L112)

`func (w *TestPkgWrapper) GetLink() string`

#### func (*TestPkgWrapper) [GetName](./test.go#L103)

`func (w *TestPkgWrapper) GetName() string`

#### func (*TestPkgWrapper) [GetVersion](./test.go#L109)

`func (w *TestPkgWrapper) GetVersion() shared.PkgVersion`

#### func (*TestPkgWrapper) [IsAbandoned](./test.go#L106)

`func (w *TestPkgWrapper) IsAbandoned() bool`

#### func (*TestPkgWrapper) [IsDevOnly](./test.go#L115)

`func (w *TestPkgWrapper) IsDevOnly() bool`

#### func (*TestPkgWrapper) [IsRootDevRequirement](./test.go#L121)

`func (w *TestPkgWrapper) IsRootDevRequirement() bool`

#### func (*TestPkgWrapper) [IsRootRequirement](./test.go#L118)

`func (w *TestPkgWrapper) IsRootRequirement() bool`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
