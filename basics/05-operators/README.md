# Operators in Go

## Description
Operators are special symbols that perform operations on operands (variables and values). Go provides a rich set of operators for arithmetic, comparison, logical operations, and more.

## Key Concepts

### Categories of Operators
1. **Arithmetic Operators** - Mathematical operations
2. **Comparison Operators** - Compare values (conditional)
3. **Logical Operators** - Combine boolean conditions
4. **Assignment Operators** - Assign values to variables
5. **Bitwise Operators** - Operate on bits
6. **Increment/Decrement** - Increase or decrease values
7. **String Operators** - Concatenate and compare strings

## 1. Arithmetic Operators

Perform mathematical operations on numeric values.

| Operator | Name | Description | Example | Result |
|----------|------|-------------|---------|--------|
| `+` | Addition | Adds two operands | `5 + 3` | `8` |
| `-` | Subtraction | Subtracts right from left | `5 - 3` | `2` |
| `*` | Multiplication | Multiplies two operands | `5 * 3` | `15` |
| `/` | Division | Divides left by right | `15 / 3` | `5` |
| `%` | Modulus | Returns remainder | `17 % 5` | `2` |

### Important Notes:
- **Integer division** truncates the decimal part: `7 / 2 = 3` (not 3.5)
- For **float division**, convert to float: `float64(7) / float64(2) = 3.5`
- **Modulus** only works with integers

### Examples:
```go
a := 20
b := 7

sum := a + b        // 27
diff := a - b       // 13
prod := a * b       // 140
quot := a / b       // 2 (integer division)
rem := a % b        // 6

// Float division
floatQuot := float64(a) / float64(b)  // 2.8571
```

## 2. Comparison Operators (Conditional)

Compare two values and return a boolean (`true` or `false`).

| Operator | Name | Description | Example | Result |
|----------|------|-------------|---------|--------|
| `==` | Equal to | Checks if equal | `5 == 5` | `true` |
| `!=` | Not equal to | Checks if not equal | `5 != 3` | `true` |
| `<` | Less than | Left is smaller | `3 < 5` | `true` |
| `>` | Greater than | Left is larger | `5 > 3` | `true` |
| `<=` | Less than or equal | Left is smaller or equal | `5 <= 5` | `true` |
| `>=` | Greater than or equal | Left is larger or equal | `5 >= 3` | `true` |

### Examples:
```go
x := 10
y := 20

x == y    // false
x != y    // true
x < y     // true
x > y     // false
x <= 10   // true
y >= 20   // true
```

### Use Cases:
- Checking conditions in `if` statements
- Loop conditions
- Validating user input
- Comparing values

## 3. Logical Operators

Combine multiple boolean conditions.

| Operator | Name | Description | Example | Result |
|----------|------|-------------|---------|--------|
| `&&` | Logical AND | True if both are true | `true && false` | `false` |
| `\|\|` | Logical OR | True if at least one is true | `true \|\| false` | `true` |
| `!` | Logical NOT | Inverts boolean value | `!true` | `false` |

### Truth Tables:

#### AND (&&)
| A | B | A && B |
|---|---|--------|
| true | true | **true** |
| true | false | false |
| false | true | false |
| false | false | false |

#### OR (||)
| A | B | A \|\| B |
|---|---|--------|
| true | true | **true** |
| true | false | **true** |
| false | true | **true** |
| false | false | false |

#### NOT (!)
| A | !A |
|---|-----|
| true | **false** |
| false | **true** |

### Examples:
```go
age := 25
hasLicense := true
hasInsurance := false

// AND - all must be true
canDrive := age >= 18 && hasLicense  // true

// OR - at least one must be true
needsUpdate := !hasLicense || !hasInsurance  // true

// NOT - inverts the value
isMinor := !( age >= 18)  // false

// Complex expression
eligible := (age >= 18 && age <= 65) || hasLicense  // true
```

### Short-Circuit Evaluation:
Go uses **short-circuit evaluation**:
- `A && B`: If `A` is false, `B` is not evaluated
- `A || B`: If `A` is true, `B` is not evaluated

```go
// If x is 0, the second part is never evaluated
if x != 0 && y/x > 2 {
    // Safe from division by zero
}
```

## 4. Assignment Operators

Assign values to variables.

