package main

import (
	"bufio"
	"io"
)

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
