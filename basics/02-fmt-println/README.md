# fmt Package - Printing Functions

## Description
Understanding the `fmt` package and its main printing functions: `Println`, `Print`, and `Printf`.

## Key Concepts

### fmt.Println()
- **Prints** values followed by a **newline**
- Automatically adds **spaces** between multiple arguments
- Most commonly used for simple output
- Syntax: `fmt.Println(value1, value2, ...)`

### fmt.Print()
- **Prints** values **without** automatic newline
- Does **not** add spaces between separate calls
- Useful for building output on the same line
- Syntax: `fmt.Print(value1, value2, ...)`

### fmt.Printf()
- **Formatted printing** with placeholders
- Requires **format specifiers** (verbs)
- Most powerful and flexible
- Syntax: `fmt.Printf("format string", values...)`

## Common Format Verbs (for Printf)

| Verb | Description | Example |
|------|-------------|---------|
| `%s` | String | `fmt.Printf("Name: %s", "John")` |
| `%d` | Decimal integer | `fmt.Printf("Age: %d", 25)` |
| `%f` | Float | `fmt.Printf("Price: %f", 19.99)` |
| `%.2f` | Float with 2 decimals | `fmt.Printf("Price: %.2f", 19.99)` |
| `%t` | Boolean | `fmt.Printf("Active: %t", true)` |
| `%v` | Default format (any value) | `fmt.Printf("Value: %v", anything)` |
| `%T` | Type of value | `fmt.Printf("Type: %T", 42)` |
| `%b` | Binary | `fmt.Printf("Binary: %b", 42)` |
| `%x` | Hexadecimal | `fmt.Printf("Hex: %x", 42)` |
| `%c` | Character (rune) | `fmt.Printf("Char: %c", 65)` |
| `%p` | Pointer address | `fmt.Printf("Address: %p", &x)` |
| `%%` | Literal percent sign | `fmt.Printf("100%% complete")` |

## Width and Precision

```go
// Width - minimum number of characters
fmt.Printf("|%5d|", 42)      // |   42| (right aligned)
fmt.Printf("|%-5d|", 42)     // |42   | (left aligned)
fmt.Printf("|%05d|", 42)     // |00042| (zero padded)

// Precision for floats
fmt.Printf("%.2f", 3.14159)  // 3.14
fmt.Printf("%.4f", 3.14159)  // 3.1416
```

## Running the Code
```bash
cd basics/02-fmt-println
go run main.go
```

## Expected Output
The program demonstrates:
1. Basic `Println` usage
2. Difference between `Print` and `Println`
3. Various `Printf` format specifiers
4. Practical formatting examples
5. Alignment and padding

## Key Differences

| Feature | Print | Println | Printf |
|---------|-------|---------|--------|
| Newline | No | Yes | No (manual `\n`) |
| Spaces between args | Yes | Yes | N/A |
| Formatting | No | No | Yes |
| Use case | Same line output | Simple output | Formatted output |

## Common Patterns

### Simple Output
```go
fmt.Println("Hello, World!")
```

### Variables
```go
name := "Alice"
age := 30
fmt.Println("Name:", name, "Age:", age)
```

### Formatted Strings
```go
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

### Currency Formatting
```go
price := 19.99
fmt.Printf("Price: $%.2f\n", price)
```

## Exercises

1. Print your name, age, and favorite programming language using `Println`
2. Use `Printf` to create a formatted student ID card with name, ID number, and GPA
3. Create a receipt printer that formats items and prices nicely aligned
4. Print a number in decimal, binary, and hexadecimal formats
5. Build a simple calculator that prints formatted equations

## Notes

- `\n` adds a newline character
- Format verbs are case-sensitive (`%d` vs `%D`)
- Wrong format verb will print the verb itself
- `%v` is a safe default when unsure of the type
- Always end `Printf` strings with `\n` for newline

## Related Topics
- Next: Variables and Data Types
- Related: String formatting with `fmt.Sprintf()`
- Related: Reading input with `fmt.Scan()`