| Operator | Name | Description | Example | Equivalent |
|----------|------|-------------|---------|------------|
| `=` | Assignment | Assigns right to left | `x = 5` | - |
| `+=` | Add and assign | Adds and assigns | `x += 3` | `x = x + 3` |
| `-=` | Subtract and assign | Subtracts and assigns | `x -= 3` | `x = x - 3` |
| `*=` | Multiply and assign | Multiplies and assigns | `x *= 3` | `x = x * 3` |
| `/=` | Divide and assign | Divides and assigns | `x /= 3` | `x = x / 3` |
| `%=` | Modulus and assign | Takes modulus and assigns | `x %= 3` | `x = x % 3` |

### Examples:
```go
num := 10
num += 5   // num = 15
num -= 3   // num = 12
num *= 2   // num = 24
num /= 4   // num = 6
num %= 4   // num = 2
```

## 5. Increment and Decrement

Special operators to increase or decrease a value by 1.

| Operator | Name | Description | Example |
|----------|------|-------------|---------|
| `++` | Increment | Increases by 1 | `x++` |
| `--` | Decrement | Decreases by 1 | `x--` |

### Important Go-Specific Rules:
1. ✅ **Only postfix** form exists: `x++` (not `++x`)
2. ✅ **Statements, not expressions**: Cannot use in assignments
3. ❌ Cannot do: `y = x++` (invalid in Go)
4. ❌ Cannot do: `++x` (prefix not allowed)

### Examples:
```go
counter := 5
counter++     // counter = 6 (valid)
counter--     // counter = 5 (valid)

// These are INVALID in Go:
// ++counter   // Error!
// y = counter++  // Error!

// Correct way:
counter++
y = counter   // OK
```

## 6. Bitwise Operators

Operate on individual bits of integers.

| Operator | Name | Description | Example |
|----------|------|-------------|---------|
| `&` | AND | Sets bit if both bits are 1 | `12 & 10 = 8` |
| `\|` | OR | Sets bit if any bit is 1 | `12 \| 10 = 14` |
| `^` | XOR | Sets bit if bits are different | `12 ^ 10 = 6` |
| `^` | NOT | Inverts all bits | `^12 = -13` |
| `<<` | Left shift | Shifts bits left | `12 << 2 = 48` |
| `>>` | Right shift | Shifts bits right | `12 >> 2 = 3` |

### Examples:
```go
p := 12  // binary: 1100
q := 10  // binary: 1010

p & q    // 8   (binary: 1000) - AND
p | q    // 14  (binary: 1110) - OR
p ^ q    // 6   (binary: 0110) - XOR
^p       // -13 (inverts all bits) - NOT
p << 2   // 48  (binary: 110000) - Left shift
p >> 2   // 3   (binary: 0011) - Right shift
```

### Use Cases:
- Flags and permissions
- Low-level programming
- Optimization
- Bit manipulation puzzles

## 7. Operator Precedence

Order in which operations are evaluated (highest to lowest):

1. **Parentheses**: `()`
2. **Unary**: `!`, `^` (NOT), `+`, `-`
3. **Multiplicative**: `*`, `/`, `%`, `<<`, `>>`, `&`, `&^`
4. **Additive**: `+`, `-`, `|`, `^` (XOR)
5. **Comparison**: `==`, `!=`, `<`, `<=`, `>`, `>=`
6. **Logical AND**: `&&`
7. **Logical OR**: `||`

### Examples:
```go
// Without parentheses
result := 2 + 3 * 4      // 14 (3*4 first, then +2)

// With parentheses
result := (2 + 3) * 4    // 20 (2+3 first, then *4)

// Complex expression
result := 10 > 5 && 3 < 7   // true (comparisons first, then &&)

// Use parentheses for clarity
result := (age >= 18) && (score >= 60)  // Clearer!
```

### Best Practice:
✅ **Use parentheses** to make your intention clear, even if not strictly needed.

## 8. String Operators

| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `+` | Concatenation | `"Hello" + " World"` | `"Hello World"` |
| `==` | Equal | `"abc" == "abc"` | `true` |
| `!=` | Not equal | `"abc" != "xyz"` | `true` |
| `<` | Less than | `"apple" < "banana"` | `true` |
| `>` | Greater than | `"zebra" > "apple"` | `true` |

