package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//nolint:varnamelen
func parseThenSolveWith(f func([][]int) int) func(io.Reader) (int, error) {
	return func(r io.Reader) (int, error) {
		parsed, err := parse(r)
		if err != nil {
			return 0, err
		}

		return f(parsed), nil
	}
}

func TestTwos(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"two", parseThenSolveWith(two)},
		{"twoNoSort", parseThenSolveWith(twoNoSort)},
		{"twoNoParse", twoNoParse},
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
						t.Fatal(err)
					}

					return f
				},
				212117,
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

func BenchmarkTwo(b *testing.B) {
	want := 212117

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	parsed, err := parse(file)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got := two(parsed)
		if got != want {
			b.Fatalf("two() mismatch: want %d, got %d", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTwoNoSort(b *testing.B) {
	want := 212117

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	parsed, err := parse(file)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got := twoNoSort(parsed)
		if got != want {
			b.Fatalf("twoNoSort() mismatch: want %d, got %d", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTwoNoParse(b *testing.B) {
	want := 212117

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got, err := twoNoParse(file)
		if err != nil {
			b.Fatalf("twoNoParse() unexpected error: %v", err)
		}

		if got != want {
			b.Fatalf("twoNoParse() mismatch: want %d, got %d", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}
