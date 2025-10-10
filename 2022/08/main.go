package main

import (
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
	lines := strings.SplitSeq(file, "\n")

	var trees [][]int
	for l := range lines {
		if len(l) == 0 {
			continue
		}

		trees = append(trees, make([]int, 0))
		for t := range strings.SplitSeq(l, "") {
			n, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			trees[len(trees)-1] = append(trees[len(trees)-1], n)
		}
	}

	res := 0
	visited := make([][]bool, len(trees))
	for i := range visited {
		visited[i] = make([]bool, len(trees[0]))
	}

	for y := range trees {
		prevHighest := 0
		for x := range trees[0] {
			if trees[y][x] > prevHighest {
				prevHighest = trees[y][x]
				if visited[y][x] {
					continue
				}
				visited[y][x] = true
				res++
			}
		}
		prevHighest = 0
		for x := len(trees[0]) - 1; x >= 0; x-- {
			if trees[y][x] > prevHighest {
				prevHighest = trees[y][x]
				if visited[y][x] {
					continue
				}
				visited[y][x] = true
				res++
			}
		}
	}

	for x := range trees[0] {
		prevHighest := 0
		for y := range trees {
			if trees[y][x] > prevHighest {
				prevHighest = trees[y][x]
				if visited[y][x] {
					continue
				}
				visited[y][x] = true
				res++
			}
		}
		prevHighest = 0
		for y := len(trees) - 1; y >= 0; y-- {
			if trees[y][x] > prevHighest {
				prevHighest = trees[y][x]
				if visited[y][x] {
					continue
				}
				visited[y][x] = true
				res++
			}
		}
	}

	for y := range visited {
		if y == 0 || y == len(visited)-1 {
			for x := range visited[y] {
				if !visited[y][x] {
					visited[y][x] = true
					res++
				}
			}
		}
		if !visited[y][0] {
			visited[y][0] = true
			res++
		}
		if !visited[y][len(visited[y])-1] {
			visited[y][len(visited[y])-1] = true
			res++
		}
	}

	return res
}

func solve2(file string) int {
	lines := strings.SplitSeq(file, "\n")

	var trees [][]int
	for l := range lines {
		if len(l) == 0 {
			continue
		}

		trees = append(trees, make([]int, 0))
		for t := range strings.SplitSeq(l, "") {
			n, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			trees[len(trees)-1] = append(trees[len(trees)-1], n)
		}
	}

	maxScore := 0
	for y := range trees {
		for x := range trees[0] {
			currTree := trees[y][x]
			partialScore := 0
			for currX := x - 1; currX >= 0; currX-- {
				partialScore++
				if currTree <= trees[y][currX] {
					break
				}
			}
			score := partialScore
			partialScore = 0
			for currX := x + 1; currX < len(trees[0]); currX++ {
				partialScore++
				if currTree <= trees[y][currX] {
					break
				}
			}
			score *= partialScore
			partialScore = 0
			for currY := y - 1; currY >= 0; currY-- {
				partialScore++
				if currTree <= trees[currY][x] {
					break
				}
			}
			score *= partialScore
			partialScore = 0
			for currY := y + 1; currY < len(trees); currY++ {
				partialScore++
				if currTree <= trees[currY][x] {
					break
				}
			}
			score *= partialScore
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}
