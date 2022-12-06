package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var seen [3]rune
	for i := 0; scanner.Scan(); i++ {
		r, _ := utf8.DecodeRune(scanner.Bytes())
		switch i {
		case 0, 1, 2:
			seen[i] = r
		default:
			if (r != seen[0] && r != seen[1] && r != seen[2]) &&
				(seen[0] != seen[1] && seen[0] != seen[2] && seen[1] != seen[2]) {
				return i + 1, nil
			}
			seen[0], seen[1], seen[2] = seen[1], seen[2], r
		}
	}

	return 0, fmt.Errorf("start of packet marker not found in datastream")
}

func oneUsingBits(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var seen [4]rune
	for i := 0; scanner.Scan(); i++ {
		seen[i%4], _ = utf8.DecodeRune(scanner.Bytes())
		if i < 3 {
			continue
		}

		var field uint64
		for _, r := range seen {
			shift := int(r) - int('A')
			shifted := uint64(1 << shift)
			field |= shifted
		}

		if countBits(field) == 4 {
			return i + 1, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("start of packet marker not found in datastream")
}
