# Go's io.Reader

Go's io.Reader is defined as:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
