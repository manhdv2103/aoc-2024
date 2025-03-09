package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

type gate struct {
	gateType   string
	inputWireA string
	inputWireB string
	outputWire string
}

func solveP1(input string) string {
	wires, outputGateMap, highestWireNo := parseInputP1(input)

	output := 0
	for i := 0; i <= highestWireNo; i++ {
		outputWire := "z" + fmt.Sprintf("%02d", i)
		g := outputGateMap[outputWire]

		bit := getBit(g, wires, outputGateMap)
		output |= bit << i
	}

	return strconv.Itoa(output)
}

func getBit(g gate, wires map[string]int, gates map[string]gate) int {
	valA, valAExists := wires[g.inputWireA]
	if !valAExists {
		valA = getBit(gates[g.inputWireA], wires, gates)
		wires[g.inputWireA] = valA
	}

	valB, valBExists := wires[g.inputWireB]
	if !valBExists {
		valB = getBit(gates[g.inputWireB], wires, gates)
		wires[g.inputWireB] = valB
	}

	switch g.gateType {
	case "AND":
		return valA & valB
	case "OR":
		return valA | valB
	case "XOR":
		return valA ^ valB
	}

	panic("invalid gate type")
}

func parseInputP1(input string) (wires map[string]int, outputGateMap map[string]gate, highestWireNo int) {
	wires = make(map[string]int)
	outputGateMap = make(map[string]gate)
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		parts := strings.Split(line, ": ")
		wires[parts[0]] = utils.Ternary(parts[1] == "0", 0, 1)
	}

	for _, line := range strings.Split(parts[1], "\n") {
		parts := strings.Fields(line)
		outputGateMap[parts[4]] = gate{parts[1], parts[0], parts[2], parts[4]}

		if parts[4][0] == 'z' {
			highestWireNo = utils.MaxInt(highestWireNo, utils.MustAtoi(parts[4][1:]))
		}
	}

	return wires, outputGateMap, highestWireNo
}

type gateInputKey struct {
	gateType  string
	inputWire string
}

// Assumes the adder is a ripple carry adder
func solveP2(input string) string {
	inputOutputMap, gates, highestWireNo := parseInputP2(input)
	highestWire := "z" + fmt.Sprintf("%02d", highestWireNo)

	wrongWires := make([]string, 0, 4)
	for _, gate := range gates {
		if gate.inputWireA == "x00" || gate.inputWireB == "x00" {
			if gate.gateType == "XOR" && gate.outputWire != "z00" {
				wrongWires = append(wrongWires, "z00")
			}

			_, outputUsedInAndGate := inputOutputMap[gateInputKey{"AND", gate.outputWire}]
			if gate.gateType == "AND" && !outputUsedInAndGate {
				wrongWires = append(wrongWires, gate.outputWire)
			}
		} else if gate.outputWire == highestWire {
			if gate.gateType != "OR" {
				wrongWires = append(wrongWires, highestWire)
			}
		} else if gate.outputWire[0] == 'z' {
			if gate.gateType != "XOR" {
				wrongWires = append(wrongWires, gate.outputWire)
			}
		} else if gate.gateType == "XOR" {
			if gate.inputWireA[0] == 'x' || gate.inputWireB[0] == 'x' {
				_, outputUsedInXorGate := inputOutputMap[gateInputKey{"XOR", gate.outputWire}]
				if !outputUsedInXorGate {
					wrongWires = append(wrongWires, gate.outputWire)
				}
			} else if gate.outputWire[0] != 'z' {
				wrongWires = append(wrongWires, gate.outputWire)
			}
		} else if gate.gateType == "AND" {
			_, outputUsedInOrGate := inputOutputMap[gateInputKey{"OR", gate.outputWire}]
			if !outputUsedInOrGate {
				wrongWires = append(wrongWires, gate.outputWire)
			}
		} else if gate.gateType == "OR" {
			_, outputUsedInXorGate := inputOutputMap[gateInputKey{"XOR", gate.outputWire}]
			_, outputUsedInAndGate := inputOutputMap[gateInputKey{"AND", gate.outputWire}]
			if !outputUsedInXorGate || !outputUsedInAndGate {
				wrongWires = append(wrongWires, gate.outputWire)
			}
		}
	}

	sort.Strings(wrongWires)

	return strings.Join(wrongWires, ",")
}

func parseInputP2(input string) (inputOutputMap map[gateInputKey]string, gates []gate, highestWireNo int) {
	inputOutputMap = make(map[gateInputKey]string)
	gates = make([]gate, 0)
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, line := range strings.Split(parts[1], "\n") {
		parts := strings.Fields(line)
		g := gate{parts[1], parts[0], parts[2], parts[4]}

		gates = append(gates, g)
		inputOutputMap[gateInputKey{g.gateType, g.inputWireA}] = g.outputWire
		inputOutputMap[gateInputKey{g.gateType, g.inputWireB}] = g.outputWire

		if parts[4][0] == 'z' {
			highestWireNo = utils.MaxInt(highestWireNo, utils.MustAtoi(parts[4][1:]))
		}
	}

	return inputOutputMap, gates, highestWireNo
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
