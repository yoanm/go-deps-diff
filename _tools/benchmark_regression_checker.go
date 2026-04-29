package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Acceptable regression (10% for instance)
	var threshold float64
	var err error

	if len(os.Args) != 1 {
		fmt.Println("Missing threshold argument. Usage: benchmark_regression_checker [threshold_percentage]")
		os.Exit(1)
	} else {
		if threshold, err = strconv.ParseFloat(os.Args[1], 64); err != nil {
			fmt.Printf("Threshold must be a valid float\n")
			os.Exit(1)
		}
		if threshold > 100 || threshold <= 0 {
			fmt.Printf("Threshold must be between 1%% and 99%%\n")
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	// Match lines like: BenchmarkAbc-42  230ns  123ns  +90.00%
	deltaRegex := regexp.MustCompile(`([+-]\d+\.?\d*)%`)

	var regList []string // Regressions list
	for scanner.Scan() {
		line := scanner.Text()
		matches := deltaRegex.FindStringSubmatch(line)

		if len(matches) > 1 {
			delta, err2 := strconv.ParseFloat(matches[1], 64)
			if err2 != nil {
				fmt.Printf("Error parsing delta from line: %s\n", line)
				continue
			}

			// Positive delta means regression (slower)
			if delta > threshold {
				regList = append(
					regList,
					fmt.Sprintf("  %s (%.2f%% slower)", strings.Fields(line)[0], delta),
				)
			}
		}
	}

	if len(regList) > 0 {
		fmt.Printf("Performance regression detected (threshold: %.1f%%):\n", threshold)
		for _, reg := range regList {
			fmt.Println(reg)
		}
		os.Exit(1)
	}

	fmt.Println("All good 🎉.")
}
