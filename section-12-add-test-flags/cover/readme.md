# Code Coverage

1. To run cover

```go
go test -cover
```

It will return total percentage of the function that have been tested

2. To show test coverage

a. Generate .out file

```go
go test -coverprofile=[filename.out]
```

Example:

```go
go test -coverprofile=cover.out
```

b. Generate in terminal

```go
go tool cover -func=[filename.out]
```

Example:

```go
go tool cover -func=cover.out
```

c. Generate in html

```go
go tool cover -html=[filename.out]
```

Example:

```go
go tool cover -html=cover.out
```

Code coverage will check the number line of the code