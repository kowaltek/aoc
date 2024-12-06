package main

import (
	"log"
	"os"
	"slices"
	"strings"
)

type Point struct {
	visited, isObstacle bool
	visitedDir          []int
}

type Guard struct {
	dir  int
	x, y int
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
	h := len(lines)
	w := len(lines[0])
	var grid [][]Point
	var guard Guard
	for y, l := range lines {
		if len(l) == 0 {
			continue
		}
		grid = append(grid, make([]Point, 0))
		cells := strings.Split(l, "")
		for x, cell := range cells {
			isObstacle := false
			if cell == "#" {
				isObstacle = true
			}
			if slices.Contains([]string{"<", ">", "^", "v"}, cell) {
				guard.x = x
				guard.y = y
				switch cell {
				case "^":
					guard.dir = 0
				case ">":
					guard.dir = 1
				case "v":
					guard.dir = 2
				case "<":
					guard.dir = 3
				}
			}
			grid[y] = append(grid[y], Point{visited: false, isObstacle: isObstacle})
		}
	}

	res := 0
	for guard.x >= 0 && guard.x < w && guard.y >= 0 && guard.y < h {
		if !grid[guard.y][guard.x].visited {
			res++
		}
		grid[guard.y][guard.x].visited = true
		isBlocked := false
		switch guard.dir {
		case 0:
			if guard.y > 0 && grid[guard.y-1][guard.x].isObstacle {
				isBlocked = true
			}
		case 1:
			if guard.x < w-1 && grid[guard.y][guard.x+1].isObstacle {
				isBlocked = true
			}
		case 2:
			if guard.y < h-1 && grid[guard.y+1][guard.x].isObstacle {
				isBlocked = true
			}
		case 3:
			if guard.x > 0 && grid[guard.y][guard.x-1].isObstacle {
				isBlocked = true
			}
		}
		if isBlocked {
			guard.dir = (guard.dir + 1) % 4
		}
		switch guard.dir {
		case 0:
			guard.y--
		case 1:
			guard.x++
		case 2:
			guard.y++
		case 3:
			guard.x--
		}
	}
	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	h := len(lines)
	w := len(lines[0])
	var grid [][]Point
	var guard Guard
	for y, l := range lines {
		if len(l) == 0 {
			continue
		}
		grid = append(grid, make([]Point, 0))
		cells := strings.Split(l, "")
		for x, cell := range cells {
			isObstacle := false
			if cell == "#" {
				isObstacle = true
			}
			if slices.Contains([]string{"<", ">", "^", "v"}, cell) {
				guard.x = x
				guard.y = y
				switch cell {
				case "^":
					guard.dir = 0
				case ">":
					guard.dir = 1
				case "v":
					guard.dir = 2
				case "<":
					guard.dir = 3
				}
			}
			grid[y] = append(grid[y], Point{visited: false, isObstacle: isObstacle})
		}
	}

	res := 0
	guardStartX := guard.x
	guardStartY := guard.y
	for {
		grid[guard.y][guard.x].visited = true
		grid[guard.y][guard.x].visitedDir = append(grid[guard.y][guard.x].visitedDir, guard.dir)
		isBlocked := false
		switch guard.dir {
		case 0:
			if guard.y > 0 && grid[guard.y-1][guard.x].isObstacle {
				isBlocked = true
			}
		case 1:
			if guard.x < w-1 && grid[guard.y][guard.x+1].isObstacle {
				isBlocked = true
			}
		case 2:
			if guard.y < h-1 && grid[guard.y+1][guard.x].isObstacle {
				isBlocked = true
			}
		case 3:
			if guard.x > 0 && grid[guard.y][guard.x-1].isObstacle {
				isBlocked = true
			}
		}
		if isBlocked {
			guard.dir = (guard.dir + 1) % 4
		}
		loops := false
		newGuard := guard
		newGuard.dir = (guard.dir + 1) % 4
		var newGrid [][]Point
		for y, row := range grid {
			newGrid = append(newGrid, make([]Point, 0))
			for _, cell := range row {
				newGrid[y] = append(newGrid[y], cell)
			}
		}
		switch guard.dir {
		case 0:
			if guard.y-1 >= 0 && !newGrid[guard.y-1][guard.x].visited {
				newGrid[guard.y-1][guard.x].isObstacle = true
				loops = inLoop(newGuard, newGrid, h, w)
			}
		case 1:
			if guard.x+1 < w && !newGrid[guard.y][guard.x+1].visited {
				newGrid[guard.y][guard.x+1].isObstacle = true
				loops = inLoop(newGuard, newGrid, h, w)
			}
		case 2:
			if guard.y+1 < h && !newGrid[guard.y+1][guard.x].visited {
				newGrid[guard.y+1][guard.x].isObstacle = true
				loops = inLoop(newGuard, newGrid, h, w)
			}
		case 3:
			if guard.x-1 >= 0 && !newGrid[guard.y][guard.x-1].visited {
				newGrid[guard.y][guard.x-1].isObstacle = true
				loops = inLoop(newGuard, newGrid, h, w)
			}
		}
		shouldBreak := false
		switch guard.dir {
		case 0:
			if guard.y-1 < 0 {
				shouldBreak = true
			}
			if guard.y-1 >= 0 && !grid[guard.y-1][guard.x].isObstacle {
				guard.y--
			}
		case 1:
			if guard.x+1 >= w {
				shouldBreak = true
			}
			if guard.x+1 < w && !grid[guard.y][guard.x+1].isObstacle {
				guard.x++
			}
		case 2:
			if guard.y+1 >= h {
				shouldBreak = true
			}
			if guard.y+1 < w && !grid[guard.y+1][guard.x].isObstacle {
				guard.y++
			}
		case 3:
			if guard.x-1 < 0 {
				shouldBreak = true
			}
			if guard.x-1 >= 0 && !grid[guard.y][guard.x-1].isObstacle {
				guard.x--
			}
		}
		if shouldBreak {
			break
		}
		if loops && !(guard.x == guardStartX && guard.y == guardStartY) {
			res++
		}
	}
	return res
}

