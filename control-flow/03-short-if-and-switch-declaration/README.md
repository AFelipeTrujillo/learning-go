# Short If and Switch Declaration

## Description
One of Go's unique and powerful features is the ability to declare variables directly in `if` and `switch` statements. This pattern promotes cleaner code by limiting variable scope to where it's actually needed.

## Key Concepts

### Short If Declaration
Declare a variable in the `if` statement itself, scoped to the if-else block.

**Syntax:**
```go
if variable := value; condition {
    // use variable here
}
```

### Short Switch Declaration
Declare a variable in the `switch` statement itself, scoped to the entire switch block.

**Syntax:**
```go
switch variable := value; variable {
case value1:
    // use variable here
}
```

### Benefits
1. **Limited Scope** - Variables don't pollute outer scope
2. **Cleaner Code** - Declaration and usage are close together
3. **Error Handling** - Perfect for checking errors immediately
4. **Readability** - Intent is clear and concise

## 1. Short If Declaration

### Basic Syntax

```go
// Traditional way
age := 25
if age >= 18 {
    fmt.Println("Adult")
}
// age is still accessible here

// Short declaration
if age := 25; age >= 18 {
    fmt.Println("Adult")
}
// age is NOT accessible here
```

### With Else

```go
if score := 85; score >= 90 {
    fmt.Println("Grade A")
} else if score >= 80 {
    fmt.Println("Grade B")
} else {
    fmt.Println("Grade C or below")
}
// score is only accessible within the if-else-if chain
```

### With Function Calls

```go
if result := calculateSum(10, 20); result > 25 {
    fmt.Printf("Sum is %d\n", result)
}
```

### Multiple Variables

```go
if x, y := 10, 20; x < y {
    fmt.Printf("x=%d is less than y=%d\n", x, y)
}
```

## 2. Error Handling Pattern

This is one of the **most common patterns** in Go:

### Basic Error Checking

```go
if err := doSomething(); err != nil {
    // Handle error
    return err
}
```

### With Multiple Return Values

```go
if value, err := strconv.Atoi("123"); err != nil {
    fmt.Println("Conversion error:", err)
} else {
    fmt.Println("Converted value:", value)
}
```

### File Operations

```go
if file, err := os.Open("data.txt"); err != nil {
    log.Fatal("Cannot open file:", err)
} else {
    defer file.Close()
    // Use file
}
```

### Real-World Example

```go
if data, err := json.Marshal(user); err != nil {
    return fmt.Errorf("failed to marshal user: %w", err)
} else {
    return string(data), nil
}
```

## 3. Short Switch Declaration

### Basic Syntax

```go
// Traditional way
day := time.Now().Weekday()
switch day {
case time.Monday:
    fmt.Println("Monday")
}

// Short declaration
switch day := time.Now().Weekday(); day {
case time.Monday:
    fmt.Println("Monday")
}
// day is NOT accessible here
```

### With Expression Switch

```go
switch hour := time.Now().Hour(); {
case hour < 12:
    fmt.Println("Morning")
case hour < 18:
    fmt.Println("Afternoon")
default:
    fmt.Println("Evening")
}
```

### With Function Calls

```go
switch result := getStatus(); result {
case "success":
    fmt.Println("Operation successful")
case "error":
    fmt.Println("Operation failed")
default:
    fmt.Println("Unknown status")
}
```

### With Calculations

```go
switch score := calculateScore(); {
case score >= 90:
    fmt.Println("Excellent")
case score >= 70:
    fmt.Println("Good")
default:
    fmt.Println("Needs improvement")
}
```

## 4. Variable Scope

### Scope Demonstration

```go
if temp := 25; temp > 20 {
    fmt.Println(temp)  // ✓ Accessible
} else {
    fmt.Println(temp)  // ✓ Accessible
}
fmt.Println(temp)  // ✗ NOT accessible (compiler error)
```

### Nested Scope

```go
if x := 10; x > 5 {
    if x > 8 {
        fmt.Println(x)  // ✓ Still accessible
    }
}
```

