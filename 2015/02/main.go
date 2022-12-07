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
	dimensionsRegexp      = regexp.MustCompile(`^(\d+)x(\d+)x(\d+)$`)
	errDoesNotMatchRegexp = errors.New("does not match regexp")
)

type dimensions struct{ l, w, h int }

func parseDimensions(r io.Reader) ([]dimensions, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var parsed []dimensions

	for scanner.Scan() {
		matches := dimensionsRegexp.FindStringSubmatch(scanner.Text())
		if len(matches) != dimensionsRegexp.NumSubexp()+1 {
			return parsed, fmt.Errorf("line %q %w %q", scanner.Text(), errDoesNotMatchRegexp, dimensionsRegexp)
		}

		length, err := strconv.Atoi(matches[1])
		if err != nil {
			return parsed, fmt.Errorf("unable to parse length %q: %w", matches[1], err)
		}

		width, err := strconv.Atoi(matches[2])
		if err != nil {
			return parsed, fmt.Errorf("unable to parse width %q: %w", matches[2], err)
		}

		height, err := strconv.Atoi(matches[3])
		if err != nil {
			return parsed, fmt.Errorf("unable to parse height %q: %w", matches[3], err)
		}

		parsed = append(parsed, dimensions{length, width, height})
	}

	return parsed, nil
}

func one(ds []dimensions) int {
	var total int

	for _, dimension := range ds {
		var smallest int

		lengthByWidth := dimension.l * dimension.w
		smallest = lengthByWidth

		widthByHeight := dimension.w * dimension.h
		if widthByHeight < smallest {
			smallest = widthByHeight
		}

		heightByLength := dimension.h * dimension.l
		if heightByLength < smallest {
			smallest = heightByLength
		}

		total += 2*lengthByWidth + 2*widthByHeight + 2*heightByLength + smallest
	}

	return total
}

func perimeter(l, w int) int {
	//nolint:gomnd
	return 2 * (l + w)
}

func two(ds []dimensions) int {
	var total int

	for _, dimension := range ds {
		var smallest int

		lengthByWidth := perimeter(dimension.l, dimension.w)
		smallest = lengthByWidth

		widthByHeight := perimeter(dimension.w, dimension.h)
		if widthByHeight < smallest {
			smallest = widthByHeight
		}

		heightByLength := perimeter(dimension.h, dimension.l)
		if heightByLength < smallest {
			smallest = heightByLength
		}

		total += smallest + dimension.l*dimension.w*dimension.h
	}

	return total
}

func main() {
	logger := log.New(os.Stderr, "", log.Llongfile)

	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}
	defer input.Close()

	parsed, err := parseDimensions(input)
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	//nolint:forbidigo
	fmt.Println(one(parsed))
	//nolint:forbidigo
	fmt.Println(two(parsed))
}
