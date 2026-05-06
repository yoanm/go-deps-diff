# depsdiff

## Functions

### func [ComposerDiff](/managers.go#L10)

`func ComposerDiff(previous, current *PkgManagerInput) (shared.DiffMap, error)`

### func [Diff](/analyzer.go#L8)

`func Diff(previous, current shared.PackageMap) (shared.DiffMap, error)`

Diff compares two packages maps and returns the differences.

## Types

### type [PkgManagerInput](/types.go#L3)

`type PkgManagerInput struct { ... }`

## Sub Packages

* [managers/composer](./managers/composer)

* [shared](./shared)

* [shared_test](./shared_test)

* [summary](./summary)

* [summary/markdown](./summary/markdown)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
