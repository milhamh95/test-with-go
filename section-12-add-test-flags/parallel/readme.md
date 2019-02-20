# Parallel flag

## Parallel
To add parallel flag. How many different test case run in parallel

```go
go test -parallel [numOfParallel]
``` 

Example:

```go
go test -parallel 2
```

## Count
Run test as count number

```go
go test -count [numOfCount]
```

Example:

```go
go test -count 2
```

## P
How many program / package go tool can run in parallel. When want to run test for all package

```go
go test -p 1
```