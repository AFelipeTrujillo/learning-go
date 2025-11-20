# Defer in Go

## Description
`defer` is a powerful keyword in Go that postpones the execution of a function until the surrounding function returns. It's commonly used for cleanup tasks, resource management, and ensuring code executes regardless of how a function exits.

## Key Concept

**Defer delays function execution until the containing function returns**

```go
func example() {
    defer fmt.Println("This executes last")
    fmt.Println("This executes first")
}
```

### When Defer Executes:
1. After the return statement
2. Before the function actually returns to caller
3. Even if a panic occurs (before panic propagates)

## 1. Basic Defer Syntax

```go
func basicExample() {
    defer fmt.Println("Deferred")
    fmt.Println("Normal")
    // Output:
    // Normal
    // Deferred
}
```

### Key Points:
- Use `defer` keyword before function call
- Executes after function returns
- Perfect for cleanup operations

## 2. Defer Execution Order (LIFO)

Multiple `defer` statements execute in **Last In, First Out** (LIFO) order:

```go
func orderExample() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    
    // Output:
    // Third defer
    // Second defer
    // First defer
}
```

### Why LIFO?
- Natural cleanup order
- Resources opened last are closed first
- Mirrors resource allocation pattern

### Visual Representation:
```
defer A()   → Executes 3rd
defer B()   → Executes 2nd
defer C()   → Executes 1st
return      ← Function returns after all defers
```

## 3. Defer with Parameters

**Important**: Parameters are evaluated immediately, not when defer executes:

```go
func parameterExample() {
    x := 10
    defer fmt.Println(x)  // Captures 10 immediately
    
    x = 20
    // Prints: 10 (not 20)
}
```

### Using Closures to Capture Current Values:

```go
func closureExample() {
    x := 10
    defer func() {
        fmt.Println(x)  // Captures x by reference
    }()
    
    x = 20
    // Prints: 20
}
```

## 4. Common Use Cases

### Use Case 1: File Handling

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Ensures file is closed
    
    // Read file operations
    return nil
}
```

### Use Case 2: Mutex Locking

```go
func criticalSection() {
    mu.Lock()
    defer mu.Unlock()  // Ensures unlock
    
    // Critical section code
}
```

### Use Case 3: Database Transactions

```go
func performTransaction(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()  // Rollback if not committed
    
    // Transaction operations
    
    return tx.Commit()
}
```

### Use Case 4: Timing Function Execution

```go
func measureTime() {
    start := time.Now()
    defer func() {
        fmt.Printf("Duration: %v\n", time.Since(start))
    }()
    
    // Function operations
}
```

### Use Case 5: Logging

```go
func operation() {
    defer log.Println("Operation finished")
    log.Println("Operation started")
    
    // Operation code
}
```

## 5. Defer with Return Values

Defer can modify **named** return values:

```go
func example() (result int) {
    defer func() {
        result = result * 2  // Modifies return value
    }()
    
    result = 21
    return  // Returns 42 (21 * 2)
}
```

### Pattern: Error Wrapping

```go
func operation() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("operation failed: %w", err)
        }
    }()
    
    // Operation that might fail
    return someOperation()
}
```

## 6. Defer in Loops

**⚠️ Warning**: Defer in loops can cause problems!

### ❌ Problem: Memory Accumulation

```go
func badExample() {
    for i := 0; i < 1000; i++ {
        file, _ := os.Open(filename)
        defer file.Close()  // Accumulates 1000 defers!
    }
    // All files stay open until function returns
}
```

### ✅ Solution: Use Function

```go
func goodExample() {
    for i := 0; i < 1000; i++ {
        func() {
            file, _ := os.Open(filename)
            defer file.Close()  // Closes immediately
            // Use file
        }()
    }
}
```

### ✅ Alternative: Close Explicitly

```go
func alternativeExample() {
    for i := 0; i < 1000; i++ {
        file, _ := os.Open(filename)
        // Use file
        file.Close()  // Close immediately, no defer
    }
}
```

## 7. Defer with Panic and Recover

Defer is essential for panic recovery:

```go
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    
    result = a / b  // May panic if b == 0
    return
}
```

### Recovery Pattern:

```go
func recoverExample() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Recovered: %v", r)
        }
    }()
    
    panic("something went wrong")
}
```

## 8. Multiple Resource Cleanup

Defer makes cleanup elegant:

```go
func processFiles() error {
    file1, err := os.Open("file1.txt")
    if err != nil {
        return err
    }
    defer file1.Close()
    
    file2, err := os.Open("file2.txt")
    if err != nil {
        return err  // file1 still closes
    }
    defer file2.Close()
    
    // Process both files
    return nil
}
```

## Best Practices

### ✅ DO:

```go
// Use defer for cleanup
func goodExample() {
    resource := acquire()
    defer release(resource)
    // Use resource
}

// Use defer immediately after resource acquisition
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()

// Use defer for symmetry
func operation() {
    log.Println("Start")
    defer log.Println("End")
    // Operation
}

// Use defer with named returns for error wrapping
func operation() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("operation failed: %w", err)
        }
    }()
    return actualOperation()
}
```

### ❌ DON'T:

```go
// Don't defer in tight loops
for i := 0; i < 10000; i++ {
    defer someFunction()  // Bad!
}

