package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//nolint:gocognit,cyclop
func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var sum int

	for line := 0; scanner.Scan(); line++ {
		ranges := strings.Split(scanner.Text(), ",")
		if expectedRanges := 2; len(ranges) != expectedRanges {
			return sum, fmt.Errorf("%d: %w but found %d in %v", line, errInvalidBounds, len(ranges), ranges)
		}

		var parsed [2][2]int

		//nolint:varnamelen
		for i, char := range ranges {
			bounds := strings.Split(char, "-")
			if expectedBounds := 2; len(bounds) != expectedBounds {
				return sum, fmt.Errorf("%d: %w but found %d in %v", line, errInvalidRanges, len(bounds), bounds)
			}

			//nolint:varnamelen
			for j, b := range bounds {
				integer, err := strconv.Atoi(b)
				if err != nil {
					return sum, fmt.Errorf("%d: %w for range %d, bound %d, %q: %v", line, errInvalidBounds, i, j, b, err)
				}

				parsed[i][j] = integer
			}
		}

		if (parsed[0][0] >= parsed[1][0] && parsed[0][0] <= parsed[1][1]) ||
			(parsed[0][1] >= parsed[1][0] && parsed[0][1] <= parsed[1][1]) ||
			(parsed[1][0] >= parsed[0][0] && parsed[1][1] <= parsed[0][1]) ||
			(parsed[1][1] >= parsed[0][0] && parsed[1][1] <= parsed[0][1]) {
			sum++
		}
	}

	return sum, nil
}
