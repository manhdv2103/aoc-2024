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

type args struct {
	stone      uint64
	blinksLeft int
}

func solveP1(input string) string {
	stones := parseInput(input)

	total := 0
	cache := make(map[args]int)
	for _, stone := range stones {
		total += blink(stone, 25, cache)
	}

	return strconv.Itoa(total)
}

func solveP2(input string) string {
	stones := parseInput(input)

	total := 0
	cache := make(map[args]int)
	for _, stone := range stones {
		total += blink(stone, 75, cache)
	}

	return strconv.Itoa(total)
}

func blink(stone uint64, blinksLeft int, cache map[args]int) (newStoneCount int) {
	if blinksLeft == 0 {
		return 1
	}

	cachedCount, ok := cache[args{stone, blinksLeft}]
	if ok {
		return cachedCount
	}

	nextArgsList := make([]args, 0)

	if stone == 0 {
		nextArgsList = append(nextArgsList, args{1, blinksLeft - 1})
	} else {
		digits := countDigit(stone)
		if digits%2 == 0 {
			s1, s2 := splitStone(stone, digits/2)

			nextArgsList = append(nextArgsList, args{s1, blinksLeft - 1})
			nextArgsList = append(nextArgsList, args{s2, blinksLeft - 1})
		} else {
			nextArgsList = append(nextArgsList, args{stone * 2024, blinksLeft - 1})
		}
	}

	total := 0
	for _, nextArgs := range nextArgsList {
		count := blink(nextArgs.stone, nextArgs.blinksLeft, cache)
		cache[nextArgs] = count
		total += count
	}

	return total
}

func countDigit(stone uint64) int {
	count := 0
	for stone > 0 {
		stone /= 10
		count++
	}

	return count
}

func splitStone(stone uint64, splitDigit int) (s1 uint64, s2 uint64) {
	count := 0
	for count < splitDigit {
		s2 += (stone % 10) * uint64(math.Pow10(count))
		stone /= 10
		count++
	}

	return stone, s2
}

func parseInput(input string) (stones []uint64) {
	for _, stone := range strings.Fields(input) {
		stones = append(stones, utils.MustAtou64(stone))
	}

	return stones
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
