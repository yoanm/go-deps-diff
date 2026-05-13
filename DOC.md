# depsdiff

Package depsdiff provides functionality to compare two dependency package maps and identify differences.

The main API is the Diff function, which compares two PackageMaps (representing lock file states
at different points in time) and returns a DiffMap containing detailed information about what changed:
added, removed, and updated packages with semantic version analysis.

Example usage:

```go
previousMap := contract.PackageMap{/* packages from old lock file */}
currentMap := contract.PackageMap{/* packages from new lock file */}

for pkgName, change := range depsdiff.Diff(previousMap, currentMap) {
	fmt.Printf("%s: %s\n", pkgName, change.Operation.Name)
}
```

## Functions

### func [Diff](/main.go#L38)

`func Diff(previous, current contract.PackageMap) contract.DiffMap`

Diff compares two package maps and returns detailed information about differences.

Parameters:

```diff
- previous: PackageMap representing the previous state (e.g., packages from old lock file)
- current: PackageMap representing the current state (e.g., packages from new lock file)
```

Returns:

```go
- DiffMap: A map where keys are package names and values contain PackageChange information
  including the operation type (added, removed, upgraded, downgraded, etc.) and semantic
  version analysis for updated packages.
```

For each package in the diff result:

```diff
- PackageChange.Package field holds a reference to the package wrapper (agnostic of the package manager).
  See contract.PkgWrapper for more information.
- PackageChange.Operation indicates what changed (ADDITION, REMOVAL, UPGRADE, DOWNGRADE, etc.)
- PackageChange.Operation.SemverType indicates the type of change (MAJOR, MINOR, PATCH, EXTRA, UNKNOWN, NONE)
- PackageChange.PreviousVersion is only populated for updated packages
```

## Sub Packages

* [contract](./contract): Package contract defines the core data types and interfaces for dependency comparison.

* [semver](./semver): Package semver provides utilities for parsing and validating semantic versions.

* [testing](./testing)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
