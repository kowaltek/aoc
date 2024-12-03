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
	if res != 161 {
		t.Errorf("expected 161, got %d", res)
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
	if res != 48 {
		t.Errorf("expected 48, got %d", res)
	}
}
