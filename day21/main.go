package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

type point = utils.Point

type cacheKey struct {
	start rune
	end   rune
}

type state struct {
	pos     point
	lastKey rune
	seqLen  int
}

func solveP1(input string) string {
	codes := parseInput(input)
	return strconv.Itoa(getTotalComplexities(codes, 3))
}

func solveP2(input string) string {
	codes := parseInput(input)
	return strconv.Itoa(getTotalComplexities(codes, 26))
}

func getTotalComplexities(codes []string, directionalKeypadNum int) int {
	numkeyMap := make(map[rune]point)
	numkeyMap['7'] = point{X: 0, Y: 0}
	numkeyMap['8'] = point{X: 1, Y: 0}
	numkeyMap['9'] = point{X: 2, Y: 0}
	numkeyMap['4'] = point{X: 0, Y: 1}
	numkeyMap['5'] = point{X: 1, Y: 1}
	numkeyMap['6'] = point{X: 2, Y: 1}
	numkeyMap['1'] = point{X: 0, Y: 2}
	numkeyMap['2'] = point{X: 1, Y: 2}
	numkeyMap['3'] = point{X: 2, Y: 2}
	numkeyMap[' '] = point{X: 0, Y: 3}
	numkeyMap['0'] = point{X: 1, Y: 3}
	numkeyMap['A'] = point{X: 2, Y: 3}

	dirkeyMap := make(map[rune]point)
	dirkeyMap[' '] = point{X: 0, Y: 0}
	dirkeyMap['^'] = point{X: 1, Y: 0}
	dirkeyMap['A'] = point{X: 2, Y: 0}
	dirkeyMap['<'] = point{X: 0, Y: 1}
	dirkeyMap['v'] = point{X: 1, Y: 1}
	dirkeyMap['>'] = point{X: 2, Y: 1}

	keyMaps := make([]map[rune]point, directionalKeypadNum)
	keyMaps[directionalKeypadNum-1] = numkeyMap
	for i := 0; i < directionalKeypadNum-1; i++ {
		keyMaps[i] = dirkeyMap
	}

	total := 0
	cache := make([]map[cacheKey]int, directionalKeypadNum)
	for _, code := range codes {
		shortestSeqLen := 0
		start := 'A'
		for _, key := range code {
			shortestSeqLen += getShortestSeqLen(start, key, directionalKeypadNum-1, keyMaps, cache)
			start = key
		}

		total += shortestSeqLen * utils.MustAtoi(code[:len(code)-1])
	}

	return total
}

func getShortestSeqLen(start rune, end rune, layer int, keyMaps []map[rune]point, cache []map[cacheKey]int) int {
	if layer == 0 {
		return utils.ManhattanDistance(keyMaps[0][start], keyMaps[0][end]) + 1 // for 'A'
	}

	if cache[layer] == nil {
		cache[layer] = make(map[cacheKey]int)
	}
	shortestSeqLen, found := cache[layer][cacheKey{start, end}]
	if found {
		return shortestSeqLen
	}

	shortestSeqLen = math.MaxInt
	keyMap := keyMaps[layer]
	gapPos := keyMap[' ']
	endPos := keyMap[end]
	queue := []state{{keyMap[start], 'A', 0}}
	for len(queue) > 0 {
		item := queue[0]
		pos := item.pos
		queue = queue[1:]

		if pos == endPos {
			seqLen := item.seqLen + getShortestSeqLen(item.lastKey, 'A', layer-1, keyMaps, cache)
			if seqLen < shortestSeqLen {
				shortestSeqLen = seqLen
			}
			continue
		}

		for _, direction := range utils.Directions {
			nextPos := utils.AddPoint(pos, direction)
			if nextPos == gapPos ||
				utils.ManhattanDistance(nextPos, endPos) >= utils.ManhattanDistance(pos, endPos) {
				continue
			}

			dirKey := '<'
			if direction == utils.Directions[0] {
				dirKey = 'v'
			} else if direction == utils.Directions[1] {
				dirKey = '>'
			} else if direction == utils.Directions[2] {
				dirKey = '^'
			}

			queue = append(queue, state{
				nextPos,
				dirKey,
				item.seqLen + getShortestSeqLen(item.lastKey, dirKey, layer-1, keyMaps, cache),
			})
		}
	}

	cache[layer][cacheKey{start, end}] = shortestSeqLen
	return shortestSeqLen
}

func parseInput(input string) (codes []string) {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
