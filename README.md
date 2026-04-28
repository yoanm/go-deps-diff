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

    // Compare
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
    for pkgName, info := range output.Changes {
        fmt.Printf("%s: %s\n", pkgName, info.Operation.Name)
    }
}
```

## Features

- ✅ Detects added, removed, and updated packages
- ✅ Semantic version parsing (MAJOR/MINOR/PATCH/EXTRA)
- ✅ Version direction detection (UP/DOWN/NONE/UNKNOWN)
- ✅ Optional requirement file integration for root requirement detection
- ✅ Support for both regular and dev dependencies
- ✅ Handles commit-based versions
- ✅ Efficient O(1) lookup-based comparison

## Sub-packages

- **composer/**: Parser for composer.lock and composer.json files

## Testing

```bash
make test
```
