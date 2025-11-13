# Variables in Go

## Description
Learn how to declare, initialize, and use variables in Go. Understanding variables is fundamental to any programming language.

## Key Concepts

### 1. Variable Declaration with `var`

The `var` keyword is used to declare variables with explicit types:

```go
var name string = "Alice"
var age int = 25
var height float64 = 1.68
var isStudent bool = true
```

### 2. Type Inference

Go can automatically infer the type from the value:

```go
var city = "New York"      // Go infers string
var temperature = 22.5     // Go infers float64
var population = 8336817   // Go infers int
```

### 3. Short Declaration (`:=`)

The **most common** way to declare variables in Go:

```go
country := "USA"           // Type inferred as string
latitude := 40.7128        // Type inferred as float64
visited := true            // Type inferred as bool
```

**Important**: `:=` can **only** be used inside functions!

### 4. Multiple Variable Declaration

Go supports declaring multiple variables at once:

```go
// With var and explicit types
var x, y, z int = 1, 2, 3

// With var and type inference
var a, b, c = 10, "hello", true

// With := (most common)
firstName, lastName, age := "John", "Doe", 30

// In a block
var (
    productName  = "Laptop"
    productPrice = 999.99
    inStock      = true
)
```

### 5. Zero Values

Variables declared without initialization get default "zero values":

| Type | Zero Value |
|------|------------|
| `int`, `int8`, `int16`, `int32`, `int64` | `0` |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `0` |
| `float32`, `float64` | `0.0` |
| `bool` | `false` |
| `string` | `""` (empty string) |
| Pointers, slices, maps, channels, functions, interfaces | `nil` |

```go
var count int        // 0
var price float64    // 0.0
var name string      // ""
var active bool      // false
```

## Data Types in Go

### Integer Types

#### Signed Integers (can be negative)
| Type | Size | Range |
|------|------|-------|
| `int8` | 8 bits | -128 to 127 |
| `int16` | 16 bits | -32,768 to 32,767 |
| `int32` | 32 bits | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64 bits | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `int` | Platform dependent | 32 or 64 bits |

#### Unsigned Integers (only positive)
| Type | Size | Range |
|------|------|-------|
| `uint8` | 8 bits | 0 to 255 |
| `uint16` | 16 bits | 0 to 65,535 |
| `uint32` | 32 bits | 0 to 4,294,967,295 |
| `uint64` | 64 bits | 0 to 18,446,744,073,709,551,615 |
| `uint` | Platform dependent | 32 or 64 bits |

#### Special Integer Types
- `byte` - alias for `uint8`
- `rune` - alias for `int32` (represents a Unicode code point)

### Floating-Point Types
| Type | Size | Precision |
|------|------|-----------|
| `float32` | 32 bits | ~7 decimal digits |
| `float64` | 64 bits | ~15 decimal digits |

**Note**: Always prefer `float64` unless you have a specific reason to use `float32`.

### Other Types
- `bool` - `true` or `false`
- `string` - sequence of characters
- `complex64` - complex numbers with float32 real and imaginary parts
- `complex128` - complex numbers with float64 real and imaginary parts

## Constants

Constants are declared with the `const` keyword and **cannot be changed**:

```go
const Pi = 3.14159
const AppName = "MyApp"
const MaxUsers = 100
```

### Multiple Constants

```go
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)
```

### iota - Constant Generator

`iota` is a special identifier used with constants to create sequential values:

```go
const (
    Sunday    = iota // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

## Variable Scope

- **Package-level variables**: Declared outside functions, accessible throughout the package
- **Function-level variables**: Declared inside functions, only accessible within that function
- **Block-level variables**: Declared inside `{}` blocks, only accessible within that block

## Common Patterns

### Swapping Variables

```go
a, b := 10, 20
a, b = b, a  // Swap without temporary variable
```

### Multiple Return Values

```go
value, err := someFunction()
if err != nil {
    // Handle error
}
```

### Ignoring Values with `_`

```go
value, _ := someFunction()  // Ignore second return value
```

## Running the Code

```bash
cd basics/04-variables
go run main.go
```

## What the Code Demonstrates

1. **var declaration** - Explicit type declaration
2. **Type inference** - Letting Go determine the type
3. **:= operator** - Short variable declaration
4. **Multiple variables** - Declaring several variables at once
5. **Zero values** - Default values for uninitialized variables
6. **Data types** - All basic types with examples
7. **Reassignment** - Changing variable values
8. **Constants** - Immutable values with const and iota
9. **Practical examples** - Real-world use cases

## Best Practices

✅ **Use `:=` for local variables** - It's more concise
```go
name := "Alice"  // Good
```

✅ **Use `var` for package-level variables** - More explicit
```go
var GlobalConfig string  // Good
```

✅ **Group related declarations**
```go
var (
    host = "localhost"
    port = 8080
    debug = true
)
```

✅ **Use meaningful names**
```go
userAge := 25           // Good
u := 25                 // Avoid
```

✅ **Prefer `int` over specific sizes** unless you need a specific size
```go
count := 10            // Use int, not int32
```

## Common Mistakes

❌ **Using := for reassignment**
```go
age := 25
age := 30  // Error! Use = for reassignment
age = 30   // Correct
```

❌ **Using := outside functions**
```go
// Outside function
name := "Alice"  // Error! Use var instead
```

❌ **Unused variables**
```go
func main() {
    name := "Alice"  // Error if never used
}
```
Go won't compile if you declare a variable but don't use it.

❌ **Type mismatch**
```go
var age int = "25"  // Error! Can't assign string to int
```

## Exercises

1. Declare variables for a person's information (name, age, email, isActive) using all three methods (var, var with inference, :=)
2. Create a program that declares multiple variables in one line
3. Demonstrate zero values for all basic types
4. Create constants for days of the week using iota
5. Write a function that returns multiple values and assign them to variables
6. Swap two variables using multiple assignment
7. Create a calculator that uses different numeric types

## Notes

- Go is **statically typed** - types are checked at compile time
- Once declared, a variable's type **cannot change**
- The `:=` operator is convenient but can only be used inside functions
- Always initialize variables to avoid unexpected zero values
- Go compiler will error on **unused variables** (except with `_`)
- Use `const` for values that never change
- Variable names should be **camelCase** in Go

## Type Conversion

Go requires **explicit type conversion**:

```go
var i int = 42
var f float64 = float64(i)  // Convert int to float64
var u uint = uint(f)        // Convert float64 to uint
```

## Related Topics
- Previous: Importing Packages
- Next: Data Types (detailed exploration)
- Related: Operators
- Related: Type Conversion and Casting

