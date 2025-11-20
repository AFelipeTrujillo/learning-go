package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Defer in Go ===\n")

	// ============================================
	// 1. BASIC DEFER
	// ============================================
	fmt.Println("--- 1. Basic Defer ---")
	
	// Defer postpones execution until function returns
	basicDeferExample()
	
	fmt.Println()

	// ============================================
	// 2. DEFER EXECUTION ORDER (LIFO)
	// ============================================
	fmt.Println("--- 2. Defer Execution Order (LIFO) ---")
	
	// Multiple defers execute in reverse order (Last In, First Out)
	deferOrderExample()
	
	fmt.Println()

	// ============================================
	// 3. DEFER WITH PARAMETERS
	// ============================================
	fmt.Println("--- 3. Defer with Parameters ---")
	
	// Parameters are evaluated immediately, but execution is deferred
	deferParametersExample()
	
	fmt.Println()

	// ============================================
	// 4. DEFER WITH ANONYMOUS FUNCTIONS
	// ============================================
	fmt.Println("--- 4. Defer with Anonymous Functions ---")
	
	// Using closures to capture current values
	deferClosureExample()
	
	fmt.Println()

	// ============================================
	// 5. DEFER FOR RESOURCE CLEANUP
	// ============================================
	fmt.Println("--- 5. Defer for Resource Cleanup ---")
	
	// Common pattern: open resource, defer close
	fmt.Println("File operation example:")
	if err := writeToFile("test.txt", "Hello, Go!"); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("File written successfully")
	}
	
	fmt.Println()

	// ============================================
	// 6. DEFER WITH RETURN VALUES
	// ============================================
	fmt.Println("--- 6. Defer with Return Values ---")
	
	// Defer can modify named return values
	result := deferWithReturn()
	fmt.Printf("Result from deferWithReturn: %d\n", result)
	
	result2 := deferModifyingReturn()
	fmt.Printf("Result from deferModifyingReturn: %d\n", result2)
	
	fmt.Println()

	// ============================================
	// 7. DEFER IN LOOPS
	// ============================================
	fmt.Println("--- 7. Defer in Loops ---")
	
	// Be careful: defer in loop accumulates
	fmt.Println("Defer inside loop (not recommended):")
	deferInLoopBad()
	
	fmt.Println("\nBetter approach - use function:")
	deferInLoopGood()
	
	fmt.Println()

	// ============================================
	// 8. DEFER FOR ERROR HANDLING
	// ============================================
	fmt.Println("--- 8. Defer for Error Handling ---")
	
	// Defer can check and handle errors
	if err := operationWithDefer(); err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	}
	
	fmt.Println()

	// ============================================
	// 9. DEFER WITH RECOVER
	// ============================================
	fmt.Println("--- 9. Defer with Recover ---")
	
	// Defer is used with recover to handle panics
	fmt.Println("Calling function that might panic:")
	safeDivide(10, 2)
	safeDivide(10, 0)
	fmt.Println("Program continues after panic recovery")
	
	fmt.Println()

	// ============================================
	// 10. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 10. Practical Examples ---")
	
	// Example 1: Timing function execution
	fmt.Println("\nExample 1: Function Timer")
	timedOperation()
	
	// Example 2: Lock/Unlock pattern
	fmt.Println("\nExample 2: Mutex Lock Pattern")
	demonstrateMutexPattern()
	
	// Example 3: Transaction rollback
	fmt.Println("\nExample 3: Transaction Pattern")
	if err := performTransaction(); err != nil {
		fmt.Printf("Transaction error: %v\n", err)
	}
	
	// Example 4: Logging function entry/exit
	fmt.Println("\nExample 4: Function Logging")
	calculateSum(5, 10)
	
	// Example 5: Cleanup multiple resources
	fmt.Println("\nExample 5: Multiple Resource Cleanup")
	if err := processMultipleFiles(); err != nil {
		fmt.Printf("Error processing files: %v\n", err)
	}
	
	// Example 6: Stack trace
	fmt.Println("\nExample 6: Stack Trace on Error")
	processWithStackTrace()
	
	fmt.Println("\n=== End of Examples ===")
	
	// Cleanup test file
	os.Remove("test.txt")
	os.Remove("file1.txt")
	os.Remove("file2.txt")
	os.Remove("output.txt")
}

// ============================================
// FUNCTION DEFINITIONS
// ============================================

// 1. Basic defer
func basicDeferExample() {
	fmt.Println("Start of function")
	
	defer fmt.Println("This executes last (deferred)")
	
	fmt.Println("Middle of function")
	fmt.Println("End of function")
	// Deferred function executes here, after "End of function"
}

// 2. Defer execution order (LIFO)
func deferOrderExample() {
	fmt.Println("Function starts")
	
	defer fmt.Println("Defer 1 (executes third)")
	defer fmt.Println("Defer 2 (executes second)")
	defer fmt.Println("Defer 3 (executes first)")
	
	fmt.Println("Function body")
	// Defers execute in reverse order: 3, 2, 1
}

// 3. Defer with parameters
func deferParametersExample() {
	x := 10
	
	// Parameters are evaluated immediately
	defer fmt.Printf("Deferred print: x = %d\n", x)
	
	x = 20
	fmt.Printf("Current value: x = %d\n", x)
	// Deferred function will print x = 10 (not 20)
}

