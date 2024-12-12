package main

import (
	"log"
	"os"
	"slices"
	"strings"
)

type side struct {
	y1, x1, y2, x2, dir int
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
	var m [][]string
	h := len(lines)
	w := len(lines[0])
	visited := make([]bool, w*h)
	for y, l := range lines {
		m = append(m, make([]string, 0))
		for _, c := range strings.Split(l, "") {
			m[y] = append(m[y], c)
		}
	}

	res := 0
	for i := 0; i > -1; i = slices.Index(visited, false) {
		area, perimeter := processArea(i/w, i%w, h, w, m, visited)
		res += area * perimeter
	}

	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	var m [][]string
	h := len(lines)
	w := len(lines[0])
	visited := make([]bool, w*h)
	for y, l := range lines {
		m = append(m, make([]string, 0))
		for _, c := range strings.Split(l, "") {
			m[y] = append(m[y], c)
		}
	}

	res := 0
	for i := 0; i > -1; i = slices.Index(visited, false) {
		var sides []side
		area, sides := processArea2(i/w, i%w, h, w, m, visited, sides)
		sides = mergeSides(sides)
		res += area * len(sides)
	}

	return res
}

func processArea(y, x, h, w int, m [][]string, visited []bool) (int, int) {
	if visited[y*w+x] {
		return 0, 0
	}
	curr := m[y][x]
	area := 1
	perimeter := 0
	visited[y*w+x] = true
	if y-1 >= 0 && m[y-1][x] != curr {
		perimeter++
	} else if y-1 >= 0 {
		area2, perimeter2 := processArea(y-1, x, h, w, m, visited)
		area += area2
		perimeter += perimeter2
	} else {
		perimeter++
	}

	if x-1 >= 0 && m[y][x-1] != curr {
		perimeter++
	} else if x-1 >= 0 {
		area2, perimeter2 := processArea(y, x-1, h, w, m, visited)
		area += area2
		perimeter += perimeter2
	} else {
		perimeter++
	}

	if y+1 < h && m[y+1][x] != curr {
		perimeter++
	} else if y+1 < h {
		area2, perimeter2 := processArea(y+1, x, h, w, m, visited)
		area += area2
		perimeter += perimeter2
	} else {
		perimeter++
	}

	if x+1 < w && m[y][x+1] != curr {
		perimeter++
	} else if x+1 < w {
		area2, perimeter2 := processArea(y, x+1, h, w, m, visited)
		area += area2
		perimeter += perimeter2
	} else {
		perimeter++
	}

	return area, perimeter
}

func processArea2(y, x, h, w int, m [][]string, visited []bool, sides []side) (int, []side) {
	if visited[y*w+x] {
		return 0, sides
	}
	curr := m[y][x]
	area := 1
	visited[y*w+x] = true

	if x-1 >= 0 && m[y][x-1] == curr {
		area2, sides2 := processArea2(y, x-1, h, w, m, visited, sides)
		area += area2
		sides = sides2
	} else {
		sides = append(sides, side{y1: y, x1: x, y2: y + 1, x2: x, dir: 0})
	}

	if y-1 >= 0 && m[y-1][x] == curr {
		area2, sides2 := processArea2(y-1, x, h, w, m, visited, sides)
		area += area2
		sides = sides2
	} else {
		sides = append(sides, side{y1: y, x1: x, y2: y, x2: x + 1, dir: 1})
	}

	if x+1 < w && m[y][x+1] == curr {
		area2, sides2 := processArea2(y, x+1, h, w, m, visited, sides)
		area += area2
		sides = sides2
	} else {
		sides = append(sides, side{y1: y, x1: x + 1, y2: y + 1, x2: x + 1, dir: 2})
	}

	if y+1 < h && m[y+1][x] == curr {
		area2, sides2 := processArea2(y+1, x, h, w, m, visited, sides)
		area += area2
		sides = sides2
	} else {
		sides = append(sides, side{y1: y + 1, x1: x, y2: y + 1, x2: x + 1, dir: 3})
	}

	return area, sides
}

func mergeSides(sides []side) []side {
	for {
		prev := len(sides)
		for i := 0; i < len(sides); i++ {
			s1 := sides[i]
			for j := i + 1; j < len(sides); j++ {
				s2 := sides[j]
				if s1.x1 == s1.x2 && s2.x1 == s2.x2 && s1.x1 == s2.x1 && s1.dir == s2.dir {
					if !((s1.y1 < s2.y1 && s1.y2 < s2.y1) || (s1.y1 > s2.y2 && s1.y2 > s2.y2)) {
						if s1.y1 > s2.y1 {
							s1.y1 = s2.y1
						}
						if s1.y2 < s2.y2 {
							s1.y2 = s2.y2
						}
						sides[i] = s1
						sidesCpy := make([]side, len(sides))
						copy(sidesCpy, sides)
						sides = sides[:j]
						sides = append(sides, sidesCpy[j+1:]...)
					}
				} else if s1.y1 == s1.y2 && s2.y1 == s2.y2 && s1.y1 == s2.y1 && s1.dir == s2.dir {
					if !((s1.x1 < s2.x1 && s1.x2 < s2.x1) || (s1.x1 > s2.x2 && s1.x2 > s2.x2)) {
						if s1.x1 > s2.x1 {
							s1.x1 = s2.x1
						}
						if s1.x2 < s2.x2 {
							s1.x2 = s2.x2
						}
						sides[i] = s1
						sidesCpy := make([]side, len(sides))
						copy(sidesCpy, sides)
						sides = sides[:j]
						sides = append(sides, sidesCpy[j+1:]...)
					}
				}
			}
		}
		if prev == len(sides) {
			break
		}
	}
	return sides
}
