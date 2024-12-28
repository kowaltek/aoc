package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var cache = make(map[string]uint)

type button struct {
	x, y int
	val  string
}

var numpadButtons = []button{
	{0, 0, "7"},
	{2, 3, "A"},
	{1, 0, "8"},
	{2, 0, "9"},
	{0, 1, "4"},
	{1, 1, "5"},
	{2, 1, "6"},
	{0, 2, "1"},
	{1, 2, "2"},
	{2, 2, "3"},
	{1, 3, "0"},
}

var dirButtons = []button{
	{2, 0, "A"},
	{1, 0, "^"},
	{0, 1, "<"},
	{1, 1, "v"},
	{2, 1, ">"},
}

var steps = []struct{ x, y int }{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type node struct {
	b        button
	d1       button
	d2       button
	path     string
	dist     int
	selected bool
}

var re = regexp.MustCompile(`\d+`)

// 145648 too high
// 143536 too high

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve1(file)
	res2 := solve2(file)
	log.Printf("Got result: %d\n", res)
	log.Printf("Got result: %d\n", res2)
}

func solve1(file string) int {
	file = strings.Trim(file, "\n")
	sequences := strings.Split(file, "\n")
	res := 0
	for _, sequence := range sequences {
		sequence = strings.Trim(sequence, "\n")
		shortestNumpad := findPaths("A"+sequence, numpadButtons)
		var shortestDirpad1 []string
		for _, path := range shortestNumpad {
			shortestDirpad1 = append(shortestDirpad1, findPaths("A"+path, dirButtons)...)
		}
		var shortestDirpad2 []string
		for _, path := range shortestDirpad1 {
			shortestDirpad2 = append(shortestDirpad2, findPaths("A"+path, dirButtons)...)
		}
		numPartStr := re.Find([]byte(sequence))
		numPart, err := strconv.Atoi(string(numPartStr))
		if err != nil {
			panic(err)
		}
		shortest := shortestDirpad2[0]
		for _, path := range shortestDirpad2 {
			if len(shortest) > len(path) {
				shortest = path
			}
		}
		fmt.Println(shortest)
		fmt.Println(len(shortest) * numPart)
		res += len(shortest) * numPart
	}
	return res
}

func findPaths(sequence string, pad []button) []string {
	var paths []string
	currPosI := slices.IndexFunc(pad, func(b button) bool {
		return b.val == string(sequence[0])
	})
	currPos := pad[currPosI]
	for _, c := range strings.Split(sequence[1:], "") {
		nextPosI := slices.IndexFunc(pad, func(b button) bool {
			return b.val == c
		})
		nextPos := pad[nextPosI]
		dx := abs(nextPos.x - currPos.x)
		dy := abs(nextPos.y - currPos.y)
		verDir := "^"
		if nextPos.y > currPos.y {
			verDir = "v"
		}
		horDir := "<"
		if nextPos.x > currPos.x {
			horDir = ">"
		}
		dirs := strings.Repeat(horDir, dx)
		dirs += strings.Repeat(verDir, dy)
		permutations := getPermutations(dirs)
		permutations = dedupStrings(permutations)
		for i := 0; i < len(permutations); i++ {
			isValid := true
			pos := currPos
			for _, move := range strings.Split(permutations[i], "") {
				switch move {
				case "^":
					pos.y--
					if !slices.ContainsFunc(pad, func(b button) bool {
						return pos.x == b.x && pos.y == b.y
					}) {
						isValid = false
					}
				case "v":
					pos.y++
					if !slices.ContainsFunc(pad, func(b button) bool {
						return pos.x == b.x && pos.y == b.y
					}) {
						isValid = false
					}
				case ">":
					pos.x++
					if !slices.ContainsFunc(pad, func(b button) bool {
						return pos.x == b.x && pos.y == b.y
					}) {
						isValid = false
					}
				case "<":
					pos.x--
					if !slices.ContainsFunc(pad, func(b button) bool {
						return pos.x == b.x && pos.y == b.y
					}) {
						isValid = false
					}
				}
				if !isValid {
					break
				}
			}
			if isValid {
				permutations[i] = permutations[i] + "A"
			} else {
				tmp := permutations[i+1:]
				permutations = permutations[:i]
				permutations = append(permutations, tmp...)
				i--
			}
		}
		var tmp []string
		if paths != nil {
			for _, path := range paths {
				for _, currPath := range permutations {
					tmp = append(tmp, path+currPath)
				}
			}
			paths = tmp
		} else {
			paths = permutations
		}
		currPos = nextPos
	}
	return paths
}

func dedupStrings(strs []string) []string {
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if strs[i] == strs[j] {
				tmp := strs[j+1:]
				strs = strs[:j]
				strs = append(strs, tmp...)
				j--
			}
		}
	}
	return strs
}

func getPermutations(s string) []string {
	if s == "" {
		return []string{""}
	}
	return permute([]rune(s), len(s))
}

func permute(s []rune, k int) []string {
	if k == 1 {
		return []string{string(s)}
	}
	var permutations []string
	permutations = append(permutations, permute(s, k-1)...)
	for i := range k - 1 {
		if k%2 == 0 {
			s[i], s[k-1] = s[k-1], s[i]
		} else {
			s[0], s[k-1] = s[k-1], s[0]
		}
		permutations = append(permutations, permute(s, k-1)...)
	}
	return permutations
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func solve2(file string) uint {
	file = strings.Trim(file, "\n")
	sequences := strings.Split(file, "\n")
	res := uint(0)
	for _, sequence := range sequences {
		sequence = strings.Trim(sequence, "\n")
		numPartStr := re.Find([]byte(sequence))
		numPart, err := strconv.Atoi(string(numPartStr))
		if err != nil {
			panic(err)
		}
		fmt.Println(sequence)
		shortest := getTotalCost(sequence)
		fmt.Println(shortest * uint(numPart))
		res += shortest * uint(numPart)
	}
	return res
}

func getCost(a, b string, numpad bool, depth int) uint {
	key := key(a, b, depth)
	res, ok := cache[key]
	if ok && !numpad {
		return res
	}
	if depth == 0 {
		paths := findPaths(a+b, dirButtons)
		shortest := uint(math.MaxUint)
		for _, path := range paths {
			if uint(len(path)) < shortest {
				shortest = uint(len(path))
			}
		}
		return shortest
	}
	var pad []button
	if numpad {
		pad = numpadButtons
	} else {
		pad = dirButtons
	}
	paths := findPaths(a+b, pad)
	shortest := uint(math.MaxUint)
	for _, path := range paths {
		path = "A" + path
		localCost := uint(0)
		for i := range len(path) - 1 {
			localCost += getCost(string(path[i]), string(path[i+1]), false, depth-1)
		}
		if localCost < shortest {
			shortest = localCost
		}
	}
	cache[key] = shortest
	return shortest
}

func key(a, b string, depth int) string {
	return a + b + strconv.Itoa(depth)
}

func getTotalCost(seq string) uint {
	path := "A" + seq
	totalCost := uint(0)
	for i := range len(path) - 1 {
		totalCost += getCost(string(path[i]), string(path[i+1]), true, 25)
	}
	return totalCost
}
