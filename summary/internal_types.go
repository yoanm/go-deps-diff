package summary

type pkgRowMode int

const (
	// versionOnlyPkgRowMode can be used when displaying ONLY one of Added/Removed/Same operations !
	//
	// Only the package name and version will be displayed:
	//
	// - `| name | version |`.
	versionOnlyPkgRowMode pkgRowMode = 2
	// withOperationPkgRowMode can be used when displaying a list containing ONLY Added/Removed/Same operations !
	//
	// The package name, version and operation will be displayed:
	//
	// - Same: `| name | version | operation |`
	// - Removed: `| name | version | operation |`
	// - Added: `| name | operation | version |`.
	withOperationPkgRowMode pkgRowMode = 3
	// fullPkgRowMode is the default mode. It can display a mix of any Operations
	//
	// The package name, version and operation will be displayed, as well as previous version when relevant:
	//
	// - Same: `| name | version | operation (colspan=2) |`
	// - Removed: `| name | version | operation (colspan=2) |`
	// - Added: `| name | operation (colspan=2) | version |`
	// - Any others: `| name | previous version | operation | current version |`.
	fullPkgRowMode pkgRowMode = 4
)
