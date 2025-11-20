# Pointers in Go

## Description
Pointers are variables that store memory addresses. They enable efficient memory management, allow functions to modify variables, and are essential for working with large data structures in Go.

## Key Concepts

### What is a Pointer?
A **pointer** is a variable that stores the memory address of another variable.

### Memory Address
A **memory address** is the location in computer memory where a value is stored.

### Go is Pass-by-Value
Go passes copies of values to functions by default. Pointers allow us to pass references to modify original values.

## Core Concepts

### 1. The `&` Operator (Address-of)

The `&` operator returns the **memory address** of a variable:

```go
age := 25
fmt.Println(&age)  // Prints something like: 0xc0000140a8
```

**Key Points:**
- `&variable` gives you the address
- The address is where the variable is stored in memory
- Addresses look like hexadecimal numbers

### 2. Pointer Declaration

A pointer stores an address and must specify what type of value it points to:

```go
var ptr *int  // Pointer to an int
var p *string // Pointer to a string
```

**Syntax**: `*Type` declares a pointer to `Type`

### 3. The `*` Operator (Dereference)

The `*` operator has two uses:

#### A. Declare Pointer Type
```go
var ptr *int  // ptr is a pointer to int
```

#### B. Dereference (Access Value)
```go
x := 42
ptr := &x
fmt.Println(*ptr)  // Prints: 42 (value at the address)
```

**Dereferencing** means accessing or modifying the value at the pointer's address.

## Basic Usage

### Getting an Address

```go
number := 100
address := &number
fmt.Printf("Value: %d\n", number)      // 100
fmt.Printf("Address: %p\n", address)   // 0xc0000140a8
```

### Storing and Using a Pointer

```go
x := 42
var ptr *int      // Declare pointer
ptr = &x          // Store address of x

fmt.Println(ptr)   // Address: 0xc0000140a8
fmt.Println(*ptr)  // Value: 42
```

### Modifying Through Pointer

```go
number := 10
ptr := &number

*ptr = 20  // Modify value through pointer

fmt.Println(number)  // Prints: 20
fmt.Println(*ptr)    // Prints: 20
```

## Pointer Types

Pointers are **type-specific** - a pointer to `int` can only point to `int`:

```go
var intPtr *int
var floatPtr *float64
var stringPtr *string

i := 42
f := 3.14
s := "Hello"

intPtr = &i
floatPtr = &f
stringPtr = &s

// This would be an error:
// intPtr = &f  // Cannot assign *float64 to *int
```

## Nil Pointers

A pointer with no assigned address has value `nil`:

```go
var p *int
fmt.Println(p)      // Prints: <nil>
fmt.Println(p == nil)  // Prints: true
```

### ⚠️ WARNING: Dereferencing Nil Panics

```go
var p *int
// fmt.Println(*p)  // PANIC! Cannot dereference nil

// Always check before dereferencing
if p != nil {
    fmt.Println(*p)
} else {
    fmt.Println("Pointer is nil")
}
```

## Pass by Value vs Pass by Reference

### Go is Pass-by-Value

By default, Go copies values when passing to functions:

```go
func modify(n int) {
    n = 99  // Only modifies local copy
}

x := 10
modify(x)
fmt.Println(x)  // Still 10 (unchanged)
```

### Pass by Reference (Using Pointers)

Pointers allow functions to modify original values:

```go
func modify(n *int) {
    *n = 99  // Modifies original value
}

x := 10
modify(&x)  // Pass address
fmt.Println(x)  // Now 99 (changed!)
```

### Comparison Table

| Aspect | Pass by Value | Pass by Pointer |
|--------|---------------|-----------------|
| What's passed | Copy of value | Copy of address |
| Original modified? | No | Yes |
| Syntax (call) | `func(x)` | `func(&x)` |
| Syntax (parameter) | `func(n int)` | `func(n *int)` |
| Use when | Small data, no modification | Large data or need modification |

## Pointers with Structs

### Creating Pointer to Struct

```go
type Person struct {
    Name string
    Age  int
}

person := Person{Name: "Alice", Age: 25}
ptr := &person
```

