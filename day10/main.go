package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

type point = utils.Point

func solveP1(input string) string {
	topMap, trailheads, width, height := parseInput(input)

	score := 0
	for _, trailhead := range trailheads {
		goals := make(map[point]bool)
		findGoals(topMap, trailhead, width, height, goals)
		score += len(goals)
	}

	return strconv.Itoa(score)
}

func solveP2(input string) string {
	topMap, trailheads, width, height := parseInput(input)

	score := 0
	for _, trailhead := range trailheads {
		score += findDistinctTrails(topMap, trailhead, width, height)
	}

	return strconv.Itoa(score)
}

var DIRECTIONS = []point{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

func findGoals(topMap map[point]int, pos point, width int, height int, goals map[point]bool) {
	if topMap[pos] == 9 {
		goals[pos] = true
	}

	for _, direction := range DIRECTIONS {
		nextPos := utils.AddPoint(pos, direction)
		if utils.InBounds(nextPos, width, height) && topMap[nextPos]-topMap[pos] == 1 {
			findGoals(topMap, nextPos, width, height, goals)
		}
	}
}

func findDistinctTrails(topMap map[point]int, pos point, width int, height int) int {
	if topMap[pos] == 9 {
		return 1
	}

	count := 0
	for _, direction := range DIRECTIONS {
		nextPos := utils.AddPoint(pos, direction)
		if utils.InBounds(nextPos, width, height) && topMap[nextPos]-topMap[pos] == 1 {
			count += findDistinctTrails(topMap, nextPos, width, height)
		}
	}

	return count
}

func parseInput(input string) (topMap map[point]int, trailheads []point, width int, height int) {
	topMap = make(map[point]int)

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, rune := range []rune(line) {
			posPoint := point{X: x, Y: y}
			pos := int(rune - '0')
			topMap[posPoint] = pos

			if pos == 0 {
				trailheads = append(trailheads, posPoint)
			}
		}

		width = len(line)
		height++
	}

	return topMap, trailheads, width, height
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
