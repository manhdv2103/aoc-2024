package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/pkg/trie"
)

const PART = "2"
const WILL_SUBMIT = true

type state struct {
	patternIndices []int
	designPointer  int
}

func solveP1(input string) string {
	patternTrie, designs := parseInput(input)

	count := 0
	cache := make(map[string]bool)
	for _, design := range designs {
		if isDesignPossible(design, patternTrie, cache) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func isDesignPossible(design string, patternTrie trie.Trie, cache map[string]bool) bool {
	if len(design) == 0 {
		return true
	}

	isPossible, cached := cache[design]
	if cached {
		return isPossible
	}

	for i := 0; i < len(design); i++ {
		if trie.IsInTrie(patternTrie, design[:i+1]) &&
			isDesignPossible(design[i+1:], patternTrie, cache) {
			cache[design] = true
			return true
		}
	}

	cache[design] = false
	return false
}

func solveP2(input string) string {
	patternTrie, designs := parseInput(input)

	total := 0
	cache := make(map[string]int)
	for _, design := range designs {
		total += countPossibleDesign(design, patternTrie, cache)
	}

	return strconv.Itoa(total)
}

func countPossibleDesign(design string, patternTrie trie.Trie, cache map[string]int) int {
	if len(design) == 0 {
		return 1
	}

	cachedCount, cached := cache[design]
	if cached {
		return cachedCount
	}

	count := 0
	for i := 0; i < len(design); i++ {
		if trie.IsInTrie(patternTrie, design[:i+1]) {
			count += countPossibleDesign(design[i+1:], patternTrie, cache)
		}
	}

	cache[design] = count

	return count
}

func parseInput(input string) (patternTrie trie.Trie, designs []string) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	patternTrie = trie.MakeTrie(strings.Split(parts[0], ", "))

	for _, design := range strings.Split(parts[1], "\n") {
		designs = append(designs, design)
	}

	return patternTrie, designs
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