### Examples:
```go
firstName := "John"
lastName := "Doe"

// Concatenation
fullName := firstName + " " + lastName  // "John Doe"

// Comparison (lexicographical order)
"apple" < "banana"   // true
"abc" == "abc"       // true
```

## Running the Code

```bash
cd basics/05-operators
go run main.go
```

## Practical Examples in Code

The `main.go` file includes 8 practical examples:
1. **Even/Odd Checker** - Using modulus operator
2. **Range Checker** - Using comparison and logical operators
3. **Leap Year Calculator** - Complex logical expression
4. **Discount Calculator** - Arithmetic with conditions
5. **Password Strength** - Multiple comparisons
6. **Grade Boundary** - Using comparison operators
7. **BMI Calculator** - Arithmetic and conditionals
8. **Time Calculator** - Division and modulus

## Exercises

1. **Calculator**: Build a simple calculator using arithmetic operators

2. **Temperature Converter**: Convert between Celsius and Fahrenheit using formula: `F = C * 9/5 + 32`

3. **Triangle Validator**: Check if three sides can form a valid triangle using comparison operators

4. **Number Properties**: Write a program that checks if a number is:
   - Even or odd
   - Positive, negative, or zero
   - Divisible by 3 and 5

5. **Eligibility Checker**: Create a program to check college admission eligibility:
   - Age between 17-25
   - Grade >= 60
   - Entrance exam score >= 70

6. **Bitwise Flags**: Create a permission system using bitwise operators (read=1, write=2, execute=4)

7. **Compound Interest**: Calculate compound interest using arithmetic operators

## Common Mistakes

### ❌ Using = instead of == for comparison
```go
if x = 5 {  // Error! Assignment, not comparison
}

// Correct:
if x == 5 {
}
```

### ❌ Integer division expecting float result
```go
result := 7 / 2  // 3, not 3.5

// Correct for float division:
result := float64(7) / float64(2)  // 3.5
```

### ❌ Modulus with floats
```go
result := 7.5 % 2.5  // Error! Modulus only works with integers

// Convert to int first, or use math.Mod()
```

### ❌ Prefix increment/decrement
```go
++counter   // Error! Go only supports postfix

// Correct:
counter++
```

### ❌ Using ++ or -- as expression
```go
x = counter++  // Error! Not an expression in Go

// Correct:
counter++
x = counter
```

## Best Practices

### ✅ DO:
- Use parentheses for clarity in complex expressions
- Use compound assignment operators (`+=`, `-=`, etc.)
- Check for division by zero before dividing
- Use comparison operators in conditions
- Use logical operators to combine conditions
- Convert types explicitly when needed

### ❌ DON'T:
- Assume integer division gives float result
- Use `=` when you mean `==`
- Create overly complex expressions
- Forget operator precedence
- Use prefix `++` or `--` (not valid in Go)

## Comparison with Other Languages

| Feature | Go | C/Java | Python |
|---------|-----|--------|--------|
| Integer division | Truncates | Truncates | Python 2: truncates, Python 3: float |
| Modulus with negatives | Same sign as dividend | Varies | Same sign as divisor |
| Increment | Postfix only | Prefix and postfix | None (`+=` only) |
| ++ as expression | No | Yes | No |
| Power operator | `math.Pow()` | `pow()` | `**` |
| String concatenation | `+` | `+` (Java only) | `+` |

## Type Compatibility

Go requires **explicit type conversion** for operations:

```go
var a int = 10
var b float64 = 3.5

// This is INVALID:
// result := a + b  // Error! Type mismatch

// Correct - explicit conversion:
result := float64(a) + b  // 13.5
```

## Related Topics
- [Variables](../04-variables/README.md) - Understanding types for operators
- [If-Else Statements](../../control-flow/01-if-and-else-statement/README.md) - Using comparison and logical operators
- [Loops](../../control-flow/03-loops/README.md) - Using operators in loop conditions
- [Functions](../../functions/README.md) - Operators in function logic

## Additional Resources
- [Go Specification: Operators](https://go.dev/ref/spec#Operators)
- [Go by Example: Operators](https://gobyexample.com/)
- [Effective Go: Control Structures](https://go.dev/doc/effective_go#control-structures)
- [Go Tour: Basic Types](https://go.dev/tour/basics/11)

