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
		c1, c2 := make(map[string]struct{}), make(map[string]struct{})
		for i, letter := range l {
			if i < len(l)/2 {
				c1[string(letter)] = struct{}{}
				continue
			}
			c2[string(letter)] = struct{}{}
		}
		common := ""
		for k := range c1 {
			if _, ok := c2[k]; ok {
				common = k
				break
			}
		}
		b := []byte(common)[0]
		if b > 96 {
			res += int(b - 96)
		} else {
			res += int(b - 64 + 26)
		}
	}

	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	res := 0
	for i := 0; i < len(lines)-1; i += 3 {
		c1, c2, c3 := make(map[string]struct{}), make(map[string]struct{}), make(map[string]struct{})
		for _, letter := range lines[i] {
			c1[string(letter)] = struct{}{}
		}
		for _, letter := range lines[i+1] {
			c2[string(letter)] = struct{}{}
		}
		for _, letter := range lines[i+2] {
			c3[string(letter)] = struct{}{}
		}
		common := ""
		for k := range c1 {
			if _, ok := c2[k]; ok {
				if _, ok := c3[k]; ok {
					common = k
					break
				}
			}
		}
		b := []byte(common)[0]
		if b > 96 {
			res += int(b - 96)
		} else {
			res += int(b - 64 + 26)
		}
	}

	return res
}
