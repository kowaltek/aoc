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
	if res != 7036 {
		t.Errorf("expected 7036, got %d", res)
	}
}

func TestSolution12(t *testing.T) {
	raw, err := os.ReadFile("testinput2.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve1(file)
	if res != 11048 {
		t.Errorf("expected 11048, got %d", res)
	}
}

func TestSolution2(t *testing.T) {
	raw, err := os.ReadFile("testinput.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve2(file)
	if res != 45 {
		t.Errorf("expected 45, got %d", res)
	}
}

func TestSolution22(t *testing.T) {
	raw, err := os.ReadFile("testinput2.txt")
	if err != nil {
		log.Println("error reading file")
		log.Println(err)
		os.Exit(1)
	}

	file := string(raw)
	res := solve2(file)
	if res != 64 {
		t.Errorf("expected 64, got %d", res)
	}
}
