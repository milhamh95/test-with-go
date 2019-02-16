# Running spesific test

To run spesific test, you can run 


```go
go test -v -run [regex pattern]
```

Example :

```go
go test -v -run TestGotcha
```

```go
go test -v -run TestGotcha/arg=3
```