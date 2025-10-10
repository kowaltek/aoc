package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	TotalSpace    = 70000000
	RequiredSpace = 30000000
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
	lines := strings.SplitSeq(file, "\n")

	var path []string
	pathSizes := make(map[string]int)
	for l := range lines {
		if len(l) == 0 {
			continue
		}

		parts := strings.Split(l, " ")
		if parts[0] == "$" {
			switch parts[1] {
			case "cd":
				if parts[2] == ".." {
					path = path[:len(path)-1]
				} else {
					path = append(path, parts[2])
				}
			}
			continue
		}

		if parts[0] == "dir" {
			continue
		}

		size, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		for i := range len(path) {
			key := strings.Join(path[:i+1], "::")
			pathSizes[key] += size
		}
	}

	res := 0
	for _, size := range pathSizes {
		if size < 100000 {
			res += size
		}
	}
	return res
}

func solve2(file string) int {
	lines := strings.SplitSeq(file, "\n")

	var path []string
	pathSizes := make(map[string]int)
	for l := range lines {
		if len(l) == 0 {
			continue
		}

		parts := strings.Split(l, " ")
		if parts[0] == "$" {
			switch parts[1] {
			case "cd":
				if parts[2] == ".." {
					path = path[:len(path)-1]
				} else {
					path = append(path, parts[2])
				}
			}
			continue
		}

		if parts[0] == "dir" {
			continue
		}

		size, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		for i := range len(path) {
			key := strings.Join(path[:i+1], "::")
			pathSizes[key] += size
		}
	}

	requiredSpace := RequiredSpace - (TotalSpace - pathSizes["/"])

	res := TotalSpace
	for _, size := range pathSizes {
		if size > requiredSpace && size < res {
			res = size
		}
	}
	return res
}
