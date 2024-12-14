package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reNums = regexp.MustCompile(`-?\d+`)

type robot struct {
	x, y, dx, dy int
}

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve2(file, 101, 103)
	log.Printf("Got result: %d\n", res)
}

func solve1(file string, w, h int) int {
	file = strings.Trim(file, "\n")
	var robots []robot
	for _, line := range strings.Split(file, "\n") {
		numsStr := reNums.FindAllString(line, -1)
		var nums []int
		for _, nStr := range numsStr {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		robots = append(robots, robot{
			x:  nums[0],
			y:  nums[1],
			dx: nums[2],
			dy: nums[3],
		})
	}
	for range 100 {
		for i, robot := range robots {
			robot.x = (robot.x + w + robot.dx) % w
			robot.y = (robot.y + h + robot.dy) % h
			robots[i] = robot
		}
	}
	quadrants := make([]int, 4)
	for _, robot := range robots {
		if robot.x < w/2 && robot.y < h/2 {
			quadrants[0]++
		} else if ((w%2 == 0 && robot.x >= w/2) || (w%2 == 1 && robot.x >= w/2+1)) && robot.y < h/2 {
			quadrants[1]++
		} else if robot.x < w/2 && ((h%2 == 0 && robot.y >= h/2) || (h%2 == 1 && robot.y >= h/2+1)) {
			quadrants[2]++
		} else if ((w%2 == 0 && robot.x >= w/2) || (w%2 == 1 && robot.x >= w/2+1)) && ((h%2 == 0 && robot.y >= h/2) || (h%2 == 1 && robot.y >= h/2+1)) {
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func solve2(file string, w, h int) int {
	file = strings.Trim(file, "\n")
	var robots []robot
	for _, line := range strings.Split(file, "\n") {
		numsStr := reNums.FindAllString(line, -1)
		var nums []int
		for _, nStr := range numsStr {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		robots = append(robots, robot{
			x:  nums[0],
			y:  nums[1],
			dx: nums[2],
			dy: nums[3],
		})
	}
	res := 0
	d := 0
	for {
		if (res-33)%101 == 0 && (res-87)%103 == 0 {
			for i, robot := range robots {
				robot.x = (robot.x + w + robot.dx*d) % w
				robot.y = (robot.y + h + robot.dy*d) % h
				robots[i] = robot
			}
			log.Println(res + 1)
			printRobots(robots, w, h)
			d = 0
			break
		}
		d++
		res++
	}
	return res
}

func makeChristmasTree(robots []robot) bool {
	return false
}

func printRobots(robots []robot, w, h int) {
	for x := range w {
		for y := range h {
			print := false
			for _, r := range robots {
				if r.x == x && r.y == y {
					print = true
					break
				}
			}
			if print {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
