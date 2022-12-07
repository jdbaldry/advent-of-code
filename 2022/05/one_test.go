package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	example = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	exampleInstruction = "move 1 from 2 to 1"
)

//nolint:dupl
func TestOnes(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (string, error)
	}{
		{"one", one},
	} {
		impl := impl

		for _, testCase := range []struct {
			name  string
			input func() io.Reader
			want  string
		}{
			{
				"example",
				func() io.Reader { return strings.NewReader(example) },
				"CMZ",
			},
		} {
			testCase := testCase

			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				got, err := impl.fn(testCase.input())
				if err != nil {
					t.Errorf("%s() unexpected errors: %v", impl.name, err)
				}

				if diff := cmp.Diff(testCase.want, got); diff != "" {
					t.Errorf("%s() mismatch (-want +got):\n%s", impl.name, diff)
				}
			})
		}
	}
}

func BenchmarkOne(b *testing.B) {
	want := "JCMHLVGMG"

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got, err := one(file)
		if err != nil {
			b.Fatalf("one() unexpected error: %v", err)
		}

		if got != want {
			b.Fatalf("one() mismatch: want %v, got %v", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLazyRegexp(b *testing.B) {
	re := regexp.MustCompile(`^.*?(\d+).*?(\d+).*?(\d+)$`)

	for i := 0; i < b.N; i++ {
		got := re.FindStringSubmatch(exampleInstruction)

		count, err := strconv.Atoi(got[1])
		if err != nil {
			b.Fatal(err)
		}

		from, err := strconv.Atoi(got[2])
		if err != nil {
			b.Fatal(err)
		}

		to, err := strconv.Atoi(got[3])
		if err != nil {
			b.Fatal(err)
		}

		if count != 1 || from != 2 || to != 1 {
			b.Errorf(`"^.*?(\\d+).*?(\\d+).*?(\\d+)$".FindStringSubmatch() must return []string{input, "1", "2", "1"}`)
		}
	}
}

func BenchmarkAccurate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := instructionRegexp.FindStringSubmatch(exampleInstruction)

		count, err := strconv.Atoi(got[1])
		if err != nil {
			b.Fatal(err)
		}

		from, err := strconv.Atoi(got[2])
		if err != nil {
			b.Fatal(err)
		}

		to, err := strconv.Atoi(got[3])
		if err != nil {
			b.Fatal(err)
		}

		if count != 1 || from != 2 || to != 1 {
			b.Errorf(`instructionRegexp.FindStringSubmatch() must return []string{input, "1", "2", "1"}`)
		}
	}
}

func BenchmarkSscanf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var count, from, to int

		fmt.Sscanf(exampleInstruction, "move %d from %d to %d", &count, &from, &to)

		if count != 1 || from != 2 || to != 1 {
			b.Errorf(`fmt.Sscanf() must set count, from, to = 1, 2, 1`)
		}
	}
}
