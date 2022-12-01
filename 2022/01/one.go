package main

import (
	"bufio"
	"io"
	"strconv"
)

func parse(r io.Reader) ([][]int, error) {
	var parsed [][]int
	scanner := bufio.NewScanner(r)
	var next []int
	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			parsed = append(parsed, next)
			next = []int{}
		default:
			int, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return parsed, err
			}
			next = append(next, int)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
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
