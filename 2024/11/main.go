package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type stone struct {
	n      int
	blinks int
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

	var stones []int
	for _, s := range strings.Split(file, " ") {
		sNum, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		stones = append(stones, sNum)
	}

	for range 25 {
		prev := make([]int, len(stones))
		copy(prev, stones)
		stones = stones[:0]
		for _, s := range prev {
			if s == 0 {
				stones = append(stones, 1)
			} else if noOfDigits(s)%2 == 0 {
				no1 := 0
				mul := 1
				for range noOfDigits(s) / 2 {
					no1 += mul * (s % 10)
					s /= 10
					mul *= 10
				}
				stones = append(stones, s, no1)
			} else {
				stones = append(stones, s*2024)
			}
		}
	}

	return len(stones)
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")

	var stones []stone
	for _, s := range strings.Split(file, " ") {
		sNum, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		stones = append(stones, stone{n: sNum, blinks: 75})
	}

	cache := make(map[stone]int)
	res := 0
	for _, s := range stones {
		res += checkStone(s, cache)
	}

	return res
}

func noOfDigits(s int) int {
	if s == 0 {
		return 1
	}
	count := 0
	for s != 0 {
		s /= 10
		count++
	}
	return count
}

func checkStone(s stone, cache map[stone]int) int {
	if res, ok := cache[s]; ok {
		return res
	}
	if s.blinks == 0 {
		return 1
	}
	if s.n == 0 {
		newS := stone{n: 1, blinks: s.blinks - 1}
		res := checkStone(newS, cache)
		cache[newS] = res
		return res
	} else if noOfDigits(s.n)%2 == 0 {
		no1 := 0
		no2 := s.n
		mul := 1
		for range noOfDigits(s.n) / 2 {
			no1 += mul * (no2 % 10)
			no2 /= 10
			mul *= 10
		}
		s1 := stone{n: no2, blinks: s.blinks - 1}
		s2 := stone{n: no1, blinks: s.blinks - 1}
		res1 := checkStone(s1, cache)
		res2 := checkStone(s2, cache)
		cache[s1] = res1
		cache[s2] = res2
		return res1 + res2
	} else {
		newS := stone{n: s.n * 2024, blinks: s.blinks - 1}
		res := checkStone(newS, cache)
		cache[newS] = res
		return res
	}
}
