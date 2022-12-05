package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestOnes(t *testing.T) {
	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"one", one},
		{"oneByRunes", oneByRunes},
	} {
		for _, tc := range []struct {
			name  string
			input func() io.Reader
			want  int
		}{
			{
				"example",
				func() io.Reader { return strings.NewReader(example) },
				2,
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
				651,
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
	want := 2
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

func BenchmarkOneByRunes(b *testing.B) {
	want := 2
	for i := 0; i < b.N; i++ {
		got, err := oneByRunes(strings.NewReader(example))
		if err != nil {
			b.Fatalf("oneByRunes() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("oneByRunes() mismatch: want %v, got %v", want, got)
		}
	}
}
