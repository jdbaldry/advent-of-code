package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

func one(r io.Reader) (int, error) {
	var sum int
	scanner := bufio.NewScanner(r)

	for line := 0; scanner.Scan(); line++ {
		ranges := strings.Split(scanner.Text(), ",")
		if len(ranges) != 2 {
			return sum, fmt.Errorf("%d: expected exactly two ranges per line but found %d in %v", line, len(ranges), ranges)
		}

		var parsed [2][2]int
		for i, r := range ranges {
			bounds := strings.Split(r, "-")
			if len(bounds) != 2 {
				return sum, fmt.Errorf("%d: expected exactly two bounds per range but found %d in %v", line, len(bounds), bounds)
			}

			for j, b := range bounds {
				int, err := strconv.Atoi(b)
				if err != nil {
					return sum, fmt.Errorf("%d: unable to parse  bound as integer but on range %d, bound %d, %q: %v", line, i, j, b, err)
				}

				parsed[i][j] = int
			}
		}
		if (parsed[0][0] >= parsed[1][0] && parsed[0][1] <= parsed[1][1]) ||
			(parsed[1][0] >= parsed[0][0] && parsed[1][1] <= parsed[0][1]) {
			sum++
		}
		parsed = [2][2]int{}
	}
	return sum, nil
}

func oneByRunes(r io.Reader) (int, error) {
	var sum int

	type state int
	var (
		scanInt = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			for {
				b := data[advance]
				switch b {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					token = append(token, b)
					advance++
				default:
					return
				}
			}
		}
	)
	const (
		stateFirstBoundLower state = iota
		stateFirstBoundDash
		stateFirstBoundUpper
		stateComma
		stateSecondBoundLower
		stateSecondBoundDash
		stateSecondBoundUpper
		stateLine
	)

	scanner := bufio.NewScanner(r)
	scanner.Split(scanInt)

	var s state
	var parsed [2][2]int
	for scanner.Scan() {
		fmt.Println(s, parsed)
		switch s {
		case stateFirstBoundLower:
			int, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return sum, err
			}
			parsed[0][0] = int
			s = stateFirstBoundDash
			scanner.Split(bufio.ScanRunes)
		case stateFirstBoundDash:
			if r, _ := utf8.DecodeRune(scanner.Bytes()); r != '-' {
				return sum, fmt.Errorf("expected rune %q but got %q", r, '-')
			}
			s = stateFirstBoundUpper
			scanner.Split(scanInt)
		case stateFirstBoundUpper:
			int, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return sum, err
			}
			parsed[0][1] = int
			s = stateComma
			scanner.Split(bufio.ScanRunes)
		case stateComma:
			if r, _ := utf8.DecodeRune(scanner.Bytes()); r != ',' {
				return sum, fmt.Errorf("expected rune %q but got %q", r, ',')
			}
			s = stateSecondBoundLower
			scanner.Split(scanInt)
		case stateSecondBoundLower:
			int, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return sum, err
			}
			parsed[1][0] = int
			s = stateSecondBoundDash
			scanner.Split(bufio.ScanRunes)
		case stateSecondBoundDash:
			if r, _ := utf8.DecodeRune(scanner.Bytes()); r != '-' {
				return sum, fmt.Errorf("expected rune %q but got %q", r, '-')
			}
			s = stateSecondBoundUpper
			scanner.Split(scanInt)
		case stateSecondBoundUpper:
			int, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return sum, err
			}
			parsed[1][1] = int
			s = stateLine
			scanner.Split(bufio.ScanRunes)
		case stateLine:
			if r, _ := utf8.DecodeRune(scanner.Bytes()); r != '\n' {
				return sum, fmt.Errorf("expected newline but got %q", r)
			}
			s = stateFirstBoundLower
			scanner.Split(scanInt)
			if (parsed[0][0] >= parsed[1][0] && parsed[0][1] <= parsed[1][1]) ||
				(parsed[1][0] >= parsed[0][0] && parsed[1][1] <= parsed[0][1]) {
				sum++
				parsed = [2][2]int{}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return sum, err
	}

	return sum, nil
}