### Accessing Fields

Go automatically dereferences struct pointers:

```go
// Both work the same way
fmt.Println(ptr.Name)        // Automatic dereference
fmt.Println((*ptr).Name)     // Manual dereference
```

### Modifying Through Pointer

```go
ptr.Age = 26  // Modifies original struct
fmt.Println(person.Age)  // Prints: 26
```

## The `new` Function

`new(T)` allocates memory for type `T` and returns a pointer:

```go
ptr := new(int)
fmt.Println(*ptr)  // Prints: 0 (zero value)

*ptr = 100
fmt.Println(*ptr)  // Prints: 100
```

### `new` vs `&`

```go
// Using new
ptr1 := new(int)
*ptr1 = 42

// Using &
value := 42
ptr2 := &value

// Both create pointers, but & requires existing variable
```

## Pointer Receivers in Methods

### Value Receiver (Doesn't Modify)

```go
func (p Person) valueMethod() {
    p.Age = 100  // Modifies copy only
}

person := Person{Name: "Bob", Age: 30}
person.valueMethod()
fmt.Println(person.Age)  // Still 30
```

### Pointer Receiver (Modifies Original)

```go
func (p *Person) pointerMethod() {
    p.Age = 100  // Modifies original
}

person := Person{Name: "Bob", Age: 30}
person.pointerMethod()
fmt.Println(person.Age)  // Now 100
```

### When to Use Pointer Receivers

Use pointer receivers when:
- ✅ Method needs to modify the receiver
- ✅ Receiver is a large struct (avoid copying)
- ✅ Consistency (some methods use pointer receivers)

Use value receivers when:
- ✅ Receiver is small (int, bool, string)
- ✅ Receiver shouldn't be modified
- ✅ Receiver is a map, function, or channel (already references)

## Pointer Arithmetic

**Go does NOT support pointer arithmetic** (unlike C/C++):

```go
arr := [3]int{10, 20, 30}
ptr := &arr[0]

// These are NOT allowed:
// ptr++           // Error!
// ptr = ptr + 1   // Error!
```

This design choice improves safety and prevents common bugs.

## Common Patterns

### Pattern 1: Swap Function

```go
func swap(x, y *int) {
    temp := *x
    *x = *y
    *y = temp
}

a, b := 5, 10
swap(&a, &b)
// a is now 10, b is now 5
```

### Pattern 2: Optional Values

```go
type User struct {
    Name  string
    Email *string  // nil means no email
}

func getUser() User {
    email := "user@example.com"
    return User{
        Name:  "Alice",
        Email: &email,  // Has email
    }
}

func getUserNoEmail() User {
    return User{
        Name:  "Bob",
        Email: nil,  // No email
    }
}
```

### Pattern 3: Factory Function

```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

p := NewPerson("Charlie", 35)
```

### Pattern 4: Linked Data Structures

```go
type Node struct {
    Value int
    Next  *Node  // Points to next node
}

node1 := &Node{Value: 1}
node2 := &Node{Value: 2}
node1.Next = node2  // Link nodes
```

### Pattern 5: Large Struct Efficiency

```go
type LargeStruct struct {
    Data [1000]int
}

// Inefficient: copies entire struct
func process(ls LargeStruct) {
    // Uses 8KB of memory
}

// Efficient: passes pointer
func processPtr(ls *LargeStruct) {
    // Uses 8 bytes (pointer size)
}
```

## Best Practices

### ✅ DO:

```go
// Check for nil before dereferencing
if ptr != nil {
    fmt.Println(*ptr)
}

// Use pointers for large structs
func process(data *LargeStruct) {
    // Efficient
}

// Use pointer receivers for modification
func (p *Person) incrementAge() {
    p.Age++
}

// Return pointers from constructor functions
func NewUser(name string) *User {
    return &User{Name: name}
}

// Use pointers for optional values
type Config struct {
    Timeout *int  // nil means use default
}
```

### ❌ DON'T:

