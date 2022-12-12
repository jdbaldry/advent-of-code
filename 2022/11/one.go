package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	monkeyRegexp        = regexp.MustCompile(`Monkey (?P<src>\d):`)
	startingItemsRegexp = regexp.MustCompile(`  Starting items: (?P<items>[^\n]*)`)
	operationRegexp     = regexp.MustCompile(`  Operation: new = old (?P<operator>[*+]) (?P<operand>(?:old|\d+))`)
	testRegexp          = regexp.MustCompile(`  Test: divisible by (?P<divisor>\d+)`)
	ifTrueRegexp        = regexp.MustCompile(`    If true: throw to monkey (?P<ifTrue>\d)`)
	ifFalseRegexp       = regexp.MustCompile(`    If false: throw to monkey (?P<ifFalse>\d)`)
	inputRegexp         = regexp.MustCompile("^" + strings.Join([]string{
		monkeyRegexp.String(),
		startingItemsRegexp.String(),
		operationRegexp.String(),
		testRegexp.String(),
		ifTrueRegexp.String(),
		ifFalseRegexp.String(),
	}, "\n"))

	errDoesNotMatchRegexp = errors.New("does not match regexp")
)

type monkey struct {
	id        int
	items     []int
	operation func(int) int
	test      func(int) bool
	ifTrue    int
	ifFalse   int
}

func (m monkey) String() string {
	var items string

	for i, n := range m.items {
		if i != 0 {
			items += ", "
		}

		items += strconv.Itoa(n)
	}

	return fmt.Sprintf("Monkey %d: %s", m.id, items)
}

func parseMonkey(token string) (monkey, error) {
	matched := inputRegexp.FindStringSubmatch(token)
	if len(matched) != inputRegexp.NumSubexp()+1 {
		return monkey{}, fmt.Errorf("%q %w %q", token, errDoesNotMatchRegexp, inputRegexp)
	}

	//nolint:varnamelen
	id, err := strconv.Atoi(matched[inputRegexp.SubexpIndex("src")])
	if err != nil {
		return monkey{}, fmt.Errorf("%q: unable to parse id: %w", token, err)
	}

	strItems := strings.Split(matched[inputRegexp.SubexpIndex("items")], ", ")
	items := make([]int, len(strItems))

	for i, s := range strItems {
		var err error

		items[i], err = strconv.Atoi(s)
		if err != nil {
			return monkey{}, fmt.Errorf("%q: unable to parse item %d: %w", token, i, err)
		}
	}

	operator := matched[inputRegexp.SubexpIndex("operator")]

	divisor, err := strconv.Atoi(matched[inputRegexp.SubexpIndex("divisor")])
	if err != nil {
		return monkey{}, fmt.Errorf("%q: unable to parse divisor: %w", token, err)
	}

	ifTrue, err := strconv.Atoi(matched[inputRegexp.SubexpIndex("ifTrue")])
	if err != nil {
		return monkey{}, fmt.Errorf("%q: unable to parse ifTrue destination monkey: %w", token, err)
	}

	ifFalse, err := strconv.Atoi(matched[inputRegexp.SubexpIndex("ifFalse")])
	if err != nil {
		return monkey{}, fmt.Errorf("%q: unable to parse ifFalse destination monkey: %w", token, err)
	}

	return monkey{
		id:    id,
		items: items,
		operation: func(i int) int { //nolint:varnamelen
			var operand int

			strOperand := matched[inputRegexp.SubexpIndex("operand")]
			switch strOperand {
			case "old":
				operand = i
			default:
				var err error

				operand, err = strconv.Atoi(strOperand)
				if err != nil {
					panic(fmt.Errorf("%q: unable to parse operand: %w", token, err))
				}
			}

			switch operator {
			case "*":
				return i * operand
			case "+":
				return i + operand
			default:
				panic(fmt.Sprintf("Unrecognized operator %q", operator))
			}
		},
		test: func(i int) bool {
			return i%divisor == 0
		},
		ifTrue:  ifTrue,
		ifFalse: ifFalse,
	}, nil
}

func one(r io.Reader) (int, error) {
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

	for i, rounds := 0, 20; i < rounds; i++ {
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
