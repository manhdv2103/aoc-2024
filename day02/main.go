package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	reports := parseInput(input)

	count := 0
	for _, report := range reports {
		if isReportSafe(report) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func solveP2(input string) string {
	reports := parseInput(input)

	count := 0
	for _, report := range reports {
		if isReportSafe(report) {
			count++
			continue
		}

		for i := range len(report) {
			if isReportSafe(utils.RemoveIndex(report, i)) {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count)
}

func isReportSafe(report []int) bool {
	direction := 1
	if report[0] > report[1] {
		direction = -1
	}

	for i := range len(report) - 1 {
		delta := (report[i+1] * direction) - (report[i] * direction)
		if delta < 1 || delta > 3 {
			return false
		}
	}

	return true
}

func parseInput(input string) [][]int {
	reports := make([][]int, 0)
	for _, reportStr := range strings.Split(strings.TrimSpace(input), "\n") {
		reports = append(reports, utils.ToInts(strings.Fields(reportStr)))
	}

	return reports
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
