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
	if res != "4,6,3,5,6,3,5,2,1,0" {
		t.Errorf("expected \"4,6,3,5,6,3,5,2,1,0\", got %s", res)
	}
}

func TestSolution2(t *testing.T) {
	raw, err := os.ReadFile("testinput2.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}
	file := string(raw)
	res := solve2(file)
	if res != 117440 {
		t.Errorf("expected 117440, got %d", res)
	}
}
