package main

import (
	"log"
	"os"
	"slices"
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

	maxCals := 0
	currCals := 0
	done := true
	for _, l := range lines {
		if len(l) == 0 {
			if currCals > maxCals {
				maxCals = currCals
			}
			done = true
			continue
		}
		num, err := strconv.Atoi(l)
		if err != nil {
			log.Println("error converting number")
			log.Println(l)
			log.Println(err)
			os.Exit(1)
		}
		if done {
			currCals = 0
		}
		currCals += num
		done = false
	}
	if currCals > maxCals {
		maxCals = currCals
	}

	return maxCals
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	var cals []int
	done := true
	for _, l := range lines {
		if len(l) == 0 {
			done = true
			continue
		}
		num, err := strconv.Atoi(l)
		if err != nil {
			log.Println("error converting number")
			log.Println(l)
			log.Println(err)
			os.Exit(1)
		}
		if done {
			cals = append(cals, 0)
		}
		cals[len(cals)-1] += num
		done = false
	}

	slices.Sort(cals)
	lastI := len(cals) - 1
	return cals[lastI] + cals[lastI-1] + cals[lastI-2]
}