### Switch Scope

```go
switch num := 10; {
case num > 5:
    fmt.Println(num)  // ✓ Accessible
default:
    fmt.Println(num)  // ✓ Accessible
}
fmt.Println(num)  // ✗ NOT accessible
```

## 5. When to Use

### ✅ Use Short Declaration When:
- Variable is only needed in that if/switch block
- Checking errors from function calls
- Processing temporary calculation results
- Validating input immediately
- Scope limitation improves code clarity

### ❌ Don't Use When:
- Variable is needed outside the block
- Makes code less readable
- Multiple uses of the variable across different blocks
- Simple condition doesn't benefit from it

## Examples Comparison

### Example 1: Traditional vs Short

```go
// Traditional - variable pollutes outer scope
password := getPassword()
if len(password) < 8 {
    fmt.Println("Password too short")
}
// password still accessible but no longer needed

// Short declaration - cleaner
if password := getPassword(); len(password) < 8 {
    fmt.Println("Password too short")
}
// password not accessible, scope is limited
```

### Example 2: Error Handling

```go
// Traditional - more verbose
data, err := fetchData()
if err != nil {
    return err
}
processData(data)

// Short declaration - cleaner for simple checks
if data, err := fetchData(); err != nil {
    return err
} else {
    processData(data)
}
```

### Example 3: Switch with Calculation

```go
// Traditional
bmi := weight / (height * height)
switch {
case bmi < 18.5:
    fmt.Println("Underweight")
case bmi < 25:
    fmt.Println("Normal")
default:
    fmt.Println("Overweight")
}

// Short declaration
switch bmi := weight / (height * height); {
case bmi < 18.5:
    fmt.Println("Underweight")
case bmi < 25:
    fmt.Println("Normal")
default:
    fmt.Println("Overweight")
}
```

## Practical Use Cases

### 1. Input Validation

```go
if username := getUserInput(); len(username) < 3 {
    return errors.New("username too short")
} else if len(username) > 20 {
    return errors.New("username too long")
}
```

### 2. HTTP Status Handling

```go
switch status := response.StatusCode; {
case status >= 200 && status < 300:
    handleSuccess()
case status >= 400 && status < 500:
    handleClientError()
case status >= 500:
    handleServerError()
}
```

### 3. File Size Classification

```go
if size := getFileSize(); size < 1024 {
    fmt.Println("Small file")
} else if size < 1024*1024 {
    fmt.Printf("%.2f KB\n", float64(size)/1024)
} else {
    fmt.Printf("%.2f MB\n", float64(size)/(1024*1024))
}
```

### 4. Time-Based Logic

```go
switch hour := time.Now().Hour(); {
case hour >= 5 && hour < 12:
    sendMorningNewsletter()
case hour >= 12 && hour < 18:
    sendAfternoonUpdate()
default:
    sendEveningDigest()
}
```

### 5. User Role Check

```go
switch role := user.GetRole(); role {
case "admin":
    grantFullAccess()
case "moderator":
    grantModeratorAccess()
default:
    grantStandardAccess()
}
```

## Running the Code

```bash
cd control-flow/03-short-if-and-switch-declaration
go run main.go
```

## Code Examples

The `main.go` file demonstrates:
1. Basic short if declaration
2. Short if with else branches
3. Short if with function calls
4. Error handling pattern (very important!)
5. Multiple variable declarations
6. Basic short switch declaration
7. Short switch with expressions
8. Short switch with function calls
9. Variable scope demonstration
10. Ten practical real-world examples

## Exercises

1. **Password Validator**: Write a password validator using short if that checks:
   - Length (8-20 characters)
   - Contains numbers
   - Contains special characters

2. **Grade Classifier**: Use short switch to classify grades and calculate GPA

3. **File Type Detector**: Use short switch with file extension to determine file type

4. **Age Category**: Use short if to categorize ages (child, teen, adult, senior)

5. **Temperature Converter**: Create a converter using short if for range-based advice

