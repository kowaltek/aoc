package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type cheat struct {
	start pos
	end   pos
	saves int
}

type pos struct {
	x, y int
}

type node struct {
	p    pos
	next *node
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
	var unvisited []pos
	var end pos
	var curr pos
	var cheats []cheat
	for y, l := range lines {
		for x, c := range strings.Split(l, "") {
			switch c {
			case "S":
				curr = pos{x, y}
				unvisited = append(unvisited, pos{x, y})
			case "E":
				end = pos{x, y}
				unvisited = append(unvisited, pos{x, y})
			case ".":
				unvisited = append(unvisited, pos{x, y})
			}
		}
	}
	for !(curr.x == end.x && curr.y == end.y) {
		if slices.Contains(unvisited, pos{curr.x + 2, curr.y}) {
			cheats = append(cheats, cheat{curr, pos{curr.x + 2, curr.y}, countCheat(unvisited, curr, pos{curr.x + 2, curr.y})})
		}
		if slices.Contains(unvisited, pos{curr.x, curr.y + 2}) {
			cheats = append(cheats, cheat{curr, pos{curr.x, curr.y + 2}, countCheat(unvisited, curr, pos{curr.x, curr.y + 2})})
		}
		if slices.Contains(unvisited, pos{curr.x - 2, curr.y}) {
			cheats = append(cheats, cheat{curr, pos{curr.x - 2, curr.y}, countCheat(unvisited, curr, pos{curr.x - 2, curr.y})})
		}
		if slices.Contains(unvisited, pos{curr.x, curr.y - 2}) {
			cheats = append(cheats, cheat{curr, pos{curr.x, curr.y - 2}, countCheat(unvisited, curr, pos{curr.x, curr.y - 2})})
		}
		toRemove := curr
		if slices.Contains(unvisited, pos{curr.x + 1, curr.y}) {
			curr.x++
		} else if slices.Contains(unvisited, pos{curr.x, curr.y + 1}) {
			curr.y++
		} else if slices.Contains(unvisited, pos{curr.x - 1, curr.y}) {
			curr.x--
		} else if slices.Contains(unvisited, pos{curr.x, curr.y - 1}) {
			curr.y--
		}
		unvisited = slices.DeleteFunc(unvisited, func(next pos) bool {
			return next.x == toRemove.x && next.y == toRemove.y
		})
	}
	res := 0
	for _, c := range cheats {
		if c.saves >= 100 {
			res++
		}
	}
	return res
}

