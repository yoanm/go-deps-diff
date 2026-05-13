// Package contract defines the core data types and interfaces for dependency comparison.
//
// This package provides:
//   - PackageMap: A map for efficient O(1) lookup of package information
//   - PackageChange: Detailed information about how a package changed
//   - PkgWrapper and PkgVersion: Abstractions for package data from lock files
//   - Operation types: Semantic representation of what changed (added, removed, upgraded, etc.)
//   - Semver: Semantic version components (major, minor, patch, extra)
//
// Types in this package are used as inputs to and outputs from the depsdiff.Diff function.
// They define the contract between package managers and the diff engine.
package contract
