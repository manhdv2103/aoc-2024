package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	secretNums := parseInput(input)

	total := 0
	for _, secretNum := range secretNums {
		for range 2000 {
			secretNum = prune(mix(secretNum*64, secretNum))
			secretNum = prune(mix(secretNum/32, secretNum))
			secretNum = prune(mix(secretNum*2048, secretNum))
		}

		total += secretNum
	}

	return strconv.Itoa(total)
}

type sequence struct {
	val0 int
	val1 int
	val2 int
	val3 int
}

func solveP2(input string) string {
	secretNums := parseInput(input)

	var wg sync.WaitGroup
	seqMaps := make(chan map[sequence]int, len(secretNums))
	for _, secretNum := range secretNums {
		wg.Add(1)
		go calculateSeqMap(secretNum, &wg, seqMaps)
	}

	go func() {
		wg.Wait()
		close(seqMaps)
	}()

	totalSeqMap := make(map[sequence]int)
	for seqMap := range seqMaps {
		for k, v := range seqMap {
			totalSeqMap[k] += v
		}
	}

	maxTotal := 0
	for _, total := range totalSeqMap {
		if total > maxTotal {
			maxTotal = total
		}
	}

	return strconv.Itoa(maxTotal)
}

func calculateSeqMap(secretNum int, wg *sync.WaitGroup, seqMaps chan<- map[sequence]int) {
	defer wg.Done()
	seqMap := make(map[sequence]int)
	changes := make([]int, 0, 4)

	prevPrice := secretNum % 10
	for range 2000 {
		secretNum = prune(mix(secretNum*64, secretNum))
		secretNum = prune(mix(secretNum/32, secretNum))
		secretNum = prune(mix(secretNum*2048, secretNum))
		price := secretNum % 10

		if len(changes) == 4 {
			changes = changes[1:]
		}
		changes = append(changes, price-prevPrice)

		if len(changes) == 4 {
			seq := sequence{changes[0], changes[1], changes[2], changes[3]}
			if _, exists := seqMap[seq]; !exists {
				seqMap[seq] = price
			}
		}

		prevPrice = price
	}

	seqMaps <- seqMap
}

func mix(value, secretNum int) int {
	return value ^ secretNum
}

func prune(secretNum int) int {
	return secretNum % 16777216
}

func parseInput(input string) (secretNums []int) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		secretNums = append(secretNums, utils.MustAtoi(line))
	}

	return secretNums
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
