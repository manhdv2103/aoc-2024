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

type state struct {
	pos   point
	steps int
}

func solveP1(input string) string {
	bytes := parseInput(input)

	fallenBytes := make(utils.Set[point])
	for i := range 1024 {
		fallenBytes[bytes[i]] = true
	}

	steps := len(bfs(fallenBytes)) - 1

	return strconv.Itoa(steps)
}

func solveP2(input string) string {
	bytes := parseInput(input)

	fallenBytes := make(utils.Set[point])
	for i := range 1024 {
		fallenBytes[bytes[i]] = true
	}

	path := bfs(fallenBytes)
	var finalByte point

	for i := 1024; i <= len(bytes); i++ {
		byte := bytes[i]
		fallenBytes[byte] = true

		if !path[byte] {
			continue
		}

		newPath := bfs(fallenBytes)
		if newPath == nil {
			finalByte = byte
			break
		}

		path = newPath
	}

	return strconv.Itoa(finalByte.X) + "," + strconv.Itoa(finalByte.Y)
}

func bfs(fallenBytes utils.Set[point]) utils.Set[point] {
	width := 71
	height := 71

	previous := make(map[point]point)

	visited := make(utils.Set[point])
	queue := []state{{point{X: 0, Y: 0}, 0}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if visited[item.pos] {
			continue
		}
		visited[item.pos] = true

		if item.pos.X == 70 && item.pos.Y == 70 {
			at := item.pos
			path := utils.Set[point]{at: true}
			for {
				nextAt, ok := previous[at]

				if !ok {
					break
				}
				path[nextAt] = true
				at = nextAt
			}

			return path
		}

		for _, direction := range utils.Directions {
			nextPos := utils.AddPoint(item.pos, direction)
			if !utils.InBounds(nextPos, width, height) || visited[nextPos] || fallenBytes[nextPos] {
				continue
			}

			previous[nextPos] = item.pos
			queue = append(queue, state{nextPos, item.steps + 1})
		}
	}

	return nil
}

func parseInput(input string) (bytes []point) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		nums := utils.ExtractNumStrings(line)
		bytes = append(bytes, point{X: utils.MustAtoi(nums[0]), Y: utils.MustAtoi(nums[1])})
	}

	return bytes
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
