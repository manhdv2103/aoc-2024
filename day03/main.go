package main

import (
	"regexp"
	"strconv"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	matches := parseInput(input)

	total := 0
	for _, match := range matches {
		if match["inst"] != "mul" {
			continue
		}

		total += utils.MustAtoi(match["x"]) * utils.MustAtoi(match["y"])
	}

	return strconv.Itoa(total)
}

func solveP2(input string) string {
	matches := parseInput(input)

	do := true
	total := 0
	for _, match := range matches {
		switch match["inst"] {
		case "do":
			do = true
		case "don't":
			do = false
		case "mul":
			if !do {
				continue
			}

			total += utils.MustAtoi(match["x"]) * utils.MustAtoi(match["y"])
		}
	}

	return strconv.Itoa(total)
}

func parseInput(input string) (matchMaps []map[string]string) {
	re := regexp.MustCompile(`(?P<inst>mul)\((?P<x>\d{1,3}),(?P<y>\d{1,3})\)|(?P<inst>do)\(\)|(?P<inst>don't)\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		matchMap := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i == 0 || name == "" || match[i] == "" {
				continue
			}

			matchMap[name] = match[i]
		}

		matchMaps = append(matchMaps, matchMap)
	}

	return matchMaps
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