// Don't defer expensive operations unnecessarily
defer time.Sleep(5 * time.Second)  // Bad idea

// Don't ignore defer's return value when it matters
defer file.Close()  // Ignores potential error
// Better:
defer func() {
    if err := file.Close(); err != nil {
        log.Printf("Error closing file: %v", err)
    }
}()

// Don't use defer for simple operations that should be immediate
x := 10
defer fmt.Println(x)  // Unnecessary complexity
fmt.Println(x)  // Just do it now
```

## Common Patterns

### Pattern 1: Function Entry/Exit Logging

```go
func loggedOperation() {
    log.Println("ENTER loggedOperation")
    defer log.Println("EXIT loggedOperation")
    
    // Operation code
}
```

### Pattern 2: Performance Measurement

```go
func measure(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}

func operation() {
    defer measure("operation")()
    // Operation code
}
```

### Pattern 3: Transaction Management

```go
func transact() error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // If we return error, transaction rolls back
    if err := step1(tx); err != nil {
        return err
    }
    if err := step2(tx); err != nil {
        return err
    }
    
    return tx.Commit()  // Commit overrides rollback
}
```

### Pattern 4: Mutex Pattern

```go
func withLock(fn func()) {
    mu.Lock()
    defer mu.Unlock()
    fn()
}
```

### Pattern 5: Cleanup Stack

```go
func setupResources() error {
    res1 := allocate1()
    defer cleanup1(res1)
    
    res2 := allocate2()
    defer cleanup2(res2)
    
    res3 := allocate3()
    defer cleanup3(res3)
    
    // Use resources
    // Cleanup happens in reverse: res3, res2, res1
    return nil
}
```

## Defer vs Finally

Coming from other languages:

| Language | Cleanup Mechanism |
|----------|-------------------|
| **Go** | `defer` |
| **Java** | `try-finally` |
| **Python** | `try-finally` or `with` |
| **C#** | `try-finally` or `using` |

### Go's Advantage:
- Cleaner syntax
- Can be placed anywhere
- Multiple defers allowed
- Natural LIFO order

## Performance Considerations

### Defer has small overhead:
- Function call overhead
- Closure allocation (for anonymous functions)
- Minimal in most cases

### When to avoid defer:
- **Hot paths**: Very frequently called code
- **Tight loops**: Use explicit cleanup or wrapper functions
- **Performance-critical code**: Measure first!

### Benchmarking:
```go
// Direct call
func BenchmarkDirect(b *testing.B) {
    for i := 0; i < b.N; i++ {
        x := 10
        // Use x
        _ = x
    }
}

// With defer
func BenchmarkDefer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        defer func(x int) {
            _ = x
        }(10)
    }
}
```

## Running the Code

```bash
cd functions/02-defer
go run main.go
```

## Code Examples

The `main.go` file demonstrates:
1. Basic defer usage
2. Defer execution order (LIFO)
3. Defer with parameters
4. Defer with anonymous functions
5. Resource cleanup patterns
6. Defer with return values
7. Defer in loops (good and bad)
8. Error handling with defer
9. Defer with panic recovery
10. Practical real-world examples

## Exercises

1. **File Operations**: Write a function that opens multiple files and ensures all are closed using defer

2. **Timer**: Create a function that measures execution time using defer

3. **Transaction**: Implement a transaction pattern with defer for rollback

4. **Cleanup Chain**: Write a function that allocates multiple resources with proper cleanup order

5. **Error Wrapping**: Use defer to wrap all errors with context information

6. **Mutex Practice**: Create a safe counter using mutex with defer

7. **Panic Recovery**: Write a function that recovers from panics using defer

8. **Log Wrapper**: Create a logging wrapper that logs function entry and exit

## Common Mistakes

### ❌ Deferring in wrong place

```go
// Wrong: defer before error check
file, err := os.Open("file.txt")
defer file.Close()  // file might be nil!
if err != nil {
    return err
}

// Correct: defer after error check
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()
```

### ❌ Not capturing current value

```go
// Wrong: prints last value only
for i := 0; i < 3; i++ {
    defer fmt.Println(i)  // Prints: 2, 1, 0
}

// Correct: capture each value
for i := 0; i < 3; i++ {
    i := i  // Shadow variable
    defer func() {
        fmt.Println(i)
    }()
}
```

### ❌ Ignoring defer errors

```go
// Problematic: ignores close error
defer file.Close()

// Better: handle close error
defer func() {
    if err := file.Close(); err != nil {
        log.Printf("Failed to close: %v", err)
    }
}()
```

## Related Topics
- [Functions](../01-using-function/README.md) - Function basics
- [Error Handling](../../error-handling/README.md) - Error patterns
- [Panic and Recover](../03-panic-recover/README.md) - Panic handling
- [Concurrency](../../concurrency/README.md) - Mutex and channels

## Additional Resources
- [Go by Example: Defer](https://gobyexample.com/defer)
- [Effective Go: Defer](https://go.dev/doc/effective_go#defer)
- [Go Blog: Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Go Specification: Defer statements](https://go.dev/ref/spec#Defer_statements)

