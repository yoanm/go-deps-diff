package composer

import "testing"

func TestIsAbandonedPkg(t *testing.T) {
	t.Parallel()

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
			pkg:      &Package{Abandoned: false}, //nolint:exhaustruct // Useless for the test purpose
			wantTrue: false,
		},
		{
			name:     "abandoned true",
			pkg:      &Package{Abandoned: true}, //nolint:exhaustruct // Useless for the test purpose
			wantTrue: true,
		},
		{
			name:     "abandoned string true",
			pkg:      &Package{Abandoned: "true"}, //nolint:exhaustruct // Useless for the test purpose
			wantTrue: true,
		},
		{
			name: "abandoned replacement string",
			//nolint:exhaustruct // Useless for the test purpose
			pkg:      &Package{Abandoned: "https://example.com/replacement"},
			wantTrue: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := isAbandonedPkg(testCase.pkg)
			if got != testCase.wantTrue {
				t.Errorf("IsAbandoned() = %v, want %v", got, testCase.wantTrue)
			}
		})
	}
}

func TestGetPkgLink(t *testing.T) {
	t.Parallel()

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
			pkg: &Package{ //nolint:exhaustruct // Useless for the test purpose
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
			pkg: &Package{ //nolint:exhaustruct // Useless for the test purpose
				Support: &Support{ //nolint:exhaustruct // Useless for the test purpose
					Docs:   "https://docs.example.com",
					Source: "https://source.example.com",
				},
				Homepage: "https://homepage.example.com",
			},
			wantLink: "https://docs.example.com",
		},
		{
			name: "source link (no wiki/docs)",
			pkg: &Package{ //nolint:exhaustruct // Useless for the test purpose
				Support: &Support{ //nolint:exhaustruct // Useless for the test purpose
					Source: "https://source.example.com",
				},
				Homepage: "https://homepage.example.com",
			},
			wantLink: "https://source.example.com",
		},
		{
			name: "homepage link (no support)",
			pkg: &Package{ //nolint:exhaustruct // Useless for the test purpose
				Homepage: "https://homepage.example.com",
			},
			wantLink: "https://homepage.example.com",
		},
		{
			name:     "no links",
			pkg:      &Package{}, //nolint:exhaustruct // Useless for the test purpose
			wantLink: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := getPkgLink(testCase.pkg)
			if got != testCase.wantLink {
				t.Errorf("GetLink() = %s, want %s", got, testCase.wantLink)
			}
		})
	}
}

func TestGetPkgRef(t *testing.T) {
	t.Parallel()

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
			pkg: &Package{ //nolint:exhaustruct // Useless for the test purpose
				Dist: &VersionReference{
					Reference: "abc123",
				},
			},
			wantRef: "abc123",
		},
		{
			name: "source reference (no dist)",
			pkg: &Package{ //nolint:exhaustruct // Useless for the test purpose
				Source: &VersionReference{
					Reference: "def456",
				},
			},
			wantRef: "def456",
		},
		{
			name: "dist preferred over source",
			pkg: &Package{ //nolint:exhaustruct // Useless for the test purpose
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
			pkg:     &Package{}, //nolint:exhaustruct // Useless for the test purpose
			wantRef: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := getPkgRef(testCase.pkg)
			if got != testCase.wantRef {
				t.Errorf("getPkgRef() = %s, want %s", got, testCase.wantRef)
			}
		})
	}
}
