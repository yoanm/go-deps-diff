package composer

import (
	"testing"
)

func TestParseLock(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		wantError bool
		checkFn   func(*ComposerLock) bool
	}{
		{
			name: "valid simple lock",
			data: []byte(`{
				"packages": [
					{"name": "vendor/pkg", "version": "1.0.0"}
				]
			}`),
			wantError: false,
			checkFn: func(lock *ComposerLock) bool {
				return lock != nil && len(lock.Packages) == 1 &&
					lock.Packages[0].Name == "vendor/pkg"
			},
		},
		{
			name: "valid lock with packages-dev",
			data: []byte(`{
				"packages": [],
				"packages-dev": [
					{"name": "vendor/test", "version": "1.0.0"}
				]
			}`),
			wantError: false,
			checkFn: func(lock *ComposerLock) bool {
				return lock != nil && len(lock.PackagesDev) == 1 &&
					lock.PackagesDev[0].Name == "vendor/test"
			},
		},
		{
			name:      "invalid json",
			data:      []byte(`{invalid}`),
			wantError: true,
		},
		{
			name:      "empty input",
			data:      []byte{},
			wantError: true,
		},
		{
			name:      "missing packages arrays",
			data:      []byte(`{"other": "field"}`),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lock, err := ParseLock(tt.data)
			if (err != nil) != tt.wantError {
				t.Errorf("ParseLock() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError && tt.checkFn != nil && !tt.checkFn(lock) {
				t.Errorf("ParseLock() result check failed")
			}
		})
	}
}

func TestParseJson(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		wantError bool
		checkFn   func(*ComposerJson) bool
	}{
		{
			name: "valid composer.json",
			data: []byte(`{
				"require": {"vendor/pkg": "^1.0"},
				"require-dev": {"vendor/test": "^1.0"}
			}`),
			wantError: false,
			checkFn: func(json *ComposerJson) bool {
				return json != nil && json.Require["vendor/pkg"] == "^1.0" &&
					json.RequireDev["vendor/test"] == "^1.0"
			},
		},
		{
			name:      "invalid json",
			data:      []byte(`{invalid}`),
			wantError: true,
		},
		{
			name:      "empty input",
			data:      []byte{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			json, err := ParseJson(tt.data)
			if (err != nil) != tt.wantError {
				t.Errorf("ParseJson() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError && tt.checkFn != nil && !tt.checkFn(json) {
				t.Errorf("ParseJson() result check failed")
			}
		})
	}
}

func TestIsAbandoned(t *testing.T) {
	tests := []struct {
		name     string
		pkg      *Package
		wantTrue bool
	}{
		{
			name:     "nil package",
			pkg:      nil,
			wantTrue: false,
		},
		{
			name:     "abandoned false",
			pkg:      &Package{Abandoned: false},
			wantTrue: false,
		},
		{
			name:     "abandoned true",
			pkg:      &Package{Abandoned: true},
			wantTrue: true,
		},
		{
			name:     "abandoned string true",
			pkg:      &Package{Abandoned: "true"},
			wantTrue: true,
		},
		{
			name:     "abandoned replacement string",
			pkg:      &Package{Abandoned: "https://example.com/replacement"},
			wantTrue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAbandoned(tt.pkg)
			if got != tt.wantTrue {
				t.Errorf("IsAbandoned() = %v, want %v", got, tt.wantTrue)
			}
		})
	}
}

func TestGetLink(t *testing.T) {
	tests := []struct {
		name     string
		pkg      *Package
		wantLink string
	}{
		{
			name:     "nil package",
			pkg:      nil,
			wantLink: "",
		},
		{
			name: "wiki link priority",
			pkg: &Package{
				Support: &Support{
					Wiki:   "https://wiki.example.com",
					Docs:   "https://docs.example.com",
					Source: "https://source.example.com",
				},
				Homepage: "https://homepage.example.com",
			},
			wantLink: "https://wiki.example.com",
		},
		{
			name: "docs link (no wiki)",
			pkg: &Package{
				Support: &Support{
					Docs:   "https://docs.example.com",
					Source: "https://source.example.com",
				},
				Homepage: "https://homepage.example.com",
			},
			wantLink: "https://docs.example.com",
		},
		{
			name: "source link (no wiki/docs)",
			pkg: &Package{
				Support: &Support{
					Source: "https://source.example.com",
				},
				Homepage: "https://homepage.example.com",
			},
			wantLink: "https://source.example.com",
		},
		{
			name: "homepage link (no support)",
			pkg: &Package{
				Homepage: "https://homepage.example.com",
			},
			wantLink: "https://homepage.example.com",
		},
		{
			name:     "no links",
			pkg:      &Package{},
			wantLink: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLink(tt.pkg)
			if got != tt.wantLink {
				t.Errorf("GetLink() = %s, want %s", got, tt.wantLink)
			}
		})
	}
}

func TestGetCommitReference(t *testing.T) {
	tests := []struct {
		name    string
		pkg     *Package
		wantRef string
	}{
		{
			name:    "nil package",
			pkg:     nil,
			wantRef: "",
		},
		{
			name: "dist reference",
			pkg: &Package{
				Dist: &VersionReference{
					Reference: "abc123",
				},
			},
			wantRef: "abc123",
		},
		{
			name: "source reference (no dist)",
			pkg: &Package{
				Source: &VersionReference{
					Reference: "def456",
				},
			},
			wantRef: "def456",
		},
		{
			name: "dist preferred over source",
			pkg: &Package{
				Dist: &VersionReference{
					Reference: "abc123",
				},
				Source: &VersionReference{
					Reference: "def456",
				},
			},
			wantRef: "abc123",
		},
		{
			name:    "no reference",
			pkg:     &Package{},
			wantRef: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCommitReference(tt.pkg)
			if got != tt.wantRef {
				t.Errorf("GetCommitReference() = %s, want %s", got, tt.wantRef)
			}
		})
	}
}
