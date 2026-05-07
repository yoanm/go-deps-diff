package summary

import (
	"github.com/yoanm/go-deps-diff/shared"
)

func buildSectionsMap(changes shared.DiffMap) sectionsMap {
	list := sectionsMap{}

	for _, change := range changes {
		categoryType, subCategoryType := getMarkdownCategoryType(change)
		itemType := getMarkdownItemType(change)
		sectionType := getMarkdownSectionType(subCategoryType, itemType, change.Package)

		switch {
		case list[sectionType] == nil:
			list[sectionType] = make(categoriesMap)

			fallthrough
		case list[sectionType][categoryType] == nil:
			list[sectionType][categoryType] = make(subCategoriesMap)

			fallthrough
		case list[sectionType][categoryType][subCategoryType] == nil:
			list[sectionType][categoryType][subCategoryType] = make(itemsMap)

			fallthrough
		case list[sectionType][categoryType][subCategoryType][itemType] == nil:
			list[sectionType][categoryType][subCategoryType][itemType] = pkgList{}
		}

		// _debugPackageList(sectionType, categoryType, subCategoryType, itemType, change)

		list[sectionType][categoryType][subCategoryType][itemType] = append(
			list[sectionType][categoryType][subCategoryType][itemType],
			change,
		)
	}

	return list
}

func getMarkdownSectionType(
	subCategoryType markdownSubCategory,
	itemType markdownItem,
	pkg shared.PkgWrapper,
) markdownSection {
	switch {
	case isCandidateForCautionSection(subCategoryType, itemType, pkg):
		return cautionSection
	case isCandidateForWarningSection(subCategoryType, itemType, pkg):
		return warningSection
	case isCandidateForImportantSection(subCategoryType, itemType, pkg):
		return importantSection
	case isCandidateForTipSection(subCategoryType, itemType):
		return tipSection
	}

	//# Note
	//## Production usage
	//### Requirements
	//- SEMVER_PATCH_UPGRADE
	//- SAME
	//### Transitives
	//- REMOVAL
	//- SEMVER_MINOR_UPGRADE
	//- SEMVER_PATCH_UPGRADE
	//- ADDITION
	//- SAME
	//## Dev-only usage
	//### Requirements
	//- SEMVER_PATCH_UPGRADE
	//- SAME
	//### Transitives
	//- REMOVAL
	//- SEMVER_MINOR_UPGRADE
	//- SEMVER_PATCH_UPGRADE
	//- ADDITION
	//- SAME

	return noteSection
}

func isCandidateForCautionSection(
	subCategoryType markdownSubCategory,
	itemType markdownItem,
	pkg shared.PkgWrapper,
) bool {
	//# Caution
	//## Production usage
	//### Requirements
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	//- ADDITION__NOT_SEMVER
	//### Transitives
	//## Dev-only usage
	//### Requirements
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	//- ADDITION__NOT_SEMVER
	//### Transitives
	// -> process prod usage and dev-only usage the same way
	return subCategoryType == requirementSubCategory &&
		(itemType == unknownUpdateItem ||
			itemType == semverMajorDowngradeItem ||
			(itemType == additionItem && (pkg.IsAbandoned() || pkg.GetVersion().Semver == nil)))
}

func isCandidateForWarningSection( //nolint:cyclop // 11 vs 10 allowed but easier to maintain that way
	subCategoryType markdownSubCategory,
	itemType markdownItem,
	pkg shared.PkgWrapper,
) bool {
	//# Warning
	//## Production usage
	//### Requirements
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	//- SAME__NOT_SEMVER
	//### Transitives
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	//- ADDITION__NOT_SEMVER
	//## Dev-only usage
	//### Requirements
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	//- SAME__NOT_SEMVER
	//### Transitives
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	//- ADDITION__NOT_SEMVER
	// -> process prod usage and dev-only usage the same way
	if subCategoryType == requirementSubCategory &&
		((itemType == semverMajorUpgradeItem || itemType == semverMinorDowngradeItem) ||
			(itemType == sameItem && pkg.GetVersion().Semver == nil)) {
		return true
	}

	return subCategoryType == transitiveSubCategory &&
		(itemType == unknownUpdateItem ||
			itemType == semverMajorDowngradeItem ||
			(itemType == additionItem && (pkg.IsAbandoned() || pkg.GetVersion().Semver == nil)))
}

