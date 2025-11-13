# Switch Statement in Go

## Description
The `switch` statement in Go is a powerful control flow structure that provides a clean way to handle multiple conditional branches. Go's switch is more flexible than in many other languages and has unique features that make it particularly useful.

## Key Concepts

### 1. Basic Switch
- **Simple syntax**: Evaluates an expression and matches against case values
- **Automatic break**: No `break` needed - each case automatically exits
- **Default case**: Optional catch-all for unmatched values

### 2. Multiple Values per Case
- **Comma-separated values**: One case can match multiple values
- **Cleaner code**: Replaces multiple if conditions with OR operators

### 3. Expression Switch
- **Cases with expressions**: Cases can contain boolean expressions
- **Range checking**: Perfect for range-based conditions

### 4. Switch without Condition
- **Acts like if-else**: Switch without an expression evaluates each case condition
- **Cleaner alternative**: More readable than long if-else chains

### 5. Type Switch
- **Type assertion**: Check the type of an interface variable
- **Type-specific handling**: Execute different code based on type

### 6. Fallthrough
- **Explicit continuation**: Forces execution to continue to next case
- **Rare usage**: Should be used sparingly as it's counter to Go's design

## Syntax

### Basic Switch
```go
switch expression {
case value1:
    // code for value1
case value2:
    // code for value2
default:
    // code if no match
}
```

### Multiple Values in Case
```go
switch expression {
case value1, value2, value3:
    // code for any of these values
case value4, value5:
    // code for these values
}
```

### Expression Switch
```go
switch {
case condition1:
    // code if condition1 is true
case condition2:
    // code if condition2 is true
default:
    // code if no condition is true
}
```

### Switch with Initialization
```go
switch variable := expression; variable {
case value1:
    // variable only exists in switch block
}
```

### Type Switch
```go
switch v := x.(type) {
case int:
    // v is an int
case string:
    // v is a string
default:
    // v is another type
}
```

### Fallthrough
```go
switch value {
case 1:
    // code
    fallthrough
case 2:
    // also executed if value is 1
}
```

## Key Differences from Other Languages

### Automatic Break
Unlike C, Java, or JavaScript, Go automatically breaks after each case:

```go
// Go - no break needed
switch x {
case 1:
    fmt.Println("one")  // automatically breaks
case 2:
    fmt.Println("two")
}

// C/Java - needs explicit break
// switch(x) {
//     case 1:
//         printf("one");
//         break;  // without this, falls through
//     case 2:
//         printf("two");
//         break;
// }
```

### No Limitations on Case Values
Go allows any comparable type and even expressions:

```go
switch {
case score >= 90:  // Expression, not just constant
    grade = "A"
case score >= 80:
    grade = "B"
}
```

### Multiple Values per Case
Cleaner than many if conditions:

```go
// Clean
switch month {
case "Dec", "Jan", "Feb":
    season = "Winter"
}

// Instead of
if month == "Dec" || month == "Jan" || month == "Feb" {
    season = "Winter"
}
```

## Code Examples
The `main.go` file demonstrates:
1. Basic switch with cases
2. Multiple values in single case
3. Switch with expressions
4. Switch without condition (like if-else)
5. Switch with initialization
6. Type switch for interface types
7. Fallthrough keyword
8. Real-world practical examples:
   - HTTP status code handler
   - File extension handler
   - User role permissions
   - Payment method processor
   - Traffic light controller
   - Grade calculator
   - Command line interface

## Running the Code
```bash
cd control-flow/02-switch-statement
go run main.go
```

## When to Use Switch vs If-Else

### Use Switch When:
- ✅ Checking **one variable** against multiple values
- ✅ Matching **exact values** (equality checks)
- ✅ Handling **many cases** (3+ conditions)
- ✅ Working with **type assertions**
- ✅ Code **readability** matters (cleaner than if-else chain)

### Use If-Else When:
- ✅ Checking **different variables** in each condition
- ✅ Complex **boolean logic** (multiple && or ||)
- ✅ **One or two** simple conditions
- ✅ **Range overlaps** (though switch can handle this too)

