package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
)

const PART = "1"
const WILL_SUBMIT = false

func solveP1(input string) string {
	locks, keys := parseInput(input)
	maxHeight := 5

	count := 0
	for _, lock := range locks {
	lockKeyLoop:
		for _, key := range keys {
			for i, keyHeight := range key {
				lockHeight := lock[i]
				if lockHeight+keyHeight > maxHeight {
					continue lockKeyLoop
				}
			}

			count++
		}
	}

	return strconv.Itoa(count)
}

func solveP2(input string) string {
	parseInput(input)

	return ""
}

func parseInput(input string) (locks [][]int, keys [][]int) {
	for _, schematic := range strings.Split(strings.TrimSpace(input), "\n\n") {
		schematicLines := strings.Split(schematic, "\n")
		heights := make([]int, 5)

		for i := 1; i < len(schematicLines)-1; i++ {
			for j, rune := range schematicLines[i] {
				if rune == '#' {
					heights[j]++
				}
			}
		}

		if schematicLines[0] == "#####" {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	return locks, keys
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
