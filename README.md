# Go data structs repo

## Provide simple implementations for absence data structs in golang

Getting started
```bash
go get -u https://github.com/brmatvey/go-data-structs
```

- ### Stack

Simple implementation of stack.
```go
func main() {
    s := stack.New[int]()
    s.Push(1)
    _ = s.Peek()
    s.Pop()
}
```

- ### Slice

Simple generic helpers for slices.

```go
func main() {
    s1, s2 := []int{1,2,3,4,5}, []int{5,4,3,2,1}
    slice.Reverse(s2)
    _ = slice.IsEqual(s1, s2)
}
```