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

func solve1(file string) uint {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	res := uint(0)
	for _, l := range lines {
		x, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		n := uint(x)
		for range 2000 {
			tmp := n
			n *= 64
			n = n ^ tmp
			n = n % 16777216
			tmp = n
			n /= 32
			n = n ^ tmp
			n = n % 16777216
			tmp = n
			n *= 2048
			n = n ^ tmp
			n = n % 16777216
		}
		res += n
	}
	return res
}

func solve2(file string) uint {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	var results []map[string]int
	for _, l := range lines {
		result := make(map[string]int)
		var diffs [4]int
		x, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		n := uint(x)
		for i := range 2000 {
			old := n % 10
			tmp := n
			n *= 64
			n = n ^ tmp
			n = n % 16777216
			tmp = n
			n /= 32
			n = n ^ tmp
			n = n % 16777216
			tmp = n
			n *= 2048
			n = n ^ tmp
			n = n % 16777216
			next := n % 10
			diffs[0], diffs[1], diffs[2] = diffs[1], diffs[2], diffs[3]
			diffs[3] = int(next) - int(old)
			if i > 2 {
				k := key(diffs)
				if _, ok := result[k]; !ok {
					result[k] = int(next)
				}
			}
		}
		results = append(results, result)
	}
	res := uint(0)
	for i := -9; i < 10; i++ {
		for j := -9; j < 10; j++ {
			for k := -9; k < 10; k++ {
				for l := -9; l < 10; l++ {
					subRes := uint(0)
					k := key([4]int{i, j, k, l})
					for _, result := range results {
						if tmp, ok := result[k]; ok {
							subRes += uint(tmp)
						}
					}
					if subRes > res {
						res = subRes
					}
				}
			}
		}
	}
	return res
}

func key(diffs [4]int) string {
	return fmt.Sprintf("%d|%d|%d|%d", diffs[0], diffs[1], diffs[2], diffs[3])
}
