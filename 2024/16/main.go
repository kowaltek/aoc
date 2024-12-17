package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	x, y, dist, dir int
	parents         []node
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
	var unvisited []node
	var end pos
	for y, l := range lines {
		for x, c := range strings.Split(l, "") {
			switch c {
			case "S":
				unvisited = append(unvisited, node{x, y, 0, 0, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 1, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 2, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 3, nil})
			case "E":
				end = pos{x, y}
				unvisited = append(unvisited, node{x, y, math.MaxInt, 0, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 1, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 2, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 3, nil})
			case ".":
				unvisited = append(unvisited, node{x, y, math.MaxInt, 0, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 1, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 2, nil})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 3, nil})
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
			if n.x == curr.x && n.y == curr.y && (n.dir == (curr.dir-1+4)%4 || n.dir == (curr.dir+1)%4) {
				if n.dist > curr.dist+1000 {
					h[i].dist = curr.dist + 1000
				}
			} else {
				switch curr.dir {
				case 0:
					if n.x == curr.x+1 && n.y == curr.y && n.dir == curr.dir {
						if n.dist > curr.dist+1 {
							h[i].dist = curr.dist + 1
						}
					}
				case 1:
					if n.x == curr.x && n.y == curr.y+1 && n.dir == curr.dir {
						if n.dist > curr.dist+1 {
							h[i].dist = curr.dist + 1
						}
					}
				case 2:
					if n.x == curr.x-1 && n.y == curr.y && n.dir == curr.dir {
						if n.dist > curr.dist+1 {
							h[i].dist = curr.dist + 1
						}
					}
				case 3:
					if n.x == curr.x && n.y == curr.y-1 && n.dir == curr.dir {
						if n.dist > curr.dist+1 {
							h[i].dist = curr.dist + 1
						}
					}
				}
			}
		}
		visited = append(visited, curr)
		heap.Init(&h)
	}
	for _, n := range visited {
		if n.x == end.x && n.y == end.y {
			return n.dist
		}
	}
	return 0
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	var unvisited []node
	var end pos
	for y, l := range lines {
		for x, c := range strings.Split(l, "") {
			switch c {
			case "S":
				unvisited = append(unvisited, node{x, y, 0, 0, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 1, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 2, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 3, make([]node, 0)})
			case "E":
				end = pos{x, y}
				unvisited = append(unvisited, node{x, y, math.MaxInt, 0, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 1, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 2, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 3, make([]node, 0)})
			case ".":
				unvisited = append(unvisited, node{x, y, math.MaxInt, 0, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 1, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 2, make([]node, 0)})
				unvisited = append(unvisited, node{x, y, math.MaxInt, 3, make([]node, 0)})
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
			if n.x == curr.x && n.y == curr.y && (n.dir == (curr.dir-1+4)%4 || n.dir == (curr.dir+1)%4) {
				if curr.dist+1000 < n.dist {
					h[i].dist = curr.dist + 1000
					h[i].parents = []node{curr}
				} else if curr.dist+1000 == n.dist {
					h[i].parents = append(h[i].parents, curr)
				}
			} else {
				switch curr.dir {
				case 0:
					if n.x == curr.x+1 && n.y == curr.y && n.dir == curr.dir {
						if curr.dist+1 < n.dist {
							h[i].dist = curr.dist + 1
							h[i].parents = []node{curr}
						} else if curr.dist+1 == n.dist {
							h[i].parents = append(h[i].parents, curr)
						}
					}
				case 1:
					if n.x == curr.x && n.y == curr.y+1 && n.dir == curr.dir {
						if curr.dist+1 < n.dist {
							h[i].dist = curr.dist + 1
							h[i].parents = []node{curr}
						} else if curr.dist+1 == n.dist {
							h[i].parents = append(h[i].parents, curr)
						}
					}
				case 2:
					if n.x == curr.x-1 && n.y == curr.y && n.dir == curr.dir {
						if curr.dist+1 < n.dist {
							h[i].dist = curr.dist + 1
							h[i].parents = []node{curr}
						} else if curr.dist+1 == n.dist {
							h[i].parents = append(h[i].parents, curr)
						}
					}
				case 3:
					if n.x == curr.x && n.y == curr.y-1 && n.dir == curr.dir {
						if curr.dist+1 < n.dist {
							h[i].dist = curr.dist + 1
							h[i].parents = []node{curr}
						} else if curr.dist+1 == n.dist {
							h[i].parents = append(h[i].parents, curr)
						}
					}
				}
			}
		}
		visited = append(visited, curr)
		heap.Init(&h)
	}
	var bestPath []pos
	bestPath = append(bestPath, end)
	var endParents []node
	for _, n := range visited {
		if n.x == end.x && n.y == end.y {
			if len(endParents) == 0 {
				endParents = append(endParents, n.parents...)
			} else {
				if endParents[0].dist > n.dist {
					endParents[0] = n
				}
			}
		}
	}
	for len(endParents) != 0 {
		var nextParents []node
		for _, n := range endParents {
			nextParents = append(nextParents, n.parents...)
			dup := false
			for _, p := range bestPath {
				if n.x == p.x && n.y == p.y {
					dup = true
				}
			}
			if !dup {
				bestPath = append(bestPath, pos{n.x, n.y})
			}
		}
		endParents = nextParents
	}
	for y := range 141 {
		for x := range 141 {
			if slices.Contains(bestPath, pos{x, y}) {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
				// fmt.Printf("%c", lines[y][x])
			}
		}
		fmt.Print("\n")
	}
	return len(bestPath)
}
