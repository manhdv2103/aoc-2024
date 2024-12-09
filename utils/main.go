package utils

import "strconv"

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func ToInts(strs []string) []int {
	ints := make([]int, 0)
	for _, str := range strs {
		int, _ := strconv.Atoi(str)
		ints = append(ints, int)
	}

	return ints
}

// https://stackoverflow.com/a/57213476
func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return n
}

func MustAtoi64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return n
}

func MustAtou64(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return n
}

type Point struct {
	X int
	Y int
}

func InBounds(point Point, width int, height int) bool {
	return point.X >= 0 && point.X < width && point.Y >= 0 && point.Y < height
}

type Set[T comparable] map[T]bool
