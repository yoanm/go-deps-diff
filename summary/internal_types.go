package summary

import "github.com/yoanm/go-deps-diff/shared"

type markdownSection string

const (
	cautionSection   markdownSection = "CAUTION"
	warningSection   markdownSection = "WARNING"
	importantSection markdownSection = "IMPORTANT"
	tipSection       markdownSection = "TIP"
	noteSection      markdownSection = "NOTE"
)

type markdownCategory string

const (
	prodUsageCategory    markdownCategory = "PROD_USAGE"
	devOnlyUsageCategory markdownCategory = "DEV_ONLY_USAGE"
)

type markdownSubCategory string

const (
	requirementSubCategory markdownSubCategory = "REQUIREMENT"
	transitiveSubCategory  markdownSubCategory = "TRANSITIVE"
)

type markdownItem string

const (
	unknownUpdateItem        markdownItem = "UNKNOWN_UPDATE"
	semverMajorUpgradeItem   markdownItem = "SEMVER_MAJOR_UPGRADE"
	semverMinorUpgradeItem   markdownItem = "SEMVER_MINOR_UPGRADE"
	semverPatchUpgradeItem   markdownItem = "SEMVER_PATCH_UPGRADE"
	semverMajorDowngradeItem markdownItem = "SEMVER_MAJOR_DOWNGRADE"
	semverMinorDowngradeItem markdownItem = "SEMVER_MINOR_DOWNGRADE"
	semverPatchDowngradeItem markdownItem = "SEMVER_PATCH_DOWNGRADE"
	removalItem              markdownItem = "REMOVAL"
	additionItem             markdownItem = "ADDITION"
	sameItem                 markdownItem = "SAME"
)

type (
	pkgList          []*shared.PackageChange
	itemsMap         map[markdownItem]pkgList
	subCategoriesMap map[markdownSubCategory]itemsMap
	categoriesMap    map[markdownCategory]subCategoriesMap
	sectionsMap      map[markdownSection]categoriesMap
)

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
