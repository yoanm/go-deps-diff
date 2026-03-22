# composer-diff

A Go library for comparing two `composer.lock` files and identifying dependency differences.

## Overview

`composer-diff` is a Go module that analyzes and compares two PHP Composer lock files to detect added, removed, and updated package dependencies. It can optionally use `composer.json` files to determine which packages are direct requirements.

## Installation

```bash
go get github.com/user/composer-diff
```

## Usage

```go
package main

import (
    "fmt"
    "os"
    
    "github.com/user/composer-diff/diff"
)

func main() {
    // Read lock files
    lockA, _ := os.ReadFile("composer-old.lock")
    lockB, _ := os.ReadFile("composer-new.lock")
    
    // Optionally read composer.json files
    jsonA, _ := os.ReadFile("composer-old.json")
    jsonB, _ := os.ReadFile("composer-new.json")
    
    // Compare
    output, err := diff.Diff(lockA, lockB, jsonA, jsonB)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    // Process results
    for _, pkg := range output.Packages {
        fmt.Printf("%s: %s\n", pkg.Name, pkg.Update.Type)
    }
}
```

## Features

- ✅ Detects added, removed, and updated packages
- ✅ Semantic version parsing (MAJOR/MINOR/PATCH)
- ✅ Version direction detection (UP/DOWN/NONE/UNKNOWN)
- ✅ Optional composer.json integration for root requirement detection
- ✅ Support for both regular and dev dependencies
- ✅ Handles commit-based versions
- ✅ Efficient O(1) lookup-based comparison

## Package Structure

- **composer/**: Parser for composer.lock and composer.json files
- **diff/**: Main comparison logic and output types

## Testing

```bash
go test ./...
```

All tests pass with 90%+ coverage.

## API Reference

### `Diff(composerLockA, composerLockB, composerJsonA, composerJsonB []byte) (*Output, error)`

Compares two composer.lock files and returns the differences.

**Parameters:**
- `composerLockA`, `composerLockB`: Required byte slices of composer.lock files
- `composerJsonA`, `composerJsonB`: Optional byte slices of corresponding composer.json files (both or neither)

**Returns:**
- `*Output`: Slice of PackageInfo entries with identified differences
- `error`: Non-nil if validation fails

### Output Structure

```go
type Output struct {
    Packages []PackageInfo
}

type PackageInfo struct {
    Name                   string
    IsRootRequirement      bool
    IsRootDevRequirement   bool
    IsAbandoned            bool
    Link                   string
    Update                 UpdateType
    Previous               PkgVersion
    Current                PkgVersion
}
```

## License

See LICENSE file for details.