package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

type point = utils.Point

type robot struct {
	position point
	velocity point
}

func solveP1(input string) string {
	robots := parseInput(input)
	width := 101
	height := 103

	for range 100 {
		for _, robot := range robots {
			robot.position = utils.WrapPoint(
				utils.AddPoint(robot.position, robot.velocity),
				width,
				height,
			)
		}
	}

	middle := point{X: width / 2, Y: height / 2}
	quadrants := []int{0, 0, 0, 0}
	for _, robot := range robots {
		p := robot.position
		if p.X == middle.X || p.Y == middle.Y {
			continue
		}

		if p.X < middle.X {
			if p.Y < middle.Y {
				quadrants[0]++
			} else {
				quadrants[1]++
			}
		} else {
			if p.Y < middle.Y {
				quadrants[2]++
			} else {
				quadrants[3]++
			}
		}
	}

	return strconv.Itoa(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])
}

func solveP2(input string) string {
	robots := parseInput(input)
	width := 101
	height := 103

	highestMaxRobotGroupCount := 0
	elapsedSecondsToMaybeDisplayChristmasTree := 0
	maybeChristmasTreeRobotMap := make(utils.Set[point])
	for i := range width * height {
		robotCountMap := make(map[point]int)
		for _, robot := range robots {
			robot.position = utils.WrapPoint(
				utils.AddPoint(robot.position, robot.velocity),
				width,
				height,
			)
			robotCountMap[robot.position]++
		}

		// Hopefully the Christmas tree is formed by a bunch of robots that are next to each other
		maxRobotGroupCount := 0
		visited := make(utils.Set[point])
		for _, robot := range robots {
			if visited[robot.position] {
				continue
			}

			robotGroupCount := 0
			queue := []point{robot.position}
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]

				if visited[p] {
					continue
				}
				visited[p] = true

				for _, direction := range utils.Directions {
					nextPos := utils.AddPoint(p, direction)
					if count, ok := robotCountMap[nextPos]; ok {
						robotGroupCount += count
						queue = append(queue, nextPos)
					}
				}
			}

			maxRobotGroupCount = utils.MaxInt(maxRobotGroupCount, robotGroupCount)
		}

		prevHighestMaxRobotGroupCount := highestMaxRobotGroupCount
		highestMaxRobotGroupCount = utils.MaxInt(highestMaxRobotGroupCount, maxRobotGroupCount)
		if highestMaxRobotGroupCount != prevHighestMaxRobotGroupCount {
			elapsedSecondsToMaybeDisplayChristmasTree = i + 1
			maybeChristmasTreeRobotMap = make(utils.Set[point])
			for _, robot := range robots {
				maybeChristmasTreeRobotMap[robot.position] = true
			}
		}
	}

	for y := range height {
		for x := range width {
			if maybeChristmasTreeRobotMap[point{X: x, Y: y}] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return strconv.Itoa(elapsedSecondsToMaybeDisplayChristmasTree)
}

func parseInput(input string) (robots []*robot) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		nums := utils.ExtractNumStrings(line)
		robots = append(robots, &robot{
			point{
				X: utils.MustAtoi(nums[0]),
				Y: utils.MustAtoi(nums[1]),
			},
			point{
				X: utils.MustAtoi(nums[2]),
				Y: utils.MustAtoi(nums[3]),
			},
		})
	}

	return robots
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
