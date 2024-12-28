package main

import (
	"container/heap"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type nodeHeap []node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *nodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1:n]
	return x
}

type pos struct {
	x, y int
}

type node struct {
	x, y, dist int
}

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
	var bytes []pos
	for _, l := range lines {
		numsStr := strings.Split(l, ",")
		x, err := strconv.Atoi(numsStr[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(numsStr[1])
		if err != nil {
			panic(err)
		}
		bytes = append(bytes, pos{x, y})
	}

	bytes = bytes[:12]
	size := 7

	var unvisited []node
	for y := range size {
		for x := range size {
			if x == 0 && y == 0 {
				unvisited = append(unvisited, node{x, y, 0})
			} else if !slices.ContainsFunc(bytes, func(b pos) bool {
				return x == b.x && y == b.y
			}) {
				unvisited = append(unvisited, node{x, y, math.MaxInt})
			}
		}
	}

	h := nodeHeap(unvisited)
	var visited []node
	heap.Init(&h)
	for h.Len() > 0 {
		x := h.Pop()
		curr := x.(node)
		for i, n := range h {
			if n.x+1 == curr.x && n.y == curr.y {
				if n.dist > curr.dist+1 {
					h[i].dist = curr.dist + 1
				}
			}
			if n.x == curr.x && n.y+1 == curr.y {
				if n.dist > curr.dist+1 {
					h[i].dist = curr.dist + 1
				}
			}
			if n.x-1 == curr.x && n.y == curr.y {
				if n.dist > curr.dist+1 {
					h[i].dist = curr.dist + 1
				}
			}
			if n.x == curr.x && n.y-1 == curr.y {
				if n.dist > curr.dist+1 {
					h[i].dist = curr.dist + 1
				}
			}
		}
		visited = append(visited, curr)
		heap.Init(&h)
	}

	res := math.MaxInt
	for _, n := range visited {
		if n.x == size-1 && n.y == size-1 && res > n.dist {
			res = n.dist
		}
	}
	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	var bytes []pos
	for _, l := range lines {
		numsStr := strings.Split(l, ",")
		x, err := strconv.Atoi(numsStr[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(numsStr[1])
		if err != nil {
			panic(err)
		}
		bytes = append(bytes, pos{x, y})
	}

	size := 71

	for x := 1024; x < len(bytes); x++ {
		log.Println(x)
		currBytes := bytes[:x]

		var unvisited []node
		for y := range size {
			for x := range size {
				if x == 0 && y == 0 {
					unvisited = append(unvisited, node{x, y, 0})
				} else if !slices.ContainsFunc(currBytes, func(b pos) bool {
					return x == b.x && y == b.y
				}) {
					unvisited = append(unvisited, node{x, y, math.MaxInt})
				}
			}
		}

		h := nodeHeap(unvisited)
		var visited []node
		heap.Init(&h)
		isConnected := false
		for h.Len() > 0 {
			x := h.Pop()
			curr := x.(node)
			if curr.x == size-1 && curr.y == size-1 && curr.dist < math.MaxInt && curr.dist > 0 {
				isConnected = true
				break
			}
			for i, n := range h {
				if n.x+1 == curr.x && n.y == curr.y {
					if n.dist > curr.dist+1 {
						h[i].dist = curr.dist + 1
					}
				}
				if n.x == curr.x && n.y+1 == curr.y {
					if n.dist > curr.dist+1 {
						h[i].dist = curr.dist + 1
					}
				}
				if n.x-1 == curr.x && n.y == curr.y {
					if n.dist > curr.dist+1 {
						h[i].dist = curr.dist + 1
					}
				}
				if n.x == curr.x && n.y-1 == curr.y {
					if n.dist > curr.dist+1 {
						h[i].dist = curr.dist + 1
					}
				}
			}
			visited = append(visited, curr)
			heap.Init(&h)
		}
		if !isConnected {
			return x
		}
	}

	return 0
}
