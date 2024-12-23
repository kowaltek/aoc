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

type button struct {
	x, y int
	val  string
}

var numpadButtons = []button{
	{0, 0, "7"},
	{1, 0, "8"},
	{2, 0, "9"},
	{0, 1, "4"},
	{1, 1, "5"},
	{2, 1, "6"},
	{0, 2, "1"},
	{1, 2, "2"},
	{2, 2, "3"},
	{1, 3, "0"},
	{2, 3, "A"},
}

var dirButtons = []button{
	{1, 0, "^"},
	{2, 0, "A"},
	{0, 1, "<"},
	{1, 1, "v"},
	{2, 1, ">"},
}

type node struct {
	b    button
	path string
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
	log.Printf("Got result: %d\n", res)
}

func solve1(file string) int {
	file = strings.Trim(file, "\n")
	sequences := strings.Split(file, "\n")
	res := 0
	for _, sequence := range sequences {
		sequence = strings.Trim(sequence, "\n")
		shortest := findPath(sequence)
		numPartStr := re.Find([]byte(sequence))
		numPart, err := strconv.Atoi(string(numPartStr))
		if err != nil {
			panic(err)
		}
		fmt.Println(len(shortest))
		fmt.Println(shortest)
		res += len(shortest) * numPart
	}
	return res
}

