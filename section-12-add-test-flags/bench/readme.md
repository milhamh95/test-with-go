# 43. Benchmark

## bench_test.go

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