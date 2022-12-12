package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//nolint:dupl
func TestTwos(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"two", two},
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
				8,
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

func BenchmarkTwo(b *testing.B) {
	want := 157320

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got, err := two(file)
		if err != nil {
			b.Fatalf("two() unexpected error: %v", err)
		}

		if got != want {
			b.Fatalf("two() mismatch: want %v, got %v", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}
