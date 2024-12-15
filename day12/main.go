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
	garden, width, height := parseInput(input)

	total := 0
	visited := make(utils.Set[point])
	for pos := range garden {
		if visited[pos] {
			continue
		}

		visited[pos] = true
		total += getRegionPrice(pos, garden, width, height, visited)
	}

	return strconv.Itoa(total)
}

func solveP2(input string) string {
	garden, width, height := parseInput(input)

	total := 0
	visited := make(utils.Set[point])
	for pos := range garden {
		if visited[pos] {
			continue
		}

		visited[pos] = true
		total += getRegionBulkPrice(pos, garden, width, height, visited)
	}

	return strconv.Itoa(total)
}

func getRegionPrice(
	pos point,
	garden map[point]rune,
	width int,
	height int,
	visited utils.Set[point],
) (price int) {
	area := 0
	perimeter := 0

	queue := []point{pos}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		plot := garden[pos]

		area++

		for _, direction := range utils.Directions {
			nextPos := utils.AddPoint(pos, direction)
			if !utils.InBounds(nextPos, width, height) || garden[nextPos] != plot {
				perimeter++
				continue
			}

			if !visited[nextPos] {
				queue = append(queue, nextPos)
				visited[nextPos] = true
			}
		}
	}

	return area * perimeter
}

func getRegionBulkPrice(
	pos point,
	garden map[point]rune,
	width int,
	height int,
	visited utils.Set[point],
) (price int) {
	area := 0
	corners := 0
	isOutside := func(p point) bool {
		return !utils.InBounds(p, width, height) || garden[pos] != garden[p]
	}

	queue := []point{pos}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		area++

		for _, direction := range utils.Directions {
			nextPos := utils.AddPoint(pos, direction)
			isNextPosOutside := isOutside(nextPos)

			adjacentDirection := utils.RotatePoint(direction)
			adjacentSidePos := utils.AddPoint(pos, adjacentDirection)
			cornerPos := utils.AddPoint(pos, utils.AddPoint(direction, adjacentDirection))

			isSidePosOutside := isNextPosOutside
			isAdjacentSidePosOutside := isOutside(adjacentSidePos)

			// Convex corner
			if isSidePosOutside && isAdjacentSidePosOutside {
				corners++
			}

			// Concave corner
			if !isSidePosOutside && !isAdjacentSidePosOutside && isOutside(cornerPos) {
				corners++
			}

			if !isNextPosOutside && !visited[nextPos] {
				queue = append(queue, nextPos)
				visited[nextPos] = true
			}
		}
	}

	// In any polygon, sides == corners (vertices)
	sides := corners

	return area * sides
}

func parseInput(input string) (garden map[point]rune, width int, height int) {
	garden = make(map[point]rune)
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, plot := range []rune(line) {
			garden[point{X: x, Y: y}] = plot
		}

		width = len(line)
		height++
	}

	return garden, width, height
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
