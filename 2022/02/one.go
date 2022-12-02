package main

import (
	"bufio"
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

var (
	oneResults = map[string]int{
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
	shapes = map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	shapesByRune = map[rune]int{
		'A': rock,
		'B': paper,
		'C': scissors,
		'X': rock,
		'Y': paper,
		'Z': scissors,
	}
)

func one(r io.Reader) (int, error) {
	var sum int
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		result, ok := oneResults[scanner.Text()]
		if !ok {
			return sum, fmt.Errorf("unexpected input on line %d: %q", i, scanner.Text())
		}
		sum += result
	}
	if err := scanner.Err(); err != nil {
		return sum, err
	}

	return sum, nil
}

// oneMod3 solves the problem using the following knowledge:
// rock < paper < scissors < rock ...
// Distance is the number of steps taken up the order to reach the opponents shape.
// When the opponent chooses a winning shape, the distance is 1.
// When the opponent chooses a losing shape, the distance is 2.
// When the opponent chooses the same shape, the distance is 0.
// The distance is rotated by one using (distance + 1) mod 3 to get results that can be mapped
// to scores using *3.
func oneMod3(r io.Reader) (int, error) {
	var sum int
	strategyRegexp := regexp.MustCompile(`^([ABC]) ([XYZ])$`)
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		matches := strategyRegexp.FindStringSubmatch(scanner.Text())
		if len(matches) != 3 {
			return sum, fmt.Errorf("unable to parse input on line %d: %q it must match the regexp %q", i, scanner.Text(), strategyRegexp.String())
		}
		opp, own := shapes[matches[1]], shapes[matches[2]]
		distance := ((own + 3) - opp) % 3
		sum += ((distance+1)%3)*3 + own
	}
	if err := scanner.Err(); err != nil {
		return sum, err
	}

	return sum, nil
}

func oneMod3ByRunes(r io.Reader) (int, error) {
	var sum int

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	const (
		stateOpp = iota
		stateSpace
		stateOwn
		stateLine
	)
	var state = stateOpp
	var opp, own int
	for line, col := 0, 0; scanner.Scan(); line, col = line+1, col+1 {
		rune, _ := utf8.DecodeRune(scanner.Bytes())

		switch state {
		case stateOpp:
			switch rune {
			case 'A', 'B', 'C':
				opp = shapesByRune[rune]
				state = stateSpace
			default:
				return sum, fmt.Errorf("%d:%d: unexpected rune %q, wanted /[ABC]/", line, col, rune)
			}
		case stateSpace:
			if rune != ' ' {
				return sum, fmt.Errorf(`%d:%d: unexpected rune %q, wanted " "`, line, col, rune)
			}
			state = stateOwn
		case stateOwn:
			switch rune {
			case 'X', 'Y', 'Z':
				own = shapesByRune[rune]
				state = stateLine
				distance := ((own + 3) - opp) % 3
				sum += ((distance+1)%3)*3 + own
			default:
				return sum, fmt.Errorf("%d:%d: unexpected rune %q, wanted /[XYZ]/", line, col, rune)
			}
		case stateLine:
			switch rune {
			case '\n':
				state = stateOpp
			default:
				return sum, fmt.Errorf("%d:%d: unexpected rune %q, wanted newline", line, col, rune)
			}
		}

	}
	if err := scanner.Err(); err != nil {
		return sum, err
	}

	return sum, nil
}
