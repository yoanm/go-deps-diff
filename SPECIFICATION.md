# Composer.lock Diff - Project Specification

## Project Overview
A Go module (library) that analyzes and compares two `composer.lock` file contents to identify differences in package dependencies. This module is designed to be imported and used by other Go projects.

## Project Goal
Provide a reusable Go library that compares two composer.lock files (previous and current) and returns structured data about the differences in package versions and availability.

## Stack
- **Language**: Go 1.23+ (latest available version)
- **Build Tool**: go build (native Go toolchain)

## Hard Constraints
- No new external dependencies may be installed without explicit approval
- Tool must work with standard library only (unless otherwise approved)

---

## Functional Requirements

### 1. Function Interface Specification

The module exports a primary function for comparing composer.lock contents:

```go
// Diff compares two composer.lock file contents and returns the differences.
// Parameters:
//   - lockPrevious: byte slice containing the previous composer.lock file content (baseline/original)
//   - lockCurrent: byte slice containing the current composer.lock file content (comparison/target)
//   - reqPrevious: (optional, only required if reqCurrent is provided) byte slice containing the CORRESPONDING composer.json file content for lockPrevious
//                    This file must be the composer.json that was used to generate lockPrevious.
//   - reqCurrent: (optional, only required if reqPrevious is provided) byte slice containing the CORRESPONDING composer.json file content for lockCurrent
//                    This file must be the composer.json that was used to generate lockCurrent.
// Returns:
//   - *Output: structured data containing identified differences (added, removed, updated packages)
//   - error: non-nil if validation fails (invalid JSON, invalid format, etc.)
func Diff(lockPrevious, lockCurrent, reqPrevious, reqCurrent []byte) (*Output, error)
```

**IMPORTANT: IsRootRequirement & IsRootDevRequirement Semantics**

These fields default to `false` and can ONLY be determined if the corresponding composer.json file is provided:

- **For added packages** (exist in current only): Use `reqCurrent` to determine flags
  - IsRootRequirement = true if package is in `reqCurrent.require`
  - IsRootDevRequirement = true if package is in `reqCurrent.require-dev`
  
- **For removed packages** (exist in previous only): Use `reqPrevious` to determine flags
  - IsRootRequirement = true if package is in `reqPrevious.require`
  - IsRootDevRequirement = true if package is in `reqPrevious.require-dev`
  
- **For updated packages** (exist in both): Use `reqCurrent` to determine flags
  - IsRootRequirement = true if package is in `reqCurrent.require`
  - IsRootDevRequirement = true if package is in `reqCurrent.require-dev`
  
- **If corresponding composer.json is NOT provided**: Both flags default to `false`
  - No fallback to composer.lock "packages" vs "packages-dev" section membership

### 2. Input Validation
- **Parameters**:
    - `lockPrevious`: Byte slice of previous composer.lock file (required)
    - `lockCurrent`: Byte slice of current composer.lock file (required)
    - `reqPrevious`: Byte slice of previous composer.json file (optional, can be nil)
    - `reqCurrent`: Byte slice of current composer.json file (optional, can be nil)
    - **Mutual dependency rule**: Both composer.json parameters must be nil together or both must be provided. Providing only one is an error.
- **Validation**:
    - Both lock files must be valid JSON
    - Both lock files must be valid composer.lock format (contain "packages" and/or "packages-dev" fields)
    - If either composer.json is provided, both must be provided (if one is nil but the other is not, return error)
    - If provided, composer.json files must be valid JSON (return error if invalid)
    - It's valid to have mismatched: (both composerJson provided or both nil), with any combination of empty packages/packages-dev sections

### 3. Output Specification
The tool must identify and categorize three types of differences:

#### 3.1 Added Packages
- **Definition**: Packages that exist in current but not in previous
- **Output Format**: List of packages with:
    - Package name
    - Version in current
    - Source (if applicable)

#### 3.2 Removed Packages
- **Definition**: Packages that exist in previous but not in current
- **Output Format**: List of packages with:
    - Package name
    - Version in previous
    - Source (if applicable)

#### 3.3 Updated Packages
- **Definition**: Packages that exist in both files but have different versions
- **Output Format**: List of packages with:
    - Package name
    - Version in previous (old)
    - Version in current (new)
    - Link, if available. Use the first available property listed below:
      - `support.wiki`
      - `support.docs`
      - `support.source` 
      - `homepage`

### 4. Composer.json File Structure

**Composer.json Structure:**
```json
{
  "require": {
    "vendor/package-name": "^1.0",
    "another/package": ">=2.0,<3.0"
  },
  "require-dev": {
    "phpunit/phpunit": "^9.0",
    "squizlabs/php_codesniffer": "^3.5"
  }
}
```

