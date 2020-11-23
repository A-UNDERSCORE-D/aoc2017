package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2017/util"
)

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func filterP1Valid(input []string) []string {
	var out []string
outer:
	for _, p := range input {
		words := make(map[string]int)
		for _, v := range strings.Split(p, " ") {
			words[v]++
		}

		for _, v := range words {
			if v > 1 {
				continue outer
			}
		}
		out = append(out, p)
	}
	return out
}

func part1(input []string) string {
	valid := 0
outer:
	for _, p := range input {
		words := make(map[string]int)
		for _, v := range strings.Split(p, " ") {
			words[v]++
		}

		for _, v := range words {
			if v > 1 {
				continue outer
			}
		}
		valid++

	}
	return fmt.Sprint(len(filterP1Valid(input)))
}

func lettersInWord(word string) map[rune]int {
	out := make(map[rune]int)
	for _, l := range word {
		out[l]++
	}
	return out
}

func cmpMap(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if res, exists := b[k]; !exists || res != v {
			return false
		}
	}

	return true
}

func part2(input []string) string {
	p1Valid := input // filterP1Valid(input)
	var p2Valid []string
outer:
	for _, v := range p1Valid {
		split := strings.Split(v, " ")
		for i, word := range split {
			letters := lettersInWord(word)
			for j, word2 := range split {
				if i == j {
					continue
				}
				w2Letters := lettersInWord(word2)
				if cmpMap(letters, w2Letters) {
					continue outer
				}
			}
		}
		p2Valid = append(p2Valid, v)
	}

	return fmt.Sprint(len(p2Valid))
}
