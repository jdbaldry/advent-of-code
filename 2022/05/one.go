package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	instructionRegexp               = regexp.MustCompile(`^move ([0-9]+) from ([0-9+]) to ([0-9]+)$`)
	errInstructionDidNotMatchRegexp = errors.New("instruction did not match regexp")
)

type cell struct {
	car rune
	cdr *cell
}

type state int

const (
	stateDrawing state = iota
	stateInstructions
)

func message(stacks []*cell) string {
	message := ""

	for i := 0; i < len(stacks); i++ {
		if stacks[i] == nil {
			continue
		}

		message += string(stacks[i].car)
	}

	return message
}

// parseStacks updates the stacks based up on the text contents of the current line.
// It only supports stacks with cells that are separated by spaces and  represented by the regexp `\[.\]`.
// An error value is returned for forwards compatibility but presently it is always nil.
//
//nolint:unparam
func parseStacks(stacks []*cell, text string) ([]*cell, state, error) {
	if len(text) == 0 {
		return stacks, stateInstructions, nil
	}

	colWidth := 4

	//nolint:varnamelen
	for i, col := 0, 1; i < len(stacks); i, col = i+1, col+colWidth {
		r := text[col]
		switch r {
		case ' ', '\n', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			if stacks[i] == nil {
				stacks[i] = &cell{rune(text[col]), nil}

				continue
			}

			var c *cell
			for c = stacks[i]; c.cdr != nil; c = c.cdr {
			}

			c.cdr = &cell{rune(text[col]), nil}
		}
	}

	return stacks, stateDrawing, nil
}

func parseInstruction(s string) (int, int, int, error) {
	var count, start, end int

	matches := instructionRegexp.FindStringSubmatch(s)
	if len(matches) != instructionRegexp.NumSubexp()+1 {
		return count, start, end, fmt.Errorf("%q %w %q", s, errInstructionDidNotMatchRegexp, instructionRegexp)
	}

	count, err := strconv.Atoi(matches[1])
	if err != nil {
		return count, start, end, fmt.Errorf("could not parse count: %w", err)
	}

	start, err = strconv.Atoi(matches[2])
	if err != nil {
		return count, start, end, fmt.Errorf("could not parse start: %w", err)
	}

	end, err = strconv.Atoi(matches[3])
	if err != nil {
		return count, start, end, fmt.Errorf("could not parse end: %w", err)
	}

	return count, start, end, nil
}

//nolint:unused,deadcode
func printStacks(stacks []*cell) {
	for i := range stacks {
		var str strings.Builder

		str.WriteString(fmt.Sprintf("%d", i))

		crate := stacks[i]
		if crate == nil {
			continue
		}

		stack := []rune{}
		for ; crate != nil; crate = crate.cdr {
			stack = append([]rune{crate.car}, stack...)
		}

		for _, crate := range stack {
			str.WriteString(fmt.Sprintf(" [%c]", crate))
		}

		//nolint:forbidigo
		fmt.Println(str.String())
	}
}

func one(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	var (
		state  state
		stacks []*cell
	)

	for line := 0; scanner.Scan(); line++ {
		text := scanner.Text()

		switch state {
		case stateDrawing:
			if line == 0 {
				stacks = make([]*cell, len(text)/4+1)
			}

			var err error

			stacks, state, err = parseStacks(stacks, text)
			if err != nil {
				return "", err
			}

		case stateInstructions:
			count, start, end, err := parseInstruction(text)
			if err != nil {
				return "", fmt.Errorf("%d: %w", line, err)
			}

			// Instructions are one-indexed by stacks are zero-indexed.
			start--
			end--

			for i := 0; i < count; i++ {
				if stacks[start] == nil {
					break
				}

				crate := stacks[start]

				stacks[start] = crate.cdr
				crate.cdr = stacks[end]
				stacks[end] = crate
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("%w during scanning", err)
	}

	return message(stacks), nil
}
