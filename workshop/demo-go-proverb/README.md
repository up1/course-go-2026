# Demo with [Go Proverbs](https://go-proverbs.github.io/)


## 1. Don't communicate by sharing memory, share memory by communicating.

Bad
```
var mu sync.Mutex
data := make(map[string]int)
go func() {
    mu.Lock()
    data["key"] = 1
    mu.Unlock()
}()
go func() {
    mu.Lock()
    fmt.Println(data["key"])
    mu.Unlock()
}()
```

Good
```
data := make(map[string]int)
ch := make(chan struct{})
go func() {
    data["key"] = 1
    ch <- struct{}{}
}()
go func() {
    <-ch
    fmt.Println(data["key"])
}()
```