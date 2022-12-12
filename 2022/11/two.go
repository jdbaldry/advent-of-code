package main

import (
	"io"
	"sort"
)

func two(r io.Reader) (int, error) {
	monkeys, err := parseMonkeys(r)
	if err != nil {
		return 0, err
	}

	inspected := make([]int, len(monkeys))

	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.divisor
	}

	for i, rounds := 0, 10000; i < rounds; i++ {
		for j, monkey := range monkeys { //nolint:varnamelen
			for _, item := range monkey.items {
				inspected[j]++

				item = monkey.operation(item)
				item %= lcm

				if monkey.test(item) {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, item)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, item)
				}
			}

			monkeys[j].items = []int{}
		}
	}

	sort.Slice(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})

	return inspected[0] * inspected[1], nil
}