// 4. Defer with closures
func deferClosureExample() {
	x := 10
	
	// Closure captures variable reference
	defer func() {
		fmt.Printf("Deferred closure: x = %d\n", x)
	}()
	
	x = 20
	fmt.Printf("Current value: x = %d\n", x)
	// Deferred closure will print x = 20 (current value)
}

// 5. Resource cleanup
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Ensures file is closed when function returns
	
	_, err = file.WriteString(content)
	return err
}

// 6. Defer with return values
func deferWithReturn() int {
	defer fmt.Println("Defer executes after return statement")
	
	fmt.Println("Before return")
	return 42
}

func deferModifyingReturn() (result int) {
	defer func() {
		result = result * 2 // Can modify named return value
	}()
	
	result = 21
	return // Returns 42 (21 * 2)
}

// 7. Defer in loops
func deferInLoopBad() {
	// Problem: all defers accumulate and execute at function end
	for i := 0; i < 3; i++ {
		defer fmt.Printf("Deferred: %d\n", i)
	}
	fmt.Println("Loop finished")
	// Prints: Loop finished, then: 2, 1, 0
}

func deferInLoopGood() {
	for i := 0; i < 3; i++ {
		func(n int) {
			defer fmt.Printf("Deferred: %d\n", n)
			fmt.Printf("Processing: %d\n", n)
		}(i)
	}
	fmt.Println("Loop finished")
}

// 8. Error handling with defer
func operationWithDefer() error {
	fmt.Println("Starting operation")
	
	defer func() {
		fmt.Println("Cleanup: operation finished")
	}()
	
	// Simulate some work
	fmt.Println("Performing work...")
	
	// Even if error occurs, defer will execute
	return nil
}

// 9. Recover from panic
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	result := a / b
	fmt.Printf("%d / %d = %d\n", a, b, result)
}

// 10. Practical examples

// Example 1: Timing function execution
func timedOperation() {
	defer trackTime("Operation")()
	
	// Simulate some work
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	fmt.Printf("Sum calculated: %d\n", sum)
}

func trackTime(name string) func() {
	fmt.Printf("%s started\n", name)
	return func() {
		fmt.Printf("%s finished\n", name)
	}
}

// Example 2: Mutex pattern (simulated)
func demonstrateMutexPattern() {
	fmt.Println("Acquiring lock...")
	defer fmt.Println("Releasing lock...")
	
	// Critical section
	fmt.Println("Performing critical operation")
	// Lock is automatically released via defer
}

// Example 3: Transaction pattern
func performTransaction() error {
	fmt.Println("Beginning transaction")
	
	committed := false
	defer func() {
		if !committed {
			fmt.Println("Rolling back transaction")
		}
	}()
	
	// Simulate transaction work
	fmt.Println("Executing transaction operations...")
	
	// If we reach here, commit
	committed = true
	fmt.Println("Committing transaction")
	return nil
}

// Example 4: Function logging
func calculateSum(a, b int) int {
	defer logExit("calculateSum")()
	logEntry("calculateSum", a, b)
	
	result := a + b
	fmt.Printf("Calculation: %d + %d = %d\n", a, b, result)
	return result
}

func logEntry(funcName string, args ...interface{}) {
	fmt.Printf("ENTER %s with args: %v\n", funcName, args)
}

func logExit(funcName string) func() {
	return func() {
		fmt.Printf("EXIT %s\n", funcName)
	}
}

// Example 5: Multiple resource cleanup
func processMultipleFiles() error {
	// Create first file
	file1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}
	defer file1.Close()
	
	// Create second file
	file2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}
	defer file2.Close()
	
	// Process files
	fmt.Println("Processing file1...")
	file1.WriteString("Content 1")
	
	fmt.Println("Processing file2...")
	file2.WriteString("Content 2")
	
	fmt.Println("Both files processed successfully")
	return nil
}

// Example 6: Stack trace
func processWithStackTrace() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occurred: %v\n", r)
			fmt.Println("Stack trace available here")
		}
	}()
	
	// Simulate some operations
	fmt.Println("Processing data...")
	// No panic in this example
	fmt.Println("Processing completed")
}

// ============================================
// ADDITIONAL HELPER FUNCTIONS
// ============================================

// Demonstrate defer with file operations
func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	
	_, err = io.Copy(destination, source)
	return err
}

// Demonstrate defer with string builder
func buildMessage() string {
	var builder strings.Builder
	
	defer func() {
		fmt.Println("Message building completed")
	}()
	
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	
	return builder.String()
}

// Demonstrate defer with error recovery
func safeOperation() (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("operation panicked: %v", r)
		}
	}()
	
	// Potentially dangerous operation
	result = "Success"
	return
}

// Demonstrate defer order with cleanup
func cleanupExample() {
	fmt.Println("Allocating resources...")
	
	defer fmt.Println("1. Final cleanup")
	defer fmt.Println("2. Releasing resource C")
	defer fmt.Println("3. Releasing resource B")
	defer fmt.Println("4. Releasing resource A")
	
	fmt.Println("Using resources...")
	fmt.Println("Operation complete")
}

