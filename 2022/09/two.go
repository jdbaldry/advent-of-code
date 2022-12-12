package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	knots := 10
	rope := make([]coord, knots)
	seen := map[coord]struct{}{{0, 0}: {}}

	for scanner.Scan() {
		line := scanner.Text()

		matched := motionRegexp.FindStringSubmatch(line)
		if len(matched) != motionRegexp.NumSubexp()+1 {
			return 0, fmt.Errorf("%q %w %q", line, errDoesNotMatchRegexp, motionRegexp)
		}

		distance, err := strconv.Atoi(matched[2])
		if err != nil {
			return 0, fmt.Errorf("could not parse distance: %w", err)
		}

		direction := matched[1]

		rope, err = move(seen, rope, direction, distance)
		if err != nil {
			return 0, err
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return len(seen), nil
}
