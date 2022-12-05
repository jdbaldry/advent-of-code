package main

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwos(t *testing.T) {
	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (string, error)
	}{
		{"two", two},
	} {
		for _, tc := range []struct {
			name  string
			input func() io.Reader
			want  string
		}{
			{
				"example",
				func() io.Reader { return strings.NewReader(example) },
				"MCD",
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

func BenchmarkTwo(b *testing.B) {
	want := "MCD"
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
