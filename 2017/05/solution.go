package main

import (
	"fmt"
	"time"

	"awesome-dragon.science/go/adventofcode2017/util"
)

func main() {
	input := /* []int{0, 3, 0, 1, -3} */ util.ReadInts("input.txt")
	p2Input := make([]int, len(input))
	if copy(p2Input, input) != len(input) {
		panic("asd")
	}
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(p2Input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func part1(input []int) string {
	curIdx := 0
	steps := 0
	for {
		if curIdx >= len(input) || curIdx < 0 {
			break
		}
		// printIdx(input, curIdx)
		oldIdx := curIdx
		curIdx += input[curIdx]
		input[oldIdx]++

		steps++
	}
	return fmt.Sprint(steps)
}

func part2(input []int) string {
	curIdx := 0
	steps := 0
	for {
		if curIdx >= len(input) || curIdx < 0 {
			break
		}
		// printIdx(input, curIdx)
		oldIdx := curIdx
		curIdx += input[curIdx]
		i := input[oldIdx]
		switch {
		case i >= 3:
			input[oldIdx]--
		default:
			input[oldIdx]++
		}

		steps++
	}
	return fmt.Sprint(steps)
}
