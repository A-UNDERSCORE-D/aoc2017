package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var testInput = `5 9 2 8
9 4 7 3
3 8 6 5`

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

func makeInts(input []string) []int {
	out := make([]int, 0, len(input))
	for _, v := range input {
		if v == "" {
			continue
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}

func splitIntoMatrix(input string, sep string) [][]int {
	var out [][]int
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		out = append(out, makeInts(strings.Split(r, sep)))
	}
	return out
}

func minMax(in []int) (int, int) {
	max := math.MinInt64
	min := math.MaxInt16

	for _, v := range in {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}

	}

	return min, max
}

func part1(input string) string {
	matrix := splitIntoMatrix(input, "\t")
	sum := 0
	for _, v := range matrix {
		if len(v) == 0 {
			continue
		}
		min, max := minMax(v)
		sum += max - min
	}
	return fmt.Sprint(sum)
}

func findDivisible(in []int) int {
	for i, v := range in {
		for j, t := range in {
			if i == j {
				continue
			}

			if v%t == 0 {
				return v / t
			}
		}
	}
	return -1
}

func part2(input string) string {
	matrix := splitIntoMatrix(input, "\t")
	sum := 0
	for _, r := range matrix {
		if len(r) == 0 {
			continue
		}
		res := findDivisible(r)
		if res < 0 {
			panic(res)
		}
		sum += res
	}

	return fmt.Sprint(sum)
}
