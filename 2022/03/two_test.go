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
	} {
		for _, tc := range []struct {
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
	want := 70
	for i := 0; i < b.N; i++ {
		got, err := two(strings.NewReader(example))
		if err != nil {
			b.Fatalf("two() unexpected error: %v", err)
		}
		if got != want {
			b.Fatalf("two() mismatch: want %v, got %v", want, got)
		}
	}
}
