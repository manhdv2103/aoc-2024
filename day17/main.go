package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/pkg/z3"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	regA, regB, regC, program := parseInput(input)
	output := make([]int, 0)

	instPointer := 0
	for instPointer < len(program) {
		opcode := program[instPointer]
		operand := program[instPointer+1]
		comboOperand := comboOperand(operand, regA, regB, regC)

		switch opcode {
		case 0:
			adv(comboOperand, &regA)
		case 1:
			bxl(operand, &regB)
		case 2:
			bst(comboOperand, &regB)
		case 3:
			newPointer, willJump := jnz(operand, &regA)
			if willJump {
				instPointer = newPointer
				continue
			}
		case 4:
			bxc(operand, &regB, &regC)
		case 5:
			output = append(output, out(comboOperand))
		case 6:
			bdv(comboOperand, &regA, &regB)
		case 7:
			cdv(comboOperand, &regA, &regC)
		}

		instPointer += 2
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output)), ","), "[]")
}

func comboOperand(operand, regA, regB, regC int) int {
	if operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	}

	panic("invalid operand")
}

func adv(comOp int, regA *int) {
	*regA = *regA / utils.PowInt(2, comOp)
}
func bxl(litOp int, regB *int) {
	*regB = *regB ^ litOp
}
func bst(comOp int, regB *int) {
	*regB = comOp % 8
}
func jnz(litOp int, regA *int) (newPointer int, willJump bool) {
	if *regA == 0 {
		return -1, false
	}
	return litOp, true
}
func bxc(_ int, regB, regC *int) {
	*regB = *regB ^ *regC
}
func out(comOp int) (output int) {
	return comOp % 8
}
func bdv(comOp int, regA, regB *int) {
	*regB = *regA / utils.PowInt(2, comOp)
}
func cdv(comOp int, regA, regC *int) {
	*regC = *regA / utils.PowInt(2, comOp)
}

func solveP2(input string) string {
	_, initRegB, initRegC, program := parseInput(input)

	minValidRegA := math.MaxInt
	for _, jumpOrder := range getValidJumpOrders(program) {
		ctx := z3.MakeContext()
		defer ctx.Delete()

		opt := z3.MakeOptimize(ctx)
		defer opt.Delete()

		bvFac := z3.MakeBVFactory(ctx, 64)
		validRegA := bvFac.MakeConst("A")

		regA := validRegA
		regB := bvFac.MakeInt(initRegB)
		regC := bvFac.MakeInt(initRegC)

		outPointer := 0
		jumpPointer := 0
		instPointer := 0
		for instPointer < len(program) {
			opcode := program[instPointer]
			operand := program[instPointer+1]
			comboOperand := comboOperandZ3(bvFac, operand, regA, regB, regC)

			switch opcode {
			case 0:
				advZ3(comboOperand, &regA)
			case 1:
				bxlZ3(operand, &regB, bvFac)
			case 2:
				bstZ3(comboOperand, &regB, bvFac)
			case 3:
				willJump := jumpOrder[jumpPointer]
				jumpPointer++

				if willJump {
					opt.Assert(regA.Ne(bvFac.MakeInt(0)))
					instPointer = operand
					continue
				}

				opt.Assert(regA.Eq(bvFac.MakeInt(0)))
			case 4:
				bxcZ3(operand, &regB, &regC)
			case 5:
				opt.Assert(bvFac.MakeInt(program[outPointer]).Eq(outZ3(comboOperand, bvFac)))
				outPointer++
			case 6:
				bdvZ3(comboOperand, &regA, &regB)
			case 7:
				cdvZ3(comboOperand, &regA, &regC)
			}

			instPointer += 2
		}

		opt.Minimize(validRegA)

		if sat, _ := opt.Check(); sat {
			var success bool
			validRegA, success = opt.Eval(validRegA)
			if success {
				minValidRegA = utils.MinInt(minValidRegA, validRegA.ToIntValue())
			}
		}
	}

	return strconv.Itoa(minValidRegA)
}

func comboOperandZ3(bvFac z3.BVFactory, operand int, regA, regB, regC z3.BV) z3.BV {
	if operand <= 3 {
		return bvFac.MakeInt(operand)
	}

	switch operand {
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	}

	panic("invalid operand")
}

func advZ3(comOp z3.BV, regA *z3.BV) {
	// a = a / (2 ^ comOp)
	*regA = (*regA).ShiftRight(comOp)
}
func bxlZ3(litOp int, regB *z3.BV, bvFac z3.BVFactory) {
	*regB = (*regB).Xor(bvFac.MakeInt(litOp))
}
func bstZ3(comOp z3.BV, regB *z3.BV, bvFac z3.BVFactory) {
	// b = comOp % 8
	*regB = comOp.And(bvFac.MakeInt(7))
}
func bxcZ3(_ int, regB, regC *z3.BV) {
	*regB = (*regB).Xor(*regC)
}
func outZ3(comOp z3.BV, bvFac z3.BVFactory) (output z3.BV) {
	// out = comOp % 8
	return comOp.And(bvFac.MakeInt(7))
}
func bdvZ3(comOp z3.BV, regA, regB *z3.BV) {
	// b = a / (2 ^ comOp)
	*regB = (*regA).ShiftRight(comOp)
}
func cdvZ3(comOp z3.BV, regA, regC *z3.BV) {
	// c = a / (2 ^ comOp)
	*regC = (*regA).ShiftRight(comOp)
}

type validJumpState struct {
	outTimes       int
	jumpList       []bool
	currentPointer int
}

func getValidJumpOrders(program []int) [][]bool {
	validStates := make([][]bool, 0)
	queue := []validJumpState{{0, []bool{}, 0}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.outTimes > len(program) {
			continue
		}

		if item.currentPointer >= len(program) {
			if item.outTimes == len(program) {
				validStates = append(validStates, item.jumpList)
			}
			continue
		}

		opcode := program[item.currentPointer]

		nextJumpList := item.jumpList
		if opcode == 3 {
			nextPointer := program[item.currentPointer+1]
			queue = append(queue, validJumpState{
				item.outTimes,
				append(slices.Clone(item.jumpList), true),
				nextPointer,
			})
			nextJumpList = append(item.jumpList, false)
		}

		queue = append(queue, validJumpState{
			item.outTimes + utils.Ternary(opcode == 5, 1, 0),
			nextJumpList,
			item.currentPointer + 2,
		})
	}

	return validStates
}

func parseInput(input string) (regA, regB, regC int, program []int) {
	lines := strings.Split(input, "\n")

	regA = utils.MustAtoi(utils.ExtractNumStrings(lines[0])[0])
	regB = utils.MustAtoi(utils.ExtractNumStrings(lines[1])[0])
	regC = utils.MustAtoi(utils.ExtractNumStrings(lines[2])[0])

	for _, num := range utils.ExtractNumStrings(lines[4]) {
		program = append(program, utils.MustAtoi(num))
	}

	return regA, regB, regC, program
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
