package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwo(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
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
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got, err := two(testCase.input())
			if err != nil {
				t.Fatalf("two() unexpected error: %v", err)
			}
			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("two() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func BenchmarkTwo(b *testing.B) {
	want := 11756

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
