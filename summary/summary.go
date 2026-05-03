package summary

import (
	"fmt"
	"github.com/yoanm/go-deps-diff/markdown"
	"github.com/yoanm/go-deps-diff/shared"
	"os"
	"strconv"
)

func Generate(mrkList SectionsMap) string {
	if len(mrkList) == 0 {
		return ""
	}

	builder := markdown.NewBuilder()

	inOrderMapIteratorHelper[MarkdownSection, CategoriesMap](
		mrkList,
		sectionOrders,
		func(sectionName MarkdownSection, categoriesMap CategoriesMap) {
			processSection(builder, categoriesMap, sectionName)
		},
	)

	processSummaryCaption(builder)

	return builder.String()
}

func processSection(builder *markdown.Builder, categoriesMap CategoriesMap, sectionName MarkdownSection) {
	if len(categoriesMap) == 0 {
		return
	}

	fmt.Fprintln(os.Stderr, "Processing section:", sectionName)
	builder.Header(sectionTitleMap[sectionName]+"<br/><sub><sup>"+sectionHeaderMap[sectionName]+"</sub></sup>", 2, 0)

	inOrderMapIteratorHelper[MarkdownCategory, SubCategoriesMap](
		categoriesMap,
		categoryOrders,
		func(categoryName MarkdownCategory, subCategoriesMap SubCategoriesMap) {
			openedDetails := categoryName == ProdUsageCategory &&
				(CautionSection == sectionName || WarningSection == sectionName || ImportantSection == sectionName)

			processCategory(builder, subCategoriesMap, categoryName, openedDetails)
		},
	)

	builder.WriteLine("<hr/>", 0)
}

func processCategory(
	builder *markdown.Builder,
	subCategoriesMap SubCategoriesMap,
	categoryName MarkdownCategory,
	openedDetails bool,
) {
	if len(subCategoriesMap) == 0 {
		return
	}

	fmt.Fprintln(os.Stderr, "Processing category:", categoryName)

	noChangePkgList, otherChangePkgList := splitItemList(subCategoriesMap)

	builder.Header(categoryTitleMap[categoryName], 3, 0)
	builder.Details(
		buildSectionSummaryMrk(subCategoriesMap),
		func(builder *markdown.Builder, indentDepth int) {
			if len(otherChangePkgList) > 0 {
				processPkgList(builder, otherChangePkgList, guessBestTableMode(otherChangePkgList), indentDepth)
			}

			if len(noChangePkgList) > 0 {
				builder.Details(
					"Unchanged packages 🟰<sup>"+strconv.Itoa(len(noChangePkgList))+"</sup>",
					func(builder *markdown.Builder, indentDepth int) {
						processPkgList(builder, noChangePkgList, versionOnlyPkgRowMode, indentDepth)
					},
					false, // Always closed by default (goal is to avoid polluting other changes)
					indentDepth,
				)
			}
		},
		openedDetails,
		0,
	)
}

func processPkgList(builder *markdown.Builder, pkgList PkgList, tableMode pkgRowMode, indentDepth int) {
	if len(pkgList) == 0 {
		return
	}

	builder.HtmlTable(
		func(yield func([]string) bool) {
			for _, item := range pkgList {
				if !yield(buildItemMrkRowCells(item, tableMode)) {
					return
				}
			}
		},
		indentDepth,
	)
}

func buildItemMrkRowCells(item *shared.PackageChange, tableMode pkgRowMode) []string {
	cellList := []string{
		buildPackageNameHtmlCell(item.Package),
	}

	pkgVersionCell := buildPackageVersionHtmlCell(item.Package.GetVersion())

	switch tableMode {
	case versionOnlyPkgRowMode:
		cellList = append(cellList, pkgVersionCell)

	case withOperationPkgRowMode:
		operationCell := buildOperationHtmlCell(item.Operation, 0)
		if item.Operation.Name != shared.AdditionOperation {
			cellList = append(cellList, pkgVersionCell, operationCell)
		} else {
			cellList = append(cellList, operationCell, pkgVersionCell)
		}

	case fullPkgRowMode:
		cellList = buildItemMrkFullPrkRowCells(item, cellList, pkgVersionCell)

	default:
		panic("Unmanaged table mode:" + strconv.Itoa(int(tableMode)))
	}

	return cellList
}

func buildItemMrkFullPrkRowCells(item *shared.PackageChange, cellList []string, pkgVersionCell string) []string {
	if item.Operation.Name != shared.AdditionOperation { // Version will be printed at the end for added package !
		cellList = append(cellList, pkgVersionCell)
	}

	colspan := 0
	if shared.NoChangeOperation == item.Operation.Name ||
		shared.AdditionOperation == item.Operation.Name ||
		shared.RemovalOperation == item.Operation.Name {
		colspan = 2
	}

	cellList = append(cellList, buildOperationHtmlCell(item.Operation, colspan))

	switch item.Operation.Name {
	case shared.AdditionOperation:
		cellList = append(cellList, pkgVersionCell)
	case shared.UnknownUpdateOperation, shared.UpgradeOperation, shared.DowngradeOperation:
		cellList = append(cellList, buildPackageVersionHtmlCell(item.PreviousVersion))
	}

	return cellList
}

func buildOperationHtmlCell(op shared.Operation, colspan int) string {
	opColspanDirective := ""
	if colspan > 1 {
		opColspanDirective = fmt.Sprintf(" colspan=\"%d\"", colspan)
	}

	return "<td align=\"center\"" + opColspanDirective + ">" + getOperationSymbol(op) + "</td>"
}

func buildPackageVersionHtmlCell(version shared.PkgVersion) string {
	return "<td align=\"right\">" + version.Label + "</td>"
}

func buildPackageNameHtmlCell(pkg shared.PkgWrapper) string {
	pkgTitle := pkg.GetName()
	if pkg.GetLink() != "" {
		pkgTitle = "<a href=\"" + pkg.GetLink() + "\">" + pkgTitle + "</a>"
	}
	// Prepend package type symbol
	pkgTitle = "<sup>" + getPackageSymbol(pkg) + "</sup>" + pkgTitle
	// Append abandoned symbol
	if pkg.IsAbandoned() {
		pkgTitle += "💀"
	}

	return "<td align=\"left\">" + pkgTitle + "</td>"
}
