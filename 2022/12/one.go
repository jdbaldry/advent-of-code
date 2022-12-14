package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

var errNoRoute = errors.New("no route found")

type node struct {
	coord  [2]int
	height int

	edges map[*node]struct{}
}

func newNode(coord [2]int, height int, edges ...*node) *node {
	node := &node{
		coord:  coord,
		height: height,
		edges:  make(map[*node]struct{}),
	}

	for _, edge := range edges {
		node.add(edge)
	}

	return node
}

// add adds an edge if it is reachable.
func (n *node) add(m *node) {
	if m.height-n.height <= 1 {
		n.edges[m] = struct{}{}
	}
}

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		start *node
		end   *node

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
		case 'S':
			start = node
		case 'E':
			node.height = int('z') - int('a')
			end = node
		default:
			node.height = int(char) - int('a')
		}

		if y > 0 {
			above := grid[y-1][x]
			node.add(above)
			above.add(node)
		}

		if x > 0 {
			left := line[x-1]
			node.add(left)
			left.add(node)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return bfs(start, end)
}

func bfs(start, end *node) (int, error) {
	distance := map[*node]int{
		start: 0,
	}

	for queue := []*node{start}; len(queue) != 0; {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return distance[curr], nil
		}

		for edge := range curr.edges {
			_, ok := distance[edge]
			if ok {
				continue
			}

			queue = append(queue, edge)
			distance[edge] = distance[curr] + 1
		}
	}

	return 0, errNoRoute
}
