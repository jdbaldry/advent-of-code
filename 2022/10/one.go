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
	instructionRegexp     = regexp.MustCompile(`^(noop|addx)(?: (-?\d+))?`)
	errDidNotMatchRegexp  = errors.New("does not match regexp")
	errUnknownInstruction = errors.New("unknown instruction")
)

func signalStrength(cycle, register int) int {
	if first, period := 20, 40; cycle%period == first {
		return cycle * register
	}

	return 0
}

func tick(cycle, register, strength, sum int) (int, int, int, int) {
	cycle++

	strength = signalStrength(cycle, register)
	if strength != 0 {
		sum += strength
	}

	return cycle, register, strength, sum
}

func addX(register, x int) int {
	return register + x
}

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var (
		cycle    int
		strength int
		sum      int

		register = 1
	)

	for scanner.Scan() {
		line := scanner.Text()

		matched := instructionRegexp.FindStringSubmatch(line)
		if len(matched) != instructionRegexp.NumSubexp()+1 {
			return sum, fmt.Errorf("%q line %w %q", line, errDidNotMatchRegexp, instructionRegexp)
		}

		switch matched[1] {
		case "noop":
			cycle, register, strength, sum = tick(cycle, register, strength, sum)
		case "addx":
			cycle, register, strength, sum = tick(cycle, register, strength, sum)
			cycle, register, strength, sum = tick(cycle, register, strength, sum)

			x, err := strconv.Atoi(matched[2])
			if err != nil {
				return sum, fmt.Errorf("unable to parse argument to addx instruction %q: %w", matched[0], err)
			}

			register = addX(register, x)
		default:
			return sum, fmt.Errorf("%q %w %q", line, errUnknownInstruction, matched[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return sum, nil
}
