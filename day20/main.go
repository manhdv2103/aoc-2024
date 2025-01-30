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
	track, start, end := parseInput(input)

	return strconv.Itoa(countBestCheats(track, start, end, 2))
}

func solveP2(input string) string {
	track, start, end := parseInput(input)

	return strconv.Itoa(countBestCheats(track, start, end, 20))
}

func countBestCheats(track utils.Set[point], start, end point, maxCheatPs int) int {
	psMap := map[point]int{start: 0}
	pos := start
	lastPos := start
	psCounter := 0
	for pos != end {
		for _, direction := range utils.Directions {
			nextPos := utils.AddPoint(pos, direction)
			if track[nextPos] && nextPos != lastPos {
				lastPos = pos
				pos = nextPos
				psCounter++
				psMap[pos] = psCounter
				break
			}
		}
	}

	bestCheats := 0
	for p1 := range track {
		for cheatEndXOffset := -maxCheatPs; cheatEndXOffset <= maxCheatPs; cheatEndXOffset++ {
			for cheatEndYOffset := -maxCheatPs + utils.AbsInt(cheatEndXOffset); cheatEndYOffset <= maxCheatPs-utils.AbsInt(cheatEndXOffset); cheatEndYOffset++ {
				if cheatEndXOffset == 0 && cheatEndYOffset == 0 {
					continue
				}

				cheatEndPos := utils.AddPoint(p1, point{X: cheatEndXOffset, Y: cheatEndYOffset})
				if track[cheatEndPos] {
					savedPs := utils.ManhattanDistance(p1, cheatEndPos)
					if savedPs <= maxCheatPs && psMap[p1]-psMap[cheatEndPos]-savedPs >= 100 {
						bestCheats++
					}
				}
			}
		}
	}

	return bestCheats
}

func parseInput(input string) (track utils.Set[point], start, end point) {
	track = make(utils.Set[point])

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, pos := range []rune(line) {
			p := point{X: x, Y: y}

			switch pos {
			case 'S':
				start = p
				track[p] = true
			case 'E':
				end = p
				track[p] = true
			case '.':
				track[p] = true
			}
		}
	}

	return track, start, end
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
