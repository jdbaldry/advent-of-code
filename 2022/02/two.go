package main

import (
	"bufio"
	"io"
)

var (
	twoResults = map[string]int{
		"A X": int(loss + scissors),
		"A Y": int(draw + rock),
		"A Z": int(win + paper),
		"B X": int(loss + rock),
		"B Y": int(draw + paper),
		"B Z": int(win + scissors),
		"C X": int(loss + paper),
		"C Y": int(draw + scissors),
		"C Z": int(win + rock),
	}
)

func two(r io.Reader) (int, error) {
	var sum int

	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		sum += twoResults[scanner.Text()]
	}
	if err := scanner.Err(); err != nil {
		return sum, err
	}
	return sum, nil
}
