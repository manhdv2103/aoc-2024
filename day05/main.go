package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	rules, _, updates := parseInput(input)

	total := 0
	for _, update := range updates {
		if isUpdateCorrect(update, rules) {
			total += utils.MustAtoi(update[len(update)/2])
		}
	}

	return strconv.Itoa(total)
}

func solveP2(input string) string {
	rules, rawRules, updates := parseInput(input)

	ruleMap := make(map[string]bool)
	for _, rawRule := range rawRules {
		ruleMap[rawRule] = true
	}

	total := 0
	for _, update := range updates {
		if !isUpdateCorrect(update, rules) {
			slices.SortFunc(update, func(a, b string) int {
				if ruleMap[a+"|"+b] {
					return -1
				}

				if ruleMap[b+"|"+a] {
					return 1
				}

				return 0
			})

			total += utils.MustAtoi(update[len(update)/2])
		}
	}

	return strconv.Itoa(total)
}

func isUpdateCorrect(update []string, rules map[string][]string) bool {
	previousPages := make(map[string]bool)

	for _, page := range update {
		afterPages := rules[page]

		for _, afterPage := range afterPages {
			if previousPages[afterPage] {
				return false
			}
		}

		previousPages[page] = true
	}

	return true
}

func parseInput(input string) (rules map[string][]string, rawRules []string, updates [][]string) {
	sections := strings.SplitN(input, "\n\n", 2)
	rawRules = strings.Split(strings.TrimSpace(sections[0]), "\n")
	rules = make(map[string][]string)

	for _, ruleStr := range rawRules {
		rule := strings.SplitN(ruleStr, "|", 2)
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	for _, update := range strings.Split(strings.TrimSpace(sections[1]), "\n") {
		updates = append(updates, strings.Split(update, ","))
	}

	return rules, rawRules, updates
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
