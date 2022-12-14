package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var (
	errNotHorizontalOrVertical = errors.New("line is not horizontal or vertical")
	coordRegexp                = regexp.MustCompile(`(\d+),(\d+)`)
)

type coord struct {
	x, y int
}

//nolint:varnamelen
func drawLine(m map[coord]struct{}, start, end coord) error {
	switch {
	case start.x == end.x:
		if start.y-end.y < 0 {
			for y := start.y; y <= end.y; y++ {
				m[coord{start.x, y}] = struct{}{}
			}
		} else {
			for y := end.y; y <= start.y; y++ {
				m[coord{start.x, y}] = struct{}{}
			}
		}
	case start.y == end.y:
		if start.x-end.x < 0 {
			for x := start.x; x <= end.x; x++ {
				m[coord{x, start.y}] = struct{}{}
			}
		} else {
			for x := end.x; x <= start.x; x++ {
				m[coord{x, start.y}] = struct{}{}
			}
		}
	default:
		return errNotHorizontalOrVertical
	}

	return nil
}

func drawMap(r io.Reader, rockMap map[coord]struct{}) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		matched := coordRegexp.FindAllStringSubmatch(scanner.Text(), -1)

		var prev coord

		for i, match := range matched { //nolint:varnamelen
			x, err := strconv.Atoi(match[1]) //nolint:varnamelen
			if err != nil {
				return fmt.Errorf("unable to parse %q as x coordinate: %w", match[0], err)
			}

			y, err := strconv.Atoi(match[2]) //nolint:varnamelen
			if err != nil {
				return fmt.Errorf("unable to parse %q as y coordinate: %w", match[1], err)
			}

			if i == 0 {
				prev = coord{x, y}

				continue
			}

			curr := coord{x, y}

			if err := drawLine(rockMap, prev, curr); err != nil {
				return err
			}

			prev = curr
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%w during scanning", err)
	}

	return nil
}

func lowestPoint(rockMap map[coord]struct{}) int {
	var bottom int

	for rock := range rockMap {
		if rock.y > bottom {
			bottom = rock.y
		}
	}

	return bottom
}

func one(r io.Reader) (int, error) {
	rockMap := map[coord]struct{}{}

	if err := drawMap(r, rockMap); err != nil {
		return 0, err
	}

	var (
		bottom = lowestPoint(rockMap)
		count  = 0
		sand   = coord{500, 0}
	)

	for _, ok := rockMap[sand]; !ok && sand.y <= bottom; {
		next := coord{sand.x, sand.y + 1}

		if _, ok := rockMap[next]; !ok {
			sand = next

			continue
		}

		next = coord{sand.x - 1, sand.y + 1}
		if _, ok := rockMap[next]; !ok {
			sand = next

			continue
		}

		next = coord{sand.x + 1, sand.y + 1}
		if _, ok := rockMap[next]; !ok {
			sand = next

			continue
		}

		rockMap[sand] = struct{}{}
		sand = coord{500, 0}

		count++
	}

	return count, nil
}
