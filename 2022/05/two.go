package main

import (
	"bufio"
	"fmt"
	"io"
)

func two(r io.Reader) (string, error) {
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

			var stack, last *cell
			for i := 0; i < count; i++ {
				if stacks[from] == nil {
					break
				}
				crate := stacks[from]
				stacks[from] = crate.cdr
				last = crate
				if i == 0 {
					stack = last
				}
			}
			last.cdr = stacks[to]
			stacks[to] = stack
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
