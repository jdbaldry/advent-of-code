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

// See twoWithXor.
func oneWithXor(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var field uint64
	var seen [4]rune
	for i := 0; scanner.Scan(); i++ {
		newest, _ := utf8.DecodeRune(scanner.Bytes())
		seen[i%4] = newest
		field ^= uint64(1 << (int(newest) - int('A')))

		if i < 3 {
			continue
		}

		if countBits(field) == 4 {
			return i + 1, nil
		}

		oldest := seen[(i+1)%4]
		field ^= uint64(1 << (int(oldest) - int('A')))

	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("start of packet marker not found in datastream")
}
