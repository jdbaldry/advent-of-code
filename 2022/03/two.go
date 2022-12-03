package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

func two(r io.Reader) (int, error) {
	var sum int

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var group [3]uint64
	for line, col := 0, 0; scanner.Scan(); col++ {
		r, _ := utf8.DecodeRune(scanner.Bytes())
		if r == '\n' {
			line, col = line+1, -1
			if line%3 == 0 {
				sum += priorities[group[0]&group[1]&group[2]]
				group = [3]uint64{}
			}
			continue
		}

		field, ok := fields[byte(r)]
		if !ok {
			return sum, fmt.Errorf("%d:%d: unexpected rune %q, wanted newline", line, col, r)
		}

		group[line%3] |= field
	}

	return sum, nil
}
