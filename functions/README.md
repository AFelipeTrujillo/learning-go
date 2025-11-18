# Functions

This section covers functions and methods in Go, which are essential for organizing and structuring code.

## Topics Covered

### 01. Using Functions
- Basic function declaration
- Functions with parameters
- Return values (single and multiple)
- Named return values
- Variadic functions (variable arguments)
- Functions as values and parameters
- Anonymous functions
- Closures
- Recursive functions
- Practical examples and patterns

### 02. Methods (Coming Soon)
- Method declarations
- Receiver types (value vs pointer)
- Methods vs functions
- Method sets
- Composition

### 03. Defer, Panic, Recover (Coming Soon)
- Defer statement
- Panic and recovery
- Resource cleanup patterns
- Error handling strategies

## Key Concepts

### Functions
Functions are **first-class citizens** in Go:
- Can be assigned to variables
- Passed as arguments
- Returned from other functions
- Created anonymously

### Function Signature
```go
func name(param1 type1, param2 type2) returnType {
    return value
}
```

Components:
- **func** - keyword
- **name** - identifier
- **parameters** - input (optional)
- **return type** - output (optional)
- **body** - implementation

## Common Patterns

### Pattern 1: Multiple Return Values (Error Handling)
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
```

### Pattern 2: Variadic Functions
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

sum(1, 2, 3)        // Can pass any number of arguments
sum(1, 2, 3, 4, 5)
```

### Pattern 3: Closures for State
```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

counter := makeCounter()
fmt.Println(counter())  // 1
fmt.Println(counter())  // 2
```

### Pattern 4: Higher-Order Functions
```go
func applyOperation(nums []int, op func(int) int) []int {
    result := make([]int, len(nums))
    for i, num := range nums {
        result[i] = op(num)
    }
    return result
}

doubled := applyOperation([]int{1, 2, 3}, func(n int) int {
    return n * 2
})
```

## Function Types

### By Purpose
- **Pure Functions** - No side effects, same input = same output
- **Impure Functions** - Have side effects (I/O, state changes)
- **Higher-Order Functions** - Take or return functions
- **Constructor Functions** - Create and return new instances

### By Complexity
- **Simple Functions** - Single responsibility
- **Complex Functions** - Multiple operations (should be refactored)
- **Helper Functions** - Support other functions
- **Utility Functions** - General-purpose tools

## Best Practices

### ✅ DO:
- **Single Responsibility** - One function, one job
- **Descriptive Names** - Clear what the function does
- **Handle Errors** - Return errors as second value
- **Keep Functions Short** - Ideally under 30 lines
- **Early Returns** - Check errors first, return early
- **Document Complex Functions** - Add comments

### ❌ DON'T:
- **God Functions** - Functions that do everything
- **Deep Nesting** - Too many nested if/loops
- **Side Effects** - Unexpected modifications
- **Magic Numbers** - Use named constants
- **Ignore Errors** - Always handle error returns

## Function Naming Conventions

### Standard Patterns
- **Getters**: `GetName()`, `Name()` (no "Get" prefix in Go)
- **Setters**: `SetName()`
- **Boolean checks**: `IsValid()`, `HasPermission()`, `CanAccess()`
- **Converters**: `ToString()`, `ToJSON()`
- **Validators**: `Validate()`, `Check()`
- **Constructors**: `NewUser()`, `MakeConfig()`

### Examples
```go
// Good naming
func CalculateTotal(price, tax float64) float64
func IsValidEmail(email string) bool
func NewConnection(host string) *Connection

// Bad naming
func calc(p, t float64) float64     // Unclear
func check(e string) bool            // Too generic
func connection(h string) *Connection // Not a constructor pattern
```

## Performance Considerations

### Function Call Overhead
- Small overhead per call
- Don't over-optimize prematurely
- Compiler may inline small functions

### Pass by Value
- Go passes everything by value (copies)
- Cheap: basic types (int, bool, string)
- Expensive: large structs
- Solution: Use pointers for large data

### Recursive Functions
- Can cause stack overflow for deep recursion
- Consider iterative alternatives
- Use memoization for repeated calculations

## Common Use Cases

### 1. Data Transformation
```go
func transform(data []string) []string {
    result := make([]string, len(data))
    for i, s := range data {
        result[i] = strings.ToUpper(s)
    }
    return result
}
```

### 2. Validation
```go
func validateInput(input string) error {
    if input == "" {
        return errors.New("input cannot be empty")
    }
    if len(input) < 3 {
        return errors.New("input too short")
    }
    return nil
}
```

### 3. Computation
```go
func calculateDiscount(price float64, percent float64) float64 {
    return price * (1 - percent/100)
}
```

### 4. Formatting
```go
func formatName(first, last string) string {
    return fmt.Sprintf("%s %s", first, last)
}
```

## Testing Functions

Functions should be testable:

```go
// Function to test
func Add(a, b int) int {
    return a + b
}

// Test
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

## Exercises

1. **Function Basics**: Create functions for basic math operations
2. **String Utils**: Write utility functions for string manipulation
3. **Validators**: Build input validation functions
4. **Converters**: Create type conversion functions
5. **Calculator**: Build a calculator using function composition
6. **Filter/Map/Reduce**: Implement functional programming patterns
7. **Memoization**: Optimize recursive functions with caching

## Comparison with Other Languages

| Feature | Go | Python | JavaScript |
|---------|-----|--------|------------|
| Multiple returns | ✅ Native | ❌ Use tuples | ❌ Use objects/arrays |
| Named returns | ✅ Yes | ❌ No | ❌ No |
| Variadic | ✅ `...type` | ✅ `*args` | ✅ `...args` |
| First-class | ✅ Yes | ✅ Yes | ✅ Yes |
| Closures | ✅ Yes | ✅ Yes | ✅ Yes |
| Overloading | ❌ No | ❌ No | ❌ No |

## Related Topics
- [Variables](../basics/04-variables/README.md) - Function scope
- [Control Flow](../control-flow/README.md) - Logic within functions
- [Pointers](../pointers/README.md) - Pass by reference
- [Interfaces](../interfaces/README.md) - Function signatures
- [Error Handling](../error-handling/README.md) - Error returns

## Additional Resources
- [Go by Example: Functions](https://gobyexample.com/functions)
- [A Tour of Go: Functions](https://go.dev/tour/basics/4)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)
- [Go Specification: Function declarations](https://go.dev/ref/spec#Function_declarations)
- [Go Blog: Error handling](https://go.dev/blog/error-handling-and-go)

