package summary

import "github.com/yoanm/go-deps-diff/summary/markdown"

type simpleSymbolHelp struct {
	symbol  string
	message string
}
type simpleSymbolHelpList []simpleSymbolHelp

const captionHeaderLevel = 4

func processSummaryCaption(builder *markdown.Builder) {
	builder.WriteEol()
	builder.WriteEol()
	// Always closed by default (goal is to avoid polluting actual changes) !
	builder.Details("Captions", createCaptionContent, false, 0)
}

var (
	//nolint:gochecknoglobals // Just to avoid multiple creation at runtime
	versionHelpList = simpleSymbolHelpList{
		{
			symbol:  "<td align=\"center\">_VERSION_❗</td>",
			message: "<td align=\"left\">Version is not semver compliant.<br/>Usually a commit ref or branch.</td>",
		},
	}
	//nolint:gochecknoglobals // Just to avoid multiple creation at runtime
	operationHelpList = simpleSymbolHelpList{
		{symbol: "<td align=\"center\">❓</td>", message: "<td align=\"left\">Unknown update</td>"},
		{symbol: "<td align=\"center\">❌</td>", message: "<td align=\"left\">Removed package</td>"},
		{symbol: "<td align=\"center\">➕️</td>", message: "<td align=\"left\">Added package</td>"},
		{symbol: "<td align=\"center\">🟰</td>", message: "<td align=\"left\">No change</td>"},
		{symbol: "<td align=\"center\"><sub><sup>🔺.🔹.🔹</sup></sub></td>", message: "<td align=\"left\">Major upgrade</td>"},
		{symbol: "<td align=\"center\"><sub><sup>🔻.🔹.🔹</sup></sub></td>", message: "<td align=\"left\">Major downgrade</td>"},
		{symbol: "<td align=\"center\"><sub><sup>🔹.🔺.🔹</sup></sub></td>", message: "<td align=\"left\">Minor upgrade</td>"},
		{symbol: "<td align=\"center\"><sub><sup>🔹.🔻.🔹</sup></sub></td>", message: "<td align=\"left\">Minor downgrade</td>"},
		{symbol: "<td align=\"center\"><sub><sup>🔹.🔹.🔺</sup></sub></td>", message: "<td align=\"left\">Patch upgrade</td>"},
		{symbol: "<td align=\"center\"><sub><sup>🔹.🔹.🔻</sup></sub></td>", message: "<td align=\"left\">Patch downgrade</td>"},
		{
			symbol:  "<td align=\"center\"><sub><sup>🔹.🔹.🔹❓</sup></sub></td>",
			message: "<td align=\"left\">Extra updated, considered as Unknown update</td>",
		},
		{symbol: "<td align=\"center\">❔</td>", message: "<td align=\"left\">Unmanaged operation</td>"},
	}
	//nolint:gochecknoglobals // Just to avoid multiple creation at runtime
	pkgTypeHelpList = simpleSymbolHelpList{
		{
			symbol:  "<td align=\"center\">🗄</td>",
			message: "<td align=\"left\">Package is explicitly required for production usage</td>",
		},
		{
			symbol:  "<td align=\"center\">🧰</td>",
			message: "<td align=\"left\">Package is explicitly required for dev-only usage</td>",
		},
		{symbol: "<td align=\"center\">🔗️</td>", message: "<td align=\"left\">Transitive dependency package</td>"},
		{
			symbol:  "<td align=\"center\">💀</td>",
			message: "<td align=\"left\">Package is declared abandoned.<br/>You should replace it.</td>",
		},
	}
	//nolint:gochecknoglobals // Just to avoid multiple creation at runtime
	prodDevOnlyHelpList = simpleSymbolHelpList{
		{symbol: "<td align=\"center\">🏭</td>", message: "<td align=\"left\">Package is available in <b>production</b></td>"},
		{
			symbol:  "<td align=\"center\">🧪</td>",
			message: "<td align=\"left\">Package is only available for <b>dev-only</b>, it won't exist in production</td>",
		},
	}
)

func createCaptionContent(builder *markdown.Builder, indentDepth int) {
	createCaptionSection(builder, "Versions", versionHelpList, indentDepth)
	createCaptionSection(builder, "Operations", operationHelpList, indentDepth)
	createCaptionSection(builder, "Package types", pkgTypeHelpList, indentDepth)

	// Prod vs dev-only
	createCaptionSection(builder, "Production vs Dev-only usage", prodDevOnlyHelpList, indentDepth)
	builder.WriteLine("There is a difference between **Usage** and **Requirement**.", indentDepth)
	builder.WriteEol()
	builder.WriteLine("- A **Requirement** can be defined as dev-only or not.", indentDepth)
	builder.WriteLine("  ", indentDepth)
	builder.WriteLine("  Depending on the manager, there might be dedicated property for each environment in "+
		"the requirement file.", indentDepth)
	builder.WriteLine(
		"- **Usage** however is whether the package is available on production or only for dev-only.",
		indentDepth,
	)
	builder.WriteLine("  ", indentDepth)
	builder.WriteLine("  Usually, it's the package lock which provides this information.", indentDepth)
	builder.WriteEol()
	builder.WriteLine("You may require a package for dev-only, but this package may actually be a dependency of "+
		"one of your requirement for production. In that case, the package you required for dev-only will be "+
		"displayed in \"Production usage\" section, because it is actually available in production.", indentDepth)
}

func createCaptionSection(
	builder *markdown.Builder,
	section string,
	helpList simpleSymbolHelpList,
	indentDepth int,
) {
	builder.Header(section, captionHeaderLevel, indentDepth)
	builder.HTMLTable(
		func(yield func([]string) bool) {
			for _, helpData := range helpList {
				if !yield([]string{helpData.symbol, helpData.message}) {
					return
				}
			}
		},
		indentDepth,
	)
}
