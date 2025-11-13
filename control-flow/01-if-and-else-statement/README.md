# If and Else Statements in Go

## Description
`if` and `else` statements are fundamental control flow structures in Go that allow you to execute code conditionally based on whether a boolean expression is true or false.

## Key Concepts

### 1. Basic If Statement
- **Simple syntax**: No parentheses required around the condition
- **Mandatory braces**: Curly braces `{}` are always required, even for a single line
- **Boolean condition**: The expression must evaluate to `true` or `false`

### 2. If-Else
- **Binary branching**: Executes one block if the condition is true, another if false
- **Clear syntax**: The `else` must be on the same line as the closing brace of the `if`

### 3. Else If
- **Multiple conditions**: Allows evaluating several conditions in sequence
- **Ordered evaluation**: Evaluated from top to bottom, stopping at the first true condition

### 4. If Statement with Initialization
- **Go's unique feature**: Allows declaring a variable before the condition
- **Limited scope**: The variable only exists within the if-else block
- **Syntax**: `if variable := value; condition { ... }`

### 5. Comparison Operators
- `==` : Equal to
- `!=` : Not equal to
- `<`  : Less than
- `<=` : Less than or equal to
- `>`  : Greater than
- `>=` : Greater than or equal to

### 6. Logical Operators
- `&&` : Logical AND
- `||` : Logical OR
- `!`  : Logical NOT (negation)

## Syntax

### Basic If
```go
if condition {
    // code to execute if condition is true
}
```

### If-Else
```go
if condition {
    // code if true
} else {
    // code if false
}
```

### If-Else If-Else
```go
if condition1 {
    // code if condition1 is true
} else if condition2 {
    // code if condition2 is true
} else {
    // code if no condition is true
}
```

### If with Initialization
```go
if variable := value; condition {
    // variable only exists here
}
```

## Code Examples
The `main.go` file demonstrates:
1. Basic if statements
2. Simple if-else
3. Multiple else if
4. If with initialization
5. Comparison operators
6. Logical operators (&&, ||, !)
7. Nested conditions
8. Real-world practical examples

## Running the Code
```bash
cd control-flow/01-if-and-else-statement
go run main.go
```

## Expected Output
The program will display various examples of if-else statements in action:
- Age verification
- Grade classification
- Temperature evaluation
- Password validation
- Authentication logic
- And more...

## Exercises

1. **BMI Calculator**: Write a program that calculates the Body Mass Index (BMI) and determines the category (underweight, normal, overweight, obese).

2. **Number Classifier**: Create a program that takes a number and determines if it is:
   - Positive, negative, or zero
   - Even or odd
   - Multiple of 3, 5, or both

3. **Discount System**: Implement a system that calculates discounts based on:
   - Product quantity (volume discount)
   - Whether the customer is a premium member
   - If it's sale season

4. **Date Validator**: Write a program that validates if a date is correct:
   - Verify the month (1-12)
   - Verify the day according to the month
   - Consider leap years

5. **Password Analyzer**: Create a validator that checks if a password is secure:
   - Minimum 8 characters
   - Contains uppercase and lowercase letters
   - Contains numbers
   - Contains special characters

## Notes

### Best Practices
- **Don't use parentheses** around conditions (not necessary in Go)
- **Always use braces** even for single-line blocks
- **Keep conditions simple** and readable
- **Use descriptive variable names** in conditions
- **Take advantage of initialization** in if when the variable is only needed in that context

### Common Mistakes
- ❌ Forgetting the mandatory braces
- ❌ Using `=` instead of `==` for comparisons
- ❌ Putting `else` on a new line (must be after `}`)
- ❌ Not understanding the scope of variables initialized in the if

### Differences from Other Languages
- **No parentheses**: Unlike C, Java, or JavaScript, Go doesn't require parentheses
- **Mandatory braces**: Unlike Python, Go always requires braces
- **Built-in initialization**: Unique feature that not all languages have
- **No ternary operator**: Go doesn't have `? :`, you must use if-else

### When to Use If vs Switch
- Use **if** for:
  - Few conditions (1-3)
  - Complex boolean expressions
  - Range comparisons
  - Different conditions per branch
  
- Use **switch** for:
  - Many conditions on the same variable
  - Value matching
  - Cleaner code for multiple cases

## Truth Table for Logical Operators

### AND (&&)
| A | B | A && B |
|---|---|--------|
| true | true | true |
| true | false | false |
| false | true | false |
| false | false | false |

### OR (||)
| A | B | A \|\| B |
|---|---|--------|
| true | true | true |
| true | false | true |
| false | true | true |
| false | false | false |

### NOT (!)
| A | !A |
|---|-----|
| true | false |
| false | true |

## Performance Tips
- **Short-circuit evaluation**: Go uses short-circuit evaluation for `&&` and `||`
  - `a && b`: If `a` is false, `b` is never evaluated
  - `a || b`: If `a` is true, `b` is never evaluated
- **Order matters**: Put the most likely condition first for better performance
- **Avoid nested ifs**: Consider using logical operators or early returns

## Related Topics
- [Variables](../../basics/04-variables/README.md) - Necessary for conditions
- [Operators](../../basics/operators/README.md) - Comparison and logical operators
- [Switch Statement](../02-switch-statement/README.md) - Alternative for multiple conditions
- [Loops](../03-loops/README.md) - Another control flow structure

## Additional Resources
- [Go by Example: If/Else](https://gobyexample.com/if-else)
- [A Tour of Go: Control Flow](https://go.dev/tour/flowcontrol/5)
- [Go Documentation: If statements](https://go.dev/ref/spec#If_statements)
- [Effective Go: Control structures](https://go.dev/doc/effective_go#control-structures)
