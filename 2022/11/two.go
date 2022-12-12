package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var (
		monkeys []monkey
		token   string
	)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			token += line + "\n"

			continue
		}

		monkey, err := parseMonkey(token)
		if err != nil {
			return 0, err
		}

		monkeys = append(monkeys, monkey)
		token = ""
	}

	monkey, err := parseMonkey(token)
	if err != nil {
		return 0, err
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	monkeys = append(monkeys, monkey)
	inspected := make([]int, len(monkeys))

	for i, rounds := 0, 10000; i < rounds; i++ {
		for j, monkey := range monkeys { //nolint:varnamelen
			for _, item := range monkey.items {
				inspected[j]++

				item = monkey.operation(item)
				item /= 3

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
