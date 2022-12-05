package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	instructionRegexp = regexp.MustCompile(`^move ([0-9]+) from ([0-9+]) to ([0-9]+)$`)
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

// parseStacks updates the stacks based up on the text contents of the current line.
// It only supports stacks with cells that are separated by spaces and  represented by the regexp `\[.\]`.
func parseStacks(stacks []*cell, text string) ([]*cell, state, error) {
	if len(text) == 0 {
		return stacks, stateInstructions, nil
	}

	for i, col := 0, 1; i < len(stacks); i, col = i+1, col+4 {
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

func parseInstruction(s string) (count, from, to int, err error) {
	matches := instructionRegexp.FindStringSubmatch(s)
	if len(matches) != 4 {
		return count, from, to, fmt.Errorf("instruction did not match regexp %q", instructionRegexp)
	}
	count, err = strconv.Atoi(matches[1])
	if err != nil {
		return count, from, to, fmt.Errorf("could not parse count %q as integer", matches[1])
	}
	from, err = strconv.Atoi(matches[2])
	if err != nil {
		return count, from, to, fmt.Errorf("could not parse from %q as integer", matches[2])
	}
	to, err = strconv.Atoi(matches[3])
	if err != nil {
		return count, from, to, fmt.Errorf("could not parse to %q as integer", matches[3])
	}

	return count, from, to, nil

}

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
		fmt.Println(str.String())
	}
}

func one(r io.Reader) (string, error) {
	var message string

	scanner := bufio.NewScanner(r)

	var s state
	var stacks []*cell
	for line := 0; scanner.Scan(); line++ {
		text := scanner.Text()
		switch s {
		case stateDrawing:
			if line == 0 {
				stacks = make([]*cell, len(text)/4+1)
			}
			var err error
			stacks, s, err = parseStacks(stacks, text)
			if err != nil {
				return message, err
			}

		case stateInstructions:
			count, from, to, err := parseInstruction(text)
			if err != nil {
				return message, fmt.Errorf("%d: %v", line, err)
			}

			// Instructions are one-indexed by stacks are zero-indexed.
			from--
			to--

			for i := 0; i < count; i++ {
				if stacks[from] == nil {
					break
				}
				crate := stacks[from]
				stacks[from] = crate.cdr
				crate.cdr = stacks[to]
				stacks[to] = crate
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return message, err
	}

	for i := 0; i < len(stacks); i++ {
		if stacks[i] == nil {
			continue
		}
		message += string(stacks[i].car)
	}
	return message, nil
}
