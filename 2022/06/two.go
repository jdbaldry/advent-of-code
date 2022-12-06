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

	var seen [14]rune
	for i := 0; scanner.Scan(); i++ {
		seen[i%14], _ = utf8.DecodeRune(scanner.Bytes())
		if i < 13 {
			continue
		}

		var field uint64
		for _, r := range seen {
			shift := int(r) - int('A')
			shifted := uint64(1 << shift)
			field |= shifted
		}

		if countBits(field) == 14 {
			return i + 1, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("start of packet marker not found in datastream")
}

func twoWithMap(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var seen [14]rune
	var counts = make(map[rune]int)
	for i := 0; scanner.Scan(); i++ {
		newest, _ := utf8.DecodeRune(scanner.Bytes())
		seen[i%14] = newest
		counts[newest]++

		if i < 13 {
			continue
		}

		oldest := seen[(i+1)%14]
		if len(counts) == 14 {
			return i + 1, nil
		}
		counts[oldest]--
		if counts[oldest] == 0 {
			delete(counts, oldest)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("start of packet marker not found in datastream")
}

// https://www.reddit.com/r/adventofcode/comments/zdw0u6/comment/iz4l6xk/?utm%255C_source=reddit&utm%255C_medium=web2x&context=3&utm_source=reddit&utm_medium=usertext&utm_name=adventofcode&utm_content=t1_iz4o78a
// If a letter is repeated an even number of times in the window xor will turn the bit off.
// If a letter is repeated an odd number of times and n > 1 the bit will be on, but will take up 3+ "slots" to only add 1 to the count.
func twoWithXor(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var field uint64
	var seen [14]rune
	for i := 0; scanner.Scan(); i++ {
		newest, _ := utf8.DecodeRune(scanner.Bytes())
		seen[i%14] = newest
		field ^= uint64(1 << (int(newest) - int('A')))

		if i < 13 {
			continue
		}

		if countBits(field) == 14 {
			return i + 1, nil
		}
		oldest := seen[(i+1)%14]
		field ^= uint64(1 << (int(oldest) - int('A')))
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("start of packet marker not found in datastream")

}
