package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func two(input [][]int) int {
	//nolint:gomnd
	highest := make([]int, 4) // three highest and the current sum

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

// insertSum will insert next into sorted array highest
// maintaining sort order.
func insertSum(highest *[3]int, next *int) {
	switch {
	case *next > highest[0]:
		highest[0], highest[1], highest[2] = *next, highest[0], highest[1]
	case *next > highest[1]:
		highest[1], highest[2] = *next, highest[1]
	case *next > highest[2]:
		highest[2] = *next
	}

	*next = 0
}

func twoNoParse(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var (
		highest [3]int
		next    int
	)

	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			insertSum(&highest, &next)
		default:
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return 0, fmt.Errorf("unable to parse input: %w", err)
			}

			next += i
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	// Handle absent empty line.
	insertSum(&highest, &next)

	var sum int
	for _, addend := range highest {
		sum += addend
	}

	return sum, nil
}
