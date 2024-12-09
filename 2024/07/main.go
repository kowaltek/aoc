package main

import (
	"fmt"
	"log"
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
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")

	res := 0
	for _, l := range lines {
		parts := strings.Split(l, ":")
		testRes, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		numsStr := strings.Split(strings.Trim(parts[1], " "), " ")
		var nums []int
		for _, nStr := range numsStr {
			num, err := strconv.Atoi(nStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		if solvable1(testRes, nums...) {
			res += testRes
		}
	}

	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")

	res := 0
	for _, l := range lines {
		parts := strings.Split(l, ":")
		testRes, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		numsStr := strings.Split(strings.Trim(parts[1], " "), " ")
		var nums []int
		for _, nStr := range numsStr {
			num, err := strconv.Atoi(nStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		if solvable2(testRes, nums...) {
			res += testRes
		}
	}

	return res
}

func solvable1(res int, nums ...int) bool {
	if len(nums) == 1 {
		if res == nums[0] {
			return true
		} else {
			return false
		}
	}

	added := []int{nums[0] + nums[1]}
	if len(nums) > 2 {
		added = append(added, nums[2:]...)
	}
	multiplied := []int{nums[0] * nums[1]}
	if len(nums) > 2 {
		multiplied = append(multiplied, nums[2:]...)
	}
	return solvable1(res, added...) || solvable1(res, multiplied...)
}

func solvable2(res int, nums ...int) bool {
	if len(nums) == 1 {
		if res == nums[0] {
			return true
		} else {
			return false
		}
	}

	added := []int{nums[0] + nums[1]}
	if len(nums) > 2 {
		added = append(added, nums[2:]...)
	}
	multiplied := []int{nums[0] * nums[1]}
	if len(nums) > 2 {
		multiplied = append(multiplied, nums[2:]...)
	}

	nStr := fmt.Sprintf("%d%d", nums[0], nums[1])
	n, err := strconv.Atoi(nStr)
	if err != nil {
		panic(err)
	}
	concatenated := []int{n}
	if len(nums) > 2 {
		concatenated = append(concatenated, nums[2:]...)
	}
	return solvable2(res, added...) || solvable2(res, multiplied...) || solvable2(res, concatenated...)
}
