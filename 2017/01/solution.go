package main

import (
	"fmt"
	"strconv"
	"strings"

	"awesome-dragon.science/go/adventofcode2017/util"
)

func main() {
	input := util.ReadLines("input.txt")
	_ = input
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func makeInts(input []string) []int {
	out := make([]int, 0, len(input))
	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}

func part1(input []string) string {
	numbers := makeInts(strings.Split(input[0], ""))
	sum := 0
	next := 0
	for i, d := range numbers {
		if len(numbers)-1 == i {
			next = numbers[0]
		} else {
			next = numbers[i+1]
		}

		if d == next {
			sum += d
		}
	}

	return strconv.Itoa(sum)
}

func part2(input []string) string {
	numbers := makeInts(strings.Split(input[0], ""))
	sum := 0
	next := 0
	for i, d := range numbers {
		next = numbers[((len(numbers) / 2) + i) % len(numbers)]
		if d == next {
			sum += d
		}
	}

	return strconv.Itoa(sum)
}
