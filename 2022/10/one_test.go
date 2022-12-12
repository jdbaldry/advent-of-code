package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`

//nolint:dupl
func TestOnes(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"one", one},
	} {
		impl := impl

		for _, testCase := range []struct {
			name  string
			input func() io.Reader
			want  int
		}{
			{
				"example",
				func() io.Reader { return strings.NewReader(example) },
				13140,
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
	want := 14240

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
