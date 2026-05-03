package summary

var (
	sectionTitleMap = map[MarkdownSection]string{
		CautionSection:   "Hazardous changes",
		WarningSection:   "Error-prone changes",
		ImportantSection: "Noteworthy changes",
		TipSection:       "Pretty safe changes",
		NoteSection:      "Note",
	}
	sectionHeaderMap = map[MarkdownSection]string{
		CautionSection:   "☣️ Changes that are likely unexpected and/or likely to cause trouble",
		WarningSection:   "⚠️ Changes that may not cause trouble in production but are likely unexpected and/or prone to cause trouble",
		ImportantSection: "🕵️ Changes unlikely to cause production issues, but worth noting if problems arise",
		TipSection:       "👀 Changes that are unlikely to cause trouble in production",
		NoteSection:      "ℹ️ All remaining changes, mostly for your information",
	}
	sectionOrders = []MarkdownSection{
		CautionSection,
		WarningSection,
		ImportantSection,
		TipSection,
		NoteSection,
	}

	categoryOrders = []MarkdownCategory{
		ProdUsageCategory,
		DevOnlyUsageCategory,
	}
	categoryTitleMap = map[MarkdownCategory]string{
		ProdUsageCategory:    "Production usage <sup>🏭</sup>",
		DevOnlyUsageCategory: "Dev-only usage <sup>🧪</sup>",
	}

	subcategoryOrders = []MarkdownSubCategory{
		RequirementSubCategory,
		TransitiveSubCategory,
	}

	itemOrders = []MarkdownItem{
		UnknownUpdateItem,
		SemverMajorDowngradeItem,
		SemverMinorDowngradeItem,
		SemverPatchDowngradeItem,
		SemverMajorUpgradeItem,
		RemovalItem,
		SemverMinorUpgradeItem,
		SemverPatchUpgradeItem,
		AdditionItem,
		SameItem,
	}
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
