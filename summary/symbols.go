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

func getOperationSymbol(op shared.Operation) string {
	switch op.Name {
	case shared.UnknownUpdateOperation:
		if op.SemverType == shared.SemverExtraUpdate {
			return "<sub><sup>🔹.🔹.🔹❓</sup></sub>"
		}

		return "❓"
	case shared.UpgradeOperation:
		return getUpdateOperationSymbol(op, false)
	case shared.DowngradeOperation:
		return getUpdateOperationSymbol(op, true)
	case shared.RemovalOperation:
		return "❌"
	case shared.AdditionOperation:
		return "➕️"
	case shared.NoChangeOperation:
		return "🟰"
	}

	return "❔"
}

func getUpdateOperationSymbol(op shared.Operation, isDowngrade bool) string {
	emojiUpdated := "🔺"
	if isDowngrade {
		emojiUpdated = "🔻"
	}

	switch op.SemverType {
	case shared.SemverMajorUpdate:
		return "<sub><sup>" + emojiUpdated + ".🔹.🔹</sup></sub>"
	case shared.SemverMinorUpdate:
		return "<sub><sup>🔹." + emojiUpdated + ".🔹</sup></sub>"
	case shared.SemverPatchUpdate:
		return "<sub><sup>🔹.🔹." + emojiUpdated + "</sup></sub>"
	}

	return "❔"
}