**Schema:**
- `require` (optional, object): Direct dependencies required for the application
  - Keys: Package name (string) in format `vendor/package`
  - Values: Version constraint (string) e.g., `^1.0`, `>=2.0`, `1.2.3`
- `require-dev` (optional, object): Development-only dependencies
  - Keys: Package name (string) in format `vendor/package`
  - Values: Version constraint (string) same format as `require`

### 5. Composer.lock File Structure
The tool must parse the following structure from composer.lock files:
```json
{
  "packages": [
    {
      "name": "vendor/package-name",
      "version": "1.0.0",
      "source": {
        "reference": "..."
      },
      "dist": {
        "reference": "..."
      },
      "support": { // Optional property
        "wiki": "...",  // Optional property
        "docs": "...",  // Optional property
        "source":  "..."  // Optional property
      },
      "homepage": "...", // Optional property
      "abandoned": false // Optional boolean property
    }
  ],
  "packages-dev": [
    // Same structure as packages
  ]
}
```

## Technical Requirements

### 1. Module Organization and Exports
The module must be importable and easy to use:

```go
package composerdiff

// Diff compares two composer.lock file contents and returns the differences.
func Diff(lockPrevious, lockCurrent, reqPrevious, reqCurrent []byte) (*Output, error)
```

**Module import:** `import "github.com/user/go-deps-diff"` 

**Module root package:** The public `Diff()` function should be exported from the module root package (or from a subpackage like `github.com/user/go-deps-diff/diff`).

**Usage expectations:**
- Consumers will read files themselves and pass byte slices to Diff()
- Diff() returns structured data (*Output) for further processing
- All errors (JSON parsing, format validation) are returned as `error` type

### 2. Composer.json Parsing (Optional)

**IMPORTANT: Mutual Dependency Rule**
- If `reqPrevious` is provided, `reqCurrent` MUST also be provided (and vice versa)
- Either provide both, or provide neither (both can be nil)
- Implementation must validate this and return an error if the rule is violated

**Parsing logic:**
- If `reqPrevious` is provided:
  - Parse JSON
  - Extract `require` object (if present) → keys are package names marked as IsRootRequirement
  - Extract `require-dev` object (if present) → keys are package names marked as IsRootDevRequirement
  - Invalid JSON in composer.json: Return error with context
- If `reqPrevious` is nil:
  - All flags default to false (IsRootRequirement=false, IsRootDevRequirement=false)
  - Do NOT use "packages" vs "packages-dev" sections as fallback
- Same logic applies for current files with `reqCurrent`

### 3. Error Handling
- **Invalid JSON**: Return error with details from encoding/json parsing
- **Invalid composer.lock format**: Return error indicating missing "packages" field or invalid structure
- **Error handling approach**:
   - Custom error types: `ErrInvalidJSON`, `ErrInvalidFormat`
   - Wrap errors with context: `fmt.Errorf("context: %w", err)`
   - Consumers check errors with `errors.Is()` or type assertions

### 4. Return Type: Output Structure
The `Diff()` function returns a pointer to an `Output` struct:
```go
type Output struct {
    Packages []PackageInfo
}

type PackageInfo struct {
    Name    string
	IsRootRequirement bool  // True if package is in `require` in composer.json (if provided); otherwise true if in `packages` section in composer.lock
    IsRootDevRequirement bool // True if package is in `require-dev` in composer.json (if provided); otherwise true if in `packages-dev` section in composer.lock
    IsAbandoned bool // Comes from `abandoned` property in composer.lock, default false if not present
	Link  string // Optional, may be empty
	Update UpdateType
	Previous *PkgVersion // Optional, may be nil (for added packages)
	Current *PkgVersion // Optional, may be nil (for removed packages)
}

type UpdateType struct {
	Type string // 'ADDED', 'REMOVED', 'UPDATED', 'NONE', 'UNKNOWN'
	SubType string // Only for Type='UPDATED': 'MAJOR', 'MINOR', 'PATCH'. For other types: 'NONE'
	Direction string // Only for Type='UPDATED': 'UP', 'DOWN', 'NONE', 'UNKNOWN'. For other types: 'NONE'
}

// PkgVersion is a discriminated union: Previous/Current fields in PackageInfo
// will contain either *PkgVersionTag or *PkgVersionCommit. Use the Type field to discriminate.
type PkgVersion struct {
    Full string  // Full version string from composer.lock
    Type string  // 'TAG' or 'COMMIT'
}

type PkgVersionTag struct {
	Full string,  // Full version string from composer.lock (`version` field)
	Type string,  // 'TAG' (always)
	Major string, // Semver Major version (as string, e.g., "1")
	Minor string, // Semver Minor version (as string, e.g., "2")
	Patch string, // Semver Patch version (as string, e.g., "3")
	Extra *string, // Optional pre-release/build metadata (e.g., "-beta.1", "+build123").
                   // Full capture from regex group 5, not parsed further.
}

type PkgVersionCommit struct {
	Full string,   // Full version string from composer.lock (commit hash)
	Type string,   // 'COMMIT' (always)
	Commit string, // Commit hash from `dist.reference` (preferred) or `source.reference`
}
```

