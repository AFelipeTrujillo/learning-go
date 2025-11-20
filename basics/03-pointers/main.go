package main

import "fmt"

func main() {
	fmt.Println("=== Pointers in Go ===\n")

	// ============================================
	// 1. UNDERSTANDING MEMORY ADDRESSES
	// ============================================
	fmt.Println("--- 1. Memory Addresses ---")
	
	// Variables are stored in memory at specific addresses
	age := 25
	name := "Alice"
	
	// Use & operator to get the address of a variable
	fmt.Printf("age value: %d\n", age)
	fmt.Printf("age address: %p\n", &age)
	fmt.Printf("name value: %s\n", name)
	fmt.Printf("name address: %p\n", &name)
	
	fmt.Println()

	// ============================================
	// 2. DECLARING POINTERS
	// ============================================
	fmt.Println("--- 2. Declaring Pointers ---")
	
	// A pointer stores a memory address
	var x int = 42
	var ptr *int  // Pointer to an int (currently nil)
	
	ptr = &x  // Store the address of x in ptr
	
	fmt.Printf("x value: %d\n", x)
	fmt.Printf("x address: %p\n", &x)
	fmt.Printf("ptr value (address it stores): %p\n", ptr)
	fmt.Printf("ptr address: %p\n", &ptr)
	
	fmt.Println()

	// ============================================
	// 3. DEREFERENCING POINTERS
	// ============================================
	fmt.Println("--- 3. Dereferencing Pointers ---")
	
	// Use * operator to access the value at the pointer's address
	number := 100
	numPtr := &number
	
	fmt.Printf("number: %d\n", number)
	fmt.Printf("numPtr (address): %p\n", numPtr)
	fmt.Printf("*numPtr (value at address): %d\n", *numPtr)
	
	// Modify value through pointer (dereferencing)
	*numPtr = 200
	fmt.Printf("After *numPtr = 200:\n")
	fmt.Printf("number: %d\n", number)
	fmt.Printf("*numPtr: %d\n", *numPtr)
	
	fmt.Println()

	// ============================================
	// 4. POINTER TYPES
	// ============================================
	fmt.Println("--- 4. Pointer Types ---")
	
	// Pointers are type-specific
	var intPtr *int
	var floatPtr *float64
	var stringPtr *string
	var boolPtr *bool
	
	i := 42
	f := 3.14
	s := "Hello"
	boolVal := true
	
	intPtr = &i
	floatPtr = &f
	stringPtr = &s
	boolPtr = &boolVal
	
	fmt.Printf("*intPtr: %d (type: %T)\n", *intPtr, intPtr)
	fmt.Printf("*floatPtr: %.2f (type: %T)\n", *floatPtr, floatPtr)
	fmt.Printf("*stringPtr: %s (type: %T)\n", *stringPtr, stringPtr)
	fmt.Printf("*boolPtr: %t (type: %T)\n", *boolPtr, boolPtr)
	
	fmt.Println()

	// ============================================
	// 5. NIL POINTERS
	// ============================================
	fmt.Println("--- 5. Nil Pointers ---")
	
	// A pointer with no assigned address is nil
	var p *int
	fmt.Printf("Nil pointer: %v\n", p)
	fmt.Printf("Is nil? %t\n", p == nil)
	
	// Dereferencing nil pointer causes panic (don't do this!)
	// fmt.Println(*p)  // This would crash the program
	
	// Check before dereferencing
	if p != nil {
		fmt.Printf("Value: %d\n", *p)
	} else {
		fmt.Println("Pointer is nil, cannot dereference")
	}
	
	fmt.Println()

	// ============================================
	// 6. PASS BY VALUE vs PASS BY REFERENCE
	// ============================================
	fmt.Println("--- 6. Pass by Value vs Pass by Reference ---")
	
	// Go is pass-by-value by default
	fmt.Println("Pass by value:")
	val := 10
	fmt.Printf("Before: val = %d\n", val)
	modifyValue(val)
	fmt.Printf("After: val = %d (unchanged)\n", val)
	
	// Pass by reference using pointers
	fmt.Println("\nPass by reference (using pointer):")
	val2 := 10
	fmt.Printf("Before: val2 = %d\n", val2)
	modifyValueWithPointer(&val2)
	fmt.Printf("After: val2 = %d (changed!)\n", val2)
	
	fmt.Println()

	// ============================================
	// 7. POINTERS WITH STRUCTS
	// ============================================
	fmt.Println("--- 7. Pointers with Structs ---")
	
	// Create instance (Person type defined at package level)
	person1 := Person{Name: "Alice", Age: 25}
	fmt.Printf("person1: %+v\n", person1)
	
	// Pointer to struct
	personPtr := &person1
	fmt.Printf("personPtr: %p\n", personPtr)
	
	// Access fields through pointer (Go automatically dereferences)
	fmt.Printf("personPtr.Name: %s\n", personPtr.Name)
	fmt.Printf("(*personPtr).Name: %s\n", (*personPtr).Name)
	
	// Modify through pointer
	personPtr.Age = 26
	fmt.Printf("After modification: %+v\n", person1)
	
	fmt.Println()

	// ============================================
	// 8. NEW FUNCTION
	// ============================================
	fmt.Println("--- 8. New Function ---")
	
	// new() allocates memory and returns a pointer
	ptr1 := new(int)
	fmt.Printf("ptr1 (address): %p\n", ptr1)
	fmt.Printf("*ptr1 (value): %d\n", *ptr1)  // Zero value
	
	*ptr1 = 100
	fmt.Printf("After assignment: *ptr1 = %d\n", *ptr1)
	
	// new with struct
	person2 := new(Person)
	person2.Name = "Bob"
	person2.Age = 30
	fmt.Printf("person2: %+v\n", *person2)
	
	fmt.Println()

	// ============================================
	// 9. POINTER RECEIVERS IN METHODS
	// ============================================
	fmt.Println("--- 9. Pointer Receivers in Methods ---")
	
	// Value receiver (doesn't modify original)
	p1 := Person{Name: "Charlie", Age: 35}
	fmt.Printf("Before valueMethod: %+v\n", p1)
	p1.valueMethod()
	fmt.Printf("After valueMethod: %+v (unchanged)\n", p1)
	
	// Pointer receiver (modifies original)
	fmt.Printf("\nBefore pointerMethod: %+v\n", p1)
	p1.pointerMethod()
	fmt.Printf("After pointerMethod: %+v (changed!)\n", p1)
	
	fmt.Println()

	// ============================================
	// 10. POINTER ARITHMETIC (NOT ALLOWED)
	// ============================================
	fmt.Println("--- 10. Pointer Arithmetic ---")
	
	// Unlike C/C++, Go does NOT allow pointer arithmetic
	arr := [3]int{10, 20, 30}
	arrPtr := &arr[0]
	fmt.Printf("First element: %d\n", *arrPtr)
	
	// This is NOT allowed in Go:
	// arrPtr++  // Compile error!
	// arrPtr = arrPtr + 1  // Compile error!
	
	fmt.Println("Note: Go does not support pointer arithmetic for safety")
	
	fmt.Println()

	// ============================================
	// 11. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("--- 11. Practical Examples ---")
	
	// Example 1: Swap function
	fmt.Println("\nExample 1: Swap Function")
	a, b := 5, 10
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)
	
	// Example 2: Update struct
	fmt.Println("\nExample 2: Update Person")
	person := Person{Name: "Diana", Age: 28}
	fmt.Printf("Before: %+v\n", person)
	updatePerson(&person, "Diana Smith", 29)
	fmt.Printf("After: %+v\n", person)
	
	// Example 3: Optional values
	fmt.Println("\nExample 3: Optional Values with Pointers")
	email1 := "alice@example.com"
	user1 := User{Name: "Alice", Email: &email1}
	user2 := User{Name: "Bob", Email: nil}
	
	printUser(user1)
	printUser(user2)
	
	// Example 4: Large struct efficiency
	fmt.Println("\nExample 4: Large Struct Efficiency")
	largeData := LargeStruct{Data: [1000]int{}}
	largeData.Data[0] = 42
	
	// Inefficient: copies entire struct
	// processLargeStruct(largeData)
	
	// Efficient: passes pointer
	processLargeStructPtr(&largeData)
	fmt.Printf("First element: %d\n", largeData.Data[0])
	
	// Example 5: Factory function
	fmt.Println("\nExample 5: Factory Function")
	newPerson := createPerson("Eve", 32)
	fmt.Printf("New person: %+v\n", *newPerson)
	
	// Example 6: Linked list node
	fmt.Println("\nExample 6: Linked List")
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	
	node1.Next = node2
	node2.Next = node3
	
	printList(node1)
	
	fmt.Println("\n=== End of Examples ===")
}

