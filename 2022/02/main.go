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
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		p1 := l[0] - 65
		p2 := l[2] - 23 - 65
		res += int(p2) + 1
		switch (p2 - p1 + 3) % 3 {
		case 0:
			res += 3
		case 1:
			res += 6
		}
	}

	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	/*
		0 0 draw 3 % 3 = 0
		0 1 win 4 % 3 = 1
		0 2 lose 5 % 3 = 2

		1 0 lose 2 % 3
		1 1 draw 0
		1 2 win 1

		2 0 win 1
		2 1 lose 2
		2 2 draw 0
	*/

	res := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		p1 := l[0] - 65
		r := l[2] - 23 - 65
		p2 := (p1 + r + 2) % 3
		res += int(p2) + 1
		switch r {
		case 1:
			res += 3
		case 2:
			res += 6
		}
	}

	return res
}
