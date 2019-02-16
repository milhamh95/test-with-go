# test-with-go
Test With Go Course from usegolang.com

## Section 11 - run test for subpackage

### 39 - running tests for subpackages
To run all test in subpackages

```go
go test ./...
```

To see list of package
```sh
for pkg in *; do echo "./$pkg"; done
```

To test package sequencely
```sh
for pkg in *; do go test "./$pkg"; done 
```

To see list of file
```sh
for pkg in */**; do echo "./$pkg"; done
```