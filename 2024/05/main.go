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

	rules := make(map[int][]int)
	var manuals [][]int
	parsingRules := true
	for _, l := range lines {
		if len(l) == 0 {
			parsingRules = false
			continue
		}
		if parsingRules {
			nums := strings.Split(l, "|")

			leftNum, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			rightNum, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			rules[leftNum] = append(rules[leftNum], rightNum)
		} else {
			numsStr := strings.Split(l, ",")
			var nums []int
			for _, nStr := range numsStr {
				n, err := strconv.Atoi(nStr)
				if err != nil {
					panic(err)
				}
				nums = append(nums, n)
			}
			manuals = append(manuals, nums)
		}
	}

	res := 0
	for _, manual := range manuals {
		valid := true
		for i, page := range manual {
			if valid == false {
				break
			}
			for _, prev := range manual[:i] {
				if slices.Contains(rules[page], prev) {
					valid = false
					break
				}
			}
		}
		if valid {
			res += manual[len(manual)/2]
		}
	}

	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	rules := make(map[int][]int)
	var manuals [][]int
	parsingRules := true
	for _, l := range lines {
		if len(l) == 0 {
			parsingRules = false
			continue
		}
		if parsingRules {
			nums := strings.Split(l, "|")

			leftNum, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			rightNum, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			rules[leftNum] = append(rules[leftNum], rightNum)
		} else {
			numsStr := strings.Split(l, ",")
			var nums []int
			for _, nStr := range numsStr {
				n, err := strconv.Atoi(nStr)
				if err != nil {
					panic(err)
				}
				nums = append(nums, n)
			}
			manuals = append(manuals, nums)
		}
	}

	res := 0
	for _, manual := range manuals {
		if !isValid(manual, rules) {
			manual = generateManual(manual, rules)
			res += manual[len(manual)/2]
		}
	}

	return res
}

func isValid(manual []int, rules map[int][]int) bool {
	valid := true
	for i, page := range manual {
		if valid == false {
			break
		}
		for _, prev := range manual[:i] {
			if slices.Contains(rules[page], prev) {
				valid = false
				break
			}
		}
	}
	return valid
}

func generateManual(pages []int, rules map[int][]int) []int {
	var res []int
	for _, page := range pages {
		nextPages := rules[page]
		inserted := false
		for i, page2 := range res {
			if slices.Contains(nextPages, page2) {
				start := res[:i]
				end := res[i:]
				res = slices.Concat(start, []int{page}, end)
				inserted = true
				break
			}
		}
		if !inserted {
			res = append(res, page)
		}
	}
	return res
}
