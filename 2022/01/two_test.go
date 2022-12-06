package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func mustParse(r io.Reader) [][]int {
	parsed, err := parse(r)
	if err != nil {
		panic(err)
	}
	return parsed
}

var tcs = []struct {
	name  string
	input func() io.Reader
	want  int
}{
	{
		"example",
		func() io.Reader { return strings.NewReader(example) },
		45000,
	},
	{
		"unordered",
		func() io.Reader { return strings.NewReader(strings.Join([]string{"4", "5", "7", "6"}, "\n\n")) },
		18,
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
		212117,
	},
}

func TestTwo(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := two(mustParse(tc.input()))
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("two() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func BenchmarkTwo(b *testing.B) {
	want := 212117
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got := two(mustParse(f))
		if got != want {
			b.Fatalf("two() mismatch: want %d, got %d", want, got)
		}
		if _, err := f.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func TestTwoNoSort(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := twoNoSort(mustParse(tc.input()))
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("twoNoSort() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func BenchmarkTwoNoSort(b *testing.B) {
	want := 212117
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got := twoNoSort(mustParse(f))
		if got != want {
			b.Fatalf("twoNoSort() mismatch: want %d, got %d", want, got)
		}
		if _, err := f.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func TestTwoNoParse(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := twoNoParse(tc.input())
			if err != nil {
				t.Fatalf("twoNoParse() unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("twoNoParse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func BenchmarkTwoNoParse(b *testing.B) {
	want := 212117
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got, err := twoNoParse(f)
		if err != nil {
			b.Fatalf("twoNoParse() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("twoNoParse() mismatch: want %d, got %d", want, got)
		}
		if _, err := f.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}
