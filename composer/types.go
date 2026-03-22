package composer

// ComposerLock represents the structure of a composer.lock file
type ComposerLock struct {
	Packages    []Package `json:"packages"`
	PackagesDev []Package `json:"packages-dev"`
}

// Package represents a single package entry in composer.lock
type Package struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Source    *VersionReference `json:"source,omitempty"`
	Dist      *VersionReference `json:"dist,omitempty"`
	Support   *Support          `json:"support,omitempty"`
	Homepage  string            `json:"homepage,omitempty"`
	Abandoned interface{}       `json:"abandoned,omitempty"` // Can be bool or string
}

// VersionReference contains the reference (commit hash or tag)
type VersionReference struct {
	Reference string `json:"reference,omitempty"`
	URL       string `json:"url,omitempty"`
}

// Support contains links to documentation and support
type Support struct {
	Wiki   string `json:"wiki,omitempty"`
	Docs   string `json:"docs,omitempty"`
	Source string `json:"source,omitempty"`
}

// ComposerReq represents the structure of a composer.json file (composer requirement)
type ComposerReq struct {
	Require    map[string]string `json:"require,omitempty"`
	RequireDev map[string]string `json:"require-dev,omitempty"`
}
