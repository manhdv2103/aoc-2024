package utils

import "strconv"

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
