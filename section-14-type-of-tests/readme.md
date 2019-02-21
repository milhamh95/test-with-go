# Type of test

## Unit Testing

Testing very small things, like a function , in isolation

Example:
```go

func Magic(a, b int) int {
    return (a + b) * (a + b)
}

func TestMagic(t *testing.T) {
    got := Magic(1,2)
    if got != 9 {
        t.Errorf("Magic() = %v, want %v", got, 9)
    }
}
```