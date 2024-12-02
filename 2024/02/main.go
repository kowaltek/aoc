package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve2(file)
	log.Printf("Got result: %d\n", res)
}

func solve1(file string) int {
	lines := strings.Split(file, "\n")
	res := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		levelsStr := strings.Split(l, " ")
		var levels []int
		for _, l := range levelsStr {
			n, err := strconv.Atoi(l)
			if err != nil {
				log.Println("error converting number")
				log.Println(err)
				os.Exit(1)
			}
			levels = append(levels, n)
		}
		if isSafe(levels) {
			res++
		}
	}

	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")
	res := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		levelsStr := strings.Split(l, " ")
		var levels []int
		for _, l := range levelsStr {
			n, err := strconv.Atoi(l)
			if err != nil {
				log.Println("error converting number")
				log.Println(err)
				os.Exit(1)
			}
			levels = append(levels, n)
		}
		if isSafe(levels) {
			res++
			continue
		}
		for i := range levels {
			var sub []int
			sub = append(sub, levels[0:i]...)
			sub = append(sub, levels[i+1:]...)
			if isSafe(sub) {
				res++
				break
			}
		}
	}
	return res
}

func isSafe(levels []int) bool {
	fmt.Println(levels)
	prev := levels[0]
	curr := levels[1]
	var increasing bool
	if curr-prev > 0 {
		increasing = true
	}
	if curr-prev == 0 || math.Abs(float64(curr-prev)) < 1 || math.Abs(float64(curr-prev)) > 3 {
		return false
	}
	safe := true
	for i := 2; i < len(levels); i++ {
		curr = levels[i]
		prev = levels[i-1]
		diff := int(math.Abs(float64(curr - prev)))
		if diff < 1 || diff > 3 {
			safe = false
			break
		}
		if curr-prev > 0 && !increasing {
			safe = false
			break
		}
		if curr-prev < 0 && increasing {
			safe = false
			break
		}
	}
	return safe
}
