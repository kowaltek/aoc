package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reNum = regexp.MustCompile(`\d+`)

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
	res := 0
	for i := 0; i < len(lines); i++ {
		first := reNum.FindAllString(lines[i], -1)
		second := reNum.FindAllString(lines[i+1], -1)
		third := reNum.FindAllString(lines[i+2], -1)
		ax, err := strconv.Atoi(first[0])
		if err != nil {
			panic(err)
		}
		ay, err := strconv.Atoi(first[1])
		if err != nil {
			panic(err)
		}
		bx, err := strconv.Atoi(second[0])
		if err != nil {
			panic(err)
		}
		by, err := strconv.Atoi(second[1])
		if err != nil {
			panic(err)
		}
		px, err := strconv.Atoi(third[0])
		if err != nil {
			panic(err)
		}
		py, err := strconv.Atoi(third[1])
		if err != nil {
			panic(err)
		}
		x, y := 0, 0
		as, bs := -1, 0
		for try := 0; try < 100; try++ {
			x, y = x+bx, y+by
			bs++
			if (px-x)%ax == 0 && (py-y)%ay == 0 && (px-x)/ax == (py-y)/ay {
				as = (px - x) / ax
				break
			}
		}
		if as != -1 {
			res += as*3 + bs
		}
		i += 3
	}
	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	res := 0
	for i := 0; i < len(lines); i++ {
		first := reNum.FindAllString(lines[i], -1)
		second := reNum.FindAllString(lines[i+1], -1)
		third := reNum.FindAllString(lines[i+2], -1)
		ax, err := strconv.Atoi(first[0])
		if err != nil {
			panic(err)
		}
		ay, err := strconv.Atoi(first[1])
		if err != nil {
			panic(err)
		}
		bx, err := strconv.Atoi(second[0])
		if err != nil {
			panic(err)
		}
		by, err := strconv.Atoi(second[1])
		if err != nil {
			panic(err)
		}
		px, err := strconv.Atoi(third[0])
		if err != nil {
			panic(err)
		}
		py, err := strconv.Atoi(third[1])
		if err != nil {
			panic(err)
		}
		px += 10000000000000
		py += 10000000000000
		bs := (px*ay - py*ax) / (ay*bx - ax*by)
		as := (py - bs*by) / ay
		if bs >= 0 && as >= 0 && px == bs*bx+as*ax && py == bs*by+as*ay {
			part := as*3 + bs
			res += part
		}
		i += 3
	}
	return res
}
