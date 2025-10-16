package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type knot struct {
	x, y int
	next *knot
}

func (p knot) key() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func (p knot) moveChildren() {
	if p.next == nil {
		return
	}

	if p.x == p.next.x && math.Abs(float64(p.y-p.next.y)) > 1 {
		if p.y > p.next.y {
			p.next.y++
		} else {
			p.next.y--
		}
	} else if p.y == p.next.y && math.Abs(float64(p.x-p.next.x)) > 1 {
		if p.x > p.next.x {
			p.next.x++
		} else {
			p.next.x--
		}
	} else if math.Abs(float64(p.y-p.next.y)) > 1 || math.Abs(float64(p.x-p.next.x)) > 1 {
		if p.y > p.next.y {
			p.next.y++
		} else {
			p.next.y--
		}
		if p.x > p.next.x {
			p.next.x++
		} else {
			p.next.x--
		}
	}

	p.next.moveChildren()
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
	lines := strings.SplitSeq(file, "\n")

	visited := make(map[string]struct{})
	h, t := knot{}, knot{}
	h.next = &t
	for l := range lines {
		if len(l) == 0 {
			continue
		}

		parts := strings.Split(l, " ")
		dir := parts[0]
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for range n {
			switch dir {
			case "U":
				h.y--
			case "D":
				h.y++
			case "L":
				h.x--
			case "R":
				h.x++
			}

			h.moveChildren()

			visited[t.key()] = struct{}{}
		}

	}

	return len(visited)
}

func solve2(file string) int {
	lines := strings.SplitSeq(file, "\n")

	visited := make(map[string]struct{})
	h := knot{}
	curr := &h
	for range 9 {
		k := &knot{}
		curr.next = k
		curr = k
	}
	t := curr

	for l := range lines {
		if len(l) == 0 {
			continue
		}

		parts := strings.Split(l, " ")
		dir := parts[0]
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for range n {
			switch dir {
			case "U":
				h.y--
			case "D":
				h.y++
			case "L":
				h.x--
			case "R":
				h.x++
			}

			h.moveChildren()

			visited[(*t).key()] = struct{}{}
		}

	}

	return len(visited)
}
