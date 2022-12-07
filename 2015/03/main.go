package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var errUnexpectedCharacter = errors.New("unexpected character in input")

type house [2]int

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		//nolint:varnamelen
		x, y    int
		visited = make(map[house]struct{})
	)

	for scanner.Scan() {
		visited[house{x, y}] = struct{}{}

		switch scanner.Text() {
		case "^":
			x--
		case "v":
			x++
		case ">":
			y++
		case "<":
			y--
		case "\n":
		default:
			return 0, fmt.Errorf("%q %w", scanner.Text(), errUnexpectedCharacter)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", scanner.Err())
	}

	return len(visited), nil
}

func two(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var (
		coordinates [2]house
		visited     = make(map[house]struct{})
	)

	//nolint:varnamelen
	for i := 0; scanner.Scan(); i++ {
		visited[coordinates[i%2]] = struct{}{}

		switch scanner.Text() {
		case "^":
			coordinates[i%2][0]--
		case "v":
			coordinates[i%2][0]++
		case ">":
			coordinates[i%2][1]--
		case "<":
			coordinates[i%2][1]++
		case "\n":
		default:
			return 0, fmt.Errorf("%q %w", scanner.Text(), errUnexpectedCharacter)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", scanner.Err())
	}

	return len(visited), nil
}

func main() {
	logger := log.New(os.Stderr, "", log.Llongfile)

	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	got, err := one(input)
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}
	//nolint:forbidigo
	fmt.Println(got)

	if _, err := input.Seek(0, 0); err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	got, err = two(input)
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}
	//nolint:forbidigo
	fmt.Println(got)
}
