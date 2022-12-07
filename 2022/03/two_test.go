package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//nolint:funlen
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
				70,
			},
			{
				"input",
				func() io.Reader {
					f, err := os.Open("input.txt")
					if err != nil {
						panic(err.Error())
					}

					return f
				},
				2817,
			},
			{
				"adams",
				func() io.Reader {
					f, err := os.Open("adams.txt")
					if err != nil {
						panic(err.Error())
					}

					return f
				},
				2821,
			},
			{
				"adams failing id 22",
				func() io.Reader {
					return strings.NewReader(`sMNnNRNrlGlsZBrGsrFQpclWlWLfpWjtzTfDtpzj
gvhPgwTgdSHtHDtpDPLp
gwhSwdvTSTbSgRrZNrrNFFNBGb
`)
				},
				46, // t
			},
			{
				"adams failing id 23",
				func() io.Reader {
					return strings.NewReader(`rtZnDHJrrDtGtGHvGHDWfdfwCjcBhjBCffwwLv
lzVlzsTRsmzVNTspVsMMsmwCLcmjmcdbBBChwfBbCW
sVTMpTpppsVMsPRPVzMNFqMFwZtQrHZDGqgHZrSQQrQQJDGn
`)
				},
				23, // w
			},
			{
				"adams failing id 50",
				func() io.Reader {
					return strings.NewReader(`VbHqLlGQlgjLjjQsNvCZTsNjMtCZvT
SJtttppwwpwBwdPvsvCvBZrvNrTrvM
JDnWJpDSSpmSwmpPzSwznhDlqGqqtqqHGHLlhblGbR
`)
				},
				20, // t
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
	want := 2817

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
