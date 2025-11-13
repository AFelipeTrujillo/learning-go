# Control Flow

This section covers control flow structures in Go that determine the order in which code is executed.

## Topics Covered

### 01. If and Else Statements
- Basic if statements
- If-else branching
- Else if for multiple conditions
- If with initialization
- Comparison operators (==, !=, <, >, <=, >=)
- Logical operators (&&, ||, !)
- Nested conditions
- Practical examples

### 02. Switch Statement
- Basic switch with cases
- Multiple values per case
- Expression-based switch
- Switch without condition
- Switch with initialization
- Type switch for interfaces
- Fallthrough keyword
- Real-world applications

### 03. For Loops
- Basic for loop
- While-style loop
- Infinite loop
- Range-based iteration
- Nested loops
- Loop control (break, continue)

### 04. Break and Continue
- Breaking out of loops
- Continuing to next iteration
- Labeled breaks
- Best practices

## Learning Approach

Each topic includes:
1. A README explaining the concepts and syntax
2. Comprehensive code examples with comments
3. Practical real-world scenarios
4. Exercises to practice
5. Common mistakes to avoid

## Key Differences from Other Languages

### If Statements
- **No parentheses** required around conditions
- **Braces always mandatory** (even for single statements)
- **Initialization in if** - unique Go feature
- **No ternary operator** - use if-else instead

### Switch Statements
- **Automatic break** - no need for explicit break
- **Multiple values** per case (comma-separated)
- **Expression cases** - not just constants
- **Type switch** - check interface types
- **No fallthrough by default** - must be explicit

### Loops
- **Only for loop** - no while or do-while
- **Range keyword** - iterate over collections easily
- **Multiple forms** - can mimic while and infinite loops

## Comparison with Other Languages

| Feature | Go | C/Java | Python |
|---------|-----|--------|--------|
| If parentheses | Not required | Required | Not used |
| If braces | Always required | Optional for single line | Indentation |
| Switch break | Automatic | Manual | No switch |
| Switch fallthrough | Explicit with keyword | Default behavior | N/A |
| While loop | Use for | Separate keyword | Separate keyword |
| Do-while | Use for | Separate keyword | Not available |
| Ternary operator | Not available | Available | Available |

## When to Use Each Structure

### Use If-Else When:
- Checking **different variables** in conditions
- **Complex boolean logic** with multiple && or ||
- **One or two conditions** only
- Conditions have **different types** of checks

### Use Switch When:
- Checking **one variable** against multiple values
- **Many cases** (3 or more)
- **Type checking** with interfaces
- **Exact value matching**
- Better **readability** compared to if-else chain

### Use For Loop When:
- **Iterating** a specific number of times
- Processing **collections** (arrays, slices, maps)
- **Repeating** until a condition is met
- Need **index and value** during iteration

## Best Practices

### Control Flow Design
1. **Keep it simple** - Avoid deeply nested structures
2. **Early returns** - Return early on error conditions
3. **Consistent style** - Follow Go conventions
4. **Guard clauses** - Check error conditions first
5. **Readable conditions** - Use meaningful variable names

### Common Patterns

#### Early Return Pattern
```go
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

#### Guard Clause Pattern
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

#### State Machine Pattern
```go
switch state {
case stateIdle:
    handleIdle()
case stateRunning:
    handleRunning()
case statePaused:
    handlePaused()
}
```

## Common Mistakes to Avoid

### ❌ Unnecessary Complexity
```go
// Bad
if condition == true {
    return true
} else {
    return false
}

// Good
return condition
```

### ❌ Comparing Booleans
```go
// Bad
if isActive == true {
    // code
}

// Good
if isActive {
    // code
}
```

### ❌ Deep Nesting
```go
// Bad
if condition1 {
    if condition2 {
        if condition3 {
            // code
        }
    }
}

// Good - use early returns
if !condition1 {
    return
}
if !condition2 {
    return
}
if !condition3 {
    return
}
// code
```

### ❌ Redundant else after return
```go
// Bad
if condition {
    return value1
} else {
    return value2
}

// Good
if condition {
    return value1
}
return value2
```

## Performance Considerations

- **Short-circuit evaluation**: Go stops evaluating as soon as result is known
- **Switch optimization**: Compiler may optimize switch into jump tables
- **Loop unrolling**: Compiler may optimize simple loops
- **Bounds check elimination**: Go eliminates redundant bounds checks in loops

## Practice Projects

1. **Calculator**: Build a calculator using switch for operations
2. **Grade System**: Create a grading system with if-else or switch
3. **Menu System**: Interactive menu using loops and switch
4. **Game Logic**: Simple game with state machine using switch
5. **Data Validator**: Validate user input with nested conditions

## Related Topics
- [Variables](../basics/04-variables/README.md) - Used in conditions
- [Functions](../functions/README.md) - Control flow within functions
- [Error Handling](../error-handling/README.md) - Conditional error checking
- [Defer, Panic, Recover](../error-handling/defer-panic-recover/README.md) - Special control flow

## Additional Resources
- [Go by Example: Control Flow](https://gobyexample.com/)
- [A Tour of Go: Flow Control](https://go.dev/tour/flowcontrol/1)
- [Effective Go: Control Structures](https://go.dev/doc/effective_go#control-structures)
- [Go Specification: Statements](https://go.dev/ref/spec#Statements)

