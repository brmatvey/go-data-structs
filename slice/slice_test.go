package slice_test

import (
	"testing"

	"github.com/brmatvey/go-data-structs/slice"
)

type testCase[T comparable] struct {
	expected []T
	actual   []T
}

func TestReverse(t *testing.T) {
	testCases := []testCase[int]{
		{
			expected: []int{3, 2, 1},
			actual:   []int{1, 2, 3},
		},
	}

	for _, c := range testCases {
		slice.Reverse(c.actual)
		if len(c.actual) != len(c.expected) {
			t.Fatal("len must be eq")
		}
		for i, expected := range c.expected {
			if expected != c.actual[i] {
				t.Fatal("must be eq")
			}
		}
	}
}

func TestIsEqual(t *testing.T) {
	testCases := []testCase[int]{
		{
			expected: []int{3, 2, 1},
			actual:   []int{3, 2, 1},
		},
	}

	wrongTestCases := []testCase[int]{
		{
			expected: []int{3, 2, 1},
			actual:   []int{1, 2, 3},
		},
	}

	for _, c := range testCases {
		if !slice.IsEqual(c.expected, c.actual) {
			t.Fatal("must be eq")
		}
	}

	for _, c := range wrongTestCases {
		if slice.IsEqual(c.expected, c.actual) {
			t.Fatal("mustn't be eq")
		}
	}
}
