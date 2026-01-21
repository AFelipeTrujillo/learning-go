package main

import "fmt"

func main() {
	// ============================================
	// fmt.Println - Prints with spaces and newline
	// ============================================
	fmt.Println("=== fmt.Println Examples ===")

	// Single value
	fmt.Println("Hello, World!")

	// Multiple values (automatically adds spaces between them)
	fmt.Println("Go", "is", "awesome!")

	// Mixing different types
	name := "Alice"
	age := 25
	fmt.Println("Name:", name, "Age:", age)

	// Numbers
	fmt.Println("Numbers:", 10, 20, 30)

	// Boolean
	isLearning := true
	fmt.Println("Currently learning Go?", isLearning)

	fmt.Println() // Empty line for separation

	// ============================================
	// fmt.Print - Prints without newline
	// ============================================
	fmt.Println("=== fmt.Print Examples ===")

	// Print without automatic newline
	fmt.Print("This ")
	fmt.Print("is ")
	fmt.Print("one ")
	fmt.Print("line")
	fmt.Println() // Add newline manually

	// Print doesn't add spaces automatically between calls
	fmt.Print("A")
	fmt.Print("B")
	fmt.Print("C")
	fmt.Println() // Add newline

	fmt.Println() // Empty line for separation

	// ============================================
	// fmt.Printf - Formatted printing
	// ============================================
	fmt.Println("=== fmt.Printf Examples ===")

	// String formatting
	firstName := "John"
	lastName := "Doe"
	fmt.Printf("Full name: %s %s\n", firstName, lastName)

	// Integer formatting
	number := 42
	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Hexadecimal: %x\n", number)
	fmt.Printf("Octal: %o\n", number)

	// Float formatting
	pi := 3.14159265359
	fmt.Printf("Default float: %f\n", pi)
	fmt.Printf("2 decimal places: %.2f\n", pi)
	fmt.Printf("Scientific notation: %e\n", pi)

	// Boolean formatting
	isTrue := true
	fmt.Printf("Boolean: %t\n", isTrue)

	// Type formatting
	var x interface{} = 42
	fmt.Printf("Type: %T, Value: %v\n", x, x)

	// Multiple values
	product := "Laptop"
	price := 999.99
	quantity := 3
	fmt.Printf("Product: %s | Price: $%.2f | Quantity: %d\n", product, price, quantity)

	// Width and alignment
	fmt.Printf("Right aligned: |%10s|\n", "Go")
	fmt.Printf("Left aligned:  |%-10s|\n", "Go")
	fmt.Printf("Padded number: |%05d|\n", 42)

	fmt.Println() // Empty line for separation

	// ============================================
	// Practical Examples
	// ============================================
	fmt.Println("=== Practical Examples ===")

	// Simple greeting
	userName := "Developer"
	fmt.Printf("Welcome back, %s!\n", userName)

	// Math calculation output
	a, b := 15, 4
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)

	// Shopping cart summary
	fmt.Println("\n--- Shopping Cart ---")
	item1, price1 := "Book", 29.99
	item2, price2 := "Pen", 2.50
	total := price1 + price2
	fmt.Printf("%-10s $%.2f\n", item1, price1)
	fmt.Printf("%-10s $%.2f\n", item2, price2)
	fmt.Println("--------------------")
	fmt.Printf("Total:     $%.2f\n", total)

	// Using fmt.Scan to get user input
	UsingScan()
}

func UsingScan() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s!\n", name)

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Printf("You are %d years old.\n", age)
}
