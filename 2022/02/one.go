package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"unicode/utf8"
)

const (
	rock = iota + 1
	paper
	scissors
)

const (
	loss = iota * 3
	draw
	win
)

const distinctShapes = 3

var (
	errUnexpectedInput = errors.New("unexpected input")
	errUnexpectedRune  = errors.New("unexpected rune")
)

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	results := map[string]int{
		"A X": draw + rock,
		"A Y": win + paper,
		"A Z": loss + scissors,
		"B X": loss + rock,
		"B Y": draw + paper,
		"B Z": win + scissors,
		"C X": win + rock,
		"C Y": loss + paper,
		"C Z": draw + scissors,
	}

	var sum int

	for line := 0; scanner.Scan(); line++ {
		result, ok := results[scanner.Text()]
		if !ok {
			return sum, fmt.Errorf("%w on line %d: %q", errUnexpectedInput, line, scanner.Text())
		}

		sum += result
	}

	if err := scanner.Err(); err != nil {
		return sum, fmt.Errorf("%w during scanning", err)
	}

	return sum, nil
}

// oneMod3 solves the problem using the following knowledge:
// rock < paper < scissors < rock ...
// Distance is the number of steps taken up the order to reach the opponents shape
// after first counting one period from one's own shape.
// When the opponent chooses a winning shape, the distance is 2.
// When the opponent chooses a losing shape, the distance is 1.
// When the opponent chooses the same shape, the distance is 0.
// The distance is rotated by one using (distance + 1) mod 3 to get results that can be mapped
// to scores using *3.
// When the opponent chooses a winning shape, (distance + 1) % 3 is 0.
// When the opponent chooses a losing shape, (distance + 1) % 3 is 2.
// When the opponent chooses the same shape, (distance + 1) % 3 is 1.
func oneMod3(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	strategyRegexp := regexp.MustCompile(`^([ABC]) ([XYZ])$`)
	shapes := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	var sum int

	for line := 0; scanner.Scan(); line++ {
		matches := strategyRegexp.FindStringSubmatch(scanner.Text())
		if len(matches) != strategyRegexp.NumSubexp()+1 {
			return sum, fmt.Errorf(
				"%w on line %d: %q does not match regexp %q",
				errUnexpectedInput,
				line,
				scanner.Text(),
				strategyRegexp.String(),
			)
		}

		opp, own := shapes[matches[1]], shapes[matches[2]]
		distance := ((own + distinctShapes) - opp) % distinctShapes
		sum += ((distance+1)%distinctShapes)*distinctShapes + own
	}

	if err := scanner.Err(); err != nil {
		return sum, fmt.Errorf("%w during scanning", err)
	}

	return sum, nil
}

//nolint:cyclop
func oneMod3ByRunes(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	const (
		stateOpp = iota
		stateSpace
		stateOwn
		stateLine
	)

	var opp, own, sum, state int

	for line, col := 0, 0; scanner.Scan(); col++ {
		char, _ := utf8.DecodeRune(scanner.Bytes())

		switch state {
		case stateOpp:
			if char < 'A' || 'C' < char {
				return sum, fmt.Errorf("%w %q on line %d, column %d, wanted /[ABC]/", errUnexpectedRune, char, line, col)
			}

			opp = int(char) - int('A')
			state = stateSpace

		case stateSpace:
			if char != ' ' {
				return sum, fmt.Errorf(`%w %q on line %d, column %d, wanted " "`, errUnexpectedRune, char, line, col)
			}

			state = stateOwn

		case stateOwn:
			if char < 'X' || 'Z' < char {
				return sum, fmt.Errorf("%w %q on line %d, column %d, wanted /[XYZ]/", errUnexpectedRune, char, line, col)
			}

			own = int(char) - int('X')
			state = stateLine
			distance := ((own + distinctShapes) - opp) % distinctShapes
			sum += ((distance+1)%distinctShapes)*distinctShapes + own + 1

		case stateLine:
			if char != '\n' {
				return sum, fmt.Errorf("%w %q on line %d, column %d, wanted newline", errUnexpectedRune, char, line, col)
			}

			state = stateOpp
			line, col = line+1, -1
		}
	}

	if err := scanner.Err(); err != nil {
		return sum, fmt.Errorf("%w during scanning", err)
	}

	return sum, nil
}
