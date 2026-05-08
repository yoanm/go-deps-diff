package summary

import "github.com/yoanm/go-deps-diff/contract"

func getPackageSymbol(pkg contract.PkgWrapper) string {
	switch {
	case pkg.IsRootRequirement():
		return "🗄️"
	case pkg.IsRootDevRequirement():
		return "🧰"
	default:
		return "🔗"
	}
}

func getOperationSymbol(operation contract.Operation) string {
	switch operation.Name {
	case contract.UnknownUpdateOperation:
		if operation.SemverType == contract.SemverExtraUpdate {
			return "<sub><sup>🔹.🔹.🔹❓</sup></sub>"
		}

		return "❓"
	case contract.UpgradeOperation:
		return getUpdateOperationSymbol(operation, false)
	case contract.DowngradeOperation:
		return getUpdateOperationSymbol(operation, true)
	case contract.RemovalOperation:
		return "❌"
	case contract.AdditionOperation:
		return "➕️"
	case contract.NoChangeOperation:
		return "🟰"
	}

	return "❔"
}

func getUpdateOperationSymbol(operation contract.Operation, isDowngrade bool) string {
	emojiUpdated := "🔺"
	if isDowngrade {
		emojiUpdated = "🔻"
	}

	switch operation.SemverType { //nolint:exhaustive // Only those cases can be managed, fallback to unknown otherwise
	case contract.SemverMajorUpdate:
		return "<sub><sup>" + emojiUpdated + ".🔹.🔹</sup></sub>"
	case contract.SemverMinorUpdate:
		return "<sub><sup>🔹." + emojiUpdated + ".🔹</sup></sub>"
	case contract.SemverPatchUpdate:
		return "<sub><sup>🔹.🔹." + emojiUpdated + "</sup></sub>"
	}

	return "❔"
}
