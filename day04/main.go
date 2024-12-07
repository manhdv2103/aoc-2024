package main

import (
	"strconv"
	"strings"

	"github.com/manhdv2103/aoc-2024/aoc"
)

const PART = "2"
const WILL_SUBMIT = false

type direction struct {
	x int
	y int
}

func solveP1(input string) string {
	wordBoard := parseInput(input)

	directions := []direction{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	count := 0
	for i, line := range wordBoard {
		for j, letter := range line {
			if letter != 'X' {
				continue
			}

			for _, d := range directions {
				if i+3*d.y < 0 || j+3*d.x < 0 || i+3*d.y >= len(wordBoard) || j+3*d.x >= len(line) {
					continue
				}

				if wordBoard[i+1*d.y][j+1*d.x] == 'M' &&
					wordBoard[i+2*d.y][j+2*d.x] == 'A' &&
					wordBoard[i+3*d.y][j+3*d.x] == 'S' {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count)
}

func solveP2(input string) string {
	wordBoard := parseInput(input)

	diagonal1Directions := []direction{
		{1, 1},
		{-1, -1},
	}

	diagonal2Directions := []direction{
		{-1, 1},
		{1, -1},
	}

	count := 0
	for i, line := range wordBoard {
		for j, letter := range line {
			if letter != 'A' {
				continue
			}

			if isValidDiagonal(wordBoard, i, j, diagonal1Directions) &&
				isValidDiagonal(wordBoard, i, j, diagonal2Directions) {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func isValidDiagonal(wordBoard [][]rune, i int, j int, diagonalDirections []direction) bool {
	for _, d := range diagonalDirections {
		if i+d.y < 0 || j+d.x < 0 || i+d.y >= len(wordBoard) || j+d.x >= len(wordBoard[i]) ||
			i-d.y < 0 || j-d.x < 0 || i-d.y >= len(wordBoard) || j-d.x >= len(wordBoard[i]) {
			continue
		}

		if wordBoard[i+d.y][j+d.x] == 'M' &&
			wordBoard[i-d.y][j-d.x] == 'S' {
			return true
		}
	}

	return false
}

func parseInput(input string) (wordBoard [][]rune) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		wordBoard = append(wordBoard, []rune(line))
	}

	return wordBoard
}

func main() {
	aoc.Process(solveP1, solveP2, PART, WILL_SUBMIT)
}
