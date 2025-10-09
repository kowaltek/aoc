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

	var letters []string
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		letters = strings.Split(l, "")
	}
	seq := letters[:4]
	res := 4
	for _, l := range letters[4:] {
		if isSeqUniq(seq) {
			return res
		}
		seq = append(seq[1:], l)
		res++
	}
	return 0
}

func solve2(file string) int {
	lines := strings.Split(file, "\n")

	var letters []string
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		letters = strings.Split(l, "")
	}
	seq := letters[:14]
	res := 14
	for _, l := range letters[14:] {
		if isSeqUniq(seq) {
			return res
		}
		seq = append(seq[1:], l)
		res++
	}
	return 0
}

func isSeqUniq(seq []string) bool {
	prev := make(map[string]struct{})
	for _, s := range seq {
		if _, ok := prev[s]; ok {
			return false
		}
		prev[s] = struct{}{}
	}
	return true
}
