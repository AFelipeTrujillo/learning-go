# Functions in Go

## Description
Functions are fundamental building blocks in Go. They allow you to organize code into reusable pieces, making programs more modular, maintainable, and easier to understand.

## Key Concepts

### Function Declaration
A function is defined using the `func` keyword followed by the name, parameters, return type, and body.

**Basic Syntax:**
```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

### Components of a Function
1. **func keyword** - Declares a function
2. **Name** - Identifier for the function
3. **Parameters** - Input values (optional)
4. **Return type** - Type of value returned (optional)
5. **Body** - Code to execute
6. **Return statement** - Returns value (if return type specified)

## 1. Basic Function Declaration

### Function without Parameters or Return

```go
func greet() {
    fmt.Println("Hello, World!")
}

// Call the function
greet()
```

### Key Points:
- No parameters needed in parentheses: `()`
- No return type specified
- Simply executes code when called

## 2. Functions with Parameters

### Single Parameter

```go
func greetPerson(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

greetPerson("Alice")  // Output: Hello, Alice!
```

### Multiple Parameters

```go
func add(a int, b int) {
    sum := a + b
    fmt.Printf("%d + %d = %d\n", a, b, sum)
}

add(10, 20)  // Output: 10 + 20 = 30
```

### Same Type Parameters (Shorthand)

```go
// Instead of: func multiply(x int, y int)
func multiply(x, y int) {
    result := x * y
    fmt.Println(result)
}
```

## 3. Functions with Return Values

### Single Return Value

```go
func add(a, b int) int {
    return a + b
}

result := add(10, 20)  // result = 30
```

### Different Return Types

```go
func calculateArea(width, height float64) float64 {
    return width * height
}

func isEven(n int) bool {
    return n % 2 == 0
}
```

## 4. Multiple Return Values

One of Go's powerful features - functions can return multiple values:

```go
func calculate(a, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}

// Capture both return values
sum, product := calculate(4, 5)
fmt.Println(sum, product)  // 9 20
```

### Common Pattern: Value and Error

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

### Ignoring Return Values

Use underscore `_` to ignore unwanted return values:

```go
sum, _ := calculate(3, 7)      // Ignore product
_, product := calculate(3, 7)  // Ignore sum
```

## 5. Named Return Values

Go allows you to name return values:

```go
func rectangleStats(width, height float64) (perimeter float64, area float64) {
    perimeter = 2 * (width + height)
    area = width * height
    return  // Naked return
}
```

### Advantages:
- Acts as documentation
- Can use "naked return" (just `return`)
- Pre-declared variables

### When to Use:
- ✅ Short functions where it improves clarity
- ❌ Long functions (use explicit returns)

## 6. Variadic Functions

Functions that accept variable number of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

sum(1, 2, 3)           // 6
sum(1, 2, 3, 4, 5)     // 15
sum(10)                // 10
```

### Key Points:
- Use `...` before type: `...int`
- Must be the last parameter
- Treated as a slice inside function
- Can pass zero or more arguments

## 7. Functions as Values

Functions are first-class citizens in Go:

### Assign Function to Variable

```go
greeting := func(name string) string {
    return "Hello, " + name
}

fmt.Println(greeting("World"))  // Hello, World
```

### Pass Function as Parameter

```go
func applyOperation(nums []int, operation func(int) int) []int {
    result := make([]int, len(nums))
    for i, num := range nums {
        result[i] = operation(num)
    }
    return result
}

double := func(n int) int {
    return n * 2
}

numbers := []int{1, 2, 3}
doubled := applyOperation(numbers, double)  // [2, 4, 6]
```

## 8. Anonymous Functions

Functions without a name:

### Immediate Execution

```go
func() {
    fmt.Println("This runs immediately")
}()
```

### With Parameters

```go
func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}("World")
```

### Assigned to Variable

```go
add := func(a, b int) int {
    return a + b
}

result := add(5, 3)  // 8
```

## 9. Closures

Functions that capture and remember variables from their surrounding scope:

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
fmt.Println(counter())  // 3
```

### Use Cases:
- Maintaining state
- Creating specialized functions
- Callbacks
- Encapsulation

## 10. Recursive Functions

Functions that call themselves:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

factorial(5)  // 120
```

### Requirements:
1. **Base case** - Condition to stop recursion
2. **Recursive case** - Function calls itself
3. **Progress** - Move toward base case

## Running the Code

```bash
cd functions/01-using-function
go run main.go
```

## Code Examples

The `main.go` file demonstrates:
1. Basic function declaration
2. Functions with parameters
3. Functions with return values
4. Multiple return values
5. Ignoring return values
6. Named return values
7. Functions as parameters
8. Anonymous functions
9. Closures
10. Variadic functions
11. Recursive functions
12. Eight practical examples

## Function Signature

The **function signature** includes:
- Function name
- Parameter types
- Return types

```go
func add(a int, b int) int  // Signature: add(int, int) int
```

## Best Practices

### ✅ DO:

```go
// Use descriptive names
func calculateTotalPrice(price, quantity float64) float64 {
    return price * quantity
}

// Keep functions focused (single responsibility)
func validateEmail(email string) bool {
    // Just validate email
    return /* validation logic */
}

// Return errors as second value
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Use early returns for error cases
func processData(data string) error {
    if data == "" {
        return errors.New("empty data")
    }
    if len(data) < 10 {
        return errors.New("data too short")
    }
    // process data
    return nil
}
```

### ❌ DON'T:

```go
// Don't make functions too long
func doEverything() {  // Bad
    // 200 lines of code
}

// Don't use unclear names
func f(x, y int) int {  // Bad
    return x + y
}

// Don't ignore errors
result, _ := divide(10, 0)  // Bad! Could panic

// Don't use naked returns in long functions
func longFunction() (a, b, c int) {
    // 50 lines of code
    return  // Hard to see what's being returned
}
```

## Common Patterns

### Pattern 1: Error Checking

```go
func doSomething() error {
    if err := validate(); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    // Continue processing
    return nil
}
```

### Pattern 2: Factory Function

```go
func NewUser(name string, age int) *User {
    return &User{
        Name: name,
        Age:  age,
        ID:   generateID(),
    }
}
```

### Pattern 3: Option Functions

```go
type Config struct {
    timeout int
    retries int
}

func WithTimeout(timeout int) func(*Config) {
    return func(c *Config) {
        c.timeout = timeout
    }
}

func NewConfig(opts ...func(*Config)) *Config {
    config := &Config{timeout: 30, retries: 3}
    for _, opt := range opts {
        opt(config)
    }
    return config
}
```

## Exercises

1. **Calculator**: Create functions for add, subtract, multiply, divide with error handling

2. **String Utils**: Write functions for:
   - Reverse string
   - Check palindrome
   - Count vowels
   - Title case

3. **Math Functions**: Implement:
   - Prime number checker
   - GCD (Greatest Common Divisor)
   - Fibonacci with memoization

4. **Array Functions**: Create:
   - Filter (keep elements matching condition)
   - Map (transform each element)
   - Reduce (combine elements into single value)

5. **Temperature Converter**: Functions to convert between Celsius, Fahrenheit, and Kelvin

6. **Factorial**: Implement both recursive and iterative versions

7. **Password Validator**: Function that checks multiple password requirements

8. **Grade Calculator**: Function that takes scores and returns letter grade and GPA

## Function Types

### By Parameters and Returns

| Type | Example |
|------|---------|
| No params, no return | `func greet()` |
| With params, no return | `func print(s string)` |
| No params, with return | `func getRandom() int` |
| With params and return | `func add(a, b int) int` |

### By Purpose

| Type | Description | Example |
|------|-------------|---------|
| **Pure** | No side effects, same input = same output | `func add(a, b int) int` |
| **Impure** | Has side effects (prints, modifies state) | `func saveToFile(data string)` |
| **Higher-order** | Takes or returns functions | `func map(fn func(int) int)` |
| **Recursive** | Calls itself | `func factorial(n int) int` |

## Common Mistakes

### ❌ Unused Return Values

```go
func getData() (string, error) {
    return "data", nil
}

getData()  // Ignoring error! Bad practice
```

### ❌ Modifying Parameters

```go
// This doesn't work as expected (pass by value)
func increment(x int) {
    x++  // Only modifies local copy
}

// Use pointers or return the value
func increment(x int) int {
    return x + 1
}
```

### ❌ Infinite Recursion

```go
func factorial(n int) int {
    return n * factorial(n-1)  // Missing base case!
}
```

## Memory and Performance

### Pass by Value
Go passes parameters by value (copies):
- Cheap for basic types (int, bool, etc.)
- Expensive for large structs
- Use pointers for large data structures

### Function Call Overhead
- Function calls have small overhead
- Don't over-optimize prematurely
- Inline small functions when needed

## Related Topics
- [Variables](../../basics/04-variables/README.md) - Understanding scope
- [Control Flow](../../control-flow/README.md) - Used within functions
- [Pointers](../../pointers/README.md) - Passing by reference
- [Methods](../02-methods/README.md) - Functions with receivers
- [Interfaces](../../interfaces/README.md) - Function signatures in interfaces

## Additional Resources
- [Go by Example: Functions](https://gobyexample.com/functions)
- [A Tour of Go: Functions](https://go.dev/tour/basics/4)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)
- [Go Specification: Function declarations](https://go.dev/ref/spec#Function_declarations)