// ============================================
// TYPES (defined at package level)
// ============================================

// Person struct
type Person struct {
	Name string
	Age  int
}

// User with optional email
type User struct {
	Name  string
	Email *string  // Optional field
}

// Large struct example
type LargeStruct struct {
	Data [1000]int
}

// Linked list node
type Node struct {
	Value int
	Next  *Node
}

// ============================================
// HELPER FUNCTIONS
// ============================================

// Pass by value (doesn't modify original)
func modifyValue(n int) {
	n = 99
	fmt.Printf("Inside modifyValue: n = %d\n", n)
}

// Pass by reference (modifies original)
func modifyValueWithPointer(n *int) {
	*n = 99
	fmt.Printf("Inside modifyValueWithPointer: *n = %d\n", *n)
}

// Value receiver - doesn't modify original
func (p Person) valueMethod() {
	p.Age = 100
	fmt.Printf("Inside valueMethod: %+v\n", p)
}

// Pointer receiver - modifies original
func (p *Person) pointerMethod() {
	p.Age = 100
	fmt.Printf("Inside pointerMethod: %+v\n", *p)
}

// Swap function using pointers
func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

// Update person
func updatePerson(p *Person, name string, age int) {
	p.Name = name
	p.Age = age
}

func printUser(u User) {
	fmt.Printf("User: %s", u.Name)
	if u.Email != nil {
		fmt.Printf(", Email: %s\n", *u.Email)
	} else {
		fmt.Println(", Email: not provided")
	}
}

