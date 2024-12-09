package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
	"github.com/manhdv2103/aoc-2024/utils"
)

const PART = "2"
const WILL_SUBMIT = true

func solveP1(input string) string {
	blocks := parseInput(input)

	checksum := 0
	count := 0
	forwardIndex := 0
	backwardIndex := (len(blocks) - 1) / 2 * 2
	for forwardIndex < backwardIndex {
		if forwardIndex%2 == 0 {
			checksum, count = updateChecksum(checksum, count, blocks[forwardIndex], forwardIndex)
		} else {
			spaceSize := blocks[forwardIndex]
			fillCount := 0

			for fillCount < spaceSize && forwardIndex < backwardIndex {
				fileSize := blocks[backwardIndex]
				newFillCount := utils.MinInt(spaceSize-fillCount, fileSize)

				blocks[backwardIndex] -= newFillCount
				fillCount += newFillCount
				checksum, count = updateChecksum(checksum, count, newFillCount, backwardIndex)

				if blocks[backwardIndex] == 0 {
					backwardIndex -= 2
				}
			}
		}

		forwardIndex++
	}

	return strconv.Itoa(checksum)
}

func solveP2(input string) string {
	blocks := parseInput(input)

	checksum := 0
	count := 0
	minSpaceSize := 0
	forwardIndex := 0
	backwardIndex := (len(blocks) - 1) / 2 * 2
	for forwardIndex <= backwardIndex {
		if forwardIndex%2 == 0 {
			fileSize := blocks[forwardIndex]

			if fileSize > 0 {
				checksum, count = updateChecksum(checksum, count, fileSize, forwardIndex)
			} else {
				// negative fileSize means file has been moved, this is the remaining spaceSize
				count -= fileSize
			}
		} else {
			spaceSize := blocks[forwardIndex]
			lastFileIndex := -1

			for i := backwardIndex; i > forwardIndex && spaceSize > minSpaceSize; i -= 2 {
				fileSize := blocks[i]
				if fileSize <= 0 {
					continue
				}

				if fileSize <= spaceSize {
					blocks[i] = -fileSize
					spaceSize -= fileSize
					checksum, count = updateChecksum(checksum, count, fileSize, i)
				} else if lastFileIndex == -1 {
					lastFileIndex = i
				}
			}

			count += spaceSize
			minSpaceSize = spaceSize
			if lastFileIndex != -1 {
				backwardIndex = lastFileIndex
			}
		}

		forwardIndex++
	}

	return strconv.Itoa(checksum)
}

func updateChecksum(checksum int, count int, fileCount int, fileIndex int) (int, int) {
	newCount := count + fileCount
	return checksum + sumRange(count, newCount-1)*(fileIndex/2), newCount
}

func sumRange(start int, end int) int {
	return int((float64(end-start+1) / 2) * float64(end+start))
}

func parseInput(input string) []int {
	input = strings.TrimSpace(input)
	blocks := make([]int, len(input))

	for i, rune := range input {
		blocks[i] = int(rune - '0')
	}

	return blocks
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
