package stack_test

import (
	"testing"

	"go-data-structs/stack"
)

func TestStack(t *testing.T) {
	expectedValues := []int{1, 2, 3, 4, 5}
	s := stack.New[int]()

	t.Run("test empty", func(t *testing.T) {
		if !s.Empty() {
			t.Fatal("must be empty")
		}
	})

	t.Run("test push and peek", func(t *testing.T) {
		validator := func(index, val int) {
			s.Push(val)
			if s.Peek() != val {
				t.Fatal("must be equal")
			}
			if s.Empty() {
				t.Fatal("must be not empty")
			}
			if s.Len() != index+1 {
				t.Fatal("illegal len")
			}
		}

		for i, val := range expectedValues {
			validator(i, val)
		}
	})

	t.Run("test pop", func(t *testing.T) {
		validator := func(index int) {
			if s.Peek() != expectedValues[index] {
				t.Fatal("must be equal")
			}
			s.Pop()
			if s.Len() != index {
				t.Fatal("illegal len")
			}
		}

		for i := len(expectedValues); i > 0; i-- {
			validator(i - 1)
		}

		if !s.Empty() {
			t.Fatal("must be not empty")
		}
	})
}
