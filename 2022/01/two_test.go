package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwo(t *testing.T) {
	want := 45000
	got := two(parsed)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("two() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkTwo(b *testing.B) {
	want := 45000
	for i := 0; i < b.N; i++ {
		got := two(parsed)
		if got != want {
			b.Fatalf("two() mismatch: want %d, got %d", want, got)
		}
	}
}

func TestTwoNoSort(t *testing.T) {
	want := 45000
	got := twoNoSort(parsed)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("twoNoSort() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkTwoNoSort(b *testing.B) {
	want := 45000
	for i := 0; i < b.N; i++ {
		got := twoNoSort(parsed)
		if got != want {
			b.Fatalf("twoNoSort() mismatch: want %d, got %d", want, got)
		}
	}
}

func TestTwoNoParse(t *testing.T) {
	want := 45000
	got := twoNoParse(strings.NewReader(example))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("twoNoParse() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkTwoNoParse(b *testing.B) {
	want := 45000
	for i := 0; i < b.N; i++ {
		got := twoNoParse(strings.NewReader(example))
		if got != want {
			b.Fatalf("twoNoParse() mismatch: want %d, got %d", want, got)
		}
	}
}

func TestTwoMinimalState(t *testing.T) {
	want := 45000
	got := twoMinimalState(strings.NewReader(example))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("twoMinimalState() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkTwoMinimalState(b *testing.B) {
	want := 45000
	for i := 0; i < b.N; i++ {
		got := twoMinimalState(strings.NewReader(example))
		if got != want {
			b.Fatalf("twoMinimalState() mismatch: want %d, got %d", want, got)
		}
	}
}
