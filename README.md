# Go data structs repo

## Provide simple implementations for absence data structs in golang

- ### Stack

```go
func main() {
    s := stack.New[int]()
    s.Push(1)
    _ = s.Peek()
    s.Pop()
}
```
