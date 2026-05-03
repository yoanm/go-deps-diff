# go-deps-diff

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
    // Read requirement and lock files
    lockPrevious, _ := os.ReadFile("previous-composer.lock")
    reqPrevious, _ := os.ReadFile("previous-composer.json")
    lockCurrent, _ := os.ReadFile("composer.lock")
    reqCurrent, _ := os.ReadFile("composer.json")

    // Compare (use appropriate function for your package manager)
    output, err := depsdiff.ComposerDiff(
        &depsdiff.Input{
            Current: depsdiff.PkgManagerInput{Lock: currentLock, Requirement: currentReq},
            Previous: depsdiff.PkgManagerInput{Lock: previousLock, Requirement: previousReq},
        },
    )
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    // Process results
    for pkgName, info := range output {
        fmt.Printf("%s: %s\n", pkgName, info.Operation.Name)
    }
}
```

## Features

- ✅ Detects added, removed and updated packages, with detailed information about the type of update
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