func isCandidateForImportantSection(
	subCategoryType markdownSubCategory,
	itemType markdownItem,
	pkg shared.PkgWrapper,
) bool {
	// # Important
	//## Production usage
	//### Requirements
	//- SEMVER_PATCH_DOWNGRADE
	//- REMOVAL
	//### Transitives
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	//- SAME__NOT_SEMVER
	//## Dev-only usage
	//### Requirements
	//- SEMVER_PATCH_DOWNGRADE
	//- REMOVAL
	//### Transitives
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	//- SAME__NOT_SEMVER
	// -> process prod usage and dev-only usage the same way
	if subCategoryType == requirementSubCategory && (itemType == semverPatchDowngradeItem || itemType == removalItem) {
		return true
	}

	return subCategoryType == transitiveSubCategory &&
		((itemType == semverMajorUpgradeItem || itemType == semverMinorDowngradeItem) ||
			(itemType == sameItem && pkg.GetVersion().Semver == nil))
}

func isCandidateForTipSection(
	subCategoryType markdownSubCategory,
	itemType markdownItem,
) bool {
	// # Tip
	//## Production usage
	//### Requirements
	//- SEMVER_MINOR_UPGRADE
	//- ADDITION
	//### Transitives
	//- SEMVER_PATCH_DOWNGRADE
	//- REMOVAL
	//## Dev-only usage
	//### Requirements
	//- SEMVER_MINOR_UPGRADE
	//- ADDITION
	//### Transitives
	//- SEMVER_PATCH_DOWNGRADE
	//- REMOVAL
	// -> process prod usage and dev-only usage the same way
	if subCategoryType == requirementSubCategory && (itemType == semverMinorUpgradeItem || itemType == additionItem) {
		return true
	}

	return subCategoryType == transitiveSubCategory && (itemType == semverPatchDowngradeItem || itemType == removalItem)
}

func getMarkdownCategoryType(change *shared.PackageChange) (markdownCategory, markdownSubCategory) {
	category := prodUsageCategory // By default, for security, better than defining a package as dev-only while it's not
	if change.Package.IsDevOnly() {
		category = devOnlyUsageCategory
	}

	if change.Package.IsRootRequirement() || change.Package.IsRootDevRequirement() {
		return category, requirementSubCategory
	} else {
		return category, transitiveSubCategory
	}
}

func getMarkdownItemType(change *shared.PackageChange) markdownItem {
	mrkItem, exists := getOperationToItemBaseMap()[change.Operation.Name]
	if exists {
		return mrkItem
	}

	//nolint:exhaustive // Only those cases are managed (others are expected to come from the map above)
	switch change.Operation.Name {
	// - SEMVER_MAJOR_UPGRADE
	// - SEMVER_MINOR_UPGRADE
	// - SEMVER_PATCH_UPGRADE
	case shared.UpgradeOperation:
		//nolint:exhaustive // SemverExtra + SemverUnknown + SemverNoUpdate managed as unknownUpdateItem !
		switch change.Operation.SemverType {
		case shared.SemverMajorUpdate:
			return semverMajorUpgradeItem
		case shared.SemverMinorUpdate:
			return semverMinorUpgradeItem
		case shared.SemverPatchUpdate:
			return semverPatchUpgradeItem
		}
	// - SEMVER_MAJOR_DOWNGRADE
	// - SEMVER_MINOR_DOWNGRADE
	// - SEMVER_PATCH_DOWNGRADE
	case shared.DowngradeOperation:
		//nolint:exhaustive // SemverExtra + SemverUnknown + SemverNoUpdate managed as unknownUpdateItem !
		switch change.Operation.SemverType {
		case shared.SemverMajorUpdate:
			return semverMajorDowngradeItem
		case shared.SemverMinorUpdate:
			return semverMinorDowngradeItem
		case shared.SemverPatchUpdate:
			return semverPatchDowngradeItem
		}
	}

	return unknownUpdateItem // Fallback on unknown
}

/*
func _debugPackageList(
	sectionType markdownSection,
	categoryType markdownCategory,
	subCategoryType markdownSubCategory,
	itemType markdownItem,
	change *shared.PackageChange,
) {
	expectedTypeKey := strings.ToLower(string(sectionType)) +
		"-" + strings.ToLower(string(categoryType)) +
		"-" + strings.ToLower(string(subCategoryType))
	if !change.Package.IsDevOnly() && change.Package.IsRootDevRequirement() {
		expectedTypeKey += "+dev_req"
	}

	expectedTypeKey += "/" + string(itemType)
	if itemType == unknownUpdateItem && change.Operation.SemverType == shared.SemverExtraUpdate {
		expectedTypeKey += "+SEMVER_EXTRA"
	} else if change.Package.IsAbandoned() {
		expectedTypeKey += "+ABANDONED"
	}

	if change.Package.GetName() != expectedTypeKey {
		slog.Error("package mismatch: got " + change.Package.GetName() + ", want " + expectedTypeKey)
	}
}
*/
