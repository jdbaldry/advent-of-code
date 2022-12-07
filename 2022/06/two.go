package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		seen                [14]rune
		startOfPacketLength = 14
	)

	//nolint:varnamelen
	for i := 0; scanner.Scan(); i++ {
		seen[i%startOfPacketLength], _ = utf8.DecodeRune(scanner.Bytes())

		if i < startOfPacketLength-1 {
			continue
		}

		var field uint64

		for _, r := range seen {
			shift := int(r) - int('A')
			shifted := uint64(1 << shift)
			field |= shifted
		}

		if countBits(field) == uint64(startOfPacketLength) {
			return i + 1, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return 0, errStartOfPacketNotFound
}

func twoWithMap(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		seen                [14]rune
		startOfPacketLength = 14
		counts              = make(map[rune]int)
	)

	//nolint:varnamelen
	for i := 0; scanner.Scan(); i++ {
		newest, _ := utf8.DecodeRune(scanner.Bytes())
		seen[i%startOfPacketLength] = newest
		counts[newest]++

		if i < startOfPacketLength-1 {
			continue
		}

		if len(counts) == startOfPacketLength {
			return i + 1, nil
		}

		oldest := seen[(i+1)%startOfPacketLength]
		counts[oldest]--

		if counts[oldest] == 0 {
			delete(counts, oldest)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return 0, errStartOfPacketNotFound
}

func twoWithXor(r io.Reader) (int, error) {
	startOfPacketLength := 14

	return withXor(r, startOfPacketLength)
}