func inLoop(guard Guard, grid [][]Point, h, w int) bool {
	for {
		if grid[guard.y][guard.x].visited && slices.Contains(grid[guard.y][guard.x].visitedDir, guard.dir) {
			return true
		}
		grid[guard.y][guard.x].visited = true
		grid[guard.y][guard.x].visitedDir = append(grid[guard.y][guard.x].visitedDir, guard.dir)
		isBlocked := false
		switch guard.dir {
		case 0:
			if guard.y > 0 && grid[guard.y-1][guard.x].isObstacle {
				isBlocked = true
			}
		case 1:
			if guard.x < w-1 && grid[guard.y][guard.x+1].isObstacle {
				isBlocked = true
			}
		case 2:
			if guard.y < h-1 && grid[guard.y+1][guard.x].isObstacle {
				isBlocked = true
			}
		case 3:
			if guard.x > 0 && grid[guard.y][guard.x-1].isObstacle {
				isBlocked = true
			}
		}
		if isBlocked {
			guard.dir = (guard.dir + 1) % 4
		}
		shouldBreak := false
		switch guard.dir {
		case 0:
			if guard.y-1 < 0 {
				shouldBreak = true
			}
			if guard.y-1 >= 0 && !grid[guard.y-1][guard.x].isObstacle {
				guard.y--
			}
		case 1:
			if guard.x+1 >= w {
				shouldBreak = true
			}
			if guard.x+1 < w && !grid[guard.y][guard.x+1].isObstacle {
				guard.x++
			}
		case 2:
			if guard.y+1 >= h {
				shouldBreak = true
			}
			if guard.y+1 < w && !grid[guard.y+1][guard.x].isObstacle {
				guard.y++
			}
		case 3:
			if guard.x-1 < 0 {
				shouldBreak = true
			}
			if guard.x-1 >= 0 && !grid[guard.y][guard.x-1].isObstacle {
				guard.x--
			}
		}
		if shouldBreak {
			break
		}
	}
	return false
}
