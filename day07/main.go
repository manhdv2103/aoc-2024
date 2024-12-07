package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = false

type equation struct {
	value   uint64
	numbers []int
}

func solveP1(input string) string {
	equations := parseInput(input)

	total := uint64(0)
	for _, equation := range equations {
		if isCorrectEquation(equation.value, 0, equation.numbers, false) {
			total += equation.value
		}
	}

	return strconv.FormatUint(total, 10)
}

func solveP2(input string) string {
	equations := parseInput(input)

	total := uint64(0)
	for _, equation := range equations {
		if isCorrectEquation(equation.value, 0, equation.numbers, true) {
			total += equation.value
		}
	}

	return strconv.FormatUint(total, 10)
}

func isCorrectEquation(value uint64, acc uint64, numbers []int, extended bool) bool {
	if len(numbers) == 0 {
		return value == acc
	}

	sumAcc := uint64(numbers[0])
	prodAcc := uint64(numbers[0])
	concatAcc := uint64(numbers[0])
	if acc != 0 {
		sumAcc += acc
		prodAcc *= acc

		if extended {
			multiplier := uint64(10)
			for multiplier <= concatAcc {
				multiplier *= 10
			}
			concatAcc += acc * multiplier
		}
	}

	return isCorrectEquation(value, sumAcc, numbers[1:], extended) ||
		isCorrectEquation(value, prodAcc, numbers[1:], extended) ||
		(extended && isCorrectEquation(value, concatAcc, numbers[1:], true))
}

func parseInput(input string) (equations []equation) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.SplitN(line, ": ", 2)

		numbers := make([]int, 0)
		for _, num := range strings.Fields(parts[1]) {
			numbers = append(numbers, utils.MustAtoi(num))
		}

		equations = append(equations, equation{utils.MustAtou64(parts[0]), numbers})
	}

	return equations
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
