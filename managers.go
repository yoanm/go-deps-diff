package depsdiff

import (
	"fmt"

	"github.com/yoanm/go-deps-diff/contract"
	"github.com/yoanm/go-deps-diff/managers/composer"
)

func ComposerDiff(previous, current *PkgManagerInput) (contract.DiffMap, error) {
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
