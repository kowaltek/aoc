package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type pos struct {
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
	var robot pos
	var boxes []pos
	var walls []pos
	var mapEnd int
	for y, l := range lines {
		if len(l) == 0 {
			mapEnd = y
			break
		}
		for x, c := range strings.Split(l, "") {
			switch c {
			case "@":
				robot.x = x
				robot.y = y
			case "O":
				boxes = append(boxes, pos{x: x, y: y})
			case "#":
				walls = append(walls, pos{x: x, y: y})
			}
		}
	}
	var moves []int
	for i := mapEnd; i < len(lines); i++ {
		for _, c := range strings.Split(lines[i], "") {
			switch c {
			case ">":
				moves = append(moves, 0)
			case "v":
				moves = append(moves, 1)
			case "<":
				moves = append(moves, 2)
			case "^":
				moves = append(moves, 3)
			}
		}
	}
	for _, move := range moves {
		// log.Printf("move %d: %d\n", i, move)
		// printMap(len(lines[0]), mapEnd, robot, walls, boxes)
		switch move {
		case 0:
			if slices.Contains(walls, pos{x: robot.x + 1, y: robot.y}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x + 1, y: robot.y}); boxI != -1 {
				newPos := moveBoxIfPossible(pos{x: robot.x + 1, y: robot.y}, boxes, walls, move)
				boxes[boxI] = newPos
				if newPos.x > robot.x+1 {
					robot.x++
				}
			} else {
				robot.x++
			}
		case 1:
			if slices.Contains(walls, pos{x: robot.x, y: robot.y + 1}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x, y: robot.y + 1}); boxI != -1 {
				newPos := moveBoxIfPossible(pos{x: robot.x, y: robot.y + 1}, boxes, walls, move)
				boxes[boxI] = newPos
				if newPos.y > robot.y+1 {
					robot.y++
				}
			} else {
				robot.y++
			}
		case 2:
			if slices.Contains(walls, pos{x: robot.x - 1, y: robot.y}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x - 1, y: robot.y}); boxI != -1 {
				newPos := moveBoxIfPossible(pos{x: robot.x - 1, y: robot.y}, boxes, walls, move)
				boxes[boxI] = newPos
				if newPos.x < robot.x-1 {
					robot.x--
				}
			} else {
				robot.x--
			}
		case 3:
			if slices.Contains(walls, pos{x: robot.x, y: robot.y - 1}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x, y: robot.y - 1}); boxI != -1 {
				newPos := moveBoxIfPossible(pos{x: robot.x, y: robot.y - 1}, boxes, walls, move)
				boxes[boxI] = newPos
				if newPos.y < robot.y-1 {
					robot.y--
				}
			} else {
				robot.y--
			}
		}
	}
	// printMap(len(lines[0]), mapEnd, robot, walls, boxes)
	res := 0
	for _, box := range boxes {
		res += box.x + box.y*100
	}
	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	var robot pos
	var boxes []pos
	var walls []pos
	var mapEnd int
	for y, l := range lines {
		if len(l) == 0 {
			mapEnd = y
			break
		}
		for x, c := range strings.Split(l, "") {
			switch c {
			case "@":
				robot.x = x * 2
				robot.y = y
			case "O":
				boxes = append(boxes, pos{x: x * 2, y: y})
			case "#":
				walls = append(walls, pos{x: x * 2, y: y})
				walls = append(walls, pos{x: x*2 + 1, y: y})
			}
		}
	}
	var moves []int
	for i := mapEnd; i < len(lines); i++ {
		for _, c := range strings.Split(lines[i], "") {
			switch c {
			case ">":
				moves = append(moves, 0)
			case "v":
				moves = append(moves, 1)
			case "<":
				moves = append(moves, 2)
			case "^":
				moves = append(moves, 3)
			}
		}
	}
	for _, move := range moves {
		// log.Printf("move %d: %d\n", loop, move)
		// printMap2(len(lines[0])*2, mapEnd, robot, walls, boxes)
		switch move {
		case 0:
			if slices.Contains(walls, pos{x: robot.x + 1, y: robot.y}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x + 1, y: robot.y}); boxI != -1 {
				newPos, newBoxes := moveBoxIfPossible2(pos{x: robot.x + 1, y: robot.y}, boxes, walls, move)
				if newPos.x > robot.x+1 {
					boxes = newBoxes
					boxes[boxI] = newPos
					robot.x++
				}
			} else {
				robot.x++
			}
		case 1:
			if slices.Contains(walls, pos{x: robot.x, y: robot.y + 1}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x, y: robot.y + 1}); boxI != -1 {
				newPos, newBoxes := moveBoxIfPossible2(pos{x: robot.x, y: robot.y + 1}, boxes, walls, move)
				if newPos.y > robot.y+1 {
					boxes = newBoxes
					boxes[boxI] = newPos
					robot.y++
				}
			} else if boxI := slices.Index(boxes, pos{x: robot.x - 1, y: robot.y + 1}); boxI != -1 {
				newPos, newBoxes := moveBoxIfPossible2(pos{x: robot.x - 1, y: robot.y + 1}, boxes, walls, move)
				if newPos.y > robot.y+1 {
					boxes = newBoxes
					boxes[boxI] = newPos
					robot.y++
				}
			} else {
				robot.y++
			}
		case 2:
			if slices.Contains(walls, pos{x: robot.x - 1, y: robot.y}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x - 2, y: robot.y}); boxI != -1 {
				newPos, newBoxes := moveBoxIfPossible2(pos{x: robot.x - 2, y: robot.y}, boxes, walls, move)
				if newPos.x < robot.x-2 {
					boxes = newBoxes
					boxes[boxI] = newPos
					robot.x--
				}
			} else {
				robot.x--
			}
		case 3:
			if slices.Contains(walls, pos{x: robot.x, y: robot.y - 1}) {
				continue
			}
			if boxI := slices.Index(boxes, pos{x: robot.x, y: robot.y - 1}); boxI != -1 {
				newPos, newBoxes := moveBoxIfPossible2(pos{x: robot.x, y: robot.y - 1}, boxes, walls, move)
				if newPos.y < robot.y-1 {
					boxes = newBoxes
					boxes[boxI] = newPos
					robot.y--
				}
			} else if boxI := slices.Index(boxes, pos{x: robot.x - 1, y: robot.y - 1}); boxI != -1 {
				newPos, newBoxes := moveBoxIfPossible2(pos{x: robot.x - 1, y: robot.y - 1}, boxes, walls, move)
				if newPos.y < robot.y-1 {
					boxes = newBoxes
					boxes[boxI] = newPos
					robot.y--
				}
			} else {
				robot.y--
			}
		}
		// input := bufio.NewScanner(os.Stdin)
		// input.Scan()
	}
	printMap2(len(lines[0])*2, mapEnd, robot, walls, boxes)
	res := 0
	for _, box := range boxes {
		res += box.x + box.y*100
	}
	return res
}

