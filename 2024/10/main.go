package main

import (
	"log"
	"os"
	"slices"
	"strconv"
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
	rows := strings.Split(file, "\n")

	var m [][]int
	var heads []Pos
	for y, row := range rows {
		m = append(m, make([]int, 0))
		for x, digit := range row {
			n, err := strconv.Atoi(string(digit))
			if err != nil {
				panic(err)
			}
			if n == 0 {
				heads = append(heads, Pos{x: x, y: y})
			}
			m[y] = append(m[y], n)
		}
	}
	res := 0
	for _, head := range heads {
		new := findTrails(head.x, head.y, m)
		var unique []Pos
		for _, p := range new {
			if !slices.Contains(unique, p) {
				unique = append(unique, p)
			}
		}
		res += len(unique)
	}
	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	rows := strings.Split(file, "\n")

	var m [][]int
	var heads []Pos
	for y, row := range rows {
		m = append(m, make([]int, 0))
		for x, digit := range row {
			n, err := strconv.Atoi(string(digit))
			if err != nil {
				panic(err)
			}
			if n == 0 {
				heads = append(heads, Pos{x: x, y: y})
			}
			m[y] = append(m[y], n)
		}
	}
	res := 0
	for _, head := range heads {
		new := findTrails(head.x, head.y, m)
		res += len(new)
	}
	return res
}

func findTrails(x, y int, m [][]int) []Pos {
	val := m[y][x]
	if val == 9 {
		return []Pos{{y: y, x: x}}
	}
	var res []Pos
	if x-1 >= 0 && m[y][x-1] == val+1 {
		res = append(res, findTrails(x-1, y, m)...)
	}
	if x+1 < len(m[y]) && m[y][x+1] == val+1 {
		res = append(res, findTrails(x+1, y, m)...)
	}
	if y-1 >= 0 && m[y-1][x] == val+1 {
		res = append(res, findTrails(x, y-1, m)...)
	}
	if y+1 < len(m) && m[y+1][x] == val+1 {
		res = append(res, findTrails(x, y+1, m)...)
	}
	return res
}
