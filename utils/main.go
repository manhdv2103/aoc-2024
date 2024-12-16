package utils

import (
	"math"
	"regexp"
	"strconv"
)

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

func MustParseFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)

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

func AddPoint(a Point, b Point) Point {
	return Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func WrapPoint(point Point, width, height int) Point {
	return Point{
		X: Mod(point.X, width),
		Y: Mod(point.Y, height),
	}
}

func RotatePoint(point Point) Point {
	return Point{X: -point.Y, Y: point.X}
}

func InversePoint(point Point) Point {
	return Point{X: -point.X, Y: -point.Y}
}

func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

type Set[T comparable] map[T]bool

var Directions = []Point{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

var numsRegex = regexp.MustCompile(`-?\d+`)

func ExtractNumStrings(s string) []string {
	return numsRegex.FindAllString(s, -1)
}

func Mod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

func Ternary[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}
