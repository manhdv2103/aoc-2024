package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = false

type point = utils.Point

func solveP1(input string) string {
	robot, warehouseMap, moves := parseInput(input)

	for _, move := range moves {
		newRobotPos := utils.AddPoint(robot, move)
		obj, isOccupied := warehouseMap[newRobotPos]
		if !isOccupied {
			robot = newRobotPos
			continue
		}

		if obj == '#' {
			continue
		}

		newBoxPos := utils.AddPoint(newRobotPos, move)
		newPosBox := warehouseMap[newBoxPos]
		for newPosBox != '#' {
			if newPosBox == 'O' {
				newBoxPos = utils.AddPoint(newBoxPos, move)
				newPosBox = warehouseMap[newBoxPos]
				continue
			}

			warehouseMap[newBoxPos] = warehouseMap[newRobotPos]
			delete(warehouseMap, newRobotPos)
			robot = newRobotPos
			break
		}
	}

	total := 0
	for pos, obj := range warehouseMap {
		if obj == 'O' {
			total += 100*pos.Y + pos.X
		}
	}

	return strconv.Itoa(total)
}

type halfBox struct {
	side rune
	pos  point
}

func solveP2(input string) string {
	robot, warehouseMap, moves := parseInputP2(input)

moveLoop:
	for _, move := range moves {
		newRobotPos := utils.AddPoint(robot, move)
		obj, isOccupied := warehouseMap[newRobotPos]
		if !isOccupied {
			robot = newRobotPos
			continue
		}

		if obj == '#' {
			continue
		}

		if move.X != 0 {
			newBoxPos := utils.AddPoint(newRobotPos, move)
			newPosBox := warehouseMap[newBoxPos]
			for newPosBox != '#' {
				if newPosBox == '[' || newPosBox == ']' {
					newBoxPos = utils.AddPoint(newBoxPos, move)
					newPosBox = warehouseMap[newBoxPos]
					continue
				}

				for newBoxPos != newRobotPos {
					nextNewBoxPos := utils.AddPoint(newBoxPos, utils.InversePoint(move))
					warehouseMap[newBoxPos] = warehouseMap[nextNewBoxPos]
					newBoxPos = nextNewBoxPos
				}
				delete(warehouseMap, newRobotPos)
				robot = newRobotPos
				break
			}
			continue
		}

		leftPos, rightPos := getFullBox(obj, newRobotPos)
		movingBoxs := make(utils.Set[halfBox])
		queue := []halfBox{
			{warehouseMap[leftPos], leftPos},
			{warehouseMap[rightPos], rightPos},
		}
		for len(queue) > 0 {
			hb := queue[0]
			queue = queue[1:]

			if movingBoxs[hb] {
				continue
			}
			movingBoxs[hb] = true

			nextHalfBoxPos := utils.AddPoint(hb.pos, move)
			nextHalfBox, isOccupied := warehouseMap[nextHalfBoxPos]

			if !isOccupied {
				continue
			}

			if nextHalfBox == '#' {
				continue moveLoop
			}

			nextLeftPos, nextRightPos := getFullBox(nextHalfBox, nextHalfBoxPos)
			queue = append(queue, halfBox{warehouseMap[nextLeftPos], nextLeftPos})
			queue = append(queue, halfBox{warehouseMap[nextRightPos], nextRightPos})
		}

		for hb := range movingBoxs {
			delete(warehouseMap, hb.pos)
		}
		for hb := range movingBoxs {
			warehouseMap[utils.AddPoint(hb.pos, move)] = hb.side
		}
		robot = newRobotPos
	}

	total := 0
	for pos, obj := range warehouseMap {
		if obj == '[' {
			total += 100*pos.Y + pos.X
		}
	}

	return strconv.Itoa(total)
}

func getFullBox(boxSide rune, boxPos point) (left point, right point) {
	if boxSide == '[' {
		return boxPos, utils.AddPoint(boxPos, point{X: 1, Y: 0})
	}
	return utils.AddPoint(boxPos, point{X: -1, Y: 0}), boxPos
}

func parseInput(input string) (point, map[point]rune, []point) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	warehouseMap := make(map[point]rune)
	robot := point{X: 0, Y: 0}
	for y, line := range strings.Split(parts[0], "\n") {
		for x, obj := range line {
			switch obj {
			case '@':
				robot.X = x
				robot.Y = y
			case 'O', '#':
				warehouseMap[point{X: x, Y: y}] = obj
			}
		}
	}

	return robot, warehouseMap, parseMoves(parts[1])
}

func parseInputP2(input string) (point, map[point]rune, []point) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	warehouseMap := make(map[point]rune)
	robot := point{X: 0, Y: 0}
	for y, line := range strings.Split(parts[0], "\n") {
		for x, obj := range line {
			switch obj {
			case '@':
				robot.X = x * 2
				robot.Y = y
			case '#':
				warehouseMap[point{X: x * 2, Y: y}] = '#'
				warehouseMap[point{X: x*2 + 1, Y: y}] = '#'
			case 'O':
				warehouseMap[point{X: x * 2, Y: y}] = '['
				warehouseMap[point{X: x*2 + 1, Y: y}] = ']'
			}
		}
	}

	return robot, warehouseMap, parseMoves(parts[1])
}

func parseMoves(s string) []point {
	moves := make([]point, len(s))
	for i, move := range []rune(s) {
		switch move {
		case 'v':
			moves[i] = utils.Directions[0]
		case '>':
			moves[i] = utils.Directions[1]
		case '^':
			moves[i] = utils.Directions[2]
		case '<':
			moves[i] = utils.Directions[3]
		}
	}

	return moves
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
