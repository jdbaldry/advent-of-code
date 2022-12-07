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

func TestOnes(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"one", one},
		{"oneMod3", oneMod3},
		{"oneMod3ByRunes", oneMod3ByRunes},
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
			testCase := testCase

			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				got, err := impl.fn(testCase.input())
				if err != nil {
					t.Fatalf("%s() unexpected error: %v", impl.name, err)
				}
				if diff := cmp.Diff(testCase.want, got); diff != "" {
					t.Errorf("%s() mismatch (-want +got):\n%s", impl.name, diff)
				}
			})
		}
	}
}

func BenchmarkOne(b *testing.B) {
	want := 12645

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

func BenchmarkOneMod3(b *testing.B) {
	want := 12645

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got, err := oneMod3(file)
		if err != nil {
			b.Fatalf("oneMod3() unexpected error: %v", err)
		}

		if got != want {
			b.Fatalf("oneMod3() mismatch: want %v, got %v", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkOneMod3ByRunes(b *testing.B) {
	want := 12645

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got, err := oneMod3ByRunes(file)
		if err != nil {
			b.Fatalf("oneMod3ByRunes() unexpected error: %v", err)
		}

		if got != want {
			b.Fatalf("oneMod3ByRunes() mismatch: want %v, got %v", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}
