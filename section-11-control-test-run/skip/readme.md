# Skip

## 40 - skipping test
To run testing.Short

```go
go test -v -short
```

## 41 - custom flags
To run custom flags

```go
go test -v -your_flag
```

example:

```go
go test -v -integration
```

it will make integration variable value become true