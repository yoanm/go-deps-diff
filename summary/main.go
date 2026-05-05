package summary

import (
	"fmt"
	"os"
	"strings"

	"github.com/yoanm/go-deps-diff/shared"
)

func GenerateForChanges(changes shared.DiffMap) string {
	return generate(buildDefaultSectionsMap(changes))
}

func buildDefaultSectionsMap(changes shared.DiffMap) sectionsMap {
	list := sectionsMap{}

	for _, change := range changes {
		categoryType, subCategoryType := getMarkdownCategoryType(change)
		itemType := getMarkdownItemType(change)
		sectionType := getMarkdownSectionType(subCategoryType, itemType, change.Package.IsAbandoned())

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
	isAbandoned bool,
) markdownSection {
	switch {
	case isCandidateForCautionSection(subCategoryType, itemType, isAbandoned):
		return cautionSection
	case isCandidateForWarningSection(subCategoryType, itemType, isAbandoned):
		return warningSection
	case isCandidateForImportantSection(subCategoryType, itemType):
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

func isCandidateForCautionSection(subCategoryType markdownSubCategory, itemType markdownItem, isAbandoned bool) bool {
	//# Caution
	//## Production usage
	//### Requirements
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	//### Transitives
	//## Dev-only usage
	//### Requirements
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	//### Transitives
	// -> process prod usage and dev-only usage the same way
	return subCategoryType == requirementSubCategory &&
		(itemType == unknownUpdateItem ||
			itemType == semverMajorDowngradeItem ||
			(itemType == additionItem && isAbandoned))
}

func isCandidateForWarningSection(subCategoryType markdownSubCategory, itemType markdownItem, isAbandoned bool) bool {
	//# Warning
	//## Production usage
	//### Requirements
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	//### Transitives
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	//## Dev-only usage
	//### Requirements
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	//### Transitives
	//- UNKNOWN_UPDATE
	//- SEMVER_MAJOR_DOWNGRADE
	//- ADDITION__ABANDONED
	// -> process prod usage and dev-only usage the same way
	if subCategoryType == requirementSubCategory &&
		(itemType == semverMajorUpgradeItem || itemType == semverMinorDowngradeItem) {
		return true
	}

	return subCategoryType == transitiveSubCategory &&
		(itemType == unknownUpdateItem ||
			itemType == semverMajorDowngradeItem ||
			(itemType == additionItem && isAbandoned))
}

func isCandidateForImportantSection(subCategoryType markdownSubCategory, itemType markdownItem) bool {
	// # Important
	//## Production usage
	//### Requirements
	//- SEMVER_PATCH_DOWNGRADE
	//- REMOVAL
	//### Transitives
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	//## Dev-only usage
	//### Requirements
	//- SEMVER_PATCH_DOWNGRADE
	//- REMOVAL
	//### Transitives
	//- SEMVER_MAJOR_UPGRADE
	//- SEMVER_MINOR_DOWNGRADE
	// -> process prod usage and dev-only usage the same way
	if subCategoryType == requirementSubCategory && (itemType == semverPatchDowngradeItem || itemType == removalItem) {
		return true
	}

	return subCategoryType == transitiveSubCategory &&
		(itemType == semverMajorUpgradeItem || itemType == semverMinorDowngradeItem)
}

func isCandidateForTipSection(subCategoryType markdownSubCategory, itemType markdownItem) bool {
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

func getMarkdownItemType(change *shared.PackageChange) markdownItem { //nolint:cyclop,lll // 13 vs 10 allowed, but 13 actual cases
	switch change.Operation.Name {
	// - UNKNOWN_UPDATE
	// - UNKNOWN_UPDATE
	case shared.UnknownUpdateOperation:
		return unknownUpdateItem
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
	// - REMOVAL
	case shared.RemovalOperation:
		return removalItem
	// - ADDITION
	case shared.AdditionOperation:
		return additionItem
	// - SAME
	case shared.NoChangeOperation:
		return sameItem
	}

	return unknownUpdateItem // Fallback on unknown
}

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
		fmt.Fprintln(os.Stderr, "package mismatch: got", change.Package.GetName(), ", want", expectedTypeKey)
	}
}
