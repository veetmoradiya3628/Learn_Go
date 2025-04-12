package main

import "fmt"

// new() - allocates memory for a variable and returns a pointer to it
// make() - allocates memory for slices, maps, and channels and returns a reference to them
// & - returns the address of a variable
// * - dereferences a pointer, giving access to the value it points to
// nil - represents a zero value for pointers, indicating that they do not point to any valid memory location
// unsafe.Pointer - a special type that can hold any pointer type, but requires conversion to use safely
// uintptr - an integer type that can hold the value of a pointer, but cannot be dereferenced directly
// reflect - a package that provides runtime reflection, allowing you to inspect types and values at runtime
// unsafe - a package that provides low-level memory manipulation functions, allowing you to bypass Go's type safety

func main() {
	var ptr *int
	fmt.Println("Pointer:", ptr)          // nil
	fmt.Println("Pointer address:", &ptr) // address of the pointer
	// fmt.Println("Pointer value:", *ptr)   // dereferencing a nil pointer will cause a runtime panic

	ptr = new(int)                             // allocate memory for an int and assign its address to ptr
	fmt.Println("Pointer value:", *ptr)        // dereferencing a pointer gives the value it points to
	*ptr = 42                                  // assign a value to the memory location pointed to by ptr
	fmt.Println("Pointer value:", *ptr)        // dereferencing a pointer gives the value it points to
	fmt.Println("Pointer address:", &ptr)      // address of the pointer
	fmt.Println("Pointer value address:", ptr) // address of the value pointed to by ptr

	var ptr1 = &ptr
	fmt.Println("Pointer1 value:", *ptr1)        // dereferencing a pointer gives the value it points to
	fmt.Println("Pointer1 address:", &ptr1)      // address of the pointer
	fmt.Println("Pointer1 value address:", ptr1) // address of the value pointed to by ptr1
	fmt.Println(**ptr1)
	**ptr1 = 100
	fmt.Println("Pointer value:", *ptr)   // dereferencing a pointer gives the value it points to
	fmt.Println("Pointer1 value:", *ptr1) // dereferencing a pointer gives the value it points to

	// Shallow copy of a pointer
	shallowCopy := ptr
	fmt.Println("Shallow Copy value:", *shallowCopy)  // dereferencing the shallow copy gives the value it points to
	fmt.Println("Shallow Copy address:", shallowCopy) // address of the value pointed to by shallow copy
	fmt.Println("Original Pointer value:", *ptr)      // original pointer value remains the same
	fmt.Println("Original Pointer address:", ptr)     // address of the value pointed to by original pointer

	*shallowCopy = 200                                        // modifying the value through shallow copy
	fmt.Println("Modified Shallow Copy value:", *shallowCopy) // value is updated
	fmt.Println("Original Pointer value:", *ptr)              // original pointer value is also updated

	// Deep copy simulation
	originalValue := *ptr      // get the value pointed to by ptr
	deepCopy := &originalValue // create a new pointer to a new memory location with the same value

	fmt.Println("Deep Copy value:", *deepCopy)    // dereferencing the deep copy gives the value it points to
	fmt.Println("Deep Copy address:", deepCopy)   // address of the value pointed to by deep copy
	fmt.Println("Original Pointer value:", *ptr)  // original pointer value remains the same
	fmt.Println("Original Pointer address:", ptr) // address of the value pointed to by original pointer

	*deepCopy = 300                                     // modifying the value through deep copy
	fmt.Println("Modified Deep Copy value:", *deepCopy) // value is updated
	fmt.Println("Original Pointer value:", *ptr)        // original pointer value is not affected
}
