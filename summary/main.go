package summary

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/yoanm/go-deps-diff/shared"
)

func GenerateForChanges(changes shared.DiffMap) string {
	return Generate(buildDefaultSectionsMap(changes))
}

func buildDefaultSectionsMap(changes shared.DiffMap) SectionsMap {
	list := SectionsMap{}
	if len(changes) == 0 {
		return list
	}

	keys := slices.Collect(maps.Keys(changes))
	sort.Strings(keys)

	for _, key := range keys {
		change := changes[key]
		categoryType, subCategoryType := getMarkdownCategoryType(change)
		itemType := getMarkdownItemType(change)
		sectionType := getMarkdownSectionType(subCategoryType, itemType, change.Package.IsAbandoned())

		switch {
		case list[sectionType] == nil:
			list[sectionType] = make(CategoriesMap)

			fallthrough
		case list[sectionType][categoryType] == nil:
			list[sectionType][categoryType] = make(SubCategoriesMap)

			fallthrough
		case list[sectionType][categoryType][subCategoryType] == nil:
			list[sectionType][categoryType][subCategoryType] = make(ItemsMap)

			fallthrough
		case list[sectionType][categoryType][subCategoryType][itemType] == nil:
			list[sectionType][categoryType][subCategoryType][itemType] = PkgList{}
		}

		_debugPackageList(sectionType, categoryType, subCategoryType, itemType, change)

		list[sectionType][categoryType][subCategoryType][itemType] = append(
			list[sectionType][categoryType][subCategoryType][itemType],
			change,
		)
	}

	return list
}

func getMarkdownSectionType(
	subCategoryType MarkdownSubCategory,
	itemType MarkdownItem,
	isAbandoned bool,
) MarkdownSection {
	switch {
	case isCandidateForCautionSection(subCategoryType, itemType, isAbandoned):
		return CautionSection
	case isCandidateForWarningSection(subCategoryType, itemType, isAbandoned):
		return WarningSection
	case isCandidateForImportantSection(subCategoryType, itemType):
		return ImportantSection
	case isCandidateForTipSection(subCategoryType, itemType):
		return TipSection
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

	return NoteSection
}

func isCandidateForCautionSection(subCategoryType MarkdownSubCategory, itemType MarkdownItem, isAbandoned bool) bool {
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
	return subCategoryType == RequirementSubCategory &&
		(itemType == UnknownUpdateItem ||
			itemType == SemverMajorDowngradeItem ||
			(itemType == AdditionItem && isAbandoned))
}

func isCandidateForWarningSection(subCategoryType MarkdownSubCategory, itemType MarkdownItem, isAbandoned bool) bool {
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
	if subCategoryType == RequirementSubCategory &&
		(itemType == SemverMajorUpgradeItem || itemType == SemverMinorDowngradeItem) {
		return true
	}
	return subCategoryType == TransitiveSubCategory &&
		(itemType == UnknownUpdateItem ||
			itemType == SemverMajorDowngradeItem ||
			(itemType == AdditionItem && isAbandoned))
}

func isCandidateForImportantSection(subCategoryType MarkdownSubCategory, itemType MarkdownItem) bool {
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
	if subCategoryType == RequirementSubCategory && (itemType == SemverPatchDowngradeItem || itemType == RemovalItem) {
		return true
	}

	return subCategoryType == TransitiveSubCategory &&
		(itemType == SemverMajorUpgradeItem || itemType == SemverMinorDowngradeItem)
}

func isCandidateForTipSection(subCategoryType MarkdownSubCategory, itemType MarkdownItem) bool {
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
	if subCategoryType == RequirementSubCategory && (itemType == SemverMinorUpgradeItem || itemType == AdditionItem) {
		return true
	}

	return subCategoryType == TransitiveSubCategory && (itemType == SemverPatchDowngradeItem || itemType == RemovalItem)
}

func getMarkdownCategoryType(change *shared.PackageChange) (MarkdownCategory, MarkdownSubCategory) {
	category := ProdUsageCategory // By default, for security, better than defining a package as dev-only while it's not
	if change.Package.IsDevOnly() {
		category = DevOnlyUsageCategory
	}

	if change.Package.IsRootRequirement() || change.Package.IsRootDevRequirement() {
		return category, RequirementSubCategory
	} else {
		return category, TransitiveSubCategory
	}
}

func getMarkdownItemType(change *shared.PackageChange) MarkdownItem { //nolint:cyclop,lll // 13 vs 10 allowed, but 13 actual cases
	switch change.Operation.Name {
	// - UNKNOWN_UPDATE
	// - UNKNOWN_UPDATE
	case shared.UnknownUpdateOperation:
		return UnknownUpdateItem
	// - SEMVER_MAJOR_UPGRADE
	// - SEMVER_MINOR_UPGRADE
	// - SEMVER_PATCH_UPGRADE
	case shared.UpgradeOperation:
		//nolint:exhaustive // SemverExtra + SemverUnknown + SemverNoUpdate managed as UnknownUpdateItem !
		switch change.Operation.SemverType {
		case shared.SemverMajorUpdate:
			return SemverMajorUpgradeItem
		case shared.SemverMinorUpdate:
			return SemverMinorUpgradeItem
		case shared.SemverPatchUpdate:
			return SemverPatchUpgradeItem
		}
	// - SEMVER_MAJOR_DOWNGRADE
	// - SEMVER_MINOR_DOWNGRADE
	// - SEMVER_PATCH_DOWNGRADE
	case shared.DowngradeOperation:
		//nolint:exhaustive // SemverExtra + SemverUnknown + SemverNoUpdate managed as UnknownUpdateItem !
		switch change.Operation.SemverType {
		case shared.SemverMajorUpdate:
			return SemverMajorDowngradeItem
		case shared.SemverMinorUpdate:
			return SemverMinorDowngradeItem
		case shared.SemverPatchUpdate:
			return SemverPatchDowngradeItem
		}
	// - REMOVAL
	case shared.RemovalOperation:
		return RemovalItem
	// - ADDITION
	case shared.AdditionOperation:
		return AdditionItem
	// - SAME
	case shared.NoChangeOperation:
		return SameItem
	}

	return UnknownUpdateItem // Fallback on unknown
}

func _debugPackageList(
	sectionType MarkdownSection,
	categoryType MarkdownCategory,
	subCategoryType MarkdownSubCategory,
	itemType MarkdownItem,
	change *shared.PackageChange,
) {
	expectedTypeKey := strings.ToLower(string(sectionType)) +
		"-" + strings.ToLower(string(categoryType)) +
		"-" + strings.ToLower(string(subCategoryType))
	if !change.Package.IsDevOnly() && change.Package.IsRootDevRequirement() {
		expectedTypeKey += "+dev_req"
	}

	expectedTypeKey += "/" + string(itemType)
	if itemType == UnknownUpdateItem && change.Operation.SemverType == shared.SemverExtraUpdate {
		expectedTypeKey += "+SEMVER_EXTRA"
	} else if change.Package.IsAbandoned() {
		expectedTypeKey += "+ABANDONED"
	}

	if change.Package.GetName() != expectedTypeKey {
		fmt.Fprintln(os.Stderr, "package mismatch: got", change.Package.GetName(), ", want", expectedTypeKey)
	}
}
