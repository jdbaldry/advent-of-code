package main

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	for _, impl := range []struct {
		name string
		fn   func(uint64) uint64
	}{
		{"countBitsNaïve", countBitsNaïve},
		{"countBitsBrianKernighan", countBitsBrianKernighan},
		{"countBitsParallel", countBitsParallel},
	} {
		for _, tc := range []struct {
			name  string
			input func() uint64
			want  uint64
		}{
			{
				name:  "16 bits",
				input: func() uint64 { return uint64(0b1010101010100100100100100100100100100100100000000000000000000000) },
				want:  16,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				got := impl.fn(tc.input())
				if got != tc.want {
					t.Errorf("countBits() mismatch want %d, got %d", tc.want, got)
				}
			})
		}
	}
}

func BenchmarkCountBitsNaïve(b *testing.B) {
	want := uint64(16)
	for i := 0; i < b.N; i++ {
		got := countBitsNaïve(uint64(0b1010101010100100100100100100100100100100100000000000000000000000))
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
