package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	for i := 0; scanner.Scan(); i++ {
		if floor == -1 {
			return i
		}

		if scanner.Text() == "(" {
			floor++
		} else {
			floor--
		}
	}

	return -1
}

func main() {
	logger := log.New(os.Stderr, "", log.Llongfile)

	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("Unable to open file input.txt: %v\n", err)
	}
	defer input.Close()

	//nolint
	fmt.Println(one(input))

	if _, err := input.Seek(0, 0); err != nil {
		logger.Println(err)
	}

	//nolint
	fmt.Println(two(input))
}
