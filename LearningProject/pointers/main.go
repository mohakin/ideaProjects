package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	//return
	*age = *age - 18
}

// ----------------------------------------------------------------
// Pointer are variables that store a value address instead of values.
// & = pointer.
// There are two main advantages of using pointers.
// 1. You can avoid unnecessary value copies.
// 2. You can directly mutate values.

// 1.
// By default, in go programs, Go creates a copy when passing values to funcs.
// For very large and complex values, this may take up too much memory space unnecessarily.
// With pointers only one value is stored in memory(and the address is passed around).

// 2.
// Pass a pointer(address) instead of a value to a function.
// The function can then directly edit the underlying value - no return value
// is required.
// This can to less code. -
// - But also to less understandable code or unexpected behaviour.
// ----------------------------------------------------------------
// Writing Code Without Pointers ----------------------------------
// ----------------------------------------------------------------
// Use the & pointer to point towards the address of a variable.
// Use * to point towards the value of a variable. Or derefrance it.
// I.E. age := 32, var agePointer *int, agePointer := &age,-
// - fmt.Println("Age:", *agePointer)
// Output == Age: 32 -- if I did not use the * it would have given me the -
// - address where it is stored instead.
// ----------------------------------------------------------------
// All values have a null value in Go. int = 0, float64 = 0.0, string = ""
// For a pointer it's nil.
// nil represents the absence of an address value.
// I.E. a pointer pointing to no address / no value in memory.
// ----------------------------------------------------------------
// Using Pointers & Passing Pointers to Functions -----------------
// ----------------------------------------------------------------
// You can not perform arithhmatics on a pointer in Go.
// ----------------------------------------------------------------
// Using Pointers for Data Mutations ------------------------------
// ----------------------------------------------------------------
//
