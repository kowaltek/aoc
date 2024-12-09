package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Pos struct {
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
	antennas := make(map[rune][]Pos)
	h := len(lines)
	w := len(lines[0])
	for y, l := range lines {
		for x, c := range l {
			if c != '.' && c != '\n' {
				antennas[c] = append(antennas[c], Pos{x: x, y: y})
			}
		}
	}
	var antinodes []Pos
	for _, a := range antennas {
		for i := 0; i < len(a)-1; i++ {
			for j := i + 1; j < len(a); j++ {
				pos1 := Pos{
					x: a[j].x + a[j].x - a[i].x,
					y: a[j].y + a[j].y - a[i].y,
				}
				pos2 := Pos{
					x: a[i].x + a[i].x - a[j].x,
					y: a[i].y + a[i].y - a[j].y,
				}
				if pos1.x >= 0 && pos1.x < w && pos1.y >= 0 && pos1.y < h && !slices.Contains(antinodes, pos1) {
					antinodes = append(antinodes, pos1)
				}
				if pos2.x >= 0 && pos2.x < w && pos2.y >= 0 && pos2.y < h && !slices.Contains(antinodes, pos2) {
					antinodes = append(antinodes, pos2)
				}
			}
		}
	}
	res := len(antinodes)
	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	antennas := make(map[rune][]Pos)
	h := len(lines)
	w := len(lines[0])
	for y, l := range lines {
		for x, c := range l {
			if c != '.' && c != '\n' {
				antennas[c] = append(antennas[c], Pos{x: x, y: y})
			}
		}
	}
	var antinodes []Pos
	for _, ant := range antennas {
		for i := 0; i < len(ant)-1; i++ {
			for j := i + 1; j < len(ant); j++ {
				diffY := ant[j].y - ant[i].y
				diffX := ant[j].x - ant[i].x
				x := ant[i].x
				y := ant[i].y
				for y >= 0 && x >= 0 && y < h && x < w {
					pos := Pos{
						x: x,
						y: y,
					}
					if !slices.Contains(antinodes, pos) {
						antinodes = append(antinodes, pos)
					}
					x += diffX
					y += diffY
				}
				x = ant[i].x
				y = ant[i].y
				for y >= 0 && x >= 0 && y < h && x < w {
					pos := Pos{
						x: x,
						y: y,
					}
					if !slices.Contains(antinodes, pos) {
						antinodes = append(antinodes, pos)
					}
					x -= diffX
					y -= diffY
				}
			}
		}
	}
	for y, l := range lines {
		for x, c := range l {
			if slices.Contains(antinodes, Pos{x: x, y: y}) && c != 'A' && c != '0' {
				fmt.Printf("#")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
	res := len(antinodes)
	return res
}