func countCheat(unvisited []pos, start, end pos) int {
	timeTaken := int(math.Abs(float64(start.x-end.x))) + int(math.Abs(float64(start.y-end.y))) + 1
	var saved []pos
	saved = append(saved, start)
	for !(start.x == end.x && start.y == end.y) {
		if slices.Contains(unvisited, pos{start.x + 1, start.y}) && !slices.Contains(saved, pos{start.x + 1, start.y}) {
			start.x++
			saved = append(saved, start)
		} else if slices.Contains(unvisited, pos{start.x, start.y + 1}) && !slices.Contains(saved, pos{start.x, start.y + 1}) {
			start.y++
			saved = append(saved, start)
		} else if slices.Contains(unvisited, pos{start.x - 1, start.y}) && !slices.Contains(saved, pos{start.x - 1, start.y}) {
			start.x--
			saved = append(saved, start)
		} else if slices.Contains(unvisited, pos{start.x, start.y - 1}) && !slices.Contains(saved, pos{start.x, start.y - 1}) {
			start.y--
			saved = append(saved, start)
		}
	}
	return len(saved) - timeTaken
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	var unvisited []pos
	var end pos
	var curr pos
	var cheats []cheat
	for y, l := range lines {
		for x, c := range strings.Split(l, "") {
			switch c {
			case "S":
				curr = pos{x, y}
				unvisited = append(unvisited, pos{x, y})
			case "E":
				end = pos{x, y}
				unvisited = append(unvisited, pos{x, y})
			case ".":
				unvisited = append(unvisited, pos{x, y})
			}
		}
	}
	h := len(lines)
	w := len(lines[0])
	start := node{
		p: curr,
	}
	prev := &start
	forwardPos := make([][]bool, w)
	for x := range w {
		forwardPos[x] = make([]bool, h)
	}
	forwardPos[start.p.x][start.p.y] = true
	forwardPos[end.x][end.y] = true
	for !(curr.x == end.x && curr.y == end.y) {
		toRemove := curr
		if slices.Contains(unvisited, pos{curr.x + 1, curr.y}) {
			curr.x++
		} else if slices.Contains(unvisited, pos{curr.x, curr.y + 1}) {
			curr.y++
		} else if slices.Contains(unvisited, pos{curr.x - 1, curr.y}) {
			curr.x--
		} else if slices.Contains(unvisited, pos{curr.x, curr.y - 1}) {
			curr.y--
		}
		forwardPos[curr.x][curr.y] = true
		next := node{
			p: curr,
		}
		prev.next = &next
		prev = &next
		unvisited = slices.DeleteFunc(unvisited, func(n pos) bool {
			return n.x == toRemove.x && n.y == toRemove.y
		})
	}
	next := node{
		p: curr,
	}
	prev.next = &next
	prev = &next
	currN := start
	curr = currN.p

	for y := range len(forwardPos[0]) {
		for x := range len(forwardPos) {
			if forwardPos[x][y] {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}

	for !(curr.x == end.x && curr.y == end.y) {
		fmt.Println(curr)
		for dx := -20; dx <= 20; dx++ {
			for dy := -20; dy <= 20; dy++ {
				delta := int(math.Abs(float64(dx))) + int(math.Abs(float64(dy)))
				if delta <= 20 && curr.x+dx >= 0 && curr.x+dx < w && curr.y+dy >= 0 && curr.y+dy < h && forwardPos[curr.x+dx][curr.y+dy] {
					cheats = append(cheats, cheat{curr, pos{curr.x + dx, curr.y + dy}, countCheat2(forwardPos, curr, pos{curr.x + dx, curr.y + dy})})
				}
			}
		}
		forwardPos[curr.x][curr.y] = false
		if forwardPos[curr.x+1][curr.y] {
			curr.x++
		} else if forwardPos[curr.x][curr.y+1] {
			curr.y++
		} else if forwardPos[curr.x-1][curr.y] {
			curr.x--
		} else if forwardPos[curr.x][curr.y-1] {
			curr.y--
		}
	}
	res := 0
	for _, c := range cheats {
		if c.saves >= 100 {
			res++
		}
	}
	return res
}

func countCheat2(forwardPos [][]bool, start, end pos) int {
	timeTaken := int(math.Abs(float64(start.x-end.x))) + int(math.Abs(float64(start.y-end.y))) + 1
	res := 1
	visited := make([][]bool, len(forwardPos))
	for x := range len(visited) {
		visited[x] = make([]bool, len(forwardPos[0]))
	}
	visited[start.x][start.y] = true
	for !(start.x == end.x && start.y == end.y) {
		if forwardPos[start.x+1][start.y] && !visited[start.x+1][start.y] {
			start.x++
			res++
		} else if forwardPos[start.x][start.y+1] && !visited[start.x][start.y+1] {
			start.y++
			res++
		} else if forwardPos[start.x-1][start.y] && !visited[start.x-1][start.y] {
			start.x--
			res++
		} else if forwardPos[start.x][start.y-1] && !visited[start.x][start.y-1] {
			start.y--
			res++
		}
		visited[start.x][start.y] = true
	}
	return res - timeTaken
}

func hasNode(head node, p pos) bool {
	for head.next != nil {
		if head.p.x == p.x && head.p.y == p.y {
			return true
		}
		head = *head.next
	}
	return false
}

// 285

// 5530 too low
// 5291
