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
	file = strings.Trim(file, "\n")
	inputStr := strings.Split(file, "")
	var input []int
	sum := 0
	for _, iStr := range inputStr {
		n, err := strconv.Atoi(iStr)
		if err != nil {
			panic(err)
		}
		input = append(input, n)
		sum += n
	}
	disk := make([]int, sum)
	for i := range disk {
		disk[i] = -1
	}
	inFile := true
	i := 0
	fileNo := 0
	for _, n := range input {
		if inFile {
			inFile = false
			for range n {
				disk[i] = fileNo
				i++
			}
			fileNo++
		} else {
			i += n
			inFile = true
		}
	}
	revI := len(disk) - 1
	for ; disk[revI] == -1; revI-- {
	}
	for i := range disk {
		if i >= revI {
			break
		}
		if disk[i] == -1 {
			for ; disk[revI] == -1; revI-- {
			}
			disk[i] = disk[revI]
			disk[revI] = -1
			revI--
		}
	}
	res := 0
	for i := range disk {
		if disk[i] != -1 {
			res += i * disk[i]
		}
	}
	return res
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	inputStr := strings.Split(file, "")
	var input []int
	sum := 0
	for _, iStr := range inputStr {
		n, err := strconv.Atoi(iStr)
		if err != nil {
			panic(err)
		}
		input = append(input, n)
		sum += n
	}
	disk := make([]int, sum)
	for i := range disk {
		disk[i] = -1
	}
	inFile := true
	i := 0
	fileNo := 0
	for _, n := range input {
		if inFile {
			inFile = false
			for range n {
				disk[i] = fileNo
				i++
			}
			fileNo++
		} else {
			i += n
			inFile = true
		}
	}
	for revI := len(disk) - 1; revI >= 0; {
		for ; revI >= 0 && disk[revI] == -1; revI-- {
		}
		if revI < 0 {
			break
		}
		fileNo := disk[revI]
		fileSize := 0
		var fileI int
		for i := revI; i >= 0 && disk[i] == fileNo; i-- {
			fileSize++
			fileI = i
		}
		for i := 0; i < len(disk) && i < fileI; i++ {
			if disk[i] != -1 {
				continue
			}
			freeSpace := 0
			for j := i; j < len(disk) && disk[j] == -1; j++ {
				freeSpace++
			}
			if freeSpace >= fileSize {
				for j := range fileSize {
					disk[i] = fileNo
					disk[fileI+j] = -1
					i++
				}
				break
			}
		}
		revI = fileI - 1
	}
	res := 0
	for i := range disk {
		if disk[i] != -1 {
			res += i * disk[i]
		}
	}
	return res
}
