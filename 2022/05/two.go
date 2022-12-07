package main

import (
	"bufio"
	"fmt"
	"io"
)

//nolint:cyclop
func two(r io.Reader) (string, error) {
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
			start, end = start-1, end-1

			var stack, last *cell

			//nolint:varnamelen
			for i := 0; i < count; i++ {
				if stacks[start] == nil {
					break
				}

				crate := stacks[start]
				stacks[start] = crate.cdr
				last = crate

				if i == 0 {
					stack = last
				}
			}

			last.cdr = stacks[end]
			stacks[end] = stack
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("%w during scanning", err)
	}

	return message(stacks), nil
}
