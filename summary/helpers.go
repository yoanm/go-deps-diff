package summary

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/yoanm/go-deps-diff/shared"
)

func splitItemList(subCategoriesMap subCategoriesMap) (pkgList, pkgList) {
	var (
		noChangePkgList    pkgList
		otherChangePkgList pkgList
	)

	inOrderMapIteratorHelper[markdownSubCategory, itemsMap](
		subCategoriesMap,
		getSubCategoriesOrder(),
		func(subCategoryName markdownSubCategory, itemsMap itemsMap) {
			inOrderMapIteratorHelper[markdownItem, pkgList](
				itemsMap,
				getItemsOrder(),
				func(itemName markdownItem, pkgList pkgList) {
					slices.SortFunc(pkgList, func(changeA, changeB *shared.PackageChange) int {
						return strings.Compare(changeA.Package.GetName(), changeB.Package.GetName())
					})

					for _, item := range pkgList {
						slog.Debug("Processing package: " + item.Package.GetName())
						// Keep track of no change items to display them at the end of the section.
						// - No change items are far less meaningful than the other items
						// Idea behind the comment is to quickly spot a change which may lead to an issue or caused an
						// issue, so more visibility to actual changes.
						if shared.NoChangeOperation == item.Operation.Name && item.Package.GetVersion().Semver != nil {
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

func guessShortestPkgRowMode(abandonedPkgList pkgList) pkgRowMode {
	needsFullTableWidth := slices.ContainsFunc(
		abandonedPkgList,
		func(item *shared.PackageChange) bool {
			// The following operations only need two cells (version + operation), any other would need one more
			// to display both the previous and current versions
			return item.Operation.Name != shared.AdditionOperation &&
				item.Operation.Name != shared.RemovalOperation &&
				item.Operation.Name != shared.NoChangeOperation
		},
	)

	if needsFullTableWidth {
		return fullPkgRowMode
	}

	return withOperationPkgRowMode
}

type itemsCounter struct {
	title string
	count int
}
type sectionSummaryCntMap map[markdownItem]*itemsCounter

func buildSectionSummaryMrk(subCategoriesMap subCategoriesMap) string {
	cntMap := buildItemsCounters(subCategoriesMap)

	partList := make([]string, len(cntMap))
	partKey := 0

	inOrderMapIteratorHelper[markdownItem, *itemsCounter](
		cntMap,
		getItemsOrder(),
		func(key markdownItem, data *itemsCounter) {
			partList[partKey] = fmt.Sprintf("%s<sup>%d</sup>", data.title, data.count)
			partKey++
		},
	)

	return strings.Join(partList, "    ")
}

func buildItemsCounters(subCategoriesMap subCategoriesMap) sectionSummaryCntMap {
	cntMap := make(sectionSummaryCntMap)

	for _, itemsMap := range subCategoriesMap {
		for itemType, pkgList := range itemsMap {
			if nil == cntMap[itemType] {
				cntMap[itemType] = &itemsCounter{title: "", count: 0}
			}

			cntMap[itemType].count += len(pkgList)
			// For unknownUpdateItem, try to catch basic unknown update rather than a SEMVER_EXTRA update.
			if itemType == unknownUpdateItem {
				if pkg := findSemverExtraUpdateChange(pkgList); pkg != nil {
					cntMap[itemType].title = getOperationSymbol(pkg.Operation)

					continue
				}
			}
			// Fallback on first available one if sample is not defined yet
			if len(pkgList) > 0 && len(cntMap[itemType].title) == 0 {
				cntMap[itemType].title = getOperationSymbol(pkgList[0].Operation)
			}
		}
	}

	return cntMap
}

func findSemverExtraUpdateChange(pkgList pkgList) *shared.PackageChange {
	for _, pkg := range pkgList {
		if pkg.Operation.SemverType != shared.SemverExtraUpdate {
			return pkg
		}
	}

	return nil
}