func moveBoxIfPossible(box pos, boxes, walls []pos, move int) pos {
	switch move {
	case 0:
		if slices.Contains(walls, pos{x: box.x + 1, y: box.y}) {
			return box
		}
		if nextBoxI := slices.Index(boxes, pos{x: box.x + 1, y: box.y}); nextBoxI != -1 {
			newPos := moveBoxIfPossible(pos{x: box.x + 1, y: box.y}, boxes, walls, move)
			boxes[nextBoxI] = newPos
			if newPos.x > box.x+1 {
				return pos{x: box.x + 1, y: box.y}
			} else {
				return box
			}
		} else {
			return pos{x: box.x + 1, y: box.y}
		}
	case 1:
		if slices.Contains(walls, pos{x: box.x, y: box.y + 1}) {
			return box
		}
		if nextBoxI := slices.Index(boxes, pos{x: box.x, y: box.y + 1}); nextBoxI != -1 {
			newPos := moveBoxIfPossible(pos{x: box.x, y: box.y + 1}, boxes, walls, move)
			boxes[nextBoxI] = newPos
			if newPos.y > box.y+1 {
				return pos{x: box.x, y: box.y + 1}
			} else {
				return box
			}
		} else {
			return pos{x: box.x, y: box.y + 1}
		}
	case 2:
		if slices.Contains(walls, pos{x: box.x - 1, y: box.y}) {
			return box
		}
		if nextBoxI := slices.Index(boxes, pos{x: box.x - 1, y: box.y}); nextBoxI != -1 {
			newPos := moveBoxIfPossible(pos{x: box.x - 1, y: box.y}, boxes, walls, move)
			boxes[nextBoxI] = newPos
			if newPos.x < box.x-1 {
				return pos{x: box.x - 1, y: box.y}
			} else {
				return box
			}
		} else {
			return pos{x: box.x - 1, y: box.y}
		}
	case 3:
		if slices.Contains(walls, pos{x: box.x, y: box.y - 1}) {
			return box
		}
		if nextBoxI := slices.Index(boxes, pos{x: box.x, y: box.y - 1}); nextBoxI != -1 {
			newPos := moveBoxIfPossible(pos{x: box.x, y: box.y - 1}, boxes, walls, move)
			boxes[nextBoxI] = newPos
			if newPos.y < box.y-1 {
				return pos{x: box.x, y: box.y - 1}
			} else {
				return box
			}
		} else {
			return pos{x: box.x, y: box.y - 1}
		}
	}
	return pos{}
}

