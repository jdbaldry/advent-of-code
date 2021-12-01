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

func one(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	var total int

	for s.Scan() {
		var prev rune
		var vowels int
		var letterAppearsTwice bool
		var containsExcluded bool
		for _, r := range s.Text() {
			if (prev == 'a' && r == 'b') ||
				(prev == 'c' && r == 'd') ||
				(prev == 'p' && r == 'q') ||
				(prev == 'x' && r == 'y') {
				containsExcluded = true
				break
			}
			if r == prev {
				letterAppearsTwice = true
			}
			if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
				vowels++
			}
			prev = r
		}
		if !containsExcluded && letterAppearsTwice && vowels >= 3 {
			total++
		}
	}
	return total
}

func two(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	var total int
	for s.Scan() {
		type pair [2]rune

		var pairs = make(map[pair]struct{})
		var prev rune
		var prevprev rune
		var nonOverlappingPair bool
		var repeatWithOneBetween bool
		for _, r := range s.Text() {
			_, ok := pairs[pair{r, prev}]
			if ok && prev != prevprev {
				nonOverlappingPair = true
			}
			pairs[pair{r, prev}] = struct{}{}

			if r == prevprev && r != prev {
				repeatWithOneBetween = true
			}
			prevprev = prev
			prev = r
		}

		if nonOverlappingPair && repeatWithOneBetween {
			total++
		}
	}
	return total
}

func main() {
	input, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("Unable to open file %s: %v\n", inputFile, err)
	}

	fmt.Println(one(input))
	input.Seek(0, 0)
	fmt.Println(two(input))
}
