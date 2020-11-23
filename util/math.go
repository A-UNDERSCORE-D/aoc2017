// Package util contains various util methods
package util

import (
	"math"
	"strconv"
)

// Min does what it says on the tin, but with ints and not float64s
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinOf(ints ...int) int {
	switch len(ints) {
	case 0:
		panic("no ints specified")
	case 1:
		return ints[0]
	case 2:
		return Min(ints[0], ints[1])
	default:
		curMin := ints[0]
		for _, i := range ints[1:] {
			curMin = Min(curMin, i)
		}
		return curMin
	}
}

// Max does what it says on the tin, but with ints and not float64s
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxOf applies Max to multiple ints, returning the largest of all of them
func MaxOf(ints ...int) (int, int) {
	switch len(ints) {
	case 0:
		panic("no ints specified")
	case 1:
		return 0, ints[0]
	case 2:
		m := Max(ints[0], ints[1])
		if m == ints[0] {
			return 0, m
		}
		return 1, m
	default:
		curMax := ints[0]
		curIdx := 0
		for idx, i := range ints[1:] {
			oldMax := curMax
			curMax = Max(curMax, i)
			if curMax == i && curMax != oldMax {
				curIdx = idx + 1
			}
		}
		return curIdx, curMax
	}
}

// Abs returns the absolute value of the given number
func Abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// GetInt returns the given string as an int, or panics if it is invalid
func GetInt(in string) int {
	res, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return res
}

// RoundDown returns the given float rounded towards 0
func RoundDown(in float64) int {
	return int(math.Trunc(in))
}

// GCD and LCM copied from https://github.com/sotsoguk/adventOfCode2019_go/blob/a10a0ca4a6f6f31770dfca1e234a08f41f0865b8/utils/mathUtils/mathUtils.go#L79 because I am lazy

// Gcd Finds the Greatest common Denominator
func Gcd(x, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Lcm finds the Least Common Multiple (LCM) via GCD
func Lcm(a, b int64, integers ...int64) int64 {
	result := a * b / Gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}
	return result
}
