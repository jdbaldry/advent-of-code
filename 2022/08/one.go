package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

//nolint:deadcode,unused
func printTrees(f [][]int) {
	for i := range f {
		for j := range f[i] {
			fmt.Printf(" %d ", f[i][j]) //nolint:forbidigo
		}

		fmt.Println() //nolint:forbidigo
	}
}

func isVisible(height int, trees []int) bool {
	for _, t := range trees {
		if t >= height {
			return false
		}
	}

	return true
}

func parseTrees(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	row := []int{}
	trees := [][]int{}

	for scanner.Scan() {
		// Take only the first byte because the input is only
		// the ASCII digits 0-9.
		s := scanner.Text()
		switch s {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			i, err := strconv.Atoi(s)
			if err != nil {
				return trees, err //nolint:wrapcheck
			}

			row = append(row, i)
		case "\n":
			trees = append(trees, row)
			row = []int{}
		}
	}

	if err := scanner.Err(); err != nil {
		return trees, fmt.Errorf("%w during scanning", err)
	}

	return trees, nil
}

//nolint:gocognit,cyclop
func one(r io.Reader) (int, error) {
	trees, err := parseTrees(r)
	if err != nil {
		return 0, err
	}

	var visible int

	for i := range trees { //nolint:varnamelen
		if i == 0 || i == len(trees)-1 {
			continue
		}

		for j, tree := range trees[i] { //nolint:varnamelen
			if j == 0 || j == len(trees[i])-1 {
				continue
			}

			if isVisible(tree, trees[i][:j]) {
				visible++

				continue
			}

			if isVisible(tree, trees[i][j+1:]) {
				visible++

				continue
			}

			vertical := make([]int, len(trees))
			for k := range trees {
				vertical[k] = trees[k][j]
			}

			if isVisible(tree, vertical[:i]) {
				visible++

				continue
			}

			if isVisible(tree, vertical[i+1:]) {
				visible++
			}
		}
	}

	corners := 4
	sides := 2
	onPerimeter := (sides * (len(trees) + len(trees[0]))) - corners

	return visible + onPerimeter, nil
}
