package composer

import (
	"encoding/json"
	"fmt"
)

// ErrInvalidJSON indicates the input JSON is malformed
type ErrInvalidJSON struct {
	message string
	err     error
}

func (e *ErrInvalidJSON) Error() string {
	return fmt.Sprintf("invalid JSON: %v", e.err)
}

// ErrInvalidFormat indicates the JSON is valid but doesn't match expected structure
type ErrInvalidFormat struct {
	message string
}

func (e *ErrInvalidFormat) Error() string {
	return fmt.Sprintf("invalid format: %s", e.message)
}

// ParseLock parses a composer.lock file from JSON bytes
func ParseLock(data []byte) (*ComposerLock, error) {
	if len(data) == 0 {
		return nil, &ErrInvalidFormat{message: "empty input"}
	}

	var lock ComposerLock
	err := json.Unmarshal(data, &lock)
	if err != nil {
		return nil, &ErrInvalidJSON{err: err}
	}

	// Validate required structure
	if lock.Packages == nil && lock.PackagesDev == nil {
		return nil, &ErrInvalidFormat{
			message: "must contain 'packages' and/or 'packages-dev' arrays",
		}
	}

	return &lock, nil
}

// ParseReq parses a composer.json (composer requirement) file from JSON bytes
func ParseReq(data []byte) (*ComposerReq, error) {
	if len(data) == 0 {
		return nil, &ErrInvalidFormat{message: "empty input"}
	}

	var result ComposerReq
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, &ErrInvalidJSON{err: err}
	}

	return &result, nil
}

// IsAbandoned safely extracts the abandoned status from the package
// Returns true if the field is explicitly set to true (boolean or string "true")
func IsAbandoned(pkg *Package) bool {
	if pkg == nil || pkg.Abandoned == nil {
		return false
	}

	switch v := pkg.Abandoned.(type) {
	case bool:
		return v
	case string:
		return v == "true" || v != ""
	default:
		return false
	}
}

// GetLink extracts the best available link from a package
// Priority: wiki -> docs -> source -> homepage
func GetLink(pkg *Package) string {
	if pkg == nil {
		return ""
	}

	if pkg.Support != nil {
		if pkg.Support.Wiki != "" {
			return pkg.Support.Wiki
		}
		if pkg.Support.Docs != "" {
			return pkg.Support.Docs
		}
		if pkg.Support.Source != "" {
			return pkg.Support.Source
		}
	}

	if pkg.Homepage != "" {
		return pkg.Homepage
	}

	return ""
}

// GetCommitReference extracts the commit hash from a package
// Prefers dist.reference over source.reference
func GetCommitReference(pkg *Package) string {
	if pkg == nil {
		return ""
	}

	if pkg.Dist != nil && pkg.Dist.Reference != "" {
		return pkg.Dist.Reference
	}

	if pkg.Source != nil && pkg.Source.Reference != "" {
		return pkg.Source.Reference
	}

	return ""
}
