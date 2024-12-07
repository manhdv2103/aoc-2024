package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
)

const PART = "2"
const WILL_SUBMIT = true

type point struct {
	x int
	y int
}

func solveP1(input string) string {
	obstructions, guard, mapWidth, mapHeight := parseInput(input)

	visited := map[point]bool{guard: true}
	direction := point{0, -1}
	for true {
		nextPos := step(guard, direction)
		if obstructions[nextPos] {
			direction = turn(direction)
			continue
		}

		if nextPos.x < 0 || nextPos.x >= mapWidth || nextPos.y < 0 || nextPos.y >= mapHeight {
			break
		}

		guard = nextPos
		visited[guard] = true
	}

	return strconv.Itoa(len(visited))
}

func solveP2(input string) string {
	obstructions, guard, mapWidth, mapHeight := parseInput(input)

	total := map[point]bool{}
	visited := map[point]bool{guard: true}
	turningPositions := map[point]bool{}
	direction := point{0, -1}
	for true {
		nextPos := step(guard, direction)
		if obstructions[nextPos] {
			turningPositions[guard] = true
			direction = turn(direction)
			continue
		}

		if nextPos.x < 0 || nextPos.x >= mapWidth || nextPos.y < 0 || nextPos.y >= mapHeight {
			break
		}

		if !visited[nextPos] && !total[nextPos] && isCyclic(
			nextPos,
			guard,
			direction,
			obstructions,
			turningPositions,
			mapWidth,
			mapHeight,
		) {
			total[nextPos] = true
		}

		guard = nextPos
		visited[guard] = true
	}

	return strconv.Itoa(len(total))
}

func isCyclic(
	newObstruction point,
	guard point,
	direction point,
	obstructions map[point]bool,
	prevTurningPositions map[point]bool,
	mapWidth int,
	mapHeight int,
) bool {
	turningPositions := map[point]bool{}
	movedAfterTurning := false
	unmovedTurningCount := 0
	for true {
		nextPos := step(guard, direction)
		if nextPos == newObstruction || obstructions[nextPos] {
			// Guard turns a full 360 deg (a loop)
			if unmovedTurningCount == 4 {
				return true
			}

			// For each turning position, there's only 1 valid direction (just try it),
			// so if the guard reaches a previously seen turning position, they're stucked in a loop
			if (turningPositions[guard] || prevTurningPositions[guard]) && movedAfterTurning {
				return true
			}

			turningPositions[guard] = true
			direction = turn(direction)
			movedAfterTurning = false
			unmovedTurningCount++
			continue
		}

		guard = nextPos
		movedAfterTurning = true
		unmovedTurningCount = 0

		if guard.x < 0 || guard.x >= mapWidth || guard.y < 0 || guard.y >= mapHeight {
			return false
		}
	}

	return false
}

func step(value point, direction point) point {
	return point{value.x + direction.x, value.y + direction.y}
}

func turn(direction point) point {
	return point{-direction.y, direction.x}
}

func parseInput(input string) (obstructions map[point]bool, startPoint point, mapWidth int, mapHeight int) {
	obstructions = make(map[point]bool)

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, pos := range line {
			if pos == '#' {
				obstructions[point{x, y}] = true
			}

			if pos == '^' {
				startPoint = point{x, y}
			}
		}

		mapWidth = len(line)
		mapHeight++
	}

	return obstructions, startPoint, mapWidth, mapHeight
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
