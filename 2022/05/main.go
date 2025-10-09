package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const StacksNum = 9

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve2(file)
	log.Printf("Got result: %s\n", res)
}

type instruction struct {
	n    int64
	src  int64
	dest int64
}

func solve1(file string) string {
	lines := strings.Split(file, "\n")

	stacks := make([][]string, StacksNum)
	var instructions []instruction
	stacksDefinition := true
	for _, l := range lines {
		if len(l) == 0 {
			stacksDefinition = false
			continue
		}
		if len(l) == 0 && !stacksDefinition {
			continue
		}
		if stacksDefinition {
			for i := range StacksNum {
				if string(l[i*4+1]) != " " {
					stacks[i] = slices.Concat([]string{string(l[i*4+1])}, stacks[i])
				}
			}
		} else {
			parts := strings.Split(l, " ")
			n, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				panic(err)
			}
			src, err := strconv.ParseInt(parts[3], 10, 64)
			if err != nil {
				panic(err)
			}
			dest, err := strconv.ParseInt(parts[5], 10, 64)
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, instruction{
				n:    n,
				src:  src - 1,
				dest: dest - 1,
			})
		}
	}
	for i := range len(stacks) {
		stacks[i] = stacks[i][1:]
	}
	for _, instruction := range instructions {
		for range instruction.n {
			toMove := stacks[instruction.src][len(stacks[instruction.src])-1]
			stacks[instruction.src] = stacks[instruction.src][:len(stacks[instruction.src])-1]
			stacks[instruction.dest] = append(stacks[instruction.dest], toMove)
		}
	}
	res := ""
	for _, stack := range stacks {
		res += stack[len(stack)-1]
	}
	return res
}

func solve2(file string) string {
	lines := strings.Split(file, "\n")

	stacks := make([][]string, StacksNum)
	var instructions []instruction
	stacksDefinition := true
	for _, l := range lines {
		if len(l) == 0 {
			stacksDefinition = false
			continue
		}
		if len(l) == 0 && !stacksDefinition {
			continue
		}
		if stacksDefinition {
			for i := range StacksNum {
				if string(l[i*4+1]) != " " {
					stacks[i] = slices.Concat([]string{string(l[i*4+1])}, stacks[i])
				}
			}
		} else {
			parts := strings.Split(l, " ")
			n, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				panic(err)
			}
			src, err := strconv.ParseInt(parts[3], 10, 64)
			if err != nil {
				panic(err)
			}
			dest, err := strconv.ParseInt(parts[5], 10, 64)
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, instruction{
				n:    n,
				src:  src - 1,
				dest: dest - 1,
			})
		}
	}
	for i := range len(stacks) {
		stacks[i] = stacks[i][1:]
	}
	for _, instruction := range instructions {
		srcLen := len(stacks[instruction.src])
		toMove := stacks[instruction.src][srcLen-int(instruction.n):]
		stacks[instruction.src] = stacks[instruction.src][:srcLen-int(instruction.n)]
		stacks[instruction.dest] = append(stacks[instruction.dest], toMove...)
	}
	res := ""
	for _, stack := range stacks {
		res += stack[len(stack)-1]
	}
	return res
}
