# Demo with [Go Proverbs](https://go-proverbs.github.io/)


## 1. Don't communicate by sharing memory, share memory by communicating

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

## 2. Clear is better than clever
```
// Using bitwise operators for simple math => a+b
result := (a ^ b) + ((a & b) << 1) 

// a*b with bitwise operators
result  := (a << 1) + (a << 2) + (a << 3) // a*2 + a*4 + a*8 = a*14

```

## 3. Don't just check errors, handle them gracefully
```
# Bad
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()  


# Good
file, err := os.Open("file.txt")
if err != nil {
    log.Printf("Failed to open file: %v", err)
    return
}
defer file.Close()
```

## 4. Don't panic, recover
* Default in web server

```
# Bad
func riskyFunction() {
    panic("Something went wrong!")
}   

# Good
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Recovered from panic: %v", r)
        }
    }()
    riskyFunction()
}
```

## 5. The bigger the interface, the weaker the abstraction
```
# Bad
type DataStore interface {
    Save(data string) error
    Load(id string) (string, error)
    Delete(id string) error
}

# Good
type Saver interface {
    Save(data string) error
}
type Loader interface {
    Load(id string) (string, error)
}
type Deleter interface {
    Delete(id string) error
}
```

## 6. Documentation is for users
```
// Package counter provides a thread-safe integer tracker.
// Documentation is for users: focus on how to use the tool, not the internal logic.
package counter

// A Counter tracks a sequence of events. The zero value is a 
// ready-to-use counter starting at zero.
type Counter struct {
    val int
}

// Increment increases the counter by one. Use this to record 
// a new occurrence of an event.
func (c *Counter) Increment() {
    c.val++
}

// Value returns the current count. It is typically used for 
// final reporting or monitoring.
func (c *Counter) Value() int {
    return c.val
}
```