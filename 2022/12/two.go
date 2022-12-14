package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"unicode/utf8"
)

func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		starts []*node
		end    *node

		grid [][]*node
		line []*node
	)

	//nolint:varnamelen
	for y, x := 0, 0; scanner.Scan(); x++ {
		char, _ := utf8.DecodeRune(scanner.Bytes())

		if char == '\n' {
			grid = append(grid, line)
			line = []*node{}
			y, x = y+1, -1

			continue
		}

		node := newNode([2]int{x, y}, 0)
		line = append(line, node)

		switch char {
		case 'S', 'a':
			starts = append(starts, node)
		case 'E':
			node.height = int('z') - int('a')
			end = node
		default:
			node.height = int(char) - int('a')
		}

		if y > 0 {
			above := grid[y-1][x]
			node.addIf(above, oneStepUp)
			above.addIf(node, oneStepUp)
		}

		if x > 0 {
			left := line[x-1]
			node.addIf(left, oneStepUp)
			left.addIf(node, oneStepUp)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	shortest := math.MaxInt

	for _, start := range starts {
		distance, err := bfs(start, func(n *node) bool { return n == end })
		if err != nil {
			if errors.Is(err, errNoRoute) {
				continue
			}

			return shortest, err
		}

		if distance < shortest {
			shortest = distance
		}
	}

	return shortest, nil
}

func twoFromE(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		start *node

		grid [][]*node
		line []*node
	)

	//nolint:varnamelen
	for y, x := 0, 0; scanner.Scan(); x++ {
		char, _ := utf8.DecodeRune(scanner.Bytes())

		if char == '\n' {
			grid = append(grid, line)
			line = []*node{}
			y, x = y+1, -1

			continue
		}

		node := newNode([2]int{x, y}, 0)
		line = append(line, node)

		switch char {
		case 'S', 'a':
			// Do nothing, default height (0) is correct.
		case 'E':
			node.height = int('z') - int('a')
			start = node
		default:
			node.height = int(char) - int('a')
		}

		if y > 0 {
			above := grid[y-1][x]
			node.addIf(above, oneStepDown)
			above.addIf(node, oneStepDown)
		}

		if x > 0 {
			left := line[x-1]
			node.addIf(left, oneStepDown)
			left.addIf(node, oneStepDown)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return bfs(start, func(n *node) bool { return n.height == 0 })
}