### 5. Code Structure
```
project-root/
├── composer/
│   ├── parser.go          # JSON parsing and validation
│   ├── types.go           # Data structures (ComposerLock, Package, etc.)
│   └── parser_test.go     # Parser tests
├── diff/
│   ├── analyzer.go        # Main Diff() function and comparison logic
│   ├── types.go           # Data structures (Output, PackageInfo, UpdateType, etc.)
│   └── analyzer_test.go   # Analyzer tests
├── testdata/
│   ├── composer-simple.lock      # Simple test fixture
│   ├── composer-complex.lock     # Complex test fixture (10k+ packages)
│   ├── composer-invalid.json     # Invalid JSON fixture
│   └── composer-invalid-format.json  # Valid JSON but invalid composer.lock format
├── go.mod                        # Go module definition
├── go.sum                        # Dependency checksums
├── README.md                     # Usage documentation and examples
├── SPECIFICATION.md              # This file
└── Makefile                      # Build targets (optional)
```

**Note**: No `main.go` or `output/formatter.go` – this is a library module, not a CLI tool. Consuming projects handle output formatting.

---

## Non-Functional Requirements

### 1. Performance
- Must handle large composer.lock files efficiently (thousands of packages)
- **Performance targets:**
  - Parse and diff files with 10,000 packages in < 500ms
  - Memory usage < 100MB for typical 10,000 package lock file
  - Output generation (struct population) < 100ms
- Single-pass parsing preferred

### 2. Robustness
- Graceful error handling with informative messages
- No panic on invalid input
- Proper error wrapping with context using `fmt.Errorf("context: %w", err)`
- Handle all documented edge cases without crashing

### 3. Portability
- Cross-platform compatible (Windows, macOS, Linux)
- Library imports work identically on all platforms
- Go 1.23+ compatibility

### 4. Code Quality
- Clear, idiomatic Go code
- Proper error handling with `error` interface
- Comments for complex logic (especially semver parsing and diff algorithm)

---

## Implementation Notes

### 1. JSON Parsing Strategy
Use Go's standard `encoding/json` package for parsing composer.lock files.

### 2. Comparison Logic
- Create a map-based index of packages from each file for O(1) lookup
- Iterate through packages to identify differences.
- Handle package name as unique identifier
- Maintain separate tracking for regular packages (from `packages` property) and dev packages (from `packages-dev` property)

### 3. Version Comparison
- Simple string comparison is acceptable for version identification
- Use semver parsing in order to detect the update direction. In case version can't be parsed (commit), use 'UNKNOWN' direction.

#### Semver Parsing Strategy
Version strings must be parsed using the following regex pattern to extract semver components:

```
^(v)?(\d+)\.(\d+)\.(\d+)(.*)$
```

**Parsing rules:**
- Optional leading `v` character (e.g., `v1.2.3` or `1.2.3` are both valid)
- Three required numeric components: MAJOR.MINOR.PATCH
- Optional extra suffix for pre-release or build metadata (e.g., `-alpha`, `-beta.1`, `+build123`)
- Capture groups: (1) optional 'v', (2) MAJOR, (3) MINOR, (4) PATCH, (5) extra suffix
- Only attempt semver parsing when PkgVersion.Type is TAG, otherwise use 'UNKNOWN' direction

**Examples:**
- `1.2.3` → MAJOR=1, MINOR=2, PATCH=3, extra=nil
- `v2.0.0` → MAJOR=2, MINOR=0, PATCH=0, extra=nil
- `1.0.0-beta` → MAJOR=1, MINOR=0, PATCH=0, extra="-beta"
- `2.5.1+build.1` → MAJOR=2, MINOR=5, PATCH=1, extra="+build.1"
- `abc123def` → Cannot parse, use 'UNKNOWN' direction

**Direction Detection:**
- Compare MAJOR first: if different, direction is UP (if B > A) or DOWN (if B < A)
- If MAJOR equal, compare MINOR; same logic
- If MAJOR and MINOR equal, compare PATCH; same logic
- If all numeric components (MAJOR, MINOR, PATCH) equal: direction is NONE (versions are equal)
- If regex match fails: direction is UNKNOWN (likely a commit hash)
- If either A or B version cannot be parsed as semver: direction is UNKNOWN