func findPath(sequence string) string {
	path := ""
	currPos := button{2, 3, "A"}
	for _, c := range strings.Split(sequence, "") {
		unvisited := []node{
			{
				b:    button{0, 0, "7"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{1, 0, "8"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 0, "9"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{0, 1, "4"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{1, 1, "5"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 1, "6"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{0, 2, "1"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{1, 2, "2"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 2, "3"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{1, 3, "0"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 3, "A"},
				path: strings.Repeat(" ", 1024),
			},
		}
		slices.SortFunc(unvisited, func(a, b node) int {
			res := 0
			if a.b.val == currPos.val {
				res = -1
			}
			return res
		})
		nextI := slices.IndexFunc(numpadButtons, func(b button) bool {
			return b.val == c
		})
		currPos = numpadButtons[nextI]
		var visited []node
		currNode := unvisited[0]
		currNode.path = ""
		unvisited = unvisited[1:]
		for len(unvisited) > 0 {
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x+1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				subpath := bestDirpadPath(currNode.path + ">")
				subpathNew := bestDirpadPath(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y+1 == n.b.y
			}); nextNodeI != -1 {
				subpath := bestDirpadPath(currNode.path + "v")
				subpathNew := bestDirpadPath(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x-1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				subpath := bestDirpadPath(currNode.path + "<")
				subpathNew := bestDirpadPath(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y-1 == n.b.y
			}); nextNodeI != -1 {
				subpath := bestDirpadPath(currNode.path + "^")
				subpathNew := bestDirpadPath(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			visited = append(visited, currNode)
			slices.SortFunc(unvisited, func(a, b node) int {
				if len(a.path) < len(b.path) {
					return -1
				} else if len(a.path) == len(b.path) {
					return 0
				} else {
					return 1
				}
			})
			currNode = unvisited[0]
			unvisited = unvisited[1:]
		}
		visited = append(visited, currNode)
		targetNodeI := slices.IndexFunc(visited, func(n node) bool {
			return n.b.val == c
		})
		path += visited[targetNodeI].path
	}
	return path
}

func bestDirpadPath(sequence string) string {
	path := ""
	currPos := button{2, 0, "A"}
	for _, c := range strings.Split(sequence, "") {
		unvisited := []node{
			{
				b:    button{1, 0, "^"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 0, "A"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{0, 1, "<"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{1, 1, "v"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 1, ">"},
				path: strings.Repeat(" ", 1024),
			},
		}
		slices.SortFunc(unvisited, func(a, b node) int {
			res := 0
			if a.b.val == currPos.val {
				res = -1
			}
			return res
		})
		nextI := slices.IndexFunc(dirButtons, func(b button) bool {
			return b.val == c
		})
		if nextI == -1 {
			return sequence
		}
		currPos = dirButtons[nextI]
		var visited []node
		currNode := unvisited[0]
		currNode.path = ""
		unvisited = unvisited[1:]
		for len(unvisited) > 0 {
			noACurrPath, _ := strings.CutSuffix(currNode.path, "A")
			subpath := bestDirpadPath2(noACurrPath + ">A")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x+1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				subpathNew := bestDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			subpath = bestDirpadPath2(noACurrPath + "vA")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y+1 == n.b.y
			}); nextNodeI != -1 {
				subpathNew := bestDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			subpath = bestDirpadPath2(noACurrPath + "<A")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x-1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				subpathNew := bestDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			subpath = bestDirpadPath2(noACurrPath + "^A")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y-1 == n.b.y
			}); nextNodeI != -1 {
				subpathNew := bestDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			visited = append(visited, currNode)
			slices.SortFunc(unvisited, func(a, b node) int {
				if len(a.path) < len(b.path) {
					return -1
				} else if len(a.path) == len(b.path) {
					return 0
				} else {
					return 1
				}
			})
			currNode = unvisited[0]
			unvisited = unvisited[1:]
		}
		visited = append(visited, currNode)
		targetNodeI := slices.IndexFunc(visited, func(n node) bool {
			return n.b.val == c
		})
		path += visited[targetNodeI].path
	}
	return path
}

func bestDirpadPath2(sequence string) string {
	path := ""
	currPos := button{2, 0, "A"}
	for _, c := range strings.Split(sequence, "") {
		unvisited := []node{
			{
				b:    button{1, 0, "^"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 0, "A"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{0, 1, "<"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{1, 1, "v"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 1, ">"},
				path: strings.Repeat(" ", 1024),
			},
		}
		slices.SortFunc(unvisited, func(a, b node) int {
			res := 0
			if a.b.val == currPos.val {
				res = -1
			}
			return res
		})
		nextI := slices.IndexFunc(dirButtons, func(b button) bool {
			return b.val == c
		})
		if nextI == -1 {
			return sequence
		}
		currPos = dirButtons[nextI]
		var visited []node
		currNode := unvisited[0]
		currNode.path = ""
		unvisited = unvisited[1:]
		for len(unvisited) > 0 {
			noACurrPath, _ := strings.CutSuffix(currNode.path, "A")
			subpath := getDirpadPath2(noACurrPath + ">A")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x+1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				subpathNew := getDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			subpath = getDirpadPath2(noACurrPath + "vA")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y+1 == n.b.y
			}); nextNodeI != -1 {
				subpathNew := getDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			subpath = getDirpadPath2(noACurrPath + "<A")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x-1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				subpathNew := getDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			subpath = getDirpadPath2(noACurrPath + "^A")
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y-1 == n.b.y
			}); nextNodeI != -1 {
				subpathNew := getDirpadPath2(unvisited[nextNodeI].path)
				if len(subpathNew) > len(subpath) {
					unvisited[nextNodeI].path = subpath
				}
			}
			visited = append(visited, currNode)
			slices.SortFunc(unvisited, func(a, b node) int {
				if len(a.path) < len(b.path) {
					return -1
				} else if len(a.path) == len(b.path) {
					return 0
				} else {
					return 1
				}
			})
			currNode = unvisited[0]
			unvisited = unvisited[1:]
		}
		visited = append(visited, currNode)
		targetNodeI := slices.IndexFunc(visited, func(n node) bool {
			return n.b.val == c
		})
		path += visited[targetNodeI].path
	}
	return path
}

func getDirpadPath2(sequence string) string {
	path := ""
	currPos := button{2, 0, "A"}
	for _, c := range strings.Split(sequence, "") {
		unvisited := []node{
			{
				b:    button{1, 0, "^"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 0, "A"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{0, 1, "<"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{1, 1, "v"},
				path: strings.Repeat(" ", 1024),
			},
			{
				b:    button{2, 1, ">"},
				path: strings.Repeat(" ", 1024),
			},
		}
		slices.SortFunc(unvisited, func(a, b node) int {
			res := 0
			if a.b.val == currPos.val {
				res = -1
			}
			return res
		})
		nextI := slices.IndexFunc(dirButtons, func(b button) bool {
			return b.val == c
		})
		if nextI == -1 {
			return sequence
		}
		currPos = dirButtons[nextI]
		var visited []node
		currNode := unvisited[0]
		currNode.path = ""
		unvisited = unvisited[1:]
		for len(unvisited) > 0 {
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x+1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				if len(unvisited[nextNodeI].path) > len(currNode.path)+1 {
					unvisited[nextNodeI].path = currNode.path + ">"
				}
			}
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y+1 == n.b.y
			}); nextNodeI != -1 {
				if len(unvisited[nextNodeI].path) > len(currNode.path)+1 {
					unvisited[nextNodeI].path = currNode.path + "v"
				}
			}
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x-1 == n.b.x && currNode.b.y == n.b.y
			}); nextNodeI != -1 {
				if len(unvisited[nextNodeI].path) > len(currNode.path)+1 {
					unvisited[nextNodeI].path = currNode.path + "<"
				}
			}
			if nextNodeI := slices.IndexFunc(unvisited, func(n node) bool {
				return currNode.b.x == n.b.x && currNode.b.y-1 == n.b.y
			}); nextNodeI != -1 {
				if len(unvisited[nextNodeI].path) > len(currNode.path)+1 {
					unvisited[nextNodeI].path = currNode.path + "^"
				}
			}
			visited = append(visited, currNode)
			slices.SortFunc(unvisited, func(a, b node) int {
				if len(a.path) < len(b.path) {
					return -1
				} else if len(a.path) == len(b.path) {
					return 0
				} else {
					return 1
				}
			})
			currNode = unvisited[0]
			unvisited = unvisited[1:]
		}
		visited = append(visited, currNode)
		targetNodeI := slices.IndexFunc(visited, func(n node) bool {
			return n.b.val == c
		})
		path += visited[targetNodeI].path + "A"

	}
	return path
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
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
