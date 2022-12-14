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
	n := &node{
		coord:  coord,
		height: height,
		edges:  make(map[*node]struct{}),
	}

	for _, edge := range edges {
		n.addIf(edge, func(*node, *node) bool { return true })
	}

	return n
}

// addIf adds an edge if it is reachable as determined by fn.
func (n *node) addIf(m *node, fn func(n, m *node) bool) {
	if fn(n, m) {
		n.edges[m] = struct{}{}
	}
}

func oneStepUp(n, m *node) bool {
	return m.height-n.height <= 1
}

func oneStepDown(n, m *node) bool {
	return n.height-m.height <= 1
}

//nolint:unused,deadcode,forbidigo
func printNode(n *node, srcFile string, distance int) {
	fmt.Printf("%s:%d:%d: (%c, %d)\n", srcFile, n.coord[1]+1, n.coord[0]+1, rune(n.height+int('a')), distance)
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

	return bfs(start, func(n *node) bool { return n == end })
}

func bfs(start *node, end func(*node) bool) (int, error) {
	distance := map[*node]int{
		start: 0,
	}

	for queue := []*node{start}; len(queue) != 0; {
		curr := queue[0]
		queue = queue[1:]

		if end(curr) {
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
