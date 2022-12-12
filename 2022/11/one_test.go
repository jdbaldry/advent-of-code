package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const example = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`

func TestMonkeyRegexp(t *testing.T) {
	t.Parallel()

	input := "Monkey 0:\n"
	matched := monkeyRegexp.FindStringSubmatch(input)

	if len(matched) != monkeyRegexp.NumSubexp()+1 {
		t.Fatalf("monkeyRegexp failed to match input %q", input)
	}

	if diff := cmp.Diff("0", matched[monkeyRegexp.SubexpIndex("src")]); diff != "" {
		t.Errorf("monkeyRegexp subexp src mismatch (-want +got):\n%s", diff)
	}
}

func TestStartingItemsRegexp(t *testing.T) {
	t.Parallel()

	input := "  Starting items: 79, 60, 97\n"
	matched := startingItemsRegexp.FindStringSubmatch(input)

	if len(matched) != startingItemsRegexp.NumSubexp()+1 {
		t.Fatalf("startingItemsRegexp failed to match input %q", input)
	}

	if diff := cmp.Diff("79, 60, 97", matched[startingItemsRegexp.SubexpIndex("items")]); diff != "" {
		t.Errorf("startingItemsRegexp subexp items mismatch (-want +got):\n%s", diff)
	}
}

func TestOperationRegexp(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name         string
		input        string
		wantOperator string
		wantOperand  string
	}{
		{
			"+ 6",
			"    Operation: new = old + 6\n",
			"+",
			"6",
		},
		{
			"* old",
			"    Operation: new = old * old\n",
			"*",
			"old",
		},
	} {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			matched := operationRegexp.FindStringSubmatch(testCase.input)

			if len(matched) != operationRegexp.NumSubexp()+1 {
				t.Fatalf("operationRegexp failed to match input %q", testCase.input)
			}

			if diff := cmp.Diff(testCase.wantOperator, matched[operationRegexp.SubexpIndex("operator")]); diff != "" {
				t.Errorf("operationRegexp subexp items mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(testCase.wantOperand, matched[operationRegexp.SubexpIndex("operand")]); diff != "" {
				t.Errorf("operationRegexp subexp items mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIfTrueRegexp(t *testing.T) {
	t.Parallel()

	input := "    If true: throw to monkey 0\n"
	matched := ifTrueRegexp.FindStringSubmatch(input)

	if len(matched) != ifTrueRegexp.NumSubexp()+1 {
		t.Fatalf("ifTrueRegexp failed to match input %q", input)
	}

	if diff := cmp.Diff("0", matched[ifTrueRegexp.SubexpIndex("ifTrue")]); diff != "" {
		t.Errorf("ifTrueRegexp subexp items mismatch (-want +got):\n%s", diff)
	}
}

func TestIfFalseRegexp(t *testing.T) {
	t.Parallel()

	input := "    If false: throw to monkey 0\n"
	matched := ifFalseRegexp.FindStringSubmatch(input)

	if len(matched) != ifFalseRegexp.NumSubexp()+1 {
		t.Fatalf("ifFalseRegexp failed to match input %q", input)
	}

	if diff := cmp.Diff("0", matched[ifFalseRegexp.SubexpIndex("ifFalse")]); diff != "" {
		t.Errorf("ifFalseRegexp subexp items mismatch (-want +got):\n%s", diff)
	}
}

//nolint:dupl
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
				10605,
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
	want := 56120

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
