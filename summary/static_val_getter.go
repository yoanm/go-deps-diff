package summary

import "github.com/yoanm/go-deps-diff/shared"

func getSectionHeaderFor(section markdownSection) string {
	switch section {
	case cautionSection:
		return "Hazardous changes"
	case warningSection:
		return "Error-prone changes"
	case importantSection:
		return "Noteworthy changes"
	case tipSection:
		return "Pretty safe changes"
	case noteSection:
		return "Note"
	}

	panic("Unknown section: " + section)
}

func getSectionDescriptionFor(section markdownSection) string {
	switch section {
	case cautionSection:
		return "☣️ Changes that are likely unexpected and/or likely to cause trouble"
	case warningSection:
		return "⚠️ Changes that may not cause trouble in production but are likely unexpected and/or prone to cause trouble" //nolint:lll // Meaningless here
	case importantSection:
		return "🕵️ Changes that are unlikely to cause production issues, but worth noting if problems arise"
	case tipSection:
		return "👀 Changes that are unlikely to cause trouble in production"
	case noteSection:
		return "ℹ️ All remaining changes, mostly for your information"
	}

	panic("Unknown section: " + section)
}

func getCategoryHeaderFor(category markdownCategory) string {
	switch category {
	case prodUsageCategory:
		return "Production usage <sup>🏭</sup>"
	case devOnlyUsageCategory:
		return "Dev-only usage <sup>🧪</sup>"
	}

	panic("Unknown category: " + category)
}

func getSectionsOrder() []markdownSection {
	return []markdownSection{
		cautionSection,
		warningSection,
		importantSection,
		tipSection,
		noteSection,
	}
}

func getCategoriesOrder() []markdownCategory {
	return []markdownCategory{
		prodUsageCategory,
		devOnlyUsageCategory,
	}
}

func getSubCategoriesOrder() []markdownSubCategory {
	return []markdownSubCategory{
		requirementSubCategory,
		transitiveSubCategory,
	}
}

func getItemsOrder() []markdownItem {
	return []markdownItem{
		unknownUpdateItem,
		semverMajorDowngradeItem,
		semverMinorDowngradeItem,
		semverPatchDowngradeItem,
		semverMajorUpgradeItem,
		removalItem,
		semverMinorUpgradeItem,
		semverPatchUpgradeItem,
		additionItem,
		sameItem,
	}
}

func getOperationToItemBaseMap() map[shared.OperationName]markdownItem {
	return map[shared.OperationName]markdownItem{ //nolint:exhaustive // Only 1-1 mapping values here !
		shared.UnknownUpdateOperation: unknownUpdateItem,
		shared.RemovalOperation:       removalItem,
		shared.AdditionOperation:      additionItem,
		shared.NoChangeOperation:      sameItem,
	}
}
