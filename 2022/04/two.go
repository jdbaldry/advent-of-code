package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func two(r io.Reader) (int, error) {
	var sum int
	scanner := bufio.NewScanner(r)

	for line := 0; scanner.Scan(); line++ {
		ranges := strings.Split(scanner.Text(), ",")
		if len(ranges) != 2 {
			return sum, fmt.Errorf("%d: expected exactly two ranges per line but found %d in %v", line, len(ranges), ranges)
		}

		var parsed [2][2]int
		for i, r := range ranges {
			bounds := strings.Split(r, "-")
			if len(bounds) != 2 {
				return sum, fmt.Errorf("%d: expected exactly two bounds per range but found %d in %v", line, len(bounds), bounds)
			}

			for j, b := range bounds {
				int, err := strconv.Atoi(b)
				if err != nil {
					return sum, fmt.Errorf("%d: unable to parse  bound as integer but on range %d, bound %d, %q: %v", line, i, j, b, err)
				}

				parsed[i][j] = int
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
