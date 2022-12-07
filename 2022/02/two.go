package main

import (
	"bufio"
	"fmt"
	"io"
)

func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	results := map[string]int{
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

	var sum int
	for i := 0; scanner.Scan(); i++ {
		sum += results[scanner.Text()]
	}

	if err := scanner.Err(); err != nil {
		return sum, fmt.Errorf("%w during scanning", err)
	}

	return sum, nil
}
