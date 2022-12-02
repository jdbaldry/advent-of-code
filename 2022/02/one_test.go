package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `A Y
B X
C Z
`

func TestOnes(t *testing.T) {
	for _, solution := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"one", one},
		{"oneMod3", oneMod3},
		{"oneMod3ByRunes", oneMod3ByRunes},
	} {
		for _, tc := range []struct {
			name  string
			input func() io.Reader
			want  int
		}{
			{
				"example",
				func() io.Reader { return strings.NewReader(example) },
				15,
			},
			{
				"input",
				func() io.Reader {
					f, err := os.Open("input.txt")
					if err != nil {
						panic(err)
					}
					return f
				},
				12645,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				got, err := solution.fn(tc.input())
				if err != nil {
					t.Fatalf("%s() unexpected error: %v", solution.name, err)
				}
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("%s() mismatch (-want +got):\n%s", solution.name, diff)
				}
			})
		}
	}
}

func BenchmarkOne(b *testing.B) {
	want := 15
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

func BenchmarkOneMod3(b *testing.B) {
	want := 15
	for i := 0; i < b.N; i++ {
		got, err := oneMod3(strings.NewReader(example))
		if err != nil {
			b.Fatalf("oneMod3() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("oneMod3() mismatch: want %v, got %v", want, got)
		}
	}
}

func BenchmarkOneMod3ByRunes(b *testing.B) {
	want := 15
	for i := 0; i < b.N; i++ {
		got, err := oneMod3ByRunes(strings.NewReader(example))
		if err != nil {
			b.Fatalf("oneMod3ByRunes() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("oneMod3ByRunes() mismatch: want %v, got %v", want, got)
		}
	}
}