```go
// Don't return pointer to local variable (Go handles this safely)
// In Go, this is actually safe (unlike C), but be aware
func makePointer() *int {
    x := 42
    return &x  // OK in Go (escapes to heap)
}

// Don't use pointers for small types unnecessarily
func process(x *int) {  // Overkill for int
    fmt.Println(*x)
}
// Better:
func process(x int) {
    fmt.Println(x)
}

// Don't dereference without nil check
func process(p *int) {
    *p = 42  // Dangerous if p is nil!
}

// Don't use pointers everywhere
// Use pointers only when needed
```

## Memory Diagram

```
Variable: x = 42
┌─────────┐
│   42    │ ← Value stored at address 0x1000
└─────────┘
    ↑
    │
┌─────────┐
│ 0x1000  │ ← Pointer (stores address)
└─────────┘
```

## When to Use Pointers

| Scenario | Use Pointer? | Reason |
|----------|--------------|--------|
| Modify function parameter | ✅ Yes | Need to change original |
| Large struct parameter | ✅ Yes | Avoid expensive copy |
| Small types (int, bool) | ❌ No | Copy is cheap |
| Method needs to modify | ✅ Yes | Use pointer receiver |
| Optional field in struct | ✅ Yes | nil represents missing |
| Return large struct | ✅ Yes | Avoid copy |
| Slices, maps, channels | ❌ No | Already reference types |

## Running the Code

```bash
cd basics/03-pointers
go run main.go
```

## Code Examples

The `main.go` file demonstrates:
1. Memory addresses with `&` operator
2. Declaring pointers
3. Dereferencing with `*` operator
4. Pointer types (type-specific)
5. Nil pointers and checking
6. Pass by value vs pass by reference
7. Pointers with structs
8. The `new` function
9. Pointer receivers in methods
10. No pointer arithmetic
11. Six practical examples

## Exercises

1. **Swap Function**: Write a function that swaps two integers using pointers

2. **Update Struct**: Create a function that updates a struct's fields using a pointer

3. **Optional Values**: Create a User struct with optional email field using pointers

4. **Linked List**: Implement a simple linked list with Node pointers

5. **Counter**: Create a counter with increment/decrement methods using pointer receivers

6. **Find and Modify**: Write a function that finds an element in a slice and modifies it via pointer

7. **Tree Node**: Create a binary tree node structure with left/right pointers

8. **Safe Dereference**: Write a function that safely dereferences a pointer with default value

## Common Mistakes

### ❌ Dereferencing Nil

```go
var p *int
fmt.Println(*p)  // PANIC!
```

### ❌ Wrong Pointer Type

```go
var intPtr *int
var floatVal float64 = 3.14
intPtr = &floatVal  // Error: type mismatch
```

### ❌ Forgetting & When Passing

```go
func modify(p *int) {
    *p = 42
}

x := 10
modify(x)  // Error: need modify(&x)
```

### ❌ Using * for Declaration and Dereference Confusion

```go
var p *int     // * declares pointer type
x := 42
p = &x
y := *p        // * dereferences (gets value)
```

## Comparison with Other Languages

| Feature | Go | C/C++ | Java |
|---------|-----|-------|------|
| Pointer arithmetic | ❌ No | ✅ Yes | ❌ No |
| Nil safety | ⚠️ Runtime | ❌ Unsafe | ✅ NullPointerException |
| Automatic dereferencing | ✅ Structs | ❌ No | N/A (references) |
| Garbage collection | ✅ Yes | ❌ No | ✅ Yes |
| Pointer syntax | `*` and `&` | `*` and `&` | No pointers (references) |

## Related Topics
- [Variables](../04-variables/README.md) - Understanding types and values
- [Functions](../../functions/01-using-function/README.md) - Passing parameters
- [Structs](../../data-structures/structs/README.md) - Pointers with structs
- [Methods](../../functions/03-methods/README.md) - Pointer receivers
- [Slices](../../data-structures/slices/README.md) - Reference types

## Additional Resources
- [Go by Example: Pointers](https://gobyexample.com/pointers)
- [A Tour of Go: Pointers](https://go.dev/tour/moretypes/1)
- [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go Specification: Address operators](https://go.dev/ref/spec#Address_operators)

