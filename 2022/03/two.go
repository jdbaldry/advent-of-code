package main

import (
	"bufio"
	"io"
	"unicode/utf8"
)

func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		group [3]uint64
		sum   int
	)

	for line, col := 0, 0; scanner.Scan(); col++ {
		char, _ := utf8.DecodeRune(scanner.Bytes())
		if char == '\n' {
			line, col = line+1, -1
			if line%3 == 0 {
				priority := group[0] & group[1] & group[2]
				for i := 1; i <= 64; i++ {
					if priority == 1 {
						sum += i

						break
					}

					priority >>= 1
				}

				group = [3]uint64{}
			}

			continue
		}

		shift := int(char) - int('a')
		if shift < 0 {
			shift += 58
		}

		group[line%3] |= uint64(1 << shift)
	}

	return sum, nil
}
