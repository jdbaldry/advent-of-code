package main

import (
	"bufio"
	"io"
)

// one returns the set intersection of the two parts of a line using a bit field.
// 58 is the magic number needed to wrap the upper case ASCII alphabet around the lower case.
// Priorities are a=1,b=2,...;z=26;A=27;B=28;Z=52.
func one(r io.Reader) (int, error) {
	var sum int

	var left, right uint64
	scanner := bufio.NewScanner(r)
	for line := 0; scanner.Scan(); line++ {
		text := scanner.Text()
		for i := 0; i < len(text)/2; i++ {
			shift := int(text[i]) - int('a')
			if shift < 0 {
				shift += 58
			}
			left |= uint64(1 << shift)
			shift = int(text[len(text)-1-i]) - int('a')
			if shift < 0 {
				shift += 58
			}
			right |= uint64(1 << shift)
		}

		priority := left & right
		for i := 1; i <= 64; i++ {
			if priority == 1 {
				sum += i
				break
			}
			priority >>= 1
		}

		left, right = 0, 0
	}

	return sum, nil
}
