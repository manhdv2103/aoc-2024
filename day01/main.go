package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	list1, list2 := parseInput(input)

	sort.Ints(list1)
	sort.Ints(list2)

	total := 0
	for i := 0; i < len(list1); i++ {
		total += utils.AbsInt(list1[i] - list2[i])
	}

	return strconv.Itoa(total)
}

func solveP2(input string) string {
	list1, list2 := parseInput(input)

	idMap := make(map[int]int)
	for _, id := range list2 {
		idMap[id] = idMap[id] + 1
	}

	score := 0
	for _, id := range list1 {
		score += id * idMap[id]
	}

	return strconv.Itoa(score)
}

func parseInput(input string) ([]int, []int) {
	pairs := strings.Split(strings.TrimSpace(input), "\n")
	list1 := make([]int, len(pairs))
	list2 := make([]int, len(pairs))

	for i, pair := range pairs {
		ids := strings.Split(pair, "   ")
		list1[i], _ = strconv.Atoi(ids[0])
		list2[i], _ = strconv.Atoi(ids[1])
	}

	return list1, list2
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
