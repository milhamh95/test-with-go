# 43. Benchmark

## bench_test.go

Function should user word ***Benchmark*** in the front of the function

To run test

```go
go test -bench .
```

# 44. Verbose
To run 

```go
go test -bench .
go test -bench [regex]
go test -bench [regex] -v
```

Example :

```go
go test -bench BenchmarkFibMemo20-4
```

# 46. Timeout

To run
```go
go test -timeout [second]s
```

Example :

```go
go test -timeout 5s
```