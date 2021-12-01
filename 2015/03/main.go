package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFile = "input.txt"

var logger = log.New(os.Stderr, "", log.Llongfile)

type house [2]int

func one(r io.Reader) int {
	var visited = make(map[house]struct{})
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for x, y := 0, 0; s.Scan(); {
		visited[house{x, y}] = struct{}{}
		switch s.Text() {
		case "^":
			x--
		case "v":
			x++
		case ">":
			y++
		case "<":
			y--
		case "\n":
		default:
			logger.Fatalf("Unexpected character in input %q\n", s.Text())
		}
	}

	var total int
	for range visited {
		total++
	}

	return total
}

func two(r io.Reader) int {
	var visited = make(map[house]struct{})
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for i, x, y, a, b := 0, 0, 0, 0, 0; s.Scan(); i++ {
		if i%2 == 0 {
			visited[house{x, y}] = struct{}{}
		} else {
			visited[house{a, b}] = struct{}{}
		}
		switch s.Text() {
		case "^":
			if i%2 == 0 {
				x--
			} else {
				a--
			}
		case "v":
			if i%2 == 0 {
				x++
			} else {
				a++
			}
		case ">":
			if i%2 == 0 {
				y++
			} else {
				b++
			}
		case "<":
			if i%2 == 0 {
				y--
			} else {
				b--
			}
		case "\n":
		default:
			logger.Fatalf("Unexpected character in input %q\n", s.Text())
		}
	}

	var total int
	for range visited {
		total++
	}
	return total
}

func main() {
	input, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("Unable to open %s: %v\n", inputFile, err)
	}

	fmt.Println(one(input))
	input.Seek(0, 0)
	fmt.Println(two(input))
}
