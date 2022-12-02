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

func TestOne(t *testing.T) {
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
			got, err := one(tc.input())
			if err != nil {
				t.Fatalf("one() unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("one() mismatch (-want +got):\n%s", diff)
			}
		})
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
			b.Fatalf("one() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("one() mismatch: want %v, got %v", want, got)
		}
	}
}
