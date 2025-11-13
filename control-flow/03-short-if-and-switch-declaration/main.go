package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Short If and Switch Declaration in Go ===\n")

	// ============================================
	// 1. BASIC SHORT IF DECLARATION
	// ============================================
	fmt.Println("--- 1. Basic Short If Declaration ---")
	
	// Traditional way
	fmt.Println("Traditional approach:")
	value1 := 42
	if value1 > 0 {
		fmt.Printf("Value %d is positive\n", value1)
	}
	
	// Short declaration (variable only exists in if block)
	fmt.Println("\nShort declaration:")
	if value2 := 42; value2 > 0 {
		fmt.Printf("Value %d is positive\n", value2)
		// value2 is accessible here
	}
	// value2 is NOT accessible here
	
	fmt.Println()

	// ============================================
	// 2. SHORT IF WITH ELSE
	// ============================================
	fmt.Println("--- 2. Short If with Else ---")
	
	// Variable is accessible in if and else blocks
	if age := 25; age >= 18 {
		fmt.Printf("Age %d: Adult\n", age)
	} else {
		fmt.Printf("Age %d: Minor\n", age)
	}
	
	// Another example
	if score := 85; score >= 90 {
		fmt.Printf("Score %d: Grade A\n", score)
	} else if score >= 80 {
		fmt.Printf("Score %d: Grade B\n", score)
	} else if score >= 70 {
		fmt.Printf("Score %d: Grade C\n", score)
	} else {
		fmt.Printf("Score %d: Grade F\n", score)
	}
	
	fmt.Println()

	// ============================================
	// 3. SHORT IF WITH FUNCTION CALLS
	// ============================================
	fmt.Println("--- 3. Short If with Function Calls ---")
	
	// Call function and check result immediately
	if result := calculateSum(10, 20); result > 25 {
		fmt.Printf("Sum result %d is greater than 25\n", result)
	}
	
	// With string processing
	if length := len("Hello, Go!"); length > 5 {
		fmt.Printf("String length %d is greater than 5\n", length)
	}
	
	// With mathematical operations
	if squareRoot := math.Sqrt(16); squareRoot == 4 {
		fmt.Printf("Square root is exactly %.0f\n", squareRoot)
	}
	
	fmt.Println()

	// ============================================
	// 4. ERROR HANDLING PATTERN
	// ============================================
	fmt.Println("--- 4. Error Handling Pattern ---")
	
	// Very common pattern in Go
	if err := validateAge(25); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Age validation passed")
	}
	
	// String conversion with error handling
	if num, err := strconv.Atoi("123"); err != nil {
		fmt.Printf("Conversion error: %v\n", err)
	} else {
		fmt.Printf("Successfully converted to number: %d\n", num)
	}
	
	// Invalid conversion
	if num, err := strconv.Atoi("abc"); err != nil {
		fmt.Printf("Conversion error: %v\n", err)
	} else {
		fmt.Printf("Successfully converted to number: %d\n", num)
	}
	
	fmt.Println()

	// ============================================
	// 5. MULTIPLE VARIABLE DECLARATIONS IN IF
	// ============================================
	fmt.Println("--- 5. Multiple Variables in Short If ---")
	
	// Can declare multiple variables
	if x, y := 10, 20; x < y {
		fmt.Printf("x=%d is less than y=%d\n", x, y)
	}
	
	// Complex calculations
	if width, height := 10, 5; width*height > 40 {
		area := width * height
		fmt.Printf("Area %d (width=%d, height=%d) is greater than 40\n", 
			area, width, height)
	}
	
	fmt.Println()

	// ============================================
	// 6. BASIC SHORT SWITCH DECLARATION
	// ============================================
	fmt.Println("--- 6. Basic Short Switch Declaration ---")
	
	// Traditional way
	fmt.Println("Traditional approach:")
	day1 := time.Now().Weekday()
	switch day1 {
	case time.Monday:
		fmt.Println("Start of work week")
	case time.Friday:
		fmt.Println("Almost weekend!")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek")
	}
	
	// Short declaration
	fmt.Println("\nShort declaration:")
	switch day2 := time.Now().Weekday(); day2 {
	case time.Monday:
		fmt.Printf("%v: Start of work week\n", day2)
	case time.Friday:
		fmt.Printf("%v: Almost weekend!\n", day2)
	case time.Saturday, time.Sunday:
		fmt.Printf("%v: Weekend!\n", day2)
	default:
		fmt.Printf("%v: Midweek\n", day2)
	}
	
	fmt.Println()

	// ============================================
	// 7. SHORT SWITCH WITH EXPRESSION
	// ============================================
	fmt.Println("--- 7. Short Switch with Expression ---")
	
	// Declare variable and use in case expressions
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Printf("Hour %d: Good morning\n", hour)
	case hour < 18:
		fmt.Printf("Hour %d: Good afternoon\n", hour)
	default:
		fmt.Printf("Hour %d: Good evening\n", hour)
	}
	
	// With calculations
	switch score := 85; {
	case score >= 90:
		fmt.Printf("Score %d: Excellent\n", score)
	case score >= 80:
		fmt.Printf("Score %d: Very Good\n", score)
	case score >= 70:
		fmt.Printf("Score %d: Good\n", score)
	default:
		fmt.Printf("Score %d: Needs improvement\n", score)
	}
	
	fmt.Println()

	// ============================================
	// 8. SHORT SWITCH WITH FUNCTION CALLS
	// ============================================
	fmt.Println("--- 8. Short Switch with Function Calls ---")
	
	// Call function and switch on result
	switch result := getGrade(87); result {
	case "A":
		fmt.Println("Outstanding performance!")
	case "B":
		fmt.Println("Great job!")
	case "C":
		fmt.Println("Good effort")
	default:
		fmt.Println("Keep trying")
	}
	
	// With string processing
	switch name := strings.ToLower("GOLANG"); name {
	case "go", "golang":
		fmt.Printf("Detected Go language: %s\n", name)
	case "python":
		fmt.Println("Python detected")
	default:
		fmt.Printf("Unknown language: %s\n", name)
	}
	
	fmt.Println()

	// ============================================
	// 9. SCOPE DEMONSTRATION
	// ============================================
	fmt.Println("--- 9. Variable Scope ---")
	
	// Variable only exists in if-else block
	if temp := 25; temp > 20 {
		fmt.Printf("Inside if: temp = %d (accessible)\n", temp)
		if temp > 22 {
			fmt.Printf("Inside nested if: temp = %d (still accessible)\n", temp)
		}
	} else {
		fmt.Printf("Inside else: temp = %d (accessible)\n", temp)
	}
	// temp is NOT accessible here
	fmt.Println("Outside if-else: temp is out of scope")
	
	// Variable only exists in switch block
	switch num := 10; {
	case num > 5:
		fmt.Printf("Inside case: num = %d (accessible)\n", num)
	default:
		fmt.Printf("Inside default: num = %d (accessible)\n", num)
	}
	// num is NOT accessible here
	fmt.Println("Outside switch: num is out of scope")
	
	fmt.Println()

	// ============================================
	// 10. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 10. Practical Examples ---")
	
	// Example 1: Input validation
	fmt.Println("\nExample 1: Input Validation")
	if password := getUserPassword(); len(password) < 8 {
		fmt.Println("Password too short (minimum 8 characters)")
	} else if len(password) > 20 {
		fmt.Println("Password too long (maximum 20 characters)")
	} else {
		fmt.Println("Password length is valid")
	}
	
	// Example 2: File size checker
	fmt.Println("\nExample 2: File Size Checker")
	if fileSize := getFileSize(); fileSize < 1024 {
		fmt.Printf("File size: %d bytes (very small)\n", fileSize)
	} else if fileSize < 1024*1024 {
		fmt.Printf("File size: %.2f KB\n", float64(fileSize)/1024)
	} else {
		fmt.Printf("File size: %.2f MB\n", float64(fileSize)/(1024*1024))
	}
	
	// Example 3: HTTP Status Handler
	fmt.Println("\nExample 3: HTTP Status Code Handler")
	switch statusCode := getHTTPStatus(); {
	case statusCode >= 200 && statusCode < 300:
		fmt.Printf("Status %d: Success\n", statusCode)
	case statusCode >= 300 && statusCode < 400:
		fmt.Printf("Status %d: Redirection\n", statusCode)
	case statusCode >= 400 && statusCode < 500:
		fmt.Printf("Status %d: Client Error\n", statusCode)
	case statusCode >= 500:
		fmt.Printf("Status %d: Server Error\n", statusCode)
	default:
		fmt.Printf("Status %d: Invalid\n", statusCode)
	}
	
	// Example 4: Temperature converter
	fmt.Println("\nExample 4: Temperature Converter")
	if celsius := 25.0; celsius < 0 {
		fmt.Printf("%.1f°C = %.1f°F (Freezing)\n", celsius, celsiusToFahrenheit(celsius))
	} else if celsius < 20 {
		fmt.Printf("%.1f°C = %.1f°F (Cold)\n", celsius, celsiusToFahrenheit(celsius))
	} else if celsius < 30 {
		fmt.Printf("%.1f°C = %.1f°F (Comfortable)\n", celsius, celsiusToFahrenheit(celsius))
	} else {
		fmt.Printf("%.1f°C = %.1f°F (Hot)\n", celsius, celsiusToFahrenheit(celsius))
	}
	
	// Example 5: User role checker
	fmt.Println("\nExample 5: User Role Permissions")
	switch role := getUserRole(); role {
	case "admin":
		fmt.Printf("Role '%s': Full access granted\n", role)
	case "moderator":
		fmt.Printf("Role '%s': Moderation access granted\n", role)
	case "user":
		fmt.Printf("Role '%s': Standard access granted\n", role)
	default:
		fmt.Printf("Role '%s': Limited access\n", role)
	}
	
	// Example 6: Discount calculator
	fmt.Println("\nExample 6: Discount Calculator")
	if quantity := 8; quantity >= 10 {
		discount := 20
		fmt.Printf("Quantity %d: %d%% discount applied\n", quantity, discount)
	} else if quantity >= 5 {
		discount := 10
		fmt.Printf("Quantity %d: %d%% discount applied\n", quantity, discount)
	} else {
		fmt.Printf("Quantity %d: No discount\n", quantity)
	}
	
	// Example 7: Time-based greeting
	fmt.Println("\nExample 7: Time-based Greeting")
	switch hour := getCurrentHour(); {
	case hour >= 5 && hour < 12:
		fmt.Printf("Time %d:00 - Good morning! ☀️\n", hour)
	case hour >= 12 && hour < 18:
		fmt.Printf("Time %d:00 - Good afternoon! 🌤️\n", hour)
	case hour >= 18 && hour < 22:
		fmt.Printf("Time %d:00 - Good evening! 🌙\n", hour)
	default:
		fmt.Printf("Time %d:00 - Good night! 😴\n", hour)
	}
	
	// Example 8: BMI calculator with categories
	fmt.Println("\nExample 8: BMI Calculator")
	if bmi := calculateBMI(70, 1.75); bmi < 18.5 {
		fmt.Printf("BMI %.2f: Underweight\n", bmi)
	} else if bmi < 25 {
		fmt.Printf("BMI %.2f: Normal weight\n", bmi)
	} else if bmi < 30 {
		fmt.Printf("BMI %.2f: Overweight\n", bmi)
	} else {
		fmt.Printf("BMI %.2f: Obese\n", bmi)
	}
	
	// Example 9: Grade with GPA
	fmt.Println("\nExample 9: Grade with GPA")
	switch percentage := 87; {
	case percentage >= 90:
		fmt.Printf("Percentage %d%%: Grade A (GPA: 4.0)\n", percentage)
	case percentage >= 80:
		fmt.Printf("Percentage %d%%: Grade B (GPA: 3.0)\n", percentage)
	case percentage >= 70:
		fmt.Printf("Percentage %d%%: Grade C (GPA: 2.0)\n", percentage)
	case percentage >= 60:
		fmt.Printf("Percentage %d%%: Grade D (GPA: 1.0)\n", percentage)
	default:
		fmt.Printf("Percentage %d%%: Grade F (GPA: 0.0)\n", percentage)
	}
	
	// Example 10: String length validator
	fmt.Println("\nExample 10: Username Validator")
	if username := "gopher_dev"; len(username) < 3 {
		fmt.Printf("Username '%s': Too short (minimum 3 characters)\n", username)
	} else if len(username) > 20 {
		fmt.Printf("Username '%s': Too long (maximum 20 characters)\n", username)
	} else if strings.ContainsAny(username, "!@#$%^&*()") {
		fmt.Printf("Username '%s': Contains invalid characters\n", username)
	} else {
		fmt.Printf("Username '%s': Valid ✓\n", username)
	}
	
	fmt.Println("\n=== End of Examples ===")
}

// ============================================
// HELPER FUNCTIONS
// ============================================

func calculateSum(a, b int) int {
	return a + b
}

func validateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age must be at least 18")
	}
	return nil
}

func getGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	}
	return "F"
}

func getUserPassword() string {
	return "SecurePass123"
}

func getFileSize() int64 {
	return 2048576 // 2 MB in bytes
}

func getHTTPStatus() int {
	return 200
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func getUserRole() string {
	return "moderator"
}

func getCurrentHour() int {
	return time.Now().Hour()
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

