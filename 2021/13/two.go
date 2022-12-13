package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func two(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	const (
		stateCoords = iota
		stateFolds
	)

	var (
		coords = make(map[coord]struct{})
		state  int
	)

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
				return "", fmt.Errorf("%q %w %q", line, errDoesNotMatchRegexp, coordRegexp)
			}

			var x, y int //nolint:varnamelen

			x, err := strconv.Atoi(matched[coordRegexp.SubexpIndex("x")])
			if err != nil {
				return "", fmt.Errorf("unable to parse x coordinate: %w", err)
			}

			y, err = strconv.Atoi(matched[coordRegexp.SubexpIndex("y")])
			if err != nil {
				return "", fmt.Errorf("unable to parse y coordinate: %w", err)
			}

			coords[coord{x, y}] = struct{}{}
		case stateFolds:
			matched := foldRegexp.FindStringSubmatch(line)
			if len(matched) != foldRegexp.NumSubexp()+1 {
				return "", fmt.Errorf("%q %w %q", line, errDoesNotMatchRegexp, foldRegexp)
			}

			line, err := strconv.Atoi(matched[foldRegexp.SubexpIndex("line")])
			if err != nil {
				return "", fmt.Errorf("unable to parse fold line: %w", err)
			}

			fold(matched[foldRegexp.SubexpIndex("axis")], line, coords)
		default:
			return "", fmt.Errorf("%q %w", state, errUnrecognizedState)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("%w during scanning", err)
	}

	var x, y int
	for coord := range coords {
		if coord.x > x {
			x = coord.x
		}

		if coord.y > y {
			y = coord.y
		}
	}

	str := strings.Builder{}
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			if _, ok := coords[coord{j, i}]; ok {
				str.WriteString("#")
			} else {
				str.WriteString(".")
			}
		}
		str.WriteString("\n")
	}

	return str.String(), nil
}
