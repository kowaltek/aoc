package main

import (
	"log"
	"os"
	"testing"
)

func TestSolution1(t *testing.T) {
	raw, err := os.ReadFile("testinput.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve1(file)
	if res != 126384 {
		t.Errorf("expected 126384, got %d", res)
	}
}

//	func TestSolution2(t *testing.T) {
//		raw, err := os.ReadFile("testinput.txt")
//		if err != nil {
//			log.Println("error reading file")
//			log.Println(err)
//			os.Exit(1)
//		}
//
//		file := string(raw)
//		res := solve2(file)
//		if res != 31 {
//			t.Errorf("expected 31, got %d", res)
//		}
//	}
// func TestSequence(t *testing.T) {
// 	tcs := map[string]string{
// 		"029A": "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A",
// 		"980A": "<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A",
// 		"179A": "<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A",
// 		"456A": "<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A",
// 		"379A": "<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A",
// 	}
// 	for input, result := range tcs {
// 		seq1 := findPath(input)
// 		seq2 := findDirpadPath(seq1)
// 		seq3 := findDirpadPath(seq2)
// 		if seq3 != result {
// 			fmt.Println("---------------------------------------")
// 			fmt.Println(input)
// 			fmt.Println(seq1)
// 			fmt.Println(seq2)
// 			fmt.Println(seq3)
// 			fmt.Println(result)
// 			fmt.Println("---------------------------------------")
// 			t.Error()
// 		}
// 	}
// }
// //
// func TestA(t *testing.T) {
// 	input := "029A"
// 	seq1 := findPath(input)
// 	if seq1 != "<A^A>^^AvvvA" {
// 		fmt.Println("---------------------------------------")
// 		fmt.Println(input)
// 		fmt.Println(seq1)
// 		fmt.Println("<A^A>^^AvvvA")
// 		fmt.Println("---------------------------------------")
// 		t.Error()
// 	}
// 	seq2 := findDirpadPath(seq1)
// 	if seq2 != "v<<A>>^A<A>AvA<^AA>A<vAAA>^A" {
// 		fmt.Println("---------------------------------------")
// 		fmt.Println(input)
// 		fmt.Println(seq1)
// 		fmt.Println(seq2)
// 		fmt.Println("v<<A>>^A<A>AvA<^AA>A<vAAA>^A")
// 		fmt.Println("---------------------------------------")
// 		t.Error()
// 	}
// 	seq3 := findDirpadPath(seq2)
// 	if seq3 != "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A" {
// 		fmt.Println("---------------------------------------")
// 		fmt.Println(input)
// 		fmt.Println(seq1)
// 		fmt.Println(seq2)
// 		fmt.Println(seq3)
// 		fmt.Println("<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A")
// 		fmt.Println("---------------------------------------")
// 		t.Error()
// 	}
// }
