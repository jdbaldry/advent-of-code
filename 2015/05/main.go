package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func isExcludedSequence(prev, curr rune) bool {
	return (prev == 'a' && curr == 'b') ||
		(prev == 'c' && curr == 'd') ||
		(prev == 'p' && curr == 'q') ||
		(prev == 'x' && curr == 'y')
}

func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func one(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var total int

	for scanner.Scan() {
		var (
			prev               rune
			vowels             int
			letterAppearsTwice bool
			containsExcluded   bool
		)

		for _, char := range scanner.Text() {
			if isExcludedSequence(prev, char) {
				containsExcluded = true

				break
			}

			if char == prev {
				letterAppearsTwice = true
			}

			if isVowel(char) {
				vowels++
			}

			prev = char
		}

		if !containsExcluded && letterAppearsTwice && vowels >= 3 {
			total++
		}
	}

	return total
}

func two(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var total int

	for scanner.Scan() {
		type pair [2]rune

		var (
			pairs                = make(map[pair]struct{})
			prev                 rune
			prevprev             rune
			nonOverlappingPair   bool
			repeatWithOneBetween bool
		)

		for _, char := range scanner.Text() {
			_, ok := pairs[pair{char, prev}]
			if ok && prev != prevprev {
				nonOverlappingPair = true
			}

			pairs[pair{char, prev}] = struct{}{}

			if char == prevprev && char != prev {
				repeatWithOneBetween = true
			}

			prevprev = prev
			prev = char
		}

		if nonOverlappingPair && repeatWithOneBetween {
			total++
		}
	}

	return total
}

func main() {
	logger := log.New(os.Stderr, "", log.Llongfile)

	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("ERROR: %v\n", err)
	}

	//nolint:forbidigo
	fmt.Println(one(input))

	if _, err := input.Seek(0, 0); err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	//nolint:forbidigo
	fmt.Println(two(input))
}
