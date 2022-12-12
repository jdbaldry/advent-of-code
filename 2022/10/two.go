package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var errDrawFailed = errors.New("failed to draw")

func draw(screen *strings.Builder, cycle, register int) error {
	width := 40
	pos := cycle % width

	if pos == 0 && cycle > 0 {
		if _, err := screen.WriteRune('\n'); err != nil {
			return errDrawFailed
		}
	}

	switch pos {
	case register - 1, register, register + 1:
		if _, err := screen.WriteRune('#'); err != nil {
			return errDrawFailed
		}
	default:
		if _, err := screen.WriteRune('.'); err != nil {
			return errDrawFailed
		}
	}

	return nil
}

func two(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	var (
		cycle    int
		register = 1
		screen   = &strings.Builder{}
	)

	for scanner.Scan() {
		line := scanner.Text()

		matched := instructionRegexp.FindStringSubmatch(line)
		if len(matched) != instructionRegexp.NumSubexp()+1 {
			return screen.String(), fmt.Errorf("%q line %w %q", line, errDidNotMatchRegexp, instructionRegexp)
		}

		switch matched[1] {
		case "noop":
			if err := draw(screen, cycle, register); err != nil {
				return screen.String(), err
			}
			cycle++

		case "addx":
			if err := draw(screen, cycle, register); err != nil {
				return screen.String(), err
			}
			cycle++

			if err := draw(screen, cycle, register); err != nil {
				return screen.String(), err
			}

			cycle++

			x, err := strconv.Atoi(matched[2])
			if err != nil {
				return screen.String(), fmt.Errorf("unable to parse argument to addx instruction %q: %w", matched[0], err)
			}

			register = addX(register, x)
		default:
			return screen.String(), fmt.Errorf("%q %w %q", line, errUnknownInstruction, matched[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return screen.String(), fmt.Errorf("%w during scanning", err)
	}

	return screen.String(), nil
}
