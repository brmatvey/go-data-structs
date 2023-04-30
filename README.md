# Go data structs repo

## Provide simple implementations for absence data structs in golang

Getting started
```bash
go get -u https://github.com/brmatvey/go-data-structs
```

- ### Stack

```go
func main() {
    s := stack.New[int]()
    s.Push(1)
    _ = s.Peek()
    s.Pop()
}
```
