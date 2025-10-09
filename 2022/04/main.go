package main

import (
	"log"
	"os"
	"strconv"
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
		ps := strings.Split(l, ",")
		e1Str := strings.Split(ps[0], "-")
		e2Str := strings.Split(ps[1], "-")
		var e1 [2]int
		var e2 [2]int
		var err error
		e1[0], err = strconv.Atoi(e1Str[0])
		if err != nil {
			panic(err)
		}
		e1[1], err = strconv.Atoi(e1Str[1])
		if err != nil {
			panic(err)
		}
		e2[0], err = strconv.Atoi(e2Str[0])
		if err != nil {
			panic(err)
		}
		e2[1], err = strconv.Atoi(e2Str[1])
		if err != nil {
			panic(err)
		}
		if (e1[0] <= e2[0] && e1[1] >= e2[1]) || (e1[0] >= e2[0] && e1[1] <= e2[1]) {
			res++
		}
	}

	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	res := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		ps := strings.Split(l, ",")
		e1Str := strings.Split(ps[0], "-")
		e2Str := strings.Split(ps[1], "-")
		var e1 [2]int
		var e2 [2]int
		var err error
		e1[0], err = strconv.Atoi(e1Str[0])
		if err != nil {
			panic(err)
		}
		e1[1], err = strconv.Atoi(e1Str[1])
		if err != nil {
			panic(err)
		}
		e2[0], err = strconv.Atoi(e2Str[0])
		if err != nil {
			panic(err)
		}
		e2[1], err = strconv.Atoi(e2Str[1])
		if err != nil {
			panic(err)
		}
		if (e1[0] <= e2[0] && e1[1] >= e2[0]) || (e1[0] <= e2[1] && e1[1] >= e2[1]) || (e1[0] >= e2[0] && e1[1] <= e2[1]) {
			res++
		}
	}

	return res
}
