package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwo(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input func() io.Reader
		want  int
	}{
		{
			"example",
			func() io.Reader { return strings.NewReader(example) },
			12,
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
			11756,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := two(tc.input())
			if err != nil {
				t.Fatalf("two() unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("two() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func BenchmarkTwo(b *testing.B) {
	want := 12
	for i := 0; i < b.N; i++ {
		got, err := two(strings.NewReader(example))
		if err != nil {
			b.Fatalf("two() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("two() mismatch: want %v, got %v", want, got)
		}
	}
}
