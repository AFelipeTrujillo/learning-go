package main

import "fmt"

func main() {
	fmt.Println("=== Variables in Go ===\n")

	// ============================================
	// 1. VARIABLE DECLARATION WITH var
	// ============================================
	fmt.Println("--- 1. Declaration with var ---")

	// Explicit type declaration
	var name string = "Alice"
	var age int = 25
	var height float64 = 1.68
	var isStudent bool = true
	var longNumberWithUnderscores int64 = 1_000_000_000

	fmt.Printf("Name: %s (type: %T)\n", name, name)
	fmt.Printf("Age: %d (type: %T)\n", age, age)
	fmt.Printf("Height: %.2f meters (type: %T)\n", height, height)
	fmt.Printf("Is student: %t (type: %T)\n", isStudent, isStudent)
	fmt.Printf("Long number: %d (type: %T)\n", longNumberWithUnderscores, longNumberWithUnderscores)
	fmt.Println()

	// ============================================
	// 2. TYPE INFERENCE
	// ============================================
	fmt.Println("--- 2. Type Inference ---")

	// Go can infer the type from the value
	var city = "New York"    // string
	var temperature = 22.5   // float64
	var population = 8336817 // int
	var isCapital = false    // bool

	fmt.Printf("City: %s (type: %T)\n", city, city)
	fmt.Printf("Temperature: %.1f°C (type: %T)\n", temperature, temperature)
	fmt.Printf("Population: %d (type: %T)\n", population, population)
	fmt.Printf("Is capital: %t (type: %T)\n", isCapital, isCapital)
	fmt.Println()

	// ============================================
	// 3. SHORT VARIABLE DECLARATION (:=)
	// ============================================
	fmt.Println("--- 3. Short Declaration with := ---")

	// Most common way to declare variables in Go
	// Can only be used INSIDE functions
	country := "USA"
	latitude := 40.7128
	longitude := -74.0060
	visited := true

	fmt.Printf("Country: %s (type: %T)\n", country, country)
	fmt.Printf("Coordinates: %.4f, %.4f (type: %T)\n", latitude, longitude, latitude)
	fmt.Printf("Visited: %t (type: %T)\n", visited, visited)
	fmt.Println()

	// ============================================
	// 4. MULTIPLE VARIABLE DECLARATION
	// ============================================
	fmt.Println("--- 4. Multiple Variables ---")

	// Multiple variables with var
	var x, y, z int = 1, 2, 3
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	// Multiple variables with type inference
	var a, b, c = 10, "hello", true
	fmt.Printf("a=%d (type: %T), b=%s (type: %T), c=%t (type: %T)\n", a, a, b, b, c, c)

	// Multiple variables with := (most common)
	firstName, lastName, fullAge := "John", "Doe", 30
	fmt.Printf("Full name: %s %s, Age: %d\n", firstName, lastName, fullAge)

	// Declaring multiple variables in a block
	var (
		productName  = "Laptop"
		productPrice = 999.99
		inStock      = true
		quantity     = 15
	)
	fmt.Printf("Product: %s | Price: $%.2f | In stock: %t | Quantity: %d\n",
		productName, productPrice, inStock, quantity)
	fmt.Println()

	// ============================================
	// 5. ZERO VALUES (Default Values)
	// ============================================
	fmt.Println("--- 5. Zero Values ---")

	// Variables declared without initialization get zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("Default int: %d\n", defaultInt)         // 0
	fmt.Printf("Default float: %f\n", defaultFloat)     // 0.000000
	fmt.Printf("Default string: '%s'\n", defaultString) // "" (empty string)
	fmt.Printf("Default bool: %t\n", defaultBool)       // false
	fmt.Println()

	// ============================================
	// 6. DATA TYPES
	// ============================================
	fmt.Println("--- 6. Common Data Types ---")

	// Integer types
	var int8Var int8 = 127                   // -128 to 127
	var int16Var int16 = 32767               // -32768 to 32767
	var int32Var int32 = 2147483647          // -2^31 to 2^31-1
	var int64Var int64 = 9223372036854775807 // -2^63 to 2^63-1
	var intVar int = 42                      // Platform dependent (32 or 64 bit)

	fmt.Printf("int8: %d\n", int8Var)
	fmt.Printf("int16: %d\n", int16Var)
	fmt.Printf("int32: %d\n", int32Var)
	fmt.Printf("int64: %d\n", int64Var)
	fmt.Printf("int: %d\n", intVar)

	// Unsigned integer types (only positive numbers)
	var uint8Var uint8 = 255          // 0 to 255
	var uint16Var uint16 = 65535      // 0 to 65535
	var uint32Var uint32 = 4294967295 // 0 to 2^32-1

	fmt.Printf("uint8: %d\n", uint8Var)
	fmt.Printf("uint16: %d\n", uint16Var)
	fmt.Printf("uint32: %d\n", uint32Var)

	// Float types
	var float32Var float32 = 3.14
	var float64Var float64 = 3.141592653589793

	fmt.Printf("float32: %.2f\n", float32Var)
	fmt.Printf("float64: %.15f\n", float64Var)

	// String and rune
	var stringVar string = "Hello, Go!"
	var runeVar rune = 'A' // rune is alias for int32, represents a Unicode code point
	var byteVar byte = 65  // byte is alias for uint8

	fmt.Printf("string: %s\n", stringVar)
	fmt.Printf("rune: %c (value: %d)\n", runeVar, runeVar)
	fmt.Printf("byte: %c (value: %d)\n", byteVar, byteVar)

	// Complex numbers
	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 3 + 4i

	fmt.Printf("complex64: %v\n", complex64Var)
	fmt.Printf("complex128: %v\n", complex128Var)
	fmt.Println()

	// ============================================
	// 7. REASSIGNING VARIABLES
	// ============================================
	fmt.Println("--- 7. Reassigning Variables ---")

	score := 100
	fmt.Printf("Initial score: %d\n", score)

	score = 150 // Reassign (use = not :=)
	fmt.Printf("Updated score: %d\n", score)

	score = score + 50
	fmt.Printf("Final score: %d\n", score)
	fmt.Println()

	// ============================================
	// 8. MULTIPLE ASSIGNMENT
	// ============================================
	fmt.Println("--- 8. Multiple Assignment ---")

	// Swap variables without temporary variable
	num1, num2 := 10, 20
	fmt.Printf("Before swap: num1=%d, num2=%d\n", num1, num2)

	num1, num2 = num2, num1 // Swap
	fmt.Printf("After swap: num1=%d, num2=%d\n", num1, num2)

	// Assigning results from a function (common pattern)
	min, max := getMinMax(15, 5)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
	fmt.Println()

	// ============================================
	// 9. CONSTANTS
	// ============================================
	fmt.Println("--- 9. Constants ---")

	// Constants cannot be changed after declaration
	const Pi = 3.14159
	const AppName = "MyApp"
	const MaxUsers = 100

	fmt.Printf("Pi: %.5f\n", Pi)
	fmt.Printf("App Name: %s\n", AppName)
	fmt.Printf("Max Users: %d\n", MaxUsers)

	// Multiple constants
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)
	fmt.Printf("HTTP Status Codes: OK=%d, NotFound=%d, Error=%d\n",
		StatusOK, StatusNotFound, StatusError)

	// iota - special constant generator
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Printf("Days: Sunday=%d, Monday=%d, Friday=%d\n", Sunday, Monday, Friday)
	fmt.Println()

	// ============================================
	// 10. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 10. Practical Examples ---")

	// User profile
	username := "gopher_dev"
	userAge := 28
	userEmail := "gopher@example.com"
	isPremium := true
	accountBalance := 150.75

	fmt.Println("=== User Profile ===")
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Age: %d\n", userAge)
	fmt.Printf("Email: %s\n", userEmail)
	fmt.Printf("Premium: %t\n", isPremium)
	fmt.Printf("Balance: $%.2f\n", accountBalance)

	// Shopping cart
	itemName := "Go Programming Book"
	itemPrice := 39.99
	itemQuantity := 2
	totalPrice := itemPrice * float64(itemQuantity)

	fmt.Println("\n=== Shopping Cart ===")
	fmt.Printf("Item: %s\n", itemName)
	fmt.Printf("Price: $%.2f x %d\n", itemPrice, itemQuantity)
	fmt.Printf("Total: $%.2f\n", totalPrice)
}

// Helper function for example 8
func getMinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