func moveBoxIfPossible2(box pos, origBoxes, walls []pos, move int) (pos, []pos) {
	boxes := make([]pos, len(origBoxes))
	copy(boxes, origBoxes)
	switch move {
	case 0:
		if slices.Contains(walls, pos{x: box.x + 2, y: box.y}) {
			return box, boxes
		}
		if nextBoxI := slices.Index(boxes, pos{x: box.x + 2, y: box.y}); nextBoxI != -1 {
			newPos, boxes := moveBoxIfPossible2(pos{x: box.x + 2, y: box.y}, boxes, walls, move)
			if newPos.x > box.x+2 {
				boxes[nextBoxI] = newPos
				return pos{x: box.x + 1, y: box.y}, boxes
			} else {
				return box, boxes
			}
		} else {
			return pos{x: box.x + 1, y: box.y}, boxes
		}
	case 1:
		if slices.Contains(walls, pos{x: box.x, y: box.y + 1}) || slices.Contains(walls, pos{x: box.x + 1, y: box.y + 1}) {
			return box, boxes
		}
		nextBoxI1 := slices.Index(boxes, pos{x: box.x, y: box.y + 1})
		nextBoxI2 := slices.Index(boxes, pos{x: box.x - 1, y: box.y + 1})
		nextBoxI3 := slices.Index(boxes, pos{x: box.x + 1, y: box.y + 1})
		canMoveBox1 := true
		canMoveBox2 := true
		canMoveBox3 := true
		var newPos1 pos
		var newPos2 pos
		var newPos3 pos
		if nextBoxI1 != -1 {
			newPos1, boxes = moveBoxIfPossible2(pos{x: box.x, y: box.y + 1}, boxes, walls, move)
			if !(newPos1.y > box.y+1) {
				canMoveBox1 = false
			}
		}
		if nextBoxI2 != -1 {
			newPos2, boxes = moveBoxIfPossible2(pos{x: box.x - 1, y: box.y + 1}, boxes, walls, move)
			if !(newPos2.y > box.y+1) {
				canMoveBox2 = false
			}
		}
		if nextBoxI3 != -1 {
			newPos3, boxes = moveBoxIfPossible2(pos{x: box.x + 1, y: box.y + 1}, boxes, walls, move)
			if !(newPos3.y > box.y+1) {
				canMoveBox3 = false
			}
		}
		if canMoveBox1 && canMoveBox2 && canMoveBox3 {
			if nextBoxI1 != -1 {
				boxes[nextBoxI1] = newPos1
			}
			if nextBoxI2 != -1 {
				boxes[nextBoxI2] = newPos2
			}
			if nextBoxI3 != -1 {
				boxes[nextBoxI3] = newPos3
			}
			return pos{x: box.x, y: box.y + 1}, boxes
		} else {
			return box, boxes
		}
	case 2:
		if slices.Contains(walls, pos{x: box.x - 1, y: box.y}) {
			return box, boxes
		}
		if nextBoxI := slices.Index(boxes, pos{x: box.x - 2, y: box.y}); nextBoxI != -1 {
			newPos, boxes := moveBoxIfPossible2(pos{x: box.x - 2, y: box.y}, boxes, walls, move)
			if newPos.x < box.x-2 {
				boxes[nextBoxI] = newPos
				return pos{x: box.x - 1, y: box.y}, boxes
			} else {
				return box, boxes
			}
		} else {
			return pos{x: box.x - 1, y: box.y}, boxes
		}
	case 3:
		if slices.Contains(walls, pos{x: box.x, y: box.y - 1}) || slices.Contains(walls, pos{x: box.x + 1, y: box.y - 1}) {
			return box, boxes
		}
		nextBoxI1 := slices.Index(boxes, pos{x: box.x, y: box.y - 1})
		nextBoxI2 := slices.Index(boxes, pos{x: box.x - 1, y: box.y - 1})
		nextBoxI3 := slices.Index(boxes, pos{x: box.x + 1, y: box.y - 1})
		canMoveBox1 := true
		canMoveBox2 := true
		canMoveBox3 := true
		var newPos1 pos
		var newPos2 pos
		var newPos3 pos
		if nextBoxI1 != -1 {
			newPos1, boxes = moveBoxIfPossible2(pos{x: box.x, y: box.y - 1}, boxes, walls, move)
			if !(newPos1.y < box.y-1) {
				canMoveBox1 = false
			}
		}
		if nextBoxI2 != -1 {
			newPos2, boxes = moveBoxIfPossible2(pos{x: box.x - 1, y: box.y - 1}, boxes, walls, move)
			if !(newPos2.y < box.y-1) {
				canMoveBox2 = false
			}
		}
		if nextBoxI3 != -1 {
			newPos3, boxes = moveBoxIfPossible2(pos{x: box.x + 1, y: box.y - 1}, boxes, walls, move)
			if !(newPos3.y < box.y-1) {
				canMoveBox3 = false
			}
		}
		if canMoveBox1 && canMoveBox2 && canMoveBox3 {
			if nextBoxI1 != -1 {
				boxes[nextBoxI1] = newPos1
			}
			if nextBoxI2 != -1 {
				boxes[nextBoxI2] = newPos2
			}
			if nextBoxI3 != -1 {
				boxes[nextBoxI3] = newPos3
			}
			return pos{x: box.x, y: box.y - 1}, boxes
		} else {
			return box, boxes
		}
	}
	return pos{}, []pos{}
}

func printMap(w, h int, robot pos, walls, boxes []pos) {
	for y := range h {
		for x := range w {
			if robot.x == x && robot.y == y {
				fmt.Print("@")
			} else if slices.Contains(boxes, pos{x: x, y: y}) {
				fmt.Print("O")
			} else if slices.Contains(walls, pos{x: x, y: y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
}

func printMap2(w, h int, robot pos, walls, boxes []pos) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if robot.x == x && robot.y == y {
				fmt.Print("@")
			} else if slices.Contains(boxes, pos{x: x, y: y}) {
				fmt.Print("[]")
				x++
			} else if slices.Contains(walls, pos{x: x, y: y}) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}
