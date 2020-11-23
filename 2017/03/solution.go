package main

import (
	"fmt"
	"time"

	"awesome-dragon.science/go/adventofcode2017/util"
)

func main() {
	input := 277678
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func mhtnDist(a, b point) int {
	return util.Abs(a.x-b.x) + util.Abs(a.y-b.y)
}

type point struct {
	x int
	y int
}

func (p point) add(other point) point {
	return point{
		p.x + other.x,
		p.y + other.y,
	}
}

func (p point) String() string {
	return p.GoString()
}

func (p point) GoString() string {
	return fmt.Sprintf("(%+02d, %+02d)", p.x, p.y)
}

func spiralCoord(spiralIdx int) point {
	out := point{}

	doingX := true
	direction := 1
	steps := 0
	remaining := 1
	length := 1
	for steps != spiralIdx-1 {
		if doingX {
			out.x += direction
			remaining -= 1
		} else {
			out.y += direction
			remaining -= 1
		}

		if !doingX && remaining == 0 {
			length += 1
			direction = -direction
		}

		if remaining == 0 {
			doingX = !doingX
			remaining = length
		}

		steps++
	}
	return point{out.y, out.x}
}

func part1(input int) string {
	return fmt.Sprint(mhtnDist(point{}, spiralCoord(277678)))
}

var dirs []point = []point{
	/*
		432
		5x1
		678
	*/
	{+0, +1},
	{+1, +1},
	{+1, +0},
	{+1, -1},
	{-1, +0},
	{-1, -1},
	{+0, -1},
	{-1, +1},
}

func part2(input int) string {
	board := map[point]int{
		{0, 0}: 1,
	}

	val := 0
	i := 2
	for val < input {
		val = 0
		coords := spiralCoord(i)
		for _, dir := range dirs {
			val += board[coords.add(dir)]
		}
		board[coords] = val
		i++
	}

	return fmt.Sprint(val)
}
