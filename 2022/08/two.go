package main

import (
	"io"
)

func scenicScore(tree int, left, right, up, down []int) int {
	var viewLeft int

	for i := len(left) - 1; i >= 0; i-- {
		viewLeft++

		if tree <= left[i] {
			break
		}
	}

	var viewRight int

	for i := 0; i < len(right); i++ {
		viewRight++

		if tree <= right[i] {
			break
		}
	}

	var viewUp int

	for i := len(up) - 1; i >= 0; i-- {
		viewUp++

		if tree <= up[i] {
			break
		}
	}

	var viewDown int

	for i := 0; i < len(down); i++ {
		viewDown++

		if tree <= down[i] {
			break
		}
	}

	return viewLeft * viewRight * viewUp * viewDown
}

func two(r io.Reader) (int, error) {
	trees, err := parseTrees(r)
	if err != nil {
		return 0, err
	}

	var highest int

	for i := range trees { //nolint:varnamelen
		if i <= 0 || i == len(trees)-1 {
			continue
		}

		for j, tree := range trees[i] { //nolint:varnamelen
			if j == 0 || j == len(trees[i])-1 {
				continue
			}

			vertical := make([]int, len(trees))
			for k := range trees {
				vertical[k] = trees[k][j]
			}

			score := scenicScore(tree, trees[i][:j], trees[i][j+1:], vertical[:i], vertical[i+1:])
			if score > highest {
				highest = score
			}
		}
	}

	return highest, nil
}
