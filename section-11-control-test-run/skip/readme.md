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

## 42 - build tags
To create custom tags, in the beginning of the files, need to add custom build tags

```go
// +build [tag_name]
```

example

```go
// +build mysql
```

to run :

```sh
go test -v -tags=mysql
```

```sh
go test -v -tags="mysql"
```