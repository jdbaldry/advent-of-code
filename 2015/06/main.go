package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	logger            = log.New(os.Stderr, "", log.Llongfile)
	instructionRegexp = regexp.MustCompile(`^(.*) (\d+),(\d+) through (\d+),(\d+)$`)
)

const (
	inputFile = "input.txt"
)

type instruction struct {
	action string
	x0     int
	y0     int
	x1     int
	y1     int
}

func parseInstructions(r io.Reader) []instruction {
	var instructions []instruction

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		matches := instructionRegexp.FindAllStringSubmatch(s.Text(), -1)
		if len(matches) != 1 {
			log.Fatalf("Line %q does not match regexp %q\n", s.Text(), instructionRegexp)
		}

		x0, err := strconv.Atoi(matches[0][2])
		if err != nil {
			log.Fatalf("Unable to parse %q as int: %v\n", matches[0][2], err)
		}
		y0, err := strconv.Atoi(matches[0][3])
		if err != nil {
			log.Fatalf("Unable to parse %q as int: %v\n", matches[0][3], err)
		}
		x1, err := strconv.Atoi(matches[0][4])
		if err != nil {
			log.Fatalf("Unable to parse %q as int: %v\n", matches[0][4], err)
		}
		y1, err := strconv.Atoi(matches[0][5])
		if err != nil {
			log.Fatalf("Unable to parse %q as int: %v\n", matches[0][5], err)
		}

		instructions = append(instructions, instruction{matches[0][1], x0, y0, x1, y1})
	}

	return instructions
}

func one(instructions []instruction) int {
	var lights [1000][1000]int
	var total int

	for _, ins := range instructions {
		switch ins.action {
		case "toggle":
			for i := ins.x0; i <= ins.x1; i++ {
				for j := ins.y0; j <= ins.y1; j++ {
					lights[i][j] = (lights[i][j] + 1) % 2
				}
			}
		case "turn on":
			for i := ins.x0; i <= ins.x1; i++ {
				for j := ins.y0; j <= ins.y1; j++ {
					lights[i][j] = 1
				}
			}
		case "turn off":
			for i := ins.x0; i <= ins.x1; i++ {
				for j := ins.y0; j <= ins.y1; j++ {
					lights[i][j] = 0
				}
			}
		default:
			log.Fatalf("Unrecognized action %q\n", ins.action)
		}
	}

	for i := range lights {
		for j := range lights[i] {
			total += lights[i][j]
		}
	}
	return total
}

func two(instructions []instruction) int {
	var lights [1000][1000]int
	var total int

	for _, ins := range instructions {
		switch ins.action {
		case "toggle":
			for i := ins.x0; i <= ins.x1; i++ {
				for j := ins.y0; j <= ins.y1; j++ {
					lights[i][j] += 2
				}
			}
		case "turn on":
			for i := ins.x0; i <= ins.x1; i++ {
				for j := ins.y0; j <= ins.y1; j++ {
					lights[i][j]++
				}
			}
		case "turn off":
			for i := ins.x0; i <= ins.x1; i++ {
				for j := ins.y0; j <= ins.y1; j++ {
					if lights[i][j] > 0 {
						lights[i][j]--
					}
				}
			}
		default:
			log.Fatalf("Unrecognized action %q\n", ins.action)
		}
	}

	for i := range lights {
		for j := range lights[i] {
			total += lights[i][j]
		}
	}
	return total
}

func main() {
	input, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("Unable to open file %s: %v\n", inputFile, err)
	}

	instructions := parseInstructions(input)
	fmt.Println(one(instructions))
	fmt.Println(two(instructions))
}
