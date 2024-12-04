package main

import (
	"log"
	"os"
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
	lines := strings.Split(file, "\n")

	res := 0
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[0 : len(lines)-1]
	}
	height := len(lines)
	for y, l := range lines {
		for x, c := range l {
			if c == 'X' {
				if x > 2 && l[x-3:x+1] == "SAMX" {
					res++
				}
				if x < len(l)-3 && l[x:x+4] == "XMAS" {
					res++
				}
				if y > 2 && lines[y-3][x] == 'S' && lines[y-2][x] == 'A' && lines[y-1][x] == 'M' {
					res++
				}
				if y < height-3 && lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
					res++
				}
				if x > 2 && y > 2 && lines[y-3][x-3] == 'S' && lines[y-2][x-2] == 'A' && lines[y-1][x-1] == 'M' {
					res++
				}
				if x > 2 && y < height-3 && lines[y+3][x-3] == 'S' && lines[y+2][x-2] == 'A' && lines[y+1][x-1] == 'M' {
					res++
				}
				if x < len(l)-3 && y > 2 && lines[y-3][x+3] == 'S' && lines[y-2][x+2] == 'A' && lines[y-1][x+1] == 'M' {
					res++
				}
				if x < len(l)-3 && y < height-3 && lines[y+3][x+3] == 'S' && lines[y+2][x+2] == 'A' && lines[y+1][x+1] == 'M' {
					res++
				}
			}
		}
	}
	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	res := 0
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[0 : len(lines)-1]
	}
	height := len(lines)
	for y, l := range lines {
		for x, c := range l {
			if !(x > 0 && x < len(l)-1 && y > 0 && y < height-1) {
				continue
			}
			if c == 'A' {
				masDiag1 := false
				masDiag2 := false

				if (lines[y-1][x-1] == 'M' && lines[y+1][x+1] == 'S') || (lines[y-1][x-1] == 'S' && lines[y+1][x+1] == 'M') {
					masDiag1 = true
				}
				if (lines[y-1][x+1] == 'M' && lines[y+1][x-1] == 'S') || (lines[y-1][x+1] == 'S' && lines[y+1][x-1] == 'M') {
					masDiag2 = true
				}
				if masDiag1 && masDiag2 {
					res++
				}
			}
		}
	}
	return res
}
