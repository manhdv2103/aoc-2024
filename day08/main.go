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
	antennas, width, height := parseInput(input)

	antinodes := make(utils.Set[point])
	for _, sameFreqAntennas := range antennas {
		for _, antenna1 := range sameFreqAntennas {
			for _, antenna2 := range sameFreqAntennas {
				if antenna1 == antenna2 {
					continue
				}

				direction := point{
					X: antenna2.X - antenna1.X,
					Y: antenna2.Y - antenna1.Y,
				}

				antinode := point{
					X: antenna1.X - direction.X,
					Y: antenna1.Y - direction.Y,
				}

				if utils.InBounds(antinode, width, height) {
					antinodes[antinode] = true
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}

func solveP2(input string) string {
	antennas, width, height := parseInput(input)

	antinodes := make(utils.Set[point])
	for _, sameFreqAntennas := range antennas {
		for _, antenna1 := range sameFreqAntennas {
			for _, antenna2 := range sameFreqAntennas {
				if antenna1 == antenna2 {
					continue
				}

				direction := point{
					X: antenna2.X - antenna1.X,
					Y: antenna2.Y - antenna1.Y,
				}

				antinode := antenna1
				for utils.InBounds(antinode, width, height) {
					antinodes[antinode] = true
					antinode = point{
						X: antinode.X - direction.X,
						Y: antinode.Y - direction.Y,
					}
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}

func parseInput(input string) (map[rune][]point, int, int) {
	antennas := make(map[rune][]point)
	width := 0
	height := 0

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, rune := range []rune(line) {
			if rune != '.' {
				antennas[rune] = append(antennas[rune], point{X: x, Y: y})
			}
		}

		width = len(line)
		height++
	}

	return antennas, width, height
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
