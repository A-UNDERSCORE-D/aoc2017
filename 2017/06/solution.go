package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2017/util"
)

func main() {
	input := util.GetInts(strings.Split(util.ReadEntireFile("input.txt"), "\t"))
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part1(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func reallocateCycle(banks []int) {
	maxIdx, count := util.MaxOf(banks...)
	banks[maxIdx] = 0

	for i := maxIdx + 1; count > 0; i++ {
		if i >= len(banks) {
			i = 0
		}
		banks[i]++
		count--
	}
}

func intsContain(source [][]int, target []int) bool {
outer:
	for _, v := range source {
		if len(v) != len(target) {
			return false
		}
		for i, num := range target {
			if num != v[i] {
				continue outer
			}
		}
		return true
	}
	return false
}

func part1(banks []int) string {
	var cycleSnapshots [][]int
	s := make([]int, len(banks))
	copy(s, banks)
	cycleSnapshots = append(cycleSnapshots, s)
	cycleCount := 0
	for {
		cycleCount++
		reallocateCycle(banks)
		if intsContain(cycleSnapshots, banks) {
			break
		}

		snapshot := make([]int, len(banks))
		copy(snapshot, banks)
		cycleSnapshots = append(cycleSnapshots, snapshot)
	}

	return fmt.Sprint(cycleCount)
}
