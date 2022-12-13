package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var (
	coordRegexp = regexp.MustCompile(`^(?P<x>\d+),(?P<y>\d+)$`)
	foldRegexp  = regexp.MustCompile(`^fold along (?P<axis>[xy])=(?P<line>\d+)$`)

	errUnrecognizedState  = errors.New("unrecognized state")
	errDoesNotMatchRegexp = errors.New("does not match regexp")
)

type coord struct {
	x int
	y int
}

func (c coord) String() string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

func mirror(n, line int) int {
	if n < line {
		return n
	}

	return line - (n - line)
}

func fold(axis string, line int, coords map[coord]struct{}) {
	for c := range coords { //nolint:varnamelen
		delete(coords, c)

		switch axis {
		case "x":
			coords[coord{mirror(c.x, line), c.y}] = struct{}{}
		case "y":
			coords[coord{c.x, mirror(c.y, line)}] = struct{}{}
		}
	}
}

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	const (
		stateCoords = iota
		stateFolds
	)

	var (
		coords = make(map[coord]struct{})
		state  int
	)

scan:
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			state = stateFolds

			continue
		}

		switch state {
		case stateCoords:
			matched := coordRegexp.FindStringSubmatch(line)
			if len(matched) != coordRegexp.NumSubexp()+1 {
				return 0, fmt.Errorf("%q %w %q", line, errDoesNotMatchRegexp, coordRegexp)
			}

			var x, y int //nolint:varnamelen

			x, err := strconv.Atoi(matched[coordRegexp.SubexpIndex("x")])
			if err != nil {
				return 0, fmt.Errorf("unable to parse x coordinate: %w", err)
			}

			y, err = strconv.Atoi(matched[coordRegexp.SubexpIndex("y")])
			if err != nil {
				return 0, fmt.Errorf("unable to parse y coordinate: %w", err)
			}

			coords[coord{x, y}] = struct{}{}
		case stateFolds:
			matched := foldRegexp.FindStringSubmatch(line)
			if len(matched) != foldRegexp.NumSubexp()+1 {
				return 0, fmt.Errorf("%q %w %q", line, errDoesNotMatchRegexp, foldRegexp)
			}

			line, err := strconv.Atoi(matched[foldRegexp.SubexpIndex("line")])
			if err != nil {
				return 0, fmt.Errorf("unable to parse fold line: %w", err)
			}

			fold(matched[foldRegexp.SubexpIndex("axis")], line, coords)
			break scan // only performing first fold
		default:
			return 0, fmt.Errorf("%q %w", state, errUnrecognizedState)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return len(coords), nil
}
