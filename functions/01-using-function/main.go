package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("=== Functions in Go ===\n")

	// ============================================
	// 1. BASIC FUNCTION DECLARATION
	// ============================================
	fmt.Println("--- 1. Basic Function Declaration ---")
	
	// Call a simple function
	greet()
	
	// Call function multiple times
	sayHello()
	sayHello()
	
	fmt.Println()

	// ============================================
	// 2. FUNCTIONS WITH PARAMETERS
	// ============================================
	fmt.Println("--- 2. Functions with Parameters ---")
	
	// Single parameter
	greetPerson("Alice")
	greetPerson("Bob")
	
	// Multiple parameters
	addNumbers(10, 20)
	addNumbers(5, 15)
	
	// Parameters of same type
	multiply(4, 5)
	
	fmt.Println()

	// ============================================
	// 3. FUNCTIONS WITH RETURN VALUES
	// ============================================
	fmt.Println("--- 3. Functions with Return Values ---")
	
	// Function returning a value
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)
	
	// Use return value directly
	fmt.Printf("5 + 15 = %d\n", add(5, 15))
	
	// Different return types
	area := calculateArea(5.0, 3.0)
	fmt.Printf("Area: %.2f\n", area)
	
	isEven := checkEven(42)
	fmt.Printf("Is 42 even? %t\n", isEven)
	
	fmt.Println()

	// ============================================
	// 4. MULTIPLE RETURN VALUES
	// ============================================
	fmt.Println("--- 4. Multiple Return Values ---")
	
	// Function returning multiple values
	sum, product := calculate(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
	
	// Swap values
	x, y := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("After swap: x=%d, y=%d\n", x, y)
	
	// Get quotient and remainder
	quot, rem := divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", quot, rem)
	
	fmt.Println()

	// ============================================
	// 5. IGNORING RETURN VALUES
	// ============================================
	fmt.Println("--- 5. Ignoring Return Values ---")
	
	// Ignore one return value with underscore
	sumOnly, _ := calculate(3, 7)
	fmt.Printf("Sum only: %d\n", sumOnly)
	
	_, productOnly := calculate(3, 7)
	fmt.Printf("Product only: %d\n", productOnly)
	
	// Ignore all return values (though not recommended)
	calculate(8, 9)
	
	fmt.Println()

	// ============================================
	// 6. NAMED RETURN VALUES
	// ============================================
	fmt.Println("--- 6. Named Return Values ---")
	
	// Function with named return values
	s, p := calculateWithNames(6, 7)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)
	
	// Rectangle dimensions
	perimeter, rectArea := rectangleStats(5.0, 3.0)
	fmt.Printf("Rectangle (5x3) - Perimeter: %.2f, Area: %.2f\n", perimeter, rectArea)
	
	fmt.Println()

	// ============================================
	// 7. FUNCTIONS AS PARAMETERS
	// ============================================
	fmt.Println("--- 7. Functions as Parameters ---")
	
	// Pass function as parameter
	numbers := []int{1, 2, 3, 4, 5}
	
	fmt.Print("Original: ")
	printSlice(numbers)
	
	fmt.Print("Doubled: ")
	result1 := applyOperation(numbers, double)
	printSlice(result1)
	
	fmt.Print("Squared: ")
	result2 := applyOperation(numbers, square)
	printSlice(result2)
	
	fmt.Println()

	// ============================================
	// 8. ANONYMOUS FUNCTIONS
	// ============================================
	fmt.Println("--- 8. Anonymous Functions ---")
	
	// Define and call immediately
	func() {
		fmt.Println("This is an anonymous function")
	}()
	
	// Anonymous function with parameters
	func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}("World")
	
	// Assign to variable
	greeting := func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(greeting("Go Developer"))
	
	fmt.Println()

	// ============================================
	// 9. CLOSURES
	// ============================================
	fmt.Println("--- 9. Closures ---")
	
	// Function returning a function (closure)
	counter := makeCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	
	// Another counter instance
	counter2 := makeCounter()
	fmt.Printf("Counter2: %d\n", counter2())
	
	// Multiplier closure
	multiplyBy2 := makeMultiplier(2)
	multiplyBy5 := makeMultiplier(5)
	fmt.Printf("10 * 2 = %d\n", multiplyBy2(10))
	fmt.Printf("10 * 5 = %d\n", multiplyBy5(10))
	
	fmt.Println()

	// ============================================
	// 10. VARIADIC FUNCTIONS
	// ============================================
	fmt.Println("--- 10. Variadic Functions ---")
	
	// Function accepting variable number of arguments
	fmt.Printf("Sum of 1, 2, 3: %d\n", sumAll(1, 2, 3))
	fmt.Printf("Sum of 1, 2, 3, 4, 5: %d\n", sumAll(1, 2, 3, 4, 5))
	fmt.Printf("Sum of 10: %d\n", sumAll(10))
	
	// String concatenation
	message := concatenate("Hello", "World", "from", "Go")
	fmt.Println(message)
	
	// Find maximum
	fmt.Printf("Max of 3, 7, 2, 9, 1: %d\n", findMax(3, 7, 2, 9, 1))
	
	fmt.Println()

	// ============================================
	// 11. RECURSIVE FUNCTIONS
	// ============================================
	fmt.Println("--- 11. Recursive Functions ---")
	
	// Calculate factorial
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Factorial of 7: %d\n", factorial(7))
	
	// Fibonacci sequence
	fmt.Print("Fibonacci(0-10): ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
	
	// Countdown
	fmt.Print("Countdown: ")
	countdown(5)
	fmt.Println()
	
	fmt.Println()

	// ============================================
	// 12. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 12. Practical Examples ---")
	
	// Example 1: Temperature converter
	fmt.Println("\nExample 1: Temperature Converter")
	celsius := 25.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)
	
	// Example 2: String utilities
	fmt.Println("\nExample 2: String Utilities")
	text := "Hello, Go!"
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Uppercase: %s\n", toUpperCase(text))
	fmt.Printf("Reversed: %s\n", reverseString(text))
	fmt.Printf("Length: %d\n", getStringLength(text))
	
	// Example 3: Math utilities
	fmt.Println("\nExample 3: Math Utilities")
	nums := []int{10, 5, 20, 15, 8}
	fmt.Printf("Numbers: %v\n", nums)
	fmt.Printf("Sum: %d\n", sumSlice(nums))
	fmt.Printf("Average: %.2f\n", average(nums))
	fmt.Printf("Min: %d\n", findMin(nums))
	fmt.Printf("Max: %d\n", findMaxSlice(nums))
	
	// Example 4: Validation functions
	fmt.Println("\nExample 4: Validation Functions")
	password := "SecurePass123"
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Is valid length: %t\n", isValidLength(password, 8, 20))
	fmt.Printf("Contains number: %t\n", containsNumber(password))
	fmt.Printf("Contains uppercase: %t\n", containsUppercase(password))
	
	// Example 5: Calculator
	fmt.Println("\nExample 5: Calculator")
	a, b := 20.0, 5.0
	fmt.Printf("a = %.1f, b = %.1f\n", a, b)
	fmt.Printf("Add: %.1f\n", calculator(a, b, "+"))
	fmt.Printf("Subtract: %.1f\n", calculator(a, b, "-"))
	fmt.Printf("Multiply: %.1f\n", calculator(a, b, "*"))
	fmt.Printf("Divide: %.1f\n", calculator(a, b, "/"))
	
	// Example 6: Grade calculator
	fmt.Println("\nExample 6: Grade Calculator")
	scores := []int{85, 90, 78, 92, 88}
	fmt.Printf("Scores: %v\n", scores)
	grade, gpa := calculateGrade(scores)
	fmt.Printf("Grade: %s, GPA: %.2f\n", grade, gpa)
	
	// Example 7: Discount calculator
	fmt.Println("\nExample 7: Discount Calculator")
	price := 100.0
	discount := 15.0
	final := applyDiscount(price, discount)
	fmt.Printf("Original: $%.2f, Discount: %.0f%%, Final: $%.2f\n", 
		price, discount, final)
	
	// Example 8: String formatter
	fmt.Println("\nExample 8: String Formatter")
	firstName := "John"
	lastName := "Doe"
	fullName := formatFullName(firstName, lastName)
	fmt.Printf("Full Name: %s\n", fullName)
	initials := getInitials(firstName, lastName)
	fmt.Printf("Initials: %s\n", initials)
	
	fmt.Println("\n=== End of Examples ===")
}

