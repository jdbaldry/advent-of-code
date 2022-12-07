package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	errInvalidRanges      = errors.New("expected exactly two ranges per line")
	errInvalidBounds      = errors.New("expected exactly two bounds per range")
	errUnableToParseBound = errors.New("unable to parse bound as integer")
)

//nolint:cyclop
func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var sum int

	for line := 0; scanner.Scan(); line++ {
		ranges := strings.Split(scanner.Text(), ",")
		if expectedRanges := 2; len(ranges) != expectedRanges {
			return sum, fmt.Errorf("%d: %w but found %d in %v", line, errInvalidRanges, len(ranges), ranges)
		}

		var parsed [2][2]int

		//nolint:varnamelen
		for i, r := range ranges {
			bounds := strings.Split(r, "-")
			if expectedBounds := 2; len(bounds) != expectedBounds {
				return sum, fmt.Errorf("%d: %w but found %d in %v", line, errInvalidBounds, len(bounds), bounds)
			}

			//nolint:varnamelen
			for j, b := range bounds {
				integer, err := strconv.Atoi(b)
				if err != nil {
					return sum, fmt.Errorf("%d: %w for range %d, bound %d, %q: %v", line, errUnableToParseBound, i, j, b, err)
				}

				parsed[i][j] = integer
			}
		}

		if (parsed[0][0] >= parsed[1][0] && parsed[0][1] <= parsed[1][1]) ||
			(parsed[1][0] >= parsed[0][0] && parsed[1][1] <= parsed[0][1]) {
			sum++
		}

		parsed = [2][2]int{}
	}

	return sum, nil
}
