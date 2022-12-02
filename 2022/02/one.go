package main

import (
	"bufio"
	"io"
)

const (
	rock uint16 = iota + 1
	paper
	scissors
)

const (
	loss = iota * 3
	draw
	win
)

var (
	oneResults = map[string]int{
		"A X": int(draw + rock),
		"A Y": int(win + paper),
		"A Z": int(loss + scissors),
		"B X": int(loss + rock),
		"B Y": int(draw + paper),
		"B Z": int(win + scissors),
		"C X": int(win + rock),
		"C Y": int(loss + paper),
		"C Z": int(draw + scissors),
	}
)

func one(r io.Reader) (int, error) {
	var sum int
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		sum += oneResults[scanner.Text()]
	}
	if err := scanner.Err(); err != nil {
		return sum, err
	}

	return sum, nil
}
