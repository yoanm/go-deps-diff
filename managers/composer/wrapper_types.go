package composer

import "github.com/yoanm/go-deps-diff/shared"

type ComposerPackageWrapper struct {
	name                 string
	isAbandoned          bool
	version              shared.PkgVersion
	link                 string
	isDevOnly            bool // true if only in lock file "packages-dev" section (dev-only dependency)
	isRootRequirement    bool // true if exists in requirement file "require" section
	isRootDevRequirement bool // true if exists in requirement file "require-dev" section
}

func (w *ComposerPackageWrapper) GetName() string {
	return w.name
}

func (w *ComposerPackageWrapper) IsAbandoned() bool {
	return w.isAbandoned
}

func (w *ComposerPackageWrapper) GetVersion() *shared.PkgVersion {
	return &w.version
}

func (w *ComposerPackageWrapper) GetLink() string {
	return w.link
}

func (w *ComposerPackageWrapper) IsDevOnly() bool {
	return w.isDevOnly
}

func (w *ComposerPackageWrapper) IsRootRequirement() bool {
	return w.isRootRequirement
}

func (w *ComposerPackageWrapper) IsRootDevRequirement() bool {
	return w.isRootDevRequirement
}
