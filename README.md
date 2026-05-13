# go-deps-diff<br/><sup><sub>Manager agnostic deps comparator</sub></sup>

[![License](https://img.shields.io/github/license/yoanm/go-deps-diff.svg)](https://github.com/yoanm/go-deps-diff)
[![Code size](https://img.shields.io/github/languages/code-size/yoanm/go-deps-diff.svg)](https://github.com/yoanm/go-deps-diff)
[![Go Reference](https://pkg.go.dev/badge/github.com/yoanm/go-deps-diff.svg)](https://pkg.go.dev/github.com/yoanm/go-deps-diff)

![Dependabot Status](https://flat.badgen.net/github/dependabot/yoanm/go-deps-diff)
![Last commit](https://badgen.net/github/last-commit/yoanm/go-deps-diff)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/ebeacd3a91a74fef8a8ed4ea879ede72)](https://app.codacy.com/gh/yoanm/go-deps-diff/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/yoanm/go-deps-diff?)](https://goreportcard.com/report/github.com/yoanm/go-deps-diff)

[![CI](https://github.com/yoanm/go-deps-diff/actions/workflows/CI.yml/badge.svg?branch=master)](https://github.com/yoanm/go-deps-diff/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/yoanm/go-deps-diff/branch/master/graph/badge.svg?token=NHdwEBUFK5)](https://codecov.io/gh/yoanm/go-deps-diff)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yoanm/go-deps-diff)


A Go library for comparing two dependency lock files and identifying differences.

## Overview

`deps-diff` is a Go module that analyzes and compares two packages maps to detect added, removed, and updated package dependencies.

Currently handles:
- `composer` - PHP

## Installation

```bash
go get github.com/yoanm/go-deps-diff
```

## Usage

```go
package main

import (
    "fmt"
    "os"
    
    "github.com/yoanm/go-deps-diff"
)

func main() {
    // Generate diff: 
    // - Use appropriate library for your package manager, like go-composer-diff for composer for instance
    // - Or directly provide two `contract.PackageMap` objects
    previousMap := contract.PackageMap{
        /* ... */
    }
    currentMap := contract.PackageMap{
        /* ... */
    }
    
    // Compare & process results
    for pkgName, info := range depsdiff.Diff(previousMap, currentMap) {
        fmt.Printf("%s: %s\n", pkgName, info.Operation.Name)
    }
}
```

## Features

- ✅ Detects added, removed and updated packages, as well as unchanged packages
- ✅ Semantic version parsing (MAJOR/MINOR/PATCH/EXTRA)
- ✅ Update direction detection (DOWNGRADE/UPGRADE) for semver compatible versions (UNKNOWN_UPDATE otherwise)
- ✅ Support for both regular and dev dependencies
- ✅ Handles commit-based versions
- ✅ Efficient O(1) lookup-based comparison

## Sub-packages

- **composer**: Parser and wrapper for composer

## Testing

```bash
make test
```

## API Reference

### Core Functions

#### `Diff(previous, current PackageMap) (DiffMap, error)`

Compares two package maps and returns detailed information about differences.

**Parameters:**
- `previous`: PackageMap representing the previous state (e.g., packages from old lock file)
- `current`: PackageMap representing the current state (e.g., packages from new lock file)

**Returns:**
- `DiffMap`: A map where keys are package names and values contain PackageChange information including operation type and semantic version analysis
- `error`: Non-nil if comparison fails (currently always returns nil)

**Example:**
```go
changes, err := depsdiff.Diff(previousMap, currentMap)
if err != nil {
    log.Fatal(err)
}

for pkgName, change := range changes {
    fmt.Printf("%s: %s (%s)\n", pkgName, change.Operation.Name, change.Operation.SemverType)
}
```

### Semantic Version Functions

#### `semver.IsValid(value string) bool`

Validates whether a string is a valid semantic version.

**Parameters:**
- `value`: The version string to validate (e.g., "1.2.3", "v1.2.3-beta")

**Returns:**
- `true` if the version matches semantic version format, `false` otherwise

**Example:**
```go
import "github.com/yoanm/go-deps-diff/semver"

if semver.IsValid("1.2.3") {
    fmt.Println("Valid semver")
}
```

#### `semver.Parse(version string) (*Semver, error)`

Parses a semantic version string into its components.

**Parameters:**
- `version`: A semantic version string (e.g., "1.2.3", "v2.0.0-beta+build.1")

**Returns:**
- `*Semver`: Pointer to parsed version with Major, Minor, Patch, and Extra fields
- `error`: InvalidVersionError if format is invalid, or InvalidComponentError if components cannot be parsed

**Example:**
```go
semver, err := semver.Parse("1.2.3-beta")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Major: %d, Minor: %d, Patch: %d\n", semver.Major, semver.Minor, semver.Patch)
```

### Core Types

#### `PackageMap`

A map type that stores package information, keyed by package name.

```go
type PackageMap map[string]PkgWrapper
```

**Usage:**
```go
var packages contract.PackageMap
packages["vendor/package"] = packageWrapper
```

#### `PackageChange`

Contains detailed information about a package difference.

**Fields:**
- `Package PkgWrapper`: The package (current version for added/updated, previous version for removed)
- `Operation Operation`: What changed (ADDITION, REMOVAL, UPGRADE, DOWNGRADE, etc.)
- `PreviousVersion PkgVersion`: Previous version (only for updated packages)

**Example:**
```go
change := changes["vendor/package"]
if change.Operation.Name == contract.UpgradeOperation {
    fmt.Printf("Upgraded from %s to %s\n", 
        change.PreviousVersion.Label, 
        change.Package.GetVersion().Label)
}
```

#### `Operation`

Describes what changed about a package.

**Fields:**
- `Name OperationName`: ADDITION, REMOVAL, UPGRADE, DOWNGRADE, UNKNOWN_UPDATE, or NONE
- `SemverType OperationSemverType`: MAJOR, MINOR, PATCH, EXTRA, UNKNOWN, or NONE

**Operation Names:**
- `ADDITION`: Package was added (new in current map)
- `REMOVAL`: Package was removed (only in previous map)
- `UPGRADE`: Package version increased
- `DOWNGRADE`: Package version decreased
- `UNKNOWN_UPDATE`: Package updated but direction unknown (non-semver)
- `NONE`: Package unchanged

**Semver Types:**
- `MAJOR`: Major version component differs
- `MINOR`: Minor version component differs
- `PATCH`: Patch version component differs
- `EXTRA`: Pre-release/build metadata differs
- `UNKNOWN`: Cannot determine difference (non-semver versions)
- `NONE`: No semver difference (added, removed, or unchanged)

#### `PkgWrapper` Interface

Provides access to package metadata from a lock file.

**Methods:**
- `GetName() string`: Package name (e.g., "vendor/package")
- `GetVersion() PkgVersion`: Version information
- `IsAbandoned() bool`: Whether package is abandoned
- `IsDevOnly() bool`: Whether package is dev-only
- `IsRootRequirement() bool`: Whether explicitly required
- `IsRootDevRequirement() bool`: Whether explicitly required for dev
- `GetLink() string`: Package documentation link

#### `PkgVersion`

Contains version information for a package.

**Fields:**
- `Raw string`: Exact version as in lock file (e.g., "1.2.3", "dev-master#abc123")
- `Label string`: Human-readable version representation
- `Semver *Semver`: Parsed semantic version (nil if not valid semver)

#### `Semver`

Represents semantic version components.

**Fields:**
- `Major int`: Major version component (e.g., 1 in "1.2.3")
- `Minor int`: Minor version component (e.g., 2 in "1.2.3")
- `Patch int`: Patch version component (e.g., 3 in "1.2.3")
- `Extra string`: Pre-release or build metadata (e.g., "-beta", "+build.1")

## Examples

### Filter Changes by Operation Type

```go
changes, err := depsdiff.Diff(previousMap, currentMap)
if err != nil {
    log.Fatal(err)
}

// Find all upgraded packages
for pkgName, change := range changes {
    if change.Operation.Name == contract.UpgradeOperation {
        fmt.Printf("Upgraded: %s\n", pkgName)
    }
}

// Find all major version updates
for pkgName, change := range changes {
    if change.Operation.SemverType == contract.SemverMajorUpdate {
        fmt.Printf("Major update: %s\n", pkgName)
    }
}
```

### Iterate Through All Changes

```go
changes, err := depsdiff.Diff(previousMap, currentMap)
if err != nil {
    log.Fatal(err)
}

for pkgName, change := range changes {
    fmt.Printf("%s: %s (%s)\n", 
        pkgName,
        change.Operation.Name,
        change.Operation.SemverType)
    
    // Show previous version for updates
    if change.PreviousVersion.Raw != "" {
        fmt.Printf("  %s -> %s\n",
            change.PreviousVersion.Label,
            change.Package.GetVersion().Label)
    }
}
```

### Check for Security-Relevant Updates

```go
changes, err := depsdiff.Diff(previousMap, currentMap)
if err != nil {
    log.Fatal(err)
}

for pkgName, change := range changes {
    switch change.Operation.SemverType {
    case contract.SemverMajorUpdate, contract.SemverMinorUpdate:
        fmt.Printf("⚠️  %s %s: %s\n", pkgName, change.Operation.SemverType, change.Package.GetVersion().Label)
    }
}
```
