package summary

import "github.com/yoanm/go-deps-diff/shared"

func getPackageSymbol(pkg shared.PkgWrapper) string {
	switch {
	case pkg.IsRootRequirement():
		return "🗄️"
	case pkg.IsRootDevRequirement():
		return "🧰"
	default:
		return "🔗"
	}
}

func getOperationSymbol(operation *shared.Operation) string {
	switch operation.Name {
	case shared.UnknownUpdateOperation:
		if operation.SemverType == shared.SemverExtraUpdate {
			return "<sub><sup>🔹.🔹.🔹❓</sup></sub>"
		}

		return "❓"
	case shared.UpgradeOperation:
		return getUpdateOperationSymbol(operation, false)
	case shared.DowngradeOperation:
		return getUpdateOperationSymbol(operation, true)
	case shared.RemovalOperation:
		return "❌"
	case shared.AdditionOperation:
		return "➕️"
	case shared.NoChangeOperation:
		return "🟰"
	}

	return "❔"
}

func getUpdateOperationSymbol(operation *shared.Operation, isDowngrade bool) string {
	emojiUpdated := "🔺"
	if isDowngrade {
		emojiUpdated = "🔻"
	}

	switch operation.SemverType { //nolint:exhaustive // Only those cases can be managed, fallback to unknown otherwise
	case shared.SemverMajorUpdate:
		return "<sub><sup>" + emojiUpdated + ".🔹.🔹</sup></sub>"
	case shared.SemverMinorUpdate:
		return "<sub><sup>🔹." + emojiUpdated + ".🔹</sup></sub>"
	case shared.SemverPatchUpdate:
		return "<sub><sup>🔹.🔹." + emojiUpdated + "</sup></sub>"
	}

	return "❔"
}
