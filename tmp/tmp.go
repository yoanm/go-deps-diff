package main

import (
	"fmt"
	"github.com/yoanm/go-deps-diff/summary"
	"os"

	depsdiff "github.com/yoanm/go-deps-diff"
)

func main() {
	previousReq, err := os.ReadFile("./tmp/testdata/previous-composer.json")
	if err != nil {
		panic(err)
	}

	currentReq, err := os.ReadFile("./tmp/testdata/current-composer.json")
	if err != nil {
		panic(err)
	}

	previousLock, err := os.ReadFile("./tmp/testdata/previous-composer.lock")
	if err != nil {
		panic(err)
	}

	currentLock, err := os.ReadFile("./tmp/testdata/current-composer.lock")
	if err != nil {
		panic(err)
	}

	out, err := depsdiff.ComposerDiff(
		&depsdiff.PkgManagerInput{Lock: previousLock, Requirement: previousReq},
		&depsdiff.PkgManagerInput{Lock: currentLock, Requirement: currentReq},
	)
	if err != nil {
		panic(err)
	}

	fmt.Print(summary.GenerateForChanges(out))
}
