package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

//nolint:dupl
func TestOnes(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"one", one},
		{"oneWithXor", oneWithXor},
	} {
		impl := impl

		for _, testCase := range []struct {
			name  string
			input func() io.Reader
			want  int
		}{
			{"example", func() io.Reader { return strings.NewReader(example) }, 7},
			{"additional example 1", func() io.Reader { return strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz") }, 5},
			{"additional example 2", func() io.Reader { return strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg") }, 6},
			{"additional example 3", func() io.Reader { return strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") }, 10},
			{"additional example 4", func() io.Reader { return strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") }, 11},
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
	want := 1655

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

func BenchmarkOneWithXor(b *testing.B) {
	want := 1655

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got, err := oneWithXor(file)
		if err != nil {
			b.Fatalf("oneWithXor() unexpected error: %v", err)
		}

		if got != want {
			b.Fatalf("oneWithXor() mismatch: want %v, got %v", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}
