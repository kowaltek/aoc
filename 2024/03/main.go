package main

import (
	"log"
	"os"
	"regexp"
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
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	r2 := regexp.MustCompile("[0-9]+")
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		m := r.FindAllString(l, -1)
		for _, exp := range m {
			nums := r2.FindAllString(exp, -1)
			n1, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Println("error converting number")
				log.Println(err)
				os.Exit(1)
			}
			n2, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Println("error converting number")
				log.Println(err)
				os.Exit(1)
			}
			res += n1 * n2
		}
	}

	return res
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")
	res := 0
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	r2 := regexp.MustCompile("[0-9]+")
	enabled := true
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		m := r.FindAllString(l, -1)
		for _, exp := range m {
			if exp == "do()" {
				enabled = true
				continue
			}
			if exp == "don't()" {
				enabled = false
				continue
			}
			if enabled {
				nums := r2.FindAllString(exp, -1)
				n1, err := strconv.Atoi(nums[0])
				if err != nil {
					log.Println("error converting number")
					log.Println(err)
					os.Exit(1)
				}
				n2, err := strconv.Atoi(nums[1])
				if err != nil {
					log.Println("error converting number")
					log.Println(err)
					os.Exit(1)
				}
				res += n1 * n2
			}
		}
	}
	return res
}
