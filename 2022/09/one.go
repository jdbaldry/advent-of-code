package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var (
	motionRegexp                = regexp.MustCompile(`^([DLRU]) (\d+)$`)
	errDoesNotMatchRegexp       = errors.New("does not match regexp")
	errUnrecognizedRelationship = errors.New("unrecognized relationship")
)

type coord struct {
	x int
	y int
}

func (c coord) String() string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

//nolint:cyclop,funlen,gomnd
func moveTail(head, tail coord) (coord, error) {
	// Scale y by 10 so that the distances are unique.
	distance := (head.x - tail.x) + 10*(head.y-tail.y)
	switch distance {
	case -22: // head is two steps down and two steps left
		tail.x--
		tail.y--
	case -21: // head is two steps down and one step left
		tail.x--
		tail.y--
	case -20: // head is two steps down
		tail.y--
	case -19: // head is two steps down and one step right
		tail.x++
		tail.y--
	case -18: // head is two steps down and two steps right
		tail.x++
		tail.y--
	case -12: // head is one step down and two steps left
		tail.x--
		tail.y--
	case -11: // head is one step down and one step left
		// do nothing
	case -10: // head is one step down
		// do nothing
	case -9: // head is one step down and one step right
		// do nothing
	case -8: // head is one step down and two steps right
		tail.x++
		tail.y--
	case -2: // head is two steps left
		tail.x--
	case -1: // head is one step left
		// do nothing
	case 0: // same location
		// do nothing
	case 1: // head is one step right
		// do nothing
	case 2: // head is two steps right
		tail.x++
	case 8: // head is one step up and two steps left
		tail.x--
		tail.y++
	case 9: // head is one step up and one step left
		// do nothing
	case 10: // head is one step up
		// do nothing
	case 11: // head is one step up and one step right
		// do nothing
	case 12: // head is one step up and two steps right
		tail.x++
		tail.y++
	case 18: // head is two steps up and two steps left
		tail.x--
		tail.y++
	case 19: // head is two steps up and one step left
		tail.x--
		tail.y++
	case 20: // head is two steps up
		tail.y++
	case 21: // head is two steps up and one step right
		tail.x++
		tail.y++
	case 22: // head is two steps up and two steps right
		tail.x++
		tail.y++
	default:
		return tail, fmt.Errorf("head %s, tail %s: %w", head, tail, errUnrecognizedRelationship)
	}

	return tail, nil
}

//nolint:cyclop
func move(seen map[coord]struct{}, rope []coord, direction string, distance int) ([]coord, error) {
distance:
	for i := 0; i < distance; i++ {
		for j := range rope { //nolint:varnamelen
			switch j {
			case 0: // head
				switch direction {
				case "L":
					rope[0].x--
				case "R":
					rope[0].x++
				case "U":
					rope[0].y++
				case "D":
					rope[0].y--
				}
			default:
				tail, err := moveTail(rope[j-1], rope[j])
				if err != nil {
					return rope, fmt.Errorf("direction %q, rope %s: %w", direction, rope, err)
				}

				if tail == rope[j] {
					continue distance
				}

				rope[j] = tail

				if j == len(rope)-1 {
					seen[rope[j]] = struct{}{}
				}
			}
		}
	}

	return rope, nil
}

func one(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	knots := 2
	rope := make([]coord, knots)
	seen := map[coord]struct{}{{0, 0}: {}}

	for scanner.Scan() {
		line := scanner.Text()

		matched := motionRegexp.FindStringSubmatch(line)
		if len(matched) != motionRegexp.NumSubexp()+1 {
			return 0, fmt.Errorf("%q %w %q", line, errDoesNotMatchRegexp, motionRegexp)
		}

		distance, err := strconv.Atoi(matched[2])
		if err != nil {
			return 0, fmt.Errorf("could not parse distance: %w", err)
		}

		direction := matched[1]

		rope, err = move(seen, rope, direction, distance)
		if err != nil {
			return 0, err
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("%w during scanning", err)
	}

	return len(seen), nil
}
