package main

import (
	"log"
	"os"
	"slices"
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
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	nodes := make(map[string][]string)
	for _, l := range lines {
		comps := strings.Split(l, "-")
		if list, ok := nodes[comps[0]]; ok {
			if !slices.Contains(list, comps[1]) {
				nodes[comps[0]] = append(nodes[comps[0]], comps[1])
			}
		} else {
			nodes[comps[0]] = []string{comps[1]}
		}
		if list, ok := nodes[comps[1]]; ok {
			if !slices.Contains(list, comps[0]) {
				nodes[comps[1]] = append(nodes[comps[1]], comps[0])
			}
		} else {
			nodes[comps[1]] = []string{comps[0]}
		}
	}
	var res [][]string
	for first, firstList := range nodes {
		for _, second := range firstList {
			secondList, _ := nodes[second]
			for _, third := range secondList {
				thirdList, _ := nodes[third]
				if slices.Contains(thirdList, first) && (first[0] == 't' || second[0] == 't' || third[0] == 't') {
					res = append(res, []string{first, second, third})
				}
			}
		}
	}
	res = dedup(res)
	return len(res)
}

func dedup(combinations [][]string) [][]string {
	for i := 0; i < len(combinations); i++ {
		curr := combinations[i]
		tmp := make([][]string, len(combinations))
		copy(tmp, combinations)
		tmp2 := tmp[i+1:]
		tmp = tmp[:i]
		tmp = append(tmp, tmp2...)
		if slices.ContainsFunc(tmp, func(combination []string) bool {
			if slices.Contains(curr, combination[0]) &&
				slices.Contains(curr, combination[1]) &&
				slices.Contains(curr, combination[2]) {
				return true
			}
			return false
		}) {
			tmp := combinations[i+1:]
			combinations = combinations[:i]
			combinations = append(combinations, tmp...)
			i--
		}
	}
	return combinations
}

func solve2(file string) string {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	nodes := make(map[string][]string)
	for _, l := range lines {
		comps := strings.Split(l, "-")
		if list, ok := nodes[comps[0]]; ok {
			if !slices.Contains(list, comps[1]) {
				nodes[comps[0]] = append(nodes[comps[0]], comps[1])
			}
		} else {
			nodes[comps[0]] = []string{comps[1]}
		}
		if list, ok := nodes[comps[1]]; ok {
			if !slices.Contains(list, comps[0]) {
				nodes[comps[1]] = append(nodes[comps[1]], comps[0])
			}
		} else {
			nodes[comps[1]] = []string{comps[0]}
		}
	}
	res := findLargestSet(nodes)
	slices.Sort(res)
	return strings.Join(res, ",")
}

func findLargestSet(nodes map[string][]string) []string {
	res := []string{}
	for first := range nodes {
		sub := findLargestSetForNode(first, nodes)
		if len(sub) > len(res) {
			res = sub
		}
	}
	return res
}

func findLargestSetForNode(node string, nodes map[string][]string) []string {
	res := []string{node}
	delete(nodes, node)
	for k, v := range nodes {
		isConnectedToAll := true
		for _, visited := range res {
			isConnected := false
			for _, connectedNode := range v {
				if visited == connectedNode {
					isConnected = true
					break
				}
			}
			if !isConnected {
				isConnectedToAll = false
				break
			}
		}
		if isConnectedToAll {
			res = append(res, k)
		}
	}
	return res
}
