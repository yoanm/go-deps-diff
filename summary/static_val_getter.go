package summary

func getSectionHeaderFor(section MarkdownSection) string {
	switch section {
	case CautionSection:
		return "Hazardous changes"
	case WarningSection:
		return "Error-prone changes"
	case ImportantSection:
		return "Noteworthy changes"
	case TipSection:
		return "Pretty safe changes"
	case NoteSection:
		return "Note"
	}

	panic("Unknown section: " + section)
}

func getSectionDescriptionFor(section MarkdownSection) string {
	switch section {
	case CautionSection:
		return "☣️ Changes that are likely unexpected and/or likely to cause trouble"
	case WarningSection:
		return "⚠️ Changes that may not cause trouble in production but are likely unexpected and/or prone to cause trouble" //nolint:lll // Meaningless here
	case ImportantSection:
		return "🕵️ Changes unlikely to cause production issues, but worth noting if problems arise"
	case TipSection:
		return "👀 Changes that are unlikely to cause trouble in production"
	case NoteSection:
		return "ℹ️ All remaining changes, mostly for your information"
	}

	panic("Unknown section: " + section)
}

func getCategoryHeaderFor(category MarkdownCategory) string {
	switch category {
	case ProdUsageCategory:
		return "Production usage <sup>🏭</sup>"
	case DevOnlyUsageCategory:
		return "Dev-only usage <sup>🧪</sup>"
	}

	panic("Unknown category: " + category)
}

func getSectionsOrder() []MarkdownSection {
	return []MarkdownSection{
		CautionSection,
		WarningSection,
		ImportantSection,
		TipSection,
		NoteSection,
	}
}

func getCategoriesOrder() []MarkdownCategory {
	return []MarkdownCategory{
		ProdUsageCategory,
		DevOnlyUsageCategory,
	}
}

func getSubCategoriesOrder() []MarkdownSubCategory {
	return []MarkdownSubCategory{
		RequirementSubCategory,
		TransitiveSubCategory,
	}
}

func getItemsOrder() []MarkdownItem {
	return []MarkdownItem{
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
}
