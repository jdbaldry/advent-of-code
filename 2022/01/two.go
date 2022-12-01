package main

import (
	"bufio"
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

func twoNoParse(r io.Reader) (int, error) {
	var highest [3]int

	scanner := bufio.NewScanner(r)

	var next int
	insertSum := func(highest *[3]int, next *int) {
		if *next > (*highest)[0] {
			*next, (*highest)[0], (*highest)[1], (*highest)[2] = 0, *next, (*highest)[0], (*highest)[1]
			return
		}
		if *next > (*highest)[1] {
			*next, (*highest)[1], (*highest)[2] = 0, *next, (*highest)[1]
			return
		}
		if *next > (*highest)[2] {
			*next, (*highest)[2] = 0, *next
		}
		*next = 0
		return
	}
	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			insertSum(&highest, &next)
		default:
			int, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return 0, err
			}
			next += int
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Handle absent empty line.
	insertSum(&highest, &next)
	var sum int
	for _, addend := range highest {
		sum += addend
	}
	return sum, nil
}