func processLargeStructPtr(ls *LargeStruct) {
	ls.Data[0] = 99
}

// Factory function
func createPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func printList(head *Node) {
	fmt.Print("List: ")
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

// ============================================
// COMPARISON FUNCTIONS
// ============================================

// Compare pass by value vs pass by reference
func demonstratePassingMechanisms() {
	fmt.Println("=== Pass by Value vs Reference ===")
	
	// Array (copied)
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Printf("Array after function: %v\n", arr)
	
	// Slice (reference to underlying array)
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Printf("Slice after function: %v\n", slice)
	
	// Pointer (reference)
	val := 10
	modifyWithPointer(&val)
	fmt.Printf("Value after function: %d\n", val)
}

func modifyArray(arr [3]int) {
	arr[0] = 99
}

func modifySlice(s []int) {
	s[0] = 99
}

func modifyWithPointer(p *int) {
	*p = 99
}

// ============================================
// BEST PRACTICES EXAMPLES
// ============================================

// Good: Return pointer for large struct
func newLargeStruct() *LargeStruct {
	return &LargeStruct{
		Data: [1000]int{},
	}
}

// Good: Use pointer receiver for methods that modify
func (p *Person) incrementAge() {
	p.Age++
}

// Good: Check for nil before dereferencing
func safeAccess(p *int) int {
	if p != nil {
		return *p
	}
	return 0  // Default value
}

// Good: Return pointer for optional values
func findPerson(name string) *Person {
	// If found
	if name == "Alice" {
		return &Person{Name: name, Age: 25}
	}
	// If not found
	return nil
}

