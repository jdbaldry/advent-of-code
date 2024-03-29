package main

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

//nolint:gochecknoglobals
var parsed = [][]int{
	{
		1000,
		2000,
		3000,
	},
	{
		4000,
	},
	{
		5000,
		6000,
	},
	{
		7000,
		8000,
		9000,
	},
	{
		10000,
	},
}

func TestParse(t *testing.T) {
	t.Parallel()

	got, err := parse(strings.NewReader(example))
	if err != nil {
		t.Fatalf("Unable to parse example input: %v", err)
	}

	if diff := cmp.Diff(parsed, got); diff != "" {
		t.Errorf("parse() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkParse(b *testing.B) {
	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_, err := parse(file)
		if err != nil {
			b.Fatalf("Unable to parse example input: %v", err)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func FuzzParse(f *testing.F) {
	f.Add(example)
	f.Fuzz(func(t *testing.T, s string) {
		_, err := parse(strings.NewReader(example))
		if err != nil {
			t.Fatalf("Unable to parse example input: %v", err)
		}
	})
}

func TestOne(t *testing.T) {
	t.Parallel()

	want := 24000

	got := one(parsed)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("one() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkOne(b *testing.B) {
	want := 72511

	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}

	parsed, err := parse(file)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		got := one(parsed)
		if got != want {
			b.Fatalf("one() mismatch: want %v, got %v", want, got)
		}

		if _, err := file.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}
