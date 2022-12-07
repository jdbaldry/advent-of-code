package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

var errStartOfPacketNotFound = errors.New("start of packet marker not found in datastream")

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		seen                [3]rune
		startOfPacketLength = 4
	)

	//nolint:varnamelen
	for i := 0; scanner.Scan(); i++ {
		char, _ := utf8.DecodeRune(scanner.Bytes())

		if i < startOfPacketLength-1 {
			seen[i] = char

			continue
		}

		if (char != seen[0] && char != seen[1] && char != seen[2]) &&
			(seen[0] != seen[1] && seen[0] != seen[2] && seen[1] != seen[2]) {
			return i + 1, nil
		}

		seen[0], seen[1], seen[2] = seen[1], seen[2], char
	}

	return 0, errStartOfPacketNotFound
}

// https://www.reddit.com/r/adventofcode/comments/zdw0u6/comment/iz4l6xk
// If a letter is repeated an even number of times in the window xor will turn the bit off.
// If a letter is repeated an odd number of times and n > 1 the bit will be on,
// but will take up 3+ "slots" to only add 1 to the count.
func withXor(r io.Reader, startOfPacketLength int) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		field uint64
		seen  = make([]rune, startOfPacketLength)
	)

	//nolint:varnamelen
	for i := 0; scanner.Scan(); i++ {
		newest, _ := utf8.DecodeRune(scanner.Bytes())
		seen[i%startOfPacketLength] = newest
		field ^= uint64(1 << (int(newest) - int('A')))

		if i < startOfPacketLength-1 {
			continue
		}

		if countBits(field) == uint64(startOfPacketLength) {
			return i + 1, nil
		}

		oldest := seen[(i+1)%startOfPacketLength]
		field ^= uint64(1 << (int(oldest) - int('A')))
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return 0, errStartOfPacketNotFound
}

func oneWithXor(r io.Reader) (int, error) {
	startOfPacketLength := 4

	return withXor(r, startOfPacketLength)
}
