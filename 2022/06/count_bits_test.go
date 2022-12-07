package main

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	t.Parallel()

	for _, impl := range []struct {
		name string
		fn   func(uint64) uint64
	}{
		{"countBitsNaïve", countBitsNaive},
		{"countBitsBrianKernighan", countBitsBrianKernighan},
		{"countBitsParallel", countBitsParallel},
	} {
		impl := impl

		for _, testCase := range []struct {
			name  string
			input func() uint64
			want  uint64
		}{
			{
				name: "16 bits",
				input: func() uint64 {
					return uint64(0b1010101010100100100100100100100100100100100000000000000000000000)
				},
				want: 16,
			},
		} {
			testCase := testCase

			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				got := impl.fn(testCase.input())
				if got != testCase.want {
					t.Errorf("countBits() mismatch want %d, got %d", testCase.want, got)
				}
			})
		}
	}
}

func BenchmarkCountBitsNaive(b *testing.B) {
	want := uint64(16)

	for i := 0; i < b.N; i++ {
		got := countBitsNaive(uint64(0b1010101010100100100100100100100100100100100000000000000000000000))
		if got != want {
			b.Errorf("countBitsNaïve() mismatch want %d, got %d", want, got)
		}
	}
}

func BenchmarkCountBitsBrianKernighan(b *testing.B) {
	want := uint64(16)

	for i := 0; i < b.N; i++ {
		got := countBitsBrianKernighan(uint64(0b1010101010100100100100100100100100100100100000000000000000000000))
		if got != want {
			b.Errorf("countBitsBrianKernighan() mismatch want %d, got %d", want, got)
		}
	}
}

func BenchmarkCountBitsParallel(b *testing.B) {
	want := uint64(16)

	for i := 0; i < b.N; i++ {
		got := countBitsParallel(uint64(0b1010101010100100100100100100100100100100100000000000000000000000))
		if got != want {
			b.Errorf("countBitsParallel() mismatch want %d, got %d", want, got)
		}
	}
}
