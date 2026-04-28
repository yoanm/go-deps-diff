package depsdiff

import (
	"fmt"

	"github.com/yoanm/go-deps-diff/composer"
)

func ComposerDiff(input *Input) (*Output, error) {
	previousMap, err := composer.BuildMapFromBytes(input.Previous.Requirement, input.Previous.Lock)
	if err != nil {
		return nil, fmt.Errorf("building previous package map: %w", err)
	}

	currentMap, err := composer.BuildMapFromBytes(input.Current.Requirement, input.Current.Lock)
	if err != nil {
		return nil, fmt.Errorf("building current package map: %w", err)
	}

	return Diff(previousMap, currentMap)
}
