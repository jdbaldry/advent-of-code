package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwos(t *testing.T) {
	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"two", two},
		{"twoWithMap", twoWithMap},
		{"twoWithXor", twoWithXor},
	} {
		for _, tc := range []struct {
			name  string
			input func() io.Reader
			want  int
		}{
			{"example", func() io.Reader { return strings.NewReader(example) }, 19},
			{"additional example 1", func() io.Reader { return strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz") }, 23},
			{"additional example 2", func() io.Reader { return strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg") }, 23},
			{"additional example 3", func() io.Reader { return strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") }, 29},
			{"additional example 4", func() io.Reader { return strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") }, 26},
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
	want := 2665
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < b.N; i++ {
		got, err := two(f)
		if err != nil {
			b.Fatalf("two() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("two() mismatch: want %v, got %v", want, got)
		}
		f.Seek(0, 0)
	}
}

func BenchmarkTwoWithMap(b *testing.B) {
	want := 2665
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < b.N; i++ {
		got, err := twoWithMap(f)
		if err != nil {
			b.Fatalf("twoWithMap() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("twoWithMap() mismatch: want %v, got %v", want, got)
		}
		f.Seek(0, 0)
	}
}

func BenchmarkTwoWithXor(b *testing.B) {
	want := 2665
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < b.N; i++ {
		got, err := twoWithXor(f)
		if err != nil {
			b.Fatalf("twoWithXor() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("twoWithXor() mismatch: want %v, got %v", want, got)
		}
		f.Seek(0, 0)
	}
}