// ============================================
// FUNCTION DEFINITIONS
// ============================================

// 1. Basic functions without parameters or return
func greet() {
	fmt.Println("Hello, World!")
}

func sayHello() {
	fmt.Println("Hi there!")
}

// 2. Functions with parameters
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func addNumbers(a int, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

func multiply(x, y int) { // Same type parameters
	result := x * y
	fmt.Printf("%d * %d = %d\n", x, y, result)
}

// 3. Functions with return values
func add(a, b int) int {
	return a + b
}

func calculateArea(width, height float64) float64 {
	return width * height
}

func checkEven(n int) bool {
	return n%2 == 0
}

// 4. Multiple return values
func calculate(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func swap(x, y int) (int, int) {
	return y, x
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 6. Named return values
func calculateWithNames(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // Naked return
}

func rectangleStats(width, height float64) (perimeter float64, area float64) {
	perimeter = 2 * (width + height)
	area = width * height
	return
}

// 7. Functions as parameters
func applyOperation(numbers []int, operation func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = operation(num)
	}
	return result
}

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}

func printSlice(nums []int) {
	fmt.Println(nums)
}

// 9. Closures
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// 10. Variadic functions
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func concatenate(words ...string) string {
	return strings.Join(words, " ")
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// 11. Recursive functions
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func countdown(n int) {
	if n <= 0 {
		fmt.Print("Go!")
		return
	}
	fmt.Printf("%d ", n)
	countdown(n - 1)
}

// 12. Practical examples
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getStringLength(s string) int {
	return len(s)
}

func sumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := sumSlice(numbers)
	return float64(sum) / float64(len(numbers))
}

func findMin(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func findMaxSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func isValidLength(s string, min, max int) bool {
	length := len(s)
	return length >= min && length <= max
}

func containsNumber(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func containsUppercase(s string) bool {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func calculator(a, b float64, operation string) float64 {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}
		return 0
	default:
		return 0
	}
}

func calculateGrade(scores []int) (string, float64) {
	avg := average(scores)
	var grade string
	var gpa float64
	
	switch {
	case avg >= 90:
		grade = "A"
		gpa = 4.0
	case avg >= 80:
		grade = "B"
		gpa = 3.0
	case avg >= 70:
		grade = "C"
		gpa = 2.0
	case avg >= 60:
		grade = "D"
		gpa = 1.0
	default:
		grade = "F"
		gpa = 0.0
	}
	
	return grade, gpa
}

func applyDiscount(price, discountPercent float64) float64 {
	discount := price * (discountPercent / 100)
	return math.Round((price-discount)*100) / 100
}

func formatFullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

func getInitials(firstName, lastName string) string {
	if len(firstName) == 0 || len(lastName) == 0 {
		return ""
	}
	return string(firstName[0]) + "." + string(lastName[0]) + "."
}

