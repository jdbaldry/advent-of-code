package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	errLineDoesNotMatchRegexp = errors.New("line does not match instruction regexp")
	errUnrecognizedAction     = errors.New("unrecognized action")
)

type instruction struct {
	action string
	x0     int
	y0     int
	x1     int
	y1     int
}

func parseInstructions(r io.Reader) ([]instruction, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var (
		instructions                   []instruction
		instructionRegexp              = regexp.MustCompile(`^(turn (?:on|off)|toggle) (\d+),(\d+) through (\d+),(\d+)$`)
		instructionRegexpCaptureGroups = 6
	)

	for scanner.Scan() {
		matches := instructionRegexp.FindStringSubmatch(scanner.Text())
		if len(matches) != instructionRegexpCaptureGroups {
			return instructions, fmt.Errorf("%q: %w", scanner.Text(), errLineDoesNotMatchRegexp)
		}

		//nolint:varnamelen
		x0, err := strconv.Atoi(matches[2])
		if err != nil {
			return instructions, fmt.Errorf("unable to parse %q as int: %w", matches[0][2], err)
		}

		//nolint:varnamelen
		y0, err := strconv.Atoi(matches[3])
		if err != nil {
			return instructions, fmt.Errorf("unable to parse %q as int: %w", matches[0][3], err)
		}

		//nolint:varnamelen
		x1, err := strconv.Atoi(matches[4])
		if err != nil {
			return instructions, fmt.Errorf("unable to parse %q as int: %w", matches[0][4], err)
		}

		y1, err := strconv.Atoi(matches[5])
		if err != nil {
			return instructions, fmt.Errorf("unable to parse %q as int: %w", matches[0][5], err)
		}

		instructions = append(instructions, instruction{matches[1], x0, y0, x1, y1})
	}

	return instructions, nil
}

func execute(ins instruction, lights *[1000][1000]int, f func(int) int) {
	for i := ins.x0; i <= ins.x1; i++ {
		for j := ins.y0; j <= ins.y1; j++ {
			lights[i][j] = f(lights[i][j])
		}
	}
}

func one(instructions []instruction) (int, error) {
	var (
		lights [1000][1000]int
		total  int
	)

	for _, ins := range instructions {
		switch ins.action {
		case "toggle":
			execute(ins, &lights, func(i int) int {
				//nolint:gomnd
				return (i + 1) % 2 // flip between 1 and 0
			})
		case "turn on":
			execute(ins, &lights, func(i int) int { return 1 })
		case "turn off":
			execute(ins, &lights, func(i int) int { return 0 })
		default:
			return total, fmt.Errorf("%q: %w", ins.action, errUnrecognizedAction)
		}
	}

	for i := range lights {
		for j := range lights[i] {
			total += lights[i][j]
		}
	}

	return total, nil
}

func two(instructions []instruction) (int, error) {
	var (
		lights [1000][1000]int
		total  int
	)

	for _, ins := range instructions {
		brightnessStep := 1

		switch ins.action {
		case "toggle":
			execute(ins, &lights, func(i int) int { return i + 2*brightnessStep })
		case "turn on":
			execute(ins, &lights, func(i int) int { return i + brightnessStep })
		case "turn off":
			execute(ins, &lights, func(i int) int {
				if i > 0 {
					return i - brightnessStep
				}

				return i
			})
		default:
			return total, fmt.Errorf("%q: %w", ins.action, errUnrecognizedAction)
		}
	}

	for i := range lights {
		for j := range lights[i] {
			total += lights[i][j]
		}
	}

	return total, nil
}

func main() {
	logger := log.New(os.Stderr, "", log.Llongfile)

	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	instructions, err := parseInstructions(input)
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	got, err := one(instructions)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	//nolint
	fmt.Println(got)

	got, err = two(instructions)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	//nolint
	fmt.Println(got)
}
