package main

import "fmt"

func main() {
	fmt.Println("=== If and Else in Go ===\n")

	// ============================================
	// 1. BASIC IF
	// ============================================
	fmt.Println("--- 1. Basic If ---")
	
	age := 20
	fmt.Printf("Age: %d\n", age)
	
	// Simple if without parentheses, but braces are mandatory
	if age >= 18 {
		fmt.Println("You are an adult")
	}
	
	temperature := 15
	fmt.Printf("Temperature: %d°C\n", temperature)
	
	if temperature < 20 {
		fmt.Println("It's cold, wear a coat")
	}
	fmt.Println()

	// ============================================
	// 2. IF-ELSE
	// ============================================
	fmt.Println("--- 2. If-Else ---")
	
	score := 75
	fmt.Printf("Score: %d\n", score)
	
	// If-else: executes one block or the other
	if score >= 60 {
		fmt.Println("Passed!")
	} else {
		fmt.Println("Failed")
	}
	
	// Another example
	isRaining := true
	fmt.Printf("Is it raining?: %t\n", isRaining)
	
	if isRaining {
		fmt.Println("Take an umbrella")
	} else {
		fmt.Println("You don't need an umbrella")
	}
	fmt.Println()

	// ============================================
	// 3. ELSE IF (MULTIPLE CONDITIONS)
	// ============================================
	fmt.Println("--- 3. Else If ---")
	
	grade := 85
	fmt.Printf("Grade: %d\n", grade)
	
	// Evaluates multiple conditions in sequence
	if grade >= 90 {
		fmt.Println("Grade: A (Excellent)")
	} else if grade >= 80 {
		fmt.Println("Grade: B (Very good)")
	} else if grade >= 70 {
		fmt.Println("Grade: C (Good)")
	} else if grade >= 60 {
		fmt.Println("Grade: D (Passed)")
	} else {
		fmt.Println("Grade: F (Failed)")
	}
	
	// Another example with different conditions
	time := 14
	fmt.Printf("Time: %d:00\n", time)
	
	if time >= 5 && time < 12 {
		fmt.Println("Good morning")
	} else if time >= 12 && time < 18 {
		fmt.Println("Good afternoon")
	} else if time >= 18 && time < 22 {
		fmt.Println("Good evening")
	} else {
		fmt.Println("You should be sleeping")
	}
	fmt.Println()

	// ============================================
	// 4. IF WITH INITIALIZATION
	// ============================================
	fmt.Println("--- 4. If with Initialization ---")
	
	// Variable declared in the if, only exists in the if-else block
	if userAge := 25; userAge >= 21 {
		fmt.Printf("User age %d: Can buy alcohol\n", userAge)
	} else {
		fmt.Printf("User age %d: Cannot buy alcohol\n", userAge)
	}
	// userAge doesn't exist here outside the if
	
	// Another example with calculation
	if total := 100 * 1.15; total > 100 {
		fmt.Printf("Total with taxes: $%.2f (more than $100)\n", total)
	} else {
		fmt.Printf("Total with taxes: $%.2f (less than $100)\n", total)
	}
	
	// Useful for validations
	if length := len("Hello, Go!"); length > 5 {
		fmt.Printf("The string has %d characters (more than 5)\n", length)
	}
	fmt.Println()

	// ============================================
	// 5. COMPARISON OPERATORS
	// ============================================
	fmt.Println("--- 5. Comparison Operators ---")
	
	x := 10
	y := 20
	fmt.Printf("x = %d, y = %d\n", x, y)
	
	// Equal to (==)
	if x == y {
		fmt.Println("x is equal to y")
	} else {
		fmt.Println("x is NOT equal to y")
	}
	
	// Not equal to (!=)
	if x != y {
		fmt.Println("x is different from y")
	}
	
	// Less than (<)
	if x < y {
		fmt.Println("x is less than y")
	}
	
	// Greater than (>)
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is NOT greater than y")
	}
	
	// Less than or equal to (<=)
	if x <= 10 {
		fmt.Println("x is less than or equal to 10")
	}
	
	// Greater than or equal to (>=)
	if y >= 20 {
		fmt.Println("y is greater than or equal to 20")
	}
	fmt.Println()

	// ============================================
	// 6. LOGICAL OPERATORS
	// ============================================
	fmt.Println("--- 6. Logical Operators ---")
	
	userName := "admin"
	password := "secret123"
	isActive := true
	
	fmt.Printf("User: %s, Active: %t\n", userName, isActive)
	
	// AND (&&) - Both conditions must be true
	if userName == "admin" && password == "secret123" {
		fmt.Println("Authentication successful (user AND password correct)")
	} else {
		fmt.Println("Authentication failed")
	}
	
	// OR (||) - At least one condition must be true
	if userName == "admin" || userName == "root" {
		fmt.Println("User with administrator permissions")
	}
	
	// NOT (!) - Negates the condition
	if !isActive {
		fmt.Println("User inactive")
	} else {
		fmt.Println("User active")
	}
	
	// Combination of operators
	hasPermission := true
	attemptCount := 3
	maxAttempts := 5
	
	if hasPermission && (attemptCount < maxAttempts) {
		fmt.Printf("Access granted (attempts: %d/%d)\n", attemptCount, maxAttempts)
	} else {
		fmt.Println("Access denied")
	}
	fmt.Println()

	// ============================================
	// 7. NESTED CONDITIONS
	// ============================================
	fmt.Println("--- 7. Nested Conditions ---")
	
	accountBalance := 500.0
	minimumPurchase := 50.0
	purchaseAmount := 75.0
	isPremiumMember := true
	
	fmt.Printf("Balance: $%.2f, Purchase: $%.2f, Premium: %t\n", 
		accountBalance, purchaseAmount, isPremiumMember)
	
	// Nested if
	if accountBalance >= purchaseAmount {
		if purchaseAmount >= minimumPurchase {
			if isPremiumMember {
				discount := purchaseAmount * 0.10
				finalPrice := purchaseAmount - discount
				fmt.Printf("Purchase approved with 10%% discount\n")
				fmt.Printf("  Original price: $%.2f\n", purchaseAmount)
				fmt.Printf("  Discount: $%.2f\n", discount)
				fmt.Printf("  Final price: $%.2f\n", finalPrice)
			} else {
				fmt.Printf("Purchase approved (price: $%.2f)\n", purchaseAmount)
				fmt.Println("  Become a premium member to get discounts")
			}
		} else {
			fmt.Printf("Minimum purchase required: $%.2f\n", minimumPurchase)
		}
	} else {
		fmt.Println("Insufficient balance")
	}
	fmt.Println()

	// ============================================
	// 8. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 8. Practical Examples ---")
	
	// Example 1: Age classifier
	fmt.Println("\nAge Classifier:")
	personAge := 35
	fmt.Printf("Age: %d years\n", personAge)
	
	if personAge < 0 {
		fmt.Println("Invalid age")
	} else if personAge <= 12 {
		fmt.Println("Category: Child")
	} else if personAge <= 17 {
		fmt.Println("Category: Teenager")
	} else if personAge <= 64 {
		fmt.Println("Category: Adult")
	} else {
		fmt.Println("Category: Senior")
	}
	
	// Example 2: Temperature evaluator
	fmt.Println("\nTemperature Evaluator:")
	celsius := 28.5
	fmt.Printf("Temperature: %.1f°C\n", celsius)
	
	if celsius < 0 {
		fmt.Println("Status: Freezing")
		fmt.Println("Recommendation: Heavy clothing and watch for ice")
	} else if celsius < 15 {
		fmt.Println("Status: Cold")
		fmt.Println("Recommendation: Wear a jacket")
	} else if celsius < 25 {
		fmt.Println("Status: Pleasant")
		fmt.Println("Recommendation: Perfect weather")
	} else if celsius < 35 {
		fmt.Println("Status: Hot")
		fmt.Println("Recommendation: Light clothing and stay hydrated")
	} else {
		fmt.Println("Status: Very hot")
		fmt.Println("Recommendation: Avoid direct sunlight and drink lots of water")
	}
	
	// Example 3: Password validator
	fmt.Println("\nPassword Validator:")
	userPassword := "Go2024!"
	minLength := 8
	maxLength := 20
	
	fmt.Printf("Password: %s\n", userPassword)
	passwordLength := len(userPassword)
	
	if passwordLength < minLength {
		fmt.Printf("Password too short (minimum %d characters)\n", minLength)
	} else if passwordLength > maxLength {
		fmt.Printf("Password too long (maximum %d characters)\n", maxLength)
	} else {
		fmt.Println("Valid password length")
		// Here you could add more validations (uppercase, numbers, etc.)
	}
	
	// Example 4: Quantity discount system
	fmt.Println("\nDiscount System:")
	quantity := 15
	pricePerUnit := 10.0
	subtotal := float64(quantity) * pricePerUnit
	
	fmt.Printf("Quantity: %d units\n", quantity)
	fmt.Printf("Price per unit: $%.2f\n", pricePerUnit)
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	
	var discountPercent float64
	
	if quantity >= 50 {
		discountPercent = 20.0
	} else if quantity >= 20 {
		discountPercent = 15.0
	} else if quantity >= 10 {
		discountPercent = 10.0
	} else if quantity >= 5 {
		discountPercent = 5.0
	} else {
		discountPercent = 0.0
	}
	
	if discountPercent > 0 {
		discountAmount := subtotal * (discountPercent / 100)
		finalTotal := subtotal - discountAmount
		fmt.Printf("Discount applied: %.0f%%\n", discountPercent)
		fmt.Printf("Savings: $%.2f\n", discountAmount)
		fmt.Printf("Final total: $%.2f\n", finalTotal)
	} else {
		fmt.Println("No discount available")
		fmt.Printf("Total: $%.2f\n", subtotal)
	}
	
	// Example 5: Weekday checker
	fmt.Println("\nDay Checker:")
	dayNumber := 3 // 1=Monday, 2=Tuesday, ..., 7=Sunday
	
	if dayNumber >= 1 && dayNumber <= 5 {
		fmt.Printf("Day %d: Weekday\n", dayNumber)
	} else if dayNumber == 6 || dayNumber == 7 {
		fmt.Printf("Day %d: Weekend\n", dayNumber)
	} else {
		fmt.Printf("Day %d: Invalid day number\n", dayNumber)
	}
	
	// Example 6: Loan eligibility
	fmt.Println("\nLoan Eligibility:")
	income := 3500.0
	creditScore := 720
	hasDebt := false
	yearsEmployed := 3
	
	fmt.Printf("Monthly income: $%.2f\n", income)
	fmt.Printf("Credit score: %d\n", creditScore)
	fmt.Printf("Has debt: %t\n", hasDebt)
	fmt.Printf("Years employed: %d\n", yearsEmployed)
	
	if income < 2000 {
		fmt.Println("Loan denied: Insufficient income")
	} else if creditScore < 650 {
		fmt.Println("Loan denied: Low credit score")
	} else if hasDebt && income < 3000 {
		fmt.Println("Loan denied: Existing debt and moderate income")
	} else if yearsEmployed < 2 {
		fmt.Println("Conditional loan: Requires guarantor (limited employment history)")
	} else {
		if creditScore >= 750 && income >= 4000 {
			fmt.Println("Loan approved: Preferred terms")
			fmt.Println("  Interest rate: 5%")
		} else if creditScore >= 700 && income >= 3000 {
			fmt.Println("Loan approved: Standard terms")
			fmt.Println("  Interest rate: 7%")
		} else {
			fmt.Println("Loan approved: Basic terms")
			fmt.Println("  Interest rate: 9%")
		}
	}
	
	fmt.Println("\n=== End of Examples ===")
}