### 4. Edge Cases to Handle
- **Empty packages list**: Both files have empty "packages" arrays – return Output with zero PackageInfo entries
- **Duplicate package names**: Should not occur in valid composer.lock, but if present, use first occurrence
- **Packages with identical names but different cases**: Treated as different packages (case-sensitive)
- **Missing version field**: If `version` field missing, attempt to use commit reference; if both missing, mark version as 'UNKNOWN'
- **Missing or empty `packages-dev` section**: Valid state; treat as no dev packages
- **Identical files (fileA == fileB)**: Return Output with empty added/removed/updated lists
- **Both files empty**: Valid state; return Output with empty Packages list
- **Special version formats**: 
  - `dev-master`: Not semver; direction is 'UNKNOWN'
  - `@dev`: Not semver; direction is 'UNKNOWN'
  - Commit hashes: Type is 'COMMIT'; direction is 'NONE' if both are commits (no version comparison possible)

### 5. Special Cases for Output
- **Packages in both `packages` and `packages-dev`**: This shouldn't occur in valid composer.lock. If it does, treat each section independently (may result in ADDED/REMOVED for same package if in different sections between files)
- **IsRootRequirement / IsRootDevRequirement**: 
  - If composer.json is provided: Check `require` and `require-dev` fields to set these flags
  - If composer.json is NOT provided: Use composer.lock sections (`packages` → IsRootRequirement=true, `packages-dev` → IsRootDevRequirement=true)
  - These fields indicate whether package is directly required (in composer.json) vs transitive
  - Note: A package can have BOTH IsRootRequirement=true AND IsRootDevRequirement=true if listed in both require and require-dev
- **IsAbandoned**: Read from `abandoned` field in composer.lock; defaults to false if field missing
- **Link priority**: Use first available: `support.wiki` → `support.docs` → `support.source` → `homepage`. If none available, Link is empty string

---

## Test Fixtures

The `testdata/` directory must contain the following test files:

### 1. `composer-simple.lock`
Minimal valid composer.lock with:
- 3 packages in `packages` array
- 1 package in `packages-dev` array
- Mix of versions with and without 'v' prefix
- All support/homepage fields present for testing link resolution

### 2. `composer-simple.json`
Composer.json paired with composer-simple.lock:
```json
{
  "require": {
    "vendor/package-a": "^1.0",
    "vendor/package-b": ">=2.0"
  },
  "require-dev": {
    "vendor/package-dev-a": "^1.0"
  }
}
```

### 3. `composer-complex.lock`
Resolved: composer-complex fixture updated to 100+ packages and 50+ dev packages.
Large realistic composer.lock with:
- 100+ packages in `packages` array
- 50+ packages in `packages-dev` array
- Mix of semver versions (1.0.0, v2.1.3, 1.0.0-beta)
- Mix of commit hashes (40-char hex strings)
- Some packages missing optional fields (no support, no homepage)
- Some packages with `abandoned: true`

### 4. `composer-complex.json`
Composer.json paired with composer-complex.lock:
```json
{
  "require": {
    "vendor/popular-lib": "^2.0",
    "another-vendor/framework": "^3.5"
  },
  "require-dev": {
    "testing/unittest-suite": "^1.0",
    "static-analysis/code-checker": "^2.0"
  }
}
```

### 5. `composer-invalid.json`
Invalid JSON file (syntax error):
- Malformed JSON (unclosed brace, trailing comma, etc.)
- Used for testing error handling

### 6. `composer-invalid-format.json`
Valid JSON but invalid composer.lock format:
- Missing `packages` array
- Or `packages` is not an array
- Or missing required `name` or `version` fields on packages

---

## Success Criteria

✓ Correctly identifies all added packages (regular and dev)

✓ Correctly identifies all removed packages (regular and dev)

✓ Correctly identifies all updated packages (regular and dev) with correct UpdateType values

✓ Clearly distinguishes between regular and dev packages (IsRootRequirement / IsRootDevRequirement fields)

✓ Proper error handling for invalid JSON and invalid composer.lock format

✓ Semver version comparison produces correct SubType (MAJOR/MINOR/PATCH) and Direction (UP/DOWN/NONE/UNKNOWN)

✓ Returns structured *Output with properly populated PackageInfo entries

✓ No external dependencies (standard library only)

✓ Handles all documented edge cases and special cases without panicking

✓ Performance targets met (< 500ms for 10k packages, < 100MB memory)

✓ Table-driven tests cover all comparison scenarios (added, removed, updated, semver parsing, commit versions)