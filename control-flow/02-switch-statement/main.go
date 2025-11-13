package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Switch Statement in Go ===\n")

	// ============================================
	// 1. BASIC SWITCH
	// ============================================
	fmt.Println("--- 1. Basic Switch ---")
	
	day := 3
	fmt.Printf("Day number: %d\n", day)
	
	// Basic switch with cases
	// No 'break' needed - Go automatically breaks after each case
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
	fmt.Println()

	// ============================================
	// 2. MULTIPLE VALUES IN CASE
	// ============================================
	fmt.Println("--- 2. Multiple Values in Case ---")
	
	month := "March"
	fmt.Printf("Month: %s\n", month)
	
	// Multiple values separated by commas
	switch month {
	case "December", "January", "February":
		fmt.Println("Season: Winter")
	case "March", "April", "May":
		fmt.Println("Season: Spring")
	case "June", "July", "August":
		fmt.Println("Season: Summer")
	case "September", "October", "November":
		fmt.Println("Season: Fall")
	default:
		fmt.Println("Unknown month")
	}
	fmt.Println()

	// ============================================
	// 3. SWITCH WITH EXPRESSIONS
	// ============================================
	fmt.Println("--- 3. Switch with Expressions ---")
	
	score := 85
	fmt.Printf("Score: %d\n", score)
	
	// Cases can contain expressions
	switch {
	case score >= 90:
		fmt.Println("Grade: A (Excellent)")
	case score >= 80:
		fmt.Println("Grade: B (Very Good)")
	case score >= 70:
		fmt.Println("Grade: C (Good)")
	case score >= 60:
		fmt.Println("Grade: D (Pass)")
	default:
		fmt.Println("Grade: F (Fail)")
	}
	fmt.Println()

	// ============================================
	// 4. SWITCH WITHOUT CONDITION (LIKE IF-ELSE)
	// ============================================
	fmt.Println("--- 4. Switch without Condition ---")
	
	temperature := 28
	fmt.Printf("Temperature: %d°C\n", temperature)
	
	// Switch without expression acts like if-else chain
	switch {
	case temperature < 0:
		fmt.Println("Freezing cold")
	case temperature >= 0 && temperature < 15:
		fmt.Println("Cold")
	case temperature >= 15 && temperature < 25:
		fmt.Println("Pleasant")
	case temperature >= 25 && temperature < 35:
		fmt.Println("Hot")
	default:
		fmt.Println("Very hot")
	}
	fmt.Println()

	// ============================================
	// 5. SWITCH WITH INITIALIZATION
	// ============================================
	fmt.Println("--- 5. Switch with Initialization ---")
	
	// Declare and initialize variable in switch statement
	switch hour := time.Now().Hour(); {
	case hour >= 5 && hour < 12:
		fmt.Printf("Time: %d:00 - Good morning\n", hour)
	case hour >= 12 && hour < 18:
		fmt.Printf("Time: %d:00 - Good afternoon\n", hour)
	case hour >= 18 && hour < 22:
		fmt.Printf("Time: %d:00 - Good evening\n", hour)
	default:
		fmt.Printf("Time: %d:00 - Good night\n", hour)
	}
	fmt.Println()

	// ============================================
	// 6. TYPE SWITCH
	// ============================================
	fmt.Println("--- 6. Type Switch ---")
	
	// Type switch checks the type of an interface variable
	var value interface{} = "Hello, Go!"
	
	fmt.Printf("Value: %v\n", value)
	
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	case string:
		fmt.Printf("String: %s (length: %d)\n", v, len(v))
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
	
	// Another type switch example
	fmt.Println("\nChecking different types:")
	checkType(42)
	checkType(3.14)
	checkType(true)
	checkType("Go")
	checkType([]int{1, 2, 3})
	fmt.Println()

	// ============================================
	// 7. FALLTHROUGH
	// ============================================
	fmt.Println("--- 7. Fallthrough ---")
	
	number := 2
	fmt.Printf("Number: %d\n", number)
	fmt.Println("Processing:")
	
	// fallthrough forces execution to continue to next case
	switch number {
	case 1:
		fmt.Println("  One")
		fallthrough
	case 2:
		fmt.Println("  Two or after one")
		fallthrough
	case 3:
		fmt.Println("  Three or after two")
	default:
		fmt.Println("  Other")
	}
	fmt.Println()

	// ============================================
	// 8. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 8. Practical Examples ---")
	
	// Example 1: HTTP Status Code Handler
	fmt.Println("\nHTTP Status Code Handler:")
	statusCode := 404
	fmt.Printf("Status Code: %d\n", statusCode)
	
	switch statusCode {
	case 200:
		fmt.Println("OK - Request successful")
	case 201:
		fmt.Println("Created - Resource created successfully")
	case 400:
		fmt.Println("Bad Request - Invalid request")
	case 401:
		fmt.Println("Unauthorized - Authentication required")
	case 403:
		fmt.Println("Forbidden - Access denied")
	case 404:
		fmt.Println("Not Found - Resource not found")
	case 500:
		fmt.Println("Internal Server Error")
	case 503:
		fmt.Println("Service Unavailable")
	default:
		fmt.Println("Unknown status code")
	}
	
	// Example 2: File Extension Handler
	fmt.Println("\nFile Extension Handler:")
	filename := "document.pdf"
	fmt.Printf("Filename: %s\n", filename)
	
	// Extract extension (simplified)
	extension := ".pdf"
	
	switch extension {
	case ".txt", ".md", ".log":
		fmt.Println("File Type: Text document")
	case ".jpg", ".jpeg", ".png", ".gif":
		fmt.Println("File Type: Image")
	case ".mp4", ".avi", ".mov":
		fmt.Println("File Type: Video")
	case ".mp3", ".wav", ".flac":
		fmt.Println("File Type: Audio")
	case ".pdf":
		fmt.Println("File Type: PDF document")
	case ".doc", ".docx":
		fmt.Println("File Type: Word document")
	case ".xls", ".xlsx":
		fmt.Println("File Type: Excel spreadsheet")
	default:
		fmt.Println("File Type: Unknown")
	}
	
	// Example 3: User Role Permissions
	fmt.Println("\nUser Role Permissions:")
	userRole := "moderator"
	fmt.Printf("User Role: %s\n", userRole)
	
	switch userRole {
	case "admin":
		fmt.Println("Permissions:")
		fmt.Println("  - Full system access")
		fmt.Println("  - User management")
		fmt.Println("  - Content management")
		fmt.Println("  - Settings management")
	case "moderator":
		fmt.Println("Permissions:")
		fmt.Println("  - Content moderation")
		fmt.Println("  - User warnings")
		fmt.Println("  - Report management")
	case "user":
		fmt.Println("Permissions:")
		fmt.Println("  - Create posts")
		fmt.Println("  - Comment")
		fmt.Println("  - Like content")
	case "guest":
		fmt.Println("Permissions:")
		fmt.Println("  - View content only")
	default:
		fmt.Println("Invalid role")
	}
	
	// Example 4: Payment Method Processor
	fmt.Println("\nPayment Method Processor:")
	paymentMethod := "credit_card"
	amount := 99.99
	
	fmt.Printf("Payment Method: %s\n", paymentMethod)
	fmt.Printf("Amount: $%.2f\n", amount)
	
	switch paymentMethod {
	case "credit_card":
		fmt.Println("Processing credit card payment...")
		fmt.Println("Fee: 2.9% + $0.30")
		fee := amount*0.029 + 0.30
		fmt.Printf("Total: $%.2f\n", amount+fee)
	case "debit_card":
		fmt.Println("Processing debit card payment...")
		fmt.Println("Fee: 1.5%")
		fee := amount * 0.015
		fmt.Printf("Total: $%.2f\n", amount+fee)
	case "paypal":
		fmt.Println("Processing PayPal payment...")
		fmt.Println("Fee: 2.5%")
		fee := amount * 0.025
		fmt.Printf("Total: $%.2f\n", amount+fee)
	case "bank_transfer":
		fmt.Println("Processing bank transfer...")
		fmt.Println("Fee: $0.00")
		fmt.Printf("Total: $%.2f\n", amount)
	case "crypto":
		fmt.Println("Processing cryptocurrency payment...")
		fmt.Println("Fee: Network fees apply")
		fmt.Printf("Amount: $%.2f\n", amount)
	default:
		fmt.Println("Invalid payment method")
	}
	
	// Example 5: Traffic Light Controller
	fmt.Println("\nTraffic Light Controller:")
	lightColor := "yellow"
	fmt.Printf("Light Color: %s\n", lightColor)
	
	switch lightColor {
	case "green":
		fmt.Println("Action: GO")
		fmt.Println("Speed: Normal")
	case "yellow":
		fmt.Println("Action: CAUTION")
		fmt.Println("Prepare to stop")
	case "red":
		fmt.Println("Action: STOP")
		fmt.Println("Wait for green light")
	default:
		fmt.Println("Error: Invalid light color")
	}
	
	// Example 6: Grade Calculator with Range
	fmt.Println("\nGrade Calculator:")
	percentage := 87
	fmt.Printf("Percentage: %d%%\n", percentage)
	
	var letterGrade string
	var gpa float64
	
	switch {
	case percentage >= 97:
		letterGrade = "A+"
		gpa = 4.0
	case percentage >= 93:
		letterGrade = "A"
		gpa = 4.0
	case percentage >= 90:
		letterGrade = "A-"
		gpa = 3.7
	case percentage >= 87:
		letterGrade = "B+"
		gpa = 3.3
	case percentage >= 83:
		letterGrade = "B"
		gpa = 3.0
	case percentage >= 80:
		letterGrade = "B-"
		gpa = 2.7
	case percentage >= 77:
		letterGrade = "C+"
		gpa = 2.3
	case percentage >= 73:
		letterGrade = "C"
		gpa = 2.0
	case percentage >= 70:
		letterGrade = "C-"
		gpa = 1.7
	case percentage >= 67:
		letterGrade = "D+"
		gpa = 1.3
	case percentage >= 65:
		letterGrade = "D"
		gpa = 1.0
	default:
		letterGrade = "F"
		gpa = 0.0
	}
	
	fmt.Printf("Letter Grade: %s\n", letterGrade)
	fmt.Printf("GPA: %.1f\n", gpa)
	
	// Example 7: Command Line Interface
	fmt.Println("\nCommand Line Interface:")
	command := "list"
	fmt.Printf("Command: %s\n", command)
	
	switch command {
	case "help", "h", "?":
		fmt.Println("Available commands:")
		fmt.Println("  list   - List all items")
		fmt.Println("  add    - Add new item")
		fmt.Println("  delete - Delete item")
		fmt.Println("  exit   - Exit program")
	case "list", "ls", "l":
		fmt.Println("Listing all items...")
		fmt.Println("  1. Item A")
		fmt.Println("  2. Item B")
		fmt.Println("  3. Item C")
	case "add", "new", "create":
		fmt.Println("Adding new item...")
	case "delete", "remove", "rm":
		fmt.Println("Deleting item...")
	case "exit", "quit", "q":
		fmt.Println("Exiting program...")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Type 'help' for available commands")
	}
	
	fmt.Println("\n=== End of Examples ===")
}

// Helper function for type switch example
func checkType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("  %v is an int\n", v)
	case float64:
		fmt.Printf("  %v is a float64\n", v)
	case bool:
		fmt.Printf("  %v is a bool\n", v)
	case string:
		fmt.Printf("  %v is a string\n", v)
	case []int:
		fmt.Printf("  %v is a slice of ints\n", v)
	default:
		fmt.Printf("  %v is of type %T\n", v, v)
	}
}

