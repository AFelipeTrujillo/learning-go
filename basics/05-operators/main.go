package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Operators in Go ===\n")

	// ============================================
	// 1. ARITHMETIC OPERATORS
	// ============================================
	fmt.Println("--- 1. Arithmetic Operators ---")
	
	a := 20
	b := 7
	fmt.Printf("a = %d, b = %d\n", a, b)
	
	// Addition
	sum := a + b
	fmt.Printf("Addition: %d + %d = %d\n", a, b, sum)
	
	// Subtraction
	difference := a - b
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, difference)
	
	// Multiplication
	product := a * b
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, product)
	
	// Division
	quotient := a / b
	fmt.Printf("Division: %d / %d = %d (integer division)\n", a, b, quotient)
	
	// Float division
	floatQuotient := float64(a) / float64(b)
	fmt.Printf("Division (float): %.2f / %.2f = %.4f\n", float64(a), float64(b), floatQuotient)
	
	// Modulus (remainder)
	remainder := a % b
	fmt.Printf("Modulus: %d %% %d = %d\n", a, b, remainder)
	fmt.Println()

	// ============================================
	// 2. COMPARISON OPERATORS (CONDITIONAL)
	// ============================================
	fmt.Println("--- 2. Comparison Operators ---")
	
	x := 10
	y := 20
	z := 10
	fmt.Printf("x = %d, y = %d, z = %d\n\n", x, y, z)
	
	// Equal to (==)
	fmt.Printf("x == y: %t (is %d equal to %d?)\n", x == y, x, y)
	fmt.Printf("x == z: %t (is %d equal to %d?)\n", x == z, x, z)
	
	// Not equal to (!=)
	fmt.Printf("x != y: %t (is %d not equal to %d?)\n", x != y, x, y)
	fmt.Printf("x != z: %t (is %d not equal to %d?)\n", x != z, x, z)
	
	// Less than (<)
	fmt.Printf("x < y: %t (is %d less than %d?)\n", x < y, x, y)
	fmt.Printf("y < x: %t (is %d less than %d?)\n", y < x, y, x)
	
	// Greater than (>)
	fmt.Printf("x > y: %t (is %d greater than %d?)\n", x > y, x, y)
	fmt.Printf("y > x: %t (is %d greater than %d?)\n", y > x, y, x)
	
	// Less than or equal to (<=)
	fmt.Printf("x <= y: %t (is %d less than or equal to %d?)\n", x <= y, x, y)
	fmt.Printf("x <= z: %t (is %d less than or equal to %d?)\n", x <= z, x, z)
	
	// Greater than or equal to (>=)
	fmt.Printf("y >= x: %t (is %d greater than or equal to %d?)\n", y >= x, y, x)
	fmt.Printf("x >= z: %t (is %d greater than or equal to %d?)\n", x >= z, x, z)
	fmt.Println()

	// ============================================
	// 3. LOGICAL OPERATORS
	// ============================================
	fmt.Println("--- 3. Logical Operators ---")
	
	isAdult := true
	hasLicense := true
	hasInsurance := false
	
	fmt.Printf("isAdult = %t\n", isAdult)
	fmt.Printf("hasLicense = %t\n", hasLicense)
	fmt.Printf("hasInsurance = %t\n\n", hasInsurance)
	
	// AND (&&) - All conditions must be true
	canDrive := isAdult && hasLicense
	fmt.Printf("Can drive? (isAdult AND hasLicense)\n")
	fmt.Printf("  %t && %t = %t\n", isAdult, hasLicense, canDrive)
	
	canDriveLegally := isAdult && hasLicense && hasInsurance
	fmt.Printf("Can drive legally? (all conditions)\n")
	fmt.Printf("  %t && %t && %t = %t\n", isAdult, hasLicense, hasInsurance, canDriveLegally)
	
	// OR (||) - At least one condition must be true
	needsDocuments := !hasLicense || !hasInsurance
	fmt.Printf("\nNeeds documents? (NOT hasLicense OR NOT hasInsurance)\n")
	fmt.Printf("  !%t || !%t = %t\n", hasLicense, hasInsurance, needsDocuments)
	
	// NOT (!) - Negates the condition
	isMinor := !isAdult
	fmt.Printf("\nIs minor? (NOT isAdult)\n")
	fmt.Printf("  !%t = %t\n", isAdult, isMinor)
	
	// Complex logical expressions
	fmt.Println("\nComplex expressions:")
	age := 25
	score := 85
	canEnter := (age >= 18 && age <= 65) || score > 90
	fmt.Printf("Can enter? ((age >= 18 AND age <= 65) OR score > 90)\n")
	fmt.Printf("  ((%d >= 18 && %d <= 65) || %d > 90) = %t\n", age, age, score, canEnter)
	fmt.Println()

	// ============================================
	// 4. ASSIGNMENT OPERATORS
	// ============================================
	fmt.Println("--- 4. Assignment Operators ---")
	
	num := 10
	fmt.Printf("Initial value: num = %d\n", num)
	
	// Simple assignment
	num = 15
	fmt.Printf("After assignment (num = 15): %d\n", num)
	
	// Add and assign (+=)
	num += 5 // equivalent to: num = num + 5
	fmt.Printf("After += 5: %d\n", num)
	
	// Subtract and assign (-=)
	num -= 3 // equivalent to: num = num - 3
	fmt.Printf("After -= 3: %d\n", num)
	
	// Multiply and assign (*=)
	num *= 2 // equivalent to: num = num * 2
	fmt.Printf("After *= 2: %d\n", num)
	
	// Divide and assign (/=)
	num /= 4 // equivalent to: num = num / 4
	fmt.Printf("After /= 4: %d\n", num)
	
	// Modulus and assign (%=)
	num %= 3 // equivalent to: num = num % 3
	fmt.Printf("After %%= 3: %d\n", num)
	fmt.Println()

	// ============================================
	// 5. INCREMENT AND DECREMENT
	// ============================================
	fmt.Println("--- 5. Increment and Decrement ---")
	
	counter := 5
	fmt.Printf("Initial counter: %d\n", counter)
	
	// Increment (++)
	counter++  // equivalent to: counter = counter + 1
	fmt.Printf("After counter++: %d\n", counter)
	
	// Decrement (--)
	counter--  // equivalent to: counter = counter - 1
	fmt.Printf("After counter--: %d\n", counter)
	
	// Note: Go only has postfix increment/decrement (no ++counter or --counter)
	// Also, ++ and -- are statements, not expressions
	// So you can't do: x = counter++  (this is invalid in Go)
	
	fmt.Println("\nLoop example with increment:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Iteration %d\n", i)
	}
	fmt.Println()

	// ============================================
	// 6. BITWISE OPERATORS
	// ============================================
	fmt.Println("--- 6. Bitwise Operators ---")
	
	p := 12  // binary: 1100
	q := 10  // binary: 1010
	
	fmt.Printf("p = %d (binary: %04b)\n", p, p)
	fmt.Printf("q = %d (binary: %04b)\n\n", q, q)
	
	// Bitwise AND (&)
	bitwiseAnd := p & q
	fmt.Printf("p & q  = %d (binary: %04b) - AND\n", bitwiseAnd, bitwiseAnd)
	
	// Bitwise OR (|)
	bitwiseOr := p | q
	fmt.Printf("p | q  = %d (binary: %04b) - OR\n", bitwiseOr, bitwiseOr)
	
	// Bitwise XOR (^)
	bitwiseXor := p ^ q
	fmt.Printf("p ^ q  = %d (binary: %04b) - XOR\n", bitwiseXor, bitwiseXor)
	
	// Bitwise NOT (^)
	bitwiseNot := ^p
	fmt.Printf("^p     = %d (binary: %064b) - NOT\n", bitwiseNot, uint64(bitwiseNot))
	
	// Left shift (<<)
	leftShift := p << 2
	fmt.Printf("p << 2 = %d (binary: %08b) - Left shift by 2\n", leftShift, leftShift)
	
	// Right shift (>>)
	rightShift := p >> 2
	fmt.Printf("p >> 2 = %d (binary: %04b) - Right shift by 2\n", rightShift, rightShift)
	fmt.Println()

	// ============================================
	// 7. OPERATOR PRECEDENCE
	// ============================================
	fmt.Println("--- 7. Operator Precedence ---")
	
	// Precedence from highest to lowest:
	// 1. () parentheses
	// 2. *, /, %, <<, >>, &, &^
	// 3. +, -, |, ^
	// 4. ==, !=, <, <=, >, >=
	// 5. &&
	// 6. ||
	
	result1 := 2 + 3 * 4
	fmt.Printf("2 + 3 * 4 = %d (multiplication first)\n", result1)
	
	result2 := (2 + 3) * 4
	fmt.Printf("(2 + 3) * 4 = %d (parentheses first)\n", result2)
	
	result3 := 10 > 5 && 3 < 7
	fmt.Printf("10 > 5 && 3 < 7 = %t (comparison, then AND)\n", result3)
	
	result4 := 5 + 3 > 7 && 10 - 2 < 9
	fmt.Printf("5 + 3 > 7 && 10 - 2 < 9 = %t (arithmetic, comparison, then AND)\n", result4)
	
	// Use parentheses for clarity
	result5 := (age >= 18) && (score >= 60)
	fmt.Printf("(age >= 18) && (score >= 60) = %t (clearer with parentheses)\n", result5)
	fmt.Println()

	// ============================================
	// 8. STRING OPERATORS
	// ============================================
	fmt.Println("--- 8. String Operators ---")
	
	firstName := "John"
	lastName := "Doe"
	
	// String concatenation with +
	fullName := firstName + " " + lastName
	fmt.Printf("firstName + \" \" + lastName = \"%s\"\n", fullName)
	
	// String comparison
	str1 := "apple"
	str2 := "banana"
	str3 := "apple"
	
	fmt.Printf("\nString comparisons:\n")
	fmt.Printf("\"%s\" == \"%s\": %t\n", str1, str2, str1 == str2)
	fmt.Printf("\"%s\" == \"%s\": %t\n", str1, str3, str1 == str3)
	fmt.Printf("\"%s\" < \"%s\": %t (alphabetical order)\n", str1, str2, str1 < str2)
	fmt.Printf("\"%s\" > \"%s\": %t\n", str2, str1, str2 > str1)
	fmt.Println()

	// ============================================
	// 9. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 9. Practical Examples ---")
	
	// Example 1: Checking even/odd
	fmt.Println("\nExample 1: Even/Odd Checker")
	number := 42
	if number%2 == 0 {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}
	
	// Example 2: Range checking
	fmt.Println("\nExample 2: Range Checker")
	temperature := 22
	if temperature >= 18 && temperature <= 25 {
		fmt.Printf("Temperature %d°C is in comfortable range (18-25°C)\n", temperature)
	} else {
		fmt.Printf("Temperature %d°C is outside comfortable range\n", temperature)
	}
	
	// Example 3: Leap year calculator
	fmt.Println("\nExample 3: Leap Year Calculator")
	year := 2024
	isLeapYear := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	if isLeapYear {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}
	
	// Example 4: Discount calculator
	fmt.Println("\nExample 4: Discount Calculator")
	originalPrice := 100.0
	quantity := 5
	var discountRate float64
	
	if quantity >= 10 {
		discountRate = 0.20
	} else if quantity >= 5 {
		discountRate = 0.10
	} else {
		discountRate = 0.0
	}
	
	discount := originalPrice * discountRate
	finalPrice := originalPrice - discount
	
	fmt.Printf("Original price: $%.2f\n", originalPrice)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Discount: %.0f%%\n", discountRate*100)
	fmt.Printf("You save: $%.2f\n", discount)
	fmt.Printf("Final price: $%.2f\n", finalPrice)
	
	// Example 5: Password strength checker
	fmt.Println("\nExample 5: Password Strength Checker")
	password := "SecurePass123"
	passwordLength := len(password)
	
	hasMinLength := passwordLength >= 8
	hasMaxLength := passwordLength <= 20
	
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Length: %d characters\n", passwordLength)
	fmt.Printf("Has minimum length (>=8): %t\n", hasMinLength)
	fmt.Printf("Has maximum length (<=20): %t\n", hasMaxLength)
	
	isValidLength := hasMinLength && hasMaxLength
	if isValidLength {
		fmt.Println("Password length is valid")
	} else {
		fmt.Println("Password length is invalid")
	}
	
	// Example 6: Grade boundary checker
	fmt.Println("\nExample 6: Grade Boundary Checker")
	studentScore := 75
	passingScore := 60
	excellenceScore := 90
	
	isPassing := studentScore >= passingScore
	isExcellent := studentScore >= excellenceScore
	needsImprovement := studentScore < passingScore
	
	fmt.Printf("Student score: %d\n", studentScore)
	fmt.Printf("Is passing (>=%d): %t\n", passingScore, isPassing)
	fmt.Printf("Is excellent (>=%d): %t\n", excellenceScore, isExcellent)
	fmt.Printf("Needs improvement (<%d): %t\n", passingScore, needsImprovement)
	
	// Example 7: BMI Calculator
	fmt.Println("\nExample 7: BMI Calculator")
	weight := 70.0  // kg
	height := 1.75  // meters
	
	bmi := weight / math.Pow(height, 2)
	fmt.Printf("Weight: %.1f kg\n", weight)
	fmt.Printf("Height: %.2f m\n", height)
	fmt.Printf("BMI: %.2f\n", bmi)
	
	if bmi < 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("Category: Normal weight")
	} else if bmi >= 25 && bmi < 30 {
		fmt.Println("Category: Overweight")
	} else {
		fmt.Println("Category: Obese")
	}
	
	// Example 8: Time calculator
	fmt.Println("\nExample 8: Time Calculator")
	totalSeconds := 3665
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	
	fmt.Printf("Total seconds: %d\n", totalSeconds)
	fmt.Printf("Converted: %d hours, %d minutes, %d seconds\n", hours, minutes, seconds)
	
	fmt.Println("\n=== End of Examples ===")
}

