package main

// ============================================
// SINGLE IMPORT
// ============================================
// import "fmt"

// ============================================
// MULTIPLE IMPORTS (Recommended style)
// ============================================
import (
	"fmt"      // Formatted I/O
	"math"     // Mathematical functions
	"strings"  // String manipulation
	"time"     // Time and date functions
	"math/rand" // Random numbers
)

// ============================================
// IMPORT WITH ALIAS
// ============================================
// You can give packages custom names
// import (
//     f "fmt"              // Now use f.Println instead of fmt.Println
//     m "math"             // Now use m.Sqrt instead of math.Sqrt
// )

// ============================================
// IMPORT WITH DOT (Not recommended in production)
// ============================================
// Imports all exported names into current namespace
// import . "fmt"
// Now you can use Println() instead of fmt.Println()
// This can cause confusion and naming conflicts

// ============================================
// BLANK IMPORT (for side effects)
// ============================================
// Used when you need the package's init() function but not its exports
// import _ "image/png"  // Registers PNG decoder
// Common in database drivers: import _ "github.com/lib/pq"

func main() {
	fmt.Println("=== Package Import Examples ===\n")

	// ============================================
	// Using fmt package
	// ============================================
	fmt.Println("--- fmt Package ---")
	name := "Gopher"
	fmt.Printf("Hello, %s!\n", name)
	fmt.Println()

	// ============================================
	// Using math package
	// ============================================
	fmt.Println("--- math Package ---")
	
	// Mathematical constants
	fmt.Printf("Pi: %.4f\n", math.Pi)
	fmt.Printf("E: %.4f\n", math.E)
	
	// Mathematical functions
	number := 16.0
	fmt.Printf("Square root of %.0f: %.2f\n", number, math.Sqrt(number))
	fmt.Printf("Power 2^10: %.0f\n", math.Pow(2, 10))
	fmt.Printf("Absolute value of -42: %.0f\n", math.Abs(-42))
	
	// Rounding
	value := 3.7
	fmt.Printf("Floor of %.1f: %.0f\n", value, math.Floor(value))
	fmt.Printf("Ceil of %.1f: %.0f\n", value, math.Ceil(value))
	fmt.Printf("Round of %.1f: %.0f\n", value, math.Round(value))
	
	// Trigonometry
	angle := 45.0
	radians := angle * math.Pi / 180
	fmt.Printf("Sin(45°): %.4f\n", math.Sin(radians))
	fmt.Printf("Cos(45°): %.4f\n", math.Cos(radians))
	fmt.Println()

	// ============================================
	// Using strings package
	// ============================================
	fmt.Println("--- strings Package ---")
	
	text := "Go Programming Language"
	
	// String manipulation
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(text))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(text))
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("Starts with 'Go': %t\n", strings.HasPrefix(text, "Go"))
	fmt.Printf("Ends with 'Language': %t\n", strings.HasSuffix(text, "Language"))
	
	// String operations
	fmt.Printf("Replace 'Go' with 'Golang': %s\n", strings.Replace(text, "Go", "Golang", 1))
	fmt.Printf("Repeat 'Go!' 3 times: %s\n", strings.Repeat("Go! ", 3))
	
	// Split and Join
	words := strings.Split(text, " ")
	fmt.Printf("Split into words: %v\n", words)
	fmt.Printf("Join with '-': %s\n", strings.Join(words, "-"))
	
	// Trim
	spaced := "   Hello, World!   "
	fmt.Printf("Original: '%s'\n", spaced)
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(spaced))
	fmt.Println()

	// ============================================
	// Using time package
	// ============================================
	fmt.Println("--- time Package ---")
	
	// Current time
	now := time.Now()
	fmt.Printf("Current time: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Current date: %s\n", now.Format("January 2, 2006"))
	fmt.Printf("Current time (12h): %s\n", now.Format("03:04:05 PM"))
	
	// Time components
	fmt.Printf("Year: %d\n", now.Year())
	fmt.Printf("Month: %s\n", now.Month())
	fmt.Printf("Day: %d\n", now.Day())
	fmt.Printf("Hour: %d\n", now.Hour())
	fmt.Printf("Weekday: %s\n", now.Weekday())
	
	// Unix timestamp
	fmt.Printf("Unix timestamp: %d\n", now.Unix())
	
	// Duration
	duration := 2 * time.Hour + 30 * time.Minute
	fmt.Printf("Duration: %s\n", duration)
	fmt.Printf("Hours: %.2f\n", duration.Hours())
	fmt.Printf("Minutes: %.0f\n", duration.Minutes())
	fmt.Println()

	// ============================================
	// Using math/rand package
	// ============================================
	fmt.Println("--- math/rand Package ---")
	
	// Seed the random number generator
	// Note: For better randomness in production, use crypto/rand
	rand.Seed(time.Now().UnixNano())
	
	// Random numbers
	fmt.Printf("Random int: %d\n", rand.Int())
	fmt.Printf("Random int (0-99): %d\n", rand.Intn(100))
	fmt.Printf("Random float64: %.4f\n", rand.Float64())
	fmt.Printf("Random float64 (0-10): %.4f\n", rand.Float64()*10)
	
	// Random selection
	colors := []string{"Red", "Green", "Blue", "Yellow", "Purple"}
	randomColor := colors[rand.Intn(len(colors))]
	fmt.Printf("Random color: %s\n", randomColor)
	fmt.Println()

	// ============================================
	// Combining packages
	// ============================================
	fmt.Println("--- Combining Multiple Packages ---")
	
	// Generate a random greeting
	greetings := []string{"Hello", "Hi", "Hey", "Greetings"}
	names := []string{"Alice", "Bob", "Charlie", "Diana"}
	
	randomGreeting := greetings[rand.Intn(len(greetings))]
	randomName := names[rand.Intn(len(names))]
	
	message := fmt.Sprintf("%s, %s!", randomGreeting, randomName)
	fmt.Printf("Random message: %s\n", message)
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(message))
	fmt.Printf("Length: %d characters\n", len(message))
	
	// Calculate and format
	price := 99.99
	discount := 0.15
	finalPrice := price * (1 - discount)
	fmt.Printf("Original price: $%.2f\n", price)
	fmt.Printf("Discount: %.0f%%\n", discount*100)
	fmt.Printf("Final price: $%.2f\n", math.Round(finalPrice*100)/100)
}

