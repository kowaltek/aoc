package main

import (
	"log"
	"math"
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

	var left, right []int
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		words := strings.Split(l, "   ")
		leftStr := words[0]
		rightStr := words[1]
		leftNum, err := strconv.Atoi(leftStr)
		if err != nil {
			log.Println("error converting number")
			log.Println(err)
			os.Exit(1)
		}
		rightNum, err := strconv.Atoi(rightStr)
		if err != nil {
			log.Println("error converting number")
			log.Println(err)
			os.Exit(1)
		}
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	res := 0
	for i := range len(left) {
		res += int(math.Abs(float64(left[i] - right[i])))
	}
	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	var left []int
	right := make(map[int]int)
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		words := strings.Split(l, "   ")
		leftStr := words[0]
		rightStr := words[1]
		leftNum, err := strconv.Atoi(leftStr)
		if err != nil {
			log.Println("error converting number")
			log.Println(err)
			os.Exit(1)
		}
		rightNum, err := strconv.Atoi(rightStr)
		if err != nil {
			log.Println("error converting number")
			log.Println(err)
			os.Exit(1)
		}
		left = append(left, leftNum)
		n := right[rightNum]
		right[rightNum] = n + 1
	}

	res := 0
	for _, leftV := range left {
		res += leftV * right[leftV]
	}
	return res
}
