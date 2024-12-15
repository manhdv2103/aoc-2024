package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

type point struct {
	x float64
	y float64
}

type machine struct {
	btnA  point
	btnB  point
	prize point
}

func solveP1(input string) string {
	machines := parseInput(input)

	total := int64(0)
	for _, machine := range machines {
		total += calculateMinToken(machine.btnA, machine.btnB, machine.prize)
	}

	return strconv.FormatInt(total, 10)
}

func solveP2(input string) string {
	machines := parseInput(input)

	total := int64(0)
	for _, machine := range machines {
		total += calculateMinToken(
			machine.btnA,
			machine.btnB,
			point{machine.prize.x + 10000000000000, machine.prize.y + 10000000000000},
		)
	}

	return strconv.FormatInt(total, 10)
}

func calculateMinToken(btnA point, btnB point, prize point) int64 {
	// btnA.x * a + btnB.x * b = prize.x
	// btnA.y * a + btnB.y * b = prize.y

	determinant := btnA.x*btnB.y - btnB.x*btnA.y
	if determinant != 0 {
		a := (prize.x*btnB.y - btnB.x*prize.y) / determinant
		b := (btnA.x*prize.y - prize.x*btnA.y) / determinant

		if a == float64(int64(a)) && b == float64(int64(b)) {
			return 3*int64(a) + int64(b)
		}
	}

	return 0
}

func parseInput(input string) (machines []machine) {
	for _, machineStr := range strings.Split(strings.TrimSpace(input), "\n\n") {
		lines := strings.Split(machineStr, "\n")

		btnANums := utils.ExtractNumStrings(lines[0])
		btnA := point{utils.MustParseFloat(btnANums[0]), utils.MustParseFloat(btnANums[1])}

		btnBNums := utils.ExtractNumStrings(lines[1])
		btnB := point{utils.MustParseFloat(btnBNums[0]), utils.MustParseFloat(btnBNums[1])}

		prizeNums := utils.ExtractNumStrings(lines[2])
		prize := point{utils.MustParseFloat(prizeNums[0]), utils.MustParseFloat(prizeNums[1])}

		machines = append(machines, machine{btnA, btnB, prize})
	}

	return machines
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
