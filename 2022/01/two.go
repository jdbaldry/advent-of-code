package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func two(input [][]int) int {
	highest := make([]int, 4)
	for _, addends := range input {
		var sum int
		for _, addend := range addends {
			sum += addend
		}
		highest[3] = sum
		sort.Sort(sort.Reverse(sort.IntSlice(highest)))
	}

	var sum int
	for _, addend := range highest[:3] {
		sum += addend
	}
	return sum
}

func twoNoSort(input [][]int) int {
	var highest [3]int
	for _, addends := range input {
		var sum int
		for _, addend := range addends {
			sum += addend
		}
		if sum > highest[0] {
			highest[0], highest[1], highest[2] = sum, highest[0], highest[1]
			continue
		}
		if sum > highest[1] {
			highest[1], highest[2] = sum, highest[1]
			continue
		}
		if sum > highest[2] {
			highest[2] = sum
		}
	}

	var sum int
	for _, addend := range highest {
		sum += addend
	}
	return sum
}

func twoNoParse(r io.Reader) int {
	var highest [3]int

	scanner := bufio.NewScanner(r)

	var next int
	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			if next > highest[0] {
				next, highest[0], highest[1], highest[2] = 0, next, highest[0], highest[1]
				continue
			}
			if next > highest[1] {
				next, highest[1], highest[2] = 0, next, highest[1]
				continue
			}
			if next > highest[2] {
				next, highest[2] = 0, next
			}
			next = 0
		default:
			int, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(fmt.Sprintf("Unable to convert %q to integer\nThis should have been an error not a panic!", scanner.Text()))
			}
			next += int
		}
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Unexpected error: %v\nThis should have been an error not a panic!", err))
	}

	// Handle absent trailing newline.
	if next > highest[0] {
		next, highest[0], highest[1], highest[2] = 0, next, highest[0], highest[1]
	}
	if next > highest[1] {
		next, highest[1], highest[2] = 0, next, highest[1]
	}
	if next > highest[2] {
		next, highest[2] = 0, next
	}

	var sum int
	for _, addend := range highest {
		sum += addend
	}
	return sum
}
