# Go Cheatsheet

Quick reference for common Go syntax and patterns.

## Basic Syntax

### Package and Imports
```go
package main

import (
    "fmt"
    "strings"
)
```

### Variables
```go
// Explicit type
var name string = "John"
var age int = 30

// Type inference
var city = "New York"

// Short declaration (inside functions only)
country := "USA"

// Multiple variables
var x, y, z int = 1, 2, 3
a, b := 10, "hello"
```

### Constants
```go
const Pi = 3.14
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

## Data Types

### Basic Types
```go
bool                    // true or false
string                  // "hello"
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64
byte                    // alias for uint8
rune                    // alias for int32 (Unicode code point)
float32 float64
complex64 complex128
```

### Composite Types
```go
// Array (fixed size)
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slice (dynamic size)
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)

// Map
m := make(map[string]int)
m["age"] = 30

// Struct
type Person struct {
    Name string
    Age  int
}
p := Person{Name: "John", Age: 30}
```

## Control Flow

### If-Else
```go
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// If with initialization
if err := doSomething(); err != nil {
    return err
}
```

### Switch
```go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("TGIF")
default:
    fmt.Println("Regular day")
}

// Type switch
switch v := x.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
}
```

### Loops
```go
// Standard for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop
for condition {
    // code
}

// Infinite loop
for {
    // code
    break
}

// Range over slice
for index, value := range slice {
    fmt.Println(index, value)
}

// Range over map
for key, value := range m {
    fmt.Println(key, value)
}
```

## Functions

### Basic Function
```go
func add(a int, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // naked return
}
```

### Variadic Functions
```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
```

### Anonymous Functions
```go
func() {
    fmt.Println("anonymous function")
}()

// Function as variable
add := func(a, b int) int {
    return a + b
}
```

## Methods
```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

## Interfaces
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Empty interface (any type)
var x interface{}
x = 42
x = "hello"
```

## Error Handling
```go
// Creating errors
err := errors.New("something went wrong")
err := fmt.Errorf("failed to process: %v", data)

// Checking errors
if err != nil {
    return err
}

// Custom error type
type MyError struct {
    Msg string
}

func (e *MyError) Error() string {
    return e.Msg
}
```

## Pointers
```go
// Declare pointer
var p *int

// Get address
x := 42
p = &x

// Dereference
fmt.Println(*p) // 42

// Modify through pointer
*p = 100
```

## Concurrency

### Goroutines
```go
// Start goroutine
go functionName()

go func() {
    // anonymous goroutine
}()
```

### Channels
```go
// Create channel
ch := make(chan int)

// Buffered channel
ch := make(chan int, 10)

// Send to channel
ch <- 42

// Receive from channel
value := <-ch

// Close channel
close(ch)

// Range over channel
for value := range ch {
    fmt.Println(value)
}
```

### Select
```go
select {
case msg1 := <-ch1:
    fmt.Println("received", msg1)
case msg2 := <-ch2:
    fmt.Println("received", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no message")
}
```

## Common Packages

### fmt
```go
fmt.Println("Hello")          // Print with newline
fmt.Printf("Value: %d\n", 42) // Formatted print
fmt.Sprintf("Value: %d", 42)  // Formatted string
```

### strings
```go
strings.Contains(s, substr)
strings.Split(s, sep)
strings.Join(slice, sep)
strings.ToLower(s)
strings.ToUpper(s)
strings.TrimSpace(s)
```

### strconv
```go
strconv.Atoi("42")           // string to int
strconv.Itoa(42)             // int to string
strconv.ParseFloat("3.14", 64)
```

## Common Patterns

### Defer
```go
func readFile() {
    f, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer f.Close() // Executed at function exit
    
    // work with file
}
```

### Error Wrapping (Go 1.13+)
```go
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

### Testing
```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

