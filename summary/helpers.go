package summary

import (
	"fmt"
	"slices"
	"strings"

	"github.com/yoanm/go-deps-diff/shared"
)

func splitItemList(subCategoriesMap SubCategoriesMap) (PkgList, PkgList) {
	var (
		noChangePkgList    PkgList
		otherChangePkgList PkgList
	)

	inOrderMapIteratorHelper[MarkdownSubCategory, ItemsMap](
		subCategoriesMap,
		subcategoryOrders,
		func(subCategoryName MarkdownSubCategory, itemsMap ItemsMap) {
			inOrderMapIteratorHelper[MarkdownItem, PkgList](
				itemsMap,
				itemOrders,
				func(itemName MarkdownItem, pkgList PkgList) {

					for _, item := range pkgList {
						// Keep track of no change items to display them at the end of the section.
						// - No change items are far less meaningful than the other items
						// Idea behind the comment is to quickly spot a change which may lead to an issue or caused an
						// issue, so more visibility to actual changes.
						if shared.NoChangeOperation == item.Operation.Name {
							noChangePkgList = append(noChangePkgList, item)
						} else {
							otherChangePkgList = append(otherChangePkgList, item)
						}
					}
				},
			)
		},
	)

	return noChangePkgList, otherChangePkgList
}

func guessBestTableMode(abandonedPkgList PkgList) pkgRowMode {
	needsFullTableWidth := slices.ContainsFunc(abandonedPkgList, func(item *shared.PackageChange) bool {
		// The following operations only need two cells (version + operation), any other would need one more
		// to display both the previous and current versions
		return item.Operation.Name != shared.AdditionOperation &&
			item.Operation.Name != shared.RemovalOperation &&
			item.Operation.Name != shared.NoChangeOperation
	})

	if needsFullTableWidth {
		return fullPkgRowMode
	}

	return withOperationPkgRowMode
}

type multiCounter struct {
	title string
	count int
}
type sectionSummaryCntMap map[MarkdownItem]*multiCounter

func buildSectionSummaryMrk(subCategoriesMap SubCategoriesMap) string {
	cntMap := buildSectionCounters(subCategoriesMap)

	partList := make([]string, len(cntMap))
	partKey := 0

	inOrderMapIteratorHelper[MarkdownItem, *multiCounter](
		cntMap,
		itemOrders,
		func(key MarkdownItem, data *multiCounter) {
			partList[partKey] = fmt.Sprintf("%s<sup>%d</sup>", data.title, data.count)
			partKey++
		},
	)

	return strings.Join(partList, "    ")
}

func buildSectionCounters(subCategoriesMap SubCategoriesMap) sectionSummaryCntMap {
	cntMap := make(sectionSummaryCntMap)

	for _, itemsMap := range subCategoriesMap {
		for itemType, pkgList := range itemsMap {
			if nil == cntMap[itemType] {
				cntMap[itemType] = &multiCounter{}
			}

			cntMap[itemType].count += len(pkgList)
			// For UnknownUpdateItem, try to catch basic unknown update rather than a SEMVER_EXTRA update.
			if itemType == UnknownUpdateItem {
				for _, pkg := range pkgList {
					if pkg.Operation.SemverType != shared.SemverExtraUpdate {
						cntMap[itemType].title = getOperationSymbol(pkg.Operation)

						break
					}
				}
			}
			// Fallback on first available one if sample is not defined yet
			if len(cntMap[itemType].title) == 0 && len(pkgList) > 0 {
				cntMap[itemType].title = getOperationSymbol(pkgList[0].Operation)
			}
		}
	}

	return cntMap
}
