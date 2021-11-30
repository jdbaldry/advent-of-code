package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "", log.Llongfile)
)

func one(r io.Reader) int {
	var floor int

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		if s.Text() == "(" {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func two(r io.Reader) int {
	var floor int

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for i := 0; s.Scan(); i++ {
		if floor == -1 {
			return i
		}
		if s.Text() == "(" {
			floor++
		} else {
			floor--
		}
	}
	return -1
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("Unable to open file input.txt: %v\n", err)
	}
	defer input.Close()

	fmt.Println(one(input))
	input.Seek(0, 0)
	fmt.Println(two(input))
}
