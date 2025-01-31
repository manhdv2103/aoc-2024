package main

import (
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	computers, network := parseInput(input)

	duplicatedCount := 0
	for v1 := range computers {
		for v2 := range network[v1] {
			for v3 := range network[v2] {
				if v1[0] != 't' && v2[0] != 't' && v3[0] != 't' {
					continue
				}

				if network[v1][v3] {
					duplicatedCount++
				}
			}
		}
	}

	return strconv.Itoa(duplicatedCount / 6)
}

func solveP2(input string) string {
	computers, network := parseInput(input)

	largestLanPartyComs := slices.Collect(maps.Keys(bronKerbosch(
		utils.Set[string]{},
		computers,
		utils.Set[string]{},
		network,
	)))
	slices.Sort(largestLanPartyComs)

	return strings.Join(largestLanPartyComs, ",")
}

// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func bronKerbosch(
	r utils.Set[string],
	p utils.Set[string],
	x utils.Set[string],
	network map[string]utils.Set[string],
) utils.Set[string] {
	if len(p) == 0 && len(x) == 0 {
		return r
	}

	var u string
	maxSize := -1
	for v := range utils.Union(p, x) {
		if len(network[v]) > maxSize {
			maxSize = len(network[v])
			u = v
		}
	}

	var res utils.Set[string]
	for v := range utils.Difference(p, network[u]) {
		newRes := bronKerbosch(
			utils.Union(r, utils.Set[string]{v: true}),
			utils.Intersect(p, network[v]),
			utils.Intersect(x, network[v]),
			network,
		)
		delete(p, v)
		x[v] = true

		if len(newRes) > len(res) {
			res = newRes
		}
	}

	return res
}

func parseInput(input string) (computers utils.Set[string], network map[string]utils.Set[string]) {
	computers = make(utils.Set[string])
	network = make(map[string]utils.Set[string])

	for _, connection := range strings.Split(strings.TrimSpace(input), "\n") {
		computerPair := strings.SplitN(connection, "-", 2)

		if _, exists := network[computerPair[0]]; !exists {
			network[computerPair[0]] = make(utils.Set[string])
		}
		network[computerPair[0]][computerPair[1]] = true

		if _, exists := network[computerPair[1]]; !exists {
			network[computerPair[1]] = make(utils.Set[string])
		}
		network[computerPair[1]][computerPair[0]] = true

		computers[computerPair[0]] = true
		computers[computerPair[1]] = true
	}

	return computers, network
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
