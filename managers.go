package depsdiff

import (
	"fmt"

	"github.com/yoanm/go-deps-diff/managers/composer"
	"github.com/yoanm/go-deps-diff/shared"
)

func ComposerDiff(previous, current *PkgManagerInput) (shared.DiffMap, error) {
	previousMap, err := composer.BuildMapFromBytes(previous.Requirement, previous.Lock)
	if err != nil {
		return nil, fmt.Errorf("building previous package map: %w", err)
	}

	currentMap, err := composer.BuildMapFromBytes(current.Requirement, current.Lock)
	if err != nil {
		return nil, fmt.Errorf("building current package map: %w", err)
	}

	return Diff(previousMap, currentMap)
}
