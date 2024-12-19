package main

import (
	"log"
	"os"
	"strings"
)

var (
	cache  = make(map[string]bool)
	cache2 = make(map[string]int)
	towels []string
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
	towels := strings.Split(lines[0], ", ")
	patterns := lines[2:]
	res := 0
	for i := range len(patterns) {
		patterns[i] = strings.Trim(patterns[i], "\n")
		if isPatternPossible(towels, patterns[i]) {
			res++
		}
	}
	return res
}

func isPatternPossible(towels []string, pattern string) bool {
	if pattern == "" {
		return true
	}
	for _, towel := range towels {
		next := strings.SplitN(pattern, towel, 2)
		if next[0] != pattern {
			pattern1Possible, ok := cache[next[0]]
			if !ok {
				pattern1Possible = isPatternPossible(towels, next[0])
			}
			if pattern1Possible {
				cache[next[0]] = true
			} else {
				cache[next[0]] = false
			}
			pattern2Possible, ok := cache[next[1]]
			if !ok {
				pattern2Possible = isPatternPossible(towels, next[1])
			}
			if pattern2Possible {
				cache[next[1]] = true
			} else {
				cache[next[1]] = false
			}
			if pattern1Possible && pattern2Possible {
				return true
			}
		}
	}
	return false
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	towels = strings.Split(lines[0], ", ")
	patterns := lines[2:]
	res := 0
	for i := range len(patterns) {
		patterns[i] = strings.Trim(patterns[i], "\n")
		combinations := isPatternPossible2(towels, patterns[i])
		res += combinations
	}
	return res
}

func isPatternPossible2(towels []string, pattern string) int {
	if pattern == "" {
		return 1
	}
	if sum, ok := cache2[pattern]; ok {
		return sum
	}
	sum := 0
	for _, towel := range towels {
		next, ok := strings.CutPrefix(pattern, towel)
		if ok {
			sum += isPatternPossible2(towels, next)
		}
	}
	cache2[pattern] = sum
	return sum
}