## Examples

### Example 1: Days of Week
```go
switch day {
case 1, 2, 3, 4, 5:
    fmt.Println("Weekday")
case 6, 7:
    fmt.Println("Weekend")
}
```

### Example 2: Grade Ranges
```go
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
default:
    grade = "F"
}
```

### Example 3: Type Checking
```go
switch v := value.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
}
```

## Exercises

1. **Calculator**: Create a calculator that uses switch to handle operations (+, -, *, /, %)

2. **Month Days**: Write a program that takes a month number (1-12) and returns the number of days in that month

3. **Rock Paper Scissors**: Implement a rock-paper-scissors game using switch to determine the winner

4. **Menu System**: Create a restaurant menu system that displays different options based on user selection

5. **Temperature Converter**: Build a converter that uses switch to convert between Celsius, Fahrenheit, and Kelvin

6. **Vowel Counter**: Write a program that checks if a character is a vowel using a switch statement

7. **Vehicle Classifier**: Create a program that classifies vehicles by number of wheels (bicycle, car, truck, etc.)

## Best Practices

### ✅ DO:
- Use switch for multiple value checks on same variable
- Take advantage of multiple values per case
- Use expression switch for range checking
- Add default case for unexpected values
- Keep cases simple and focused

### ❌ DON'T:
- Overuse fallthrough (it's rarely needed)
- Create overly complex case expressions
- Use switch when simple if-else is clearer
- Forget the default case for error handling
- Duplicate logic across cases

## Common Patterns

### Pattern 1: Command Dispatcher
```go
switch command {
case "start", "run":
    startServer()
case "stop", "halt":
    stopServer()
case "restart":
    stopServer()
    startServer()
}
```

### Pattern 2: State Machine
```go
switch currentState {
case "idle":
    // handle idle state
case "running":
    // handle running state
case "paused":
    // handle paused state
}
```

### Pattern 3: Error Code Handler
```go
switch err := doSomething(); err {
case nil:
    // success
case ErrNotFound:
    // handle not found
case ErrPermission:
    // handle permission error
default:
    // handle other errors
}
```

## Performance Considerations

- **Jump table optimization**: Compiler may optimize switch with integer cases into a jump table
- **Order doesn't matter**: Unlike if-else, case order doesn't affect performance
- **Type switch overhead**: Minimal overhead for type assertions
- **Expression evaluation**: Each case expression is evaluated only if previous cases didn't match

## Common Mistakes

### ❌ Forgetting that switch breaks automatically
```go
// This is WRONG - no break needed
switch x {
case 1:
    fmt.Println("one")
    break  // Unnecessary!
}
```

### ❌ Expecting fallthrough by default
```go
// This WON'T fall through to case 2
switch x {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")  // Only prints if x == 2
}
```

### ❌ Using fallthrough incorrectly
```go
// fallthrough can't have a condition
switch x {
case 1:
    fallthrough case 2:  // WRONG!
}

// Correct:
switch x {
case 1:
    fallthrough
case 2:
    // code
}
```

## Comparison: Switch vs If-Else

| Feature | Switch | If-Else |
|---------|--------|---------|
| Readability | Better for many cases | Better for few cases |
| Performance | Optimized by compiler | Linear evaluation |
| Flexibility | Best for equality | Best for complex conditions |
| Type checking | Supports type switch | Needs type assertion |
| Multiple values | Easy (comma-separated) | Needs OR operators |

## Related Topics
- [If-Else Statement](../01-if-and-else-statement/README.md) - Alternative control flow
- [Loops](../03-loops/README.md) - Iteration control
- [Variables](../../basics/04-variables/README.md) - Understanding types for type switch
- [Interfaces](../../interfaces/README.md) - Used with type switch

## Additional Resources
- [Go by Example: Switch](https://gobyexample.com/switch)
- [A Tour of Go: Switch](https://go.dev/tour/flowcontrol/9)
- [Go Specification: Switch statements](https://go.dev/ref/spec#Switch_statements)
- [Effective Go: Switch](https://go.dev/doc/effective_go#switch)

