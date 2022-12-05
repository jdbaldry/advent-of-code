package main

import (
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestOnes(t *testing.T) {
	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (string, error)
	}{
		{"one", one},
	} {
		for _, tc := range []struct {
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
			t.Run(tc.name, func(t *testing.T) {
				got, err := impl.fn(tc.input())
				if err != nil {
					t.Errorf("%s() unexpected errors: %v", impl.name, err)
				}
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("%s() mismatch (-want +got):\n%s", impl.name, diff)
				}
			})
		}
	}
}

func BenchmarkOne(b *testing.B) {
	want := "CMZ"
	for i := 0; i < b.N; i++ {
		got, err := one(strings.NewReader(example))
		if err != nil {
			b.Fatalf("one() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("one() mismatch: want %v, got %v", want, got)
		}
	}
}

func BenchmarkLazyRegexp(b *testing.B) {
	input := "move 1 from 2 to 1"
	re := regexp.MustCompile(`^.*?(\d+).*?(\d+).*?(\d+)$`)
	for i := 0; i < b.N; i++ {
		got := re.MatchString(input)
		if got != true {
			b.Errorf("`^.*?(\\d+).*?(\\d+).*?(\\d+)$`.MatchString(input) must return true")
		}
	}
}

func BenchmarkAccurate(b *testing.B) {
	input := "move 1 from 2 to 1"
	for i := 0; i < b.N; i++ {
		got := instructionRegexp.MatchString(input)
		if got != true {
			b.Errorf("instructionRegexp.MatchString(input) must return true")
		}
	}
}
