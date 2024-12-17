package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type proc struct {
	regs struct {
		a, b, c int
	}
	ip int
}

var cache = make(map[string]int)

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

func solve1(file string) string {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	inProg := false
	var prog []int
	dev := proc{}
	for _, l := range lines {
		if len(l) == 0 {
			inProg = true
			continue
		}
		if !inProg {
			sub := strings.Split(l, ": ")
			val, err := strconv.Atoi(sub[1])
			if err != nil {
				panic(err)
			}
			switch l[9] {
			case 'A':
				dev.regs.a = val
			case 'B':
				dev.regs.b = val
			case 'C':
				dev.regs.c = val
			}
		} else {
			sub := strings.Split(l, ": ")
			ops := strings.Split(sub[1], ",")
			for _, opStr := range ops {
				op, err := strconv.Atoi(opStr)
				if err != nil {
					panic(err)
				}
				prog = append(prog, op)
			}
		}
	}

	var res []string
	for dev.ip >= 0 && dev.ip < len(prog) {
		op := prog[dev.ip]
		operand := prog[dev.ip+1]
		hasJumped := false
		switch op {
		case 0:
			dev.regs.a = dev.regs.a / int(math.Pow(2, float64(comboOperand(operand, dev))))
		case 1:
			dev.regs.b ^= operand
		case 2:
			dev.regs.b = comboOperand(operand, dev) % 8
		case 3:
			if dev.regs.a != 0 {
				dev.ip = operand
				hasJumped = true
			}
		case 4:
			dev.regs.b ^= dev.regs.c
		case 5:
			res = append(res, strconv.Itoa(comboOperand(operand, dev)%8))
		case 6:
			dev.regs.b = dev.regs.a / int(math.Pow(2, float64(comboOperand(operand, dev))))
		case 7:
			dev.regs.c = dev.regs.a / int(math.Pow(2, float64(comboOperand(operand, dev))))
		}
		if !hasJumped {
			dev.ip += 2
		}
	}
	return strings.Join(res, ",")
}

func solve2(file string) int {
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n")
	inProg := false
	var prog []int
	dev := proc{}
	for _, l := range lines {
		if len(l) == 0 {
			inProg = true
			continue
		}
		if !inProg {
			sub := strings.Split(l, ": ")
			val, err := strconv.Atoi(sub[1])
			if err != nil {
				panic(err)
			}
			switch l[9] {
			case 'A':
				dev.regs.a = val
			case 'B':
				dev.regs.b = val
			case 'C':
				dev.regs.c = val
			}
		} else {
			sub := strings.Split(l, ": ")
			ops := strings.Split(sub[1], ",")
			for _, opStr := range ops {
				op, err := strconv.Atoi(opStr)
				if err != nil {
					panic(err)
				}
				prog = append(prog, op)
			}
		}
	}

	// var origA int
	// var a int
	a := 190384609508360
	for {
		log.Println(a)
		dev.regs.a = a
		res := runProg(dev, prog)
		if res == nil {
			gotIt := true
			for i := range prog {
				if prog[i] != res[i] {
					gotIt = false
					break
				}
			}
			if gotIt {
				log.Println(res)
				log.Println(prog)
				return a
			}
		}
		d := 0
		for i := 0; i < len(prog); i++ {
			if len(res)-i > 0 && prog[len(prog)-1-i] != res[len(res)-1-i] {
				// d = int(math.Pow(2, float64(len(res)-i)*3))
				d = 1
				break
			} else if len(res)-i == 0 && prog[len(prog)-i] == res[len(res)-i] {
				d = 0
				a = (a+1)*8 - 8
				break
			} else {
				d = 1
			}
		}
		a += d
	}

	// var as []int
	// for len(as) < len(prog) {
	// 	i := len(as)
	// 	a := 0
	// 	for j, n := range as {
	// 		a += int(math.Pow(8, float64(j))) * n
	// 	}
	// 	log.Println(a)
	// 	for {
	// 		log.Println(a)
	// 		dev.regs.a = a
	// 		res := runProg(dev, prog)
	// 		if res[0] == prog[i] {
	// 			as = append(as, a)
	// 			break
	// 		}
	// 		a += 8
	// 	}
	// }
	return 0
	// for i := 0; i < math.MaxInt; i++ {
	// 	a = int(math.Pow(2, float64(i)))
	// 	dev.regs.a = a
	// 	lenRes := runProg(dev, prog)
	// 	if lastRes < lenRes {
	// 		log.Println(a)
	// 		log.Println(lenRes)
	// 		lastRes = lenRes
	// 	}
	// 	if lenRes >= len(prog) || lenRes == -1 {
	// 		origA = a
	// 		break
	// 	}
	// }
	// log.Println(origA)
	// for a := origA; a < origA+8; {
	// 	// if a%1000000 == 0 {
	// 	log.Println(a)
	// 	// }
	// 	dev.regs.a = a
	// 	if runProg(dev, prog) == -1 {
	// 		return a
	// 	}
	// 	a++
	// }
	return 0
}

func comboOperand(operand int, dev proc) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return dev.regs.a
	case 5:
		return dev.regs.b
	case 6:
		return dev.regs.c
	}
	return -1
}

func runProg(dev proc, prog []int) []int {
	var res []int
	for dev.ip >= 0 && dev.ip < len(prog) {
		// log.Printf("%+v\n", dev)
		// if c, ok := cache[key]; ok {
		// return c + len(res)
		// }
		op := prog[dev.ip]
		operand := prog[dev.ip+1]
		hasJumped := false
		switch op {
		case 0:
			dev.regs.a = dev.regs.a / int(math.Pow(2, float64(comboOperand(operand, dev))))
		case 1:
			dev.regs.b ^= operand
		case 2:
			dev.regs.b = comboOperand(operand, dev) % 8
		case 3:
			if dev.regs.a != 0 {
				dev.ip = operand
				hasJumped = true
			}
		case 4:
			dev.regs.b ^= dev.regs.c
		case 5:
			res = append(res, comboOperand(operand, dev)%8)
			if len(res) > 0 && res[len(res)-1] != prog[len(res)-1] {
			}
		case 6:
			dev.regs.b = dev.regs.a / int(math.Pow(2, float64(comboOperand(operand, dev))))
		case 7:
			dev.regs.c = dev.regs.a / int(math.Pow(2, float64(comboOperand(operand, dev))))
		}
		if !hasJumped {
			dev.ip += 2
		}
	}
	if len(prog) != len(res) {
		log.Printf("%+v\n", res)
		return res
	}
	for i := range prog {
		if prog[i] != res[i] {
			log.Printf("%+v\n", res)
			return res
		}
	}
	log.Printf("%+v\n", res)
	return nil
}

func getKey(dev proc, res []int) string {
	keyArr := []int{dev.regs.a, dev.regs.b, dev.regs.c, dev.ip}
	keyArr = append(keyArr, res...)
	var keyStr []string
	for _, n := range keyArr {
		keyStr = append(keyStr, strconv.Itoa(n))
	}
	return strings.Join(keyStr, ",")
}
