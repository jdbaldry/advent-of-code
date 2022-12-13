package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
`

func TestCoordRegexp(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid",
			input: "1,4",
			want:  true,
		},
	} {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := coordRegexp.MatchString(testCase.input)

			if diff := cmp.Diff(got, testCase.want); diff != "" {
				t.Errorf("coordRegexp mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFoldRegexp(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "fold along x=7",
			input: "fold along x=7",
			want:  true,
		},
	} {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := foldRegexp.MatchString(testCase.input)

			if diff := cmp.Diff(got, testCase.want); diff != "" {
				t.Errorf("foldRegexp mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOnes(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(io.Reader) (int, error)
	}{
		{"one", one},
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
				17,
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

func BenchmarkOne(b *testing.B) {
	want := 781

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

func TestMirror(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name string
		n    int
		line int

		want int
	}{
		{
			name: "mirror(14, 7)",
			n:    14,
			line: 7,

			want: 0,
		},
		{
			name: "mirror(0, 7)",
			n:    0,
			line: 7,

			want: 0,
		},
		{
			name: "mirror(13, 7)",
			n:    13,
			line: 7,

			want: 1,
		},
		{
			name: "mirror(1, 7)",
			n:    1,
			line: 7,

			want: 1,
		},
	} {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := mirror(testCase.n, testCase.line)

			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("mirror() mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