6. **HTTP Response Handler**: Use short switch to handle different HTTP status codes

7. **BMI Calculator**: Calculate BMI using short if and provide health advice

8. **Number Analyzer**: Use short if to check if a number is even/odd, positive/negative, prime

## Best Practices

### ✅ DO:

```go
// Use for error checking
if err := doSomething(); err != nil {
    return err
}

// Use when variable is only needed in block
if score := calculateScore(); score > 90 {
    awardBonus()
}

// Use for immediate validation
if age := getAge(); age < 18 {
    return errors.New("too young")
}
```

### ❌ DON'T:

```go
// Don't use if variable is needed later
if data := fetchData(); isValid(data) {
    // process data
}
// Need data here too? Use traditional declaration instead

// Don't make it too complex
if x, y, z := calc1(), calc2(), calc3(); x+y*z > someComplexFunc(x, y, z) {
    // Too complicated, hard to read
}

// Don't nest too deeply
if a := getA(); a > 0 {
    if b := getB(); b > 0 {
        if c := getC(); c > 0 {
            // Too many levels
        }
    }
}
```

## Common Patterns

### Pattern 1: Immediate Error Check

```go
if err := validateInput(input); err != nil {
    return fmt.Errorf("validation failed: %w", err)
}
```

### Pattern 2: Range Validation

```go
if value := getValue(); value < min || value > max {
    return errors.New("out of range")
}
```

### Pattern 3: State Check

```go
switch state := getState(); state {
case stateReady:
    start()
case stateRunning:
    continue
case stateError:
    handleError()
}
```

### Pattern 4: Conversion with Check

```go
if num, err := strconv.Atoi(str); err == nil {
    processNumber(num)
}
```

## Common Mistakes

### ❌ Trying to access variable outside scope

```go
if age := 25; age >= 18 {
    fmt.Println("Adult")
}
fmt.Println(age)  // Error: age not defined
```

### ❌ Shadowing outer variables unintentionally

```go
age := 20
if age := 25; age >= 18 {  // This shadows outer age
    fmt.Println(age)  // Prints 25, not 20
}
fmt.Println(age)  // Prints 20
```

### ❌ Overly complex initialization

```go
// Too complex
if x, y, z := func1(), func2(), func3(); x > y && y > z && someComplexCheck(x, y, z) {
    // Hard to read
}

// Better
x, y, z := func1(), func2(), func3()
if x > y && y > z && someComplexCheck(x, y, z) {
    // Easier to read
}
```

## Advantages over Traditional Approach

| Aspect | Traditional | Short Declaration |
|--------|------------|-------------------|
| Scope | Wider (outer scope) | Limited (block only) |
| Memory | Variable persists | Variable freed sooner |
| Clarity | Declaration separate | Declaration + use together |
| Error handling | More verbose | More concise |
| Namespace pollution | Higher risk | Lower risk |

## Comparison with Other Languages

Go's short declaration in if/switch is relatively unique:

| Language | Feature |
|----------|---------|
| **Go** | ✅ if/switch with initialization |
| **C++** | ✅ if with initialization (C++17+) |
| **Rust** | ✅ if let and match guards |
| **Python** | ❌ No direct equivalent |
| **Java** | ❌ No direct equivalent |
| **JavaScript** | ❌ No direct equivalent |

## Related Topics
- [If-Else Statements](../01-if-and-else-statement/README.md) - Basic conditional logic
- [Switch Statement](../02-switch-statement/README.md) - Multi-way branching
- [Variables](../../basics/04-variables/README.md) - Variable scope and declaration
- [Error Handling](../../error-handling/README.md) - Error checking patterns

## Additional Resources
- [Go Specification: If statements](https://go.dev/ref/spec#If_statements)
- [Go Specification: Switch statements](https://go.dev/ref/spec#Switch_statements)
- [Effective Go: Control structures](https://go.dev/doc/effective_go#control-structures)
- [Go Blog: Error handling](https://go.dev/blog/error-handling-and-go)

