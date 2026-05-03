package summary

import "github.com/yoanm/go-deps-diff/shared"

type MarkdownSection string

const (
	CautionSection   MarkdownSection = "CAUTION"
	WarningSection   MarkdownSection = "WARNING"
	ImportantSection MarkdownSection = "IMPORTANT"
	TipSection       MarkdownSection = "TIP"
	NoteSection      MarkdownSection = "NOTE"
)

type MarkdownCategory string

const (
	ProdUsageCategory    MarkdownCategory = "PROD_USAGE"
	DevOnlyUsageCategory MarkdownCategory = "DEV_ONLY_USAGE"
)

type MarkdownSubCategory string

const (
	RequirementSubCategory MarkdownSubCategory = "REQUIREMENT"
	TransitiveSubCategory  MarkdownSubCategory = "TRANSITIVE"
)

type MarkdownItem string

const (
	UnknownUpdateItem        MarkdownItem = "UNKNOWN_UPDATE"
	SemverMajorUpgradeItem   MarkdownItem = "SEMVER_MAJOR_UPGRADE"
	SemverMinorUpgradeItem   MarkdownItem = "SEMVER_MINOR_UPGRADE"
	SemverPatchUpgradeItem   MarkdownItem = "SEMVER_PATCH_UPGRADE"
	SemverMajorDowngradeItem MarkdownItem = "SEMVER_MAJOR_DOWNGRADE"
	SemverMinorDowngradeItem MarkdownItem = "SEMVER_MINOR_DOWNGRADE"
	SemverPatchDowngradeItem MarkdownItem = "SEMVER_PATCH_DOWNGRADE"
	RemovalItem              MarkdownItem = "REMOVAL"
	AdditionItem             MarkdownItem = "ADDITION"
	SameItem                 MarkdownItem = "SAME"
)

type (
	PkgList          []*shared.PackageChange
	ItemsMap         map[MarkdownItem]PkgList
	SubCategoriesMap map[MarkdownSubCategory]ItemsMap
	CategoriesMap    map[MarkdownCategory]SubCategoriesMap
	SectionsMap      map[MarkdownSection]CategoriesMap
)
