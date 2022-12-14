package main

import (
	"io"
)

func two(r io.Reader) (int, error) {
	rockMap := map[coord]struct{}{}

	if err := drawMap(r, rockMap); err != nil {
		return 0, err
	}

	var (
		bottom = lowestPoint(rockMap) + 2
		count  = 0
		sand   = coord{500, 0}
	)

	for _, ok := rockMap[sand]; !ok; {
		next := coord{sand.x, sand.y + 1}

		if next.y == bottom {
			rockMap[sand] = struct{}{}
			sand = coord{500, 0}
			count++

			continue
		}

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

		if sand.x == 500 && sand.y == 0 {
			count++

			break
		}

		rockMap[sand] = struct{}{}
		sand = coord{500, 0}

		count++
	}

	return count, nil
}
