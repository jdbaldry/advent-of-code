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
	logger           = log.New(os.Stderr, "", log.Llongfile)
	dimensionsRegexp = regexp.MustCompile(`^(\d+)x(\d+)x(\d+)$`)
)

type dimensions struct{ l, w, h int }

func parseDimensions(r io.Reader) []dimensions {
	var ds []dimensions

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		matches := dimensionsRegexp.FindAllStringSubmatch(s.Text(), -1)
		if len(matches) == 0 {
			logger.Fatalf("Line %q does not match regexp %q", s.Text(), dimensionsRegexp)
		}

		l, err := strconv.Atoi(matches[0][1])
		if err != nil {
			logger.Fatalf("Unable to parse length %q: %v", matches[0][1], err)
		}
		w, err := strconv.Atoi(matches[0][2])
		if err != nil {
			logger.Fatalf("Unable to parse width %q: %v", matches[0][2], err)
		}
		h, err := strconv.Atoi(matches[0][3])
		if err != nil {
			logger.Fatalf("Unable to parse height %q: %v", matches[0][3], err)
		}

		ds = append(ds, dimensions{l, w, h})
	}

	return ds
}

func one(ds []dimensions) int {
	var total int

	for _, d := range ds {
		var smallest int

		lw := d.l * d.w
		smallest = lw
		wh := d.w * d.h
		if wh < smallest {
			smallest = wh
		}
		hl := d.h * d.l
		if hl < smallest {
			smallest = hl
		}

		total += 2*lw + 2*wh + 2*hl + smallest
	}
	return total
}

func two(ds []dimensions) int {
	var total int

	for _, d := range ds {
		var smallest int

		lw := 2 * (d.l + d.w)
		smallest = lw
		wh := 2 * (d.w + d.h)
		if wh < smallest {
			smallest = wh
		}
		hl := 2 * (d.h + d.l)
		if hl < smallest {
			smallest = hl
		}

		total += smallest + d.l*d.w*d.h
	}

	return total
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("Unable to open input.txt: %v\n", err)
	}
	defer input.Close()

	dimensions := parseDimensions(input)

	fmt.Println(one(dimensions))
	fmt.Println(two(dimensions))
}
