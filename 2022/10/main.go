package main

import (
	"fmt"
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
	lines := strings.SplitSeq(file, "\n")

	ins := []string{}
	for l := range lines {
		if len(l) == 0 {
			continue
		}

		ins = append(ins, l)
	}

	x := 1
	cycle := 1
	ip := 0
	res := 0

	for cycle < 221 {
		parts := strings.Split(ins[(ip)%len(ins)], " ")
		if (cycle+20)%40 == 0 {
			res += cycle * x
		}
		switch parts[0] {
		case "addx":
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			cycle++
			if (cycle+20)%40 == 0 {
				res += cycle * x
			}
			cycle++
			x += v
		case "noop":
			cycle++
		default:
			panic("wrong op code")
		}
		ip++
	}

	return res
}

func solve2(file string) int {
	lines := strings.SplitSeq(file, "\n")

	ins := []string{}
	for l := range lines {
		if len(l) == 0 {
			continue
		}

		ins = append(ins, l)
	}

	x := 1
	cycle := 0
	ip := 0
	res := 0
	img := make([]bool, 240)

	for cycle < 240 {
		parts := strings.Split(ins[(ip)%len(ins)], " ")
		if horPos := cycle % 40; horPos == x-1 || horPos == x || horPos == x+1 {
			img[cycle] = true
		}
		switch parts[0] {
		case "addx":
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			cycle++
			if horPos := cycle % 40; horPos == x-1 || horPos == x || horPos == x+1 {
				img[cycle] = true
			}
			cycle++
			x += v
		case "noop":
			cycle++
		default:
			panic("wrong op code")
		}
		ip++
	}

	for i, p := range img {
		if i%40 == 0 {
			fmt.Println()
		}
		if p {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}

	return res
}
