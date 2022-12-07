package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func parse(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)

	var (
		parsed [][]int
		next   []int
	)

	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			parsed = append(parsed, next)
			next = []int{}
		default:
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return parsed, fmt.Errorf("unable to parse input: %w", err)
			}

			next = append(next, i)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("%w during scanning", err)
	}

	// Handle absent trailing newline.
	parsed = append(parsed, next)

	return parsed, nil
}

func one(input [][]int) int {
	var highest int

	for _, addends := range input {
		var sum int
		for _, addend := range addends {
			sum += addend
		}

		if sum > highest {
			highest = sum
		}
	}

	return highest
}
