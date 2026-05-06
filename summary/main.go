package summary

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/summary/markdown"
)

const (
	sectionHeaderLevel  = 2
	categoryHeaderLevel = 3
)

func GenerateForChanges(changes shared.DiffMap) string {
	mrkMap := buildSectionsMap(changes)
	builder := markdown.NewBuilder()

	inOrderMapIteratorHelper[markdownSection, categoriesMap](
		mrkMap,
		getSectionsOrder(),
		func(sectionName markdownSection, categoriesMap categoriesMap) {
			processSection(builder, categoriesMap, sectionName)
		},
	)

	processSummaryCaption(builder)

	return builder.String()
}

func processSection(builder *markdown.Builder, categoriesMap categoriesMap, sectionName markdownSection) {
	slog.Debug("Processing section: " + string(sectionName))
	builder.Header(
		getSectionHeaderFor(sectionName)+"<br/><sub><sup>"+getSectionDescriptionFor(sectionName)+"</sub></sup>",
		sectionHeaderLevel,
		0,
	)

	inOrderMapIteratorHelper[markdownCategory, subCategoriesMap](
		categoriesMap,
		getCategoriesOrder(),
		func(categoryName markdownCategory, subCategoriesMap subCategoriesMap) {
			openedDetails := categoryName == prodUsageCategory &&
				(cautionSection == sectionName || warningSection == sectionName || importantSection == sectionName)

			processCategory(builder, subCategoriesMap, categoryName, openedDetails)
		},
	)

	builder.WriteLine("<hr/>", 0)
}

func processCategory(
	builder *markdown.Builder,
	subCategoriesMap subCategoriesMap,
	categoryName markdownCategory,
	openedDetails bool,
) {
	slog.Debug("Processing category: " + string(categoryName))

	noChangePkgList, otherChangePkgList := splitItemList(subCategoriesMap)

	builder.Header(getCategoryHeaderFor(categoryName), categoryHeaderLevel, 0)
	builder.Details(
		buildSectionSummaryMrk(subCategoriesMap),
		func(builder *markdown.Builder, indentDepth int) {
			if len(otherChangePkgList) > 0 {
				processPkgList(builder, otherChangePkgList, guessShortestPkgRowMode(otherChangePkgList), indentDepth)
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

func processPkgList(builder *markdown.Builder, pkgList pkgList, tableMode pkgRowMode, indentDepth int) {
	builder.HTMLTable(
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
		buildPackageNameHTMLCell(item.Package),
	}

	pkgVersionCell := buildPackageVersionHTMLCell(item.Package.GetVersion())

	switch tableMode {
	case versionOnlyPkgRowMode:
		cellList = append(cellList, pkgVersionCell)

	case withOperationPkgRowMode:
		operationCell := buildOperationHTMLCell(item.Operation, 0)
		if item.Operation.Name != shared.AdditionOperation {
			cellList = append(cellList, pkgVersionCell, operationCell)
		} else {
			cellList = append(cellList, operationCell, pkgVersionCell)
		}

	case fullPkgRowMode:
		cellList = buildItemMrkFullPkgRowCells(item, cellList, pkgVersionCell)

	default:
		panic("Unmanaged table mode:" + strconv.Itoa(int(tableMode)))
	}

	return cellList
}

func buildItemMrkFullPkgRowCells(item *shared.PackageChange, cellList []string, pkgVersionCell string) []string {
	if item.Operation.Name != shared.AdditionOperation { // Version will be printed at the end for added package !
		switch item.Operation.Name {
		case shared.UnknownUpdateOperation, shared.UpgradeOperation, shared.DowngradeOperation:
			cellList = append(cellList, buildPackageVersionHTMLCell(item.PreviousVersion))
		default:
			cellList = append(cellList, pkgVersionCell)
		}
	}

	colspan := 0
	if shared.NoChangeOperation == item.Operation.Name ||
		shared.AdditionOperation == item.Operation.Name ||
		shared.RemovalOperation == item.Operation.Name {
		colspan = 2
	}

	cellList = append(cellList, buildOperationHTMLCell(item.Operation, colspan))

	switch item.Operation.Name { //nolint:exhaustive // Only those cases should be handled here !
	case shared.AdditionOperation, shared.UnknownUpdateOperation, shared.UpgradeOperation, shared.DowngradeOperation:
		cellList = append(cellList, pkgVersionCell)
	}

	return cellList
}

func buildOperationHTMLCell(operation shared.Operation, colspan int) string {
	opColspanDirective := ""
	if colspan > 1 {
		opColspanDirective = fmt.Sprintf(" colspan=\"%d\"", colspan)
	}

	return "<td align=\"center\"" + opColspanDirective + ">" + getOperationSymbol(operation) + "</td>"
}

func buildPackageVersionHTMLCell(version shared.PkgVersion) string {
	return "<td align=\"right\">" + version.Label + "</td>"
}

func buildPackageNameHTMLCell(pkg shared.PkgWrapper) string {
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
