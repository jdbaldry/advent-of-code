package main

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func TestOnes(t *testing.T) {
	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"one", one},
		{"oneUsingBits", oneUsingBits},
	} {
		for _, tc := range []struct {
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
	want := 7
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
