package main

import "fmt"

func main() {
	// Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// unint uint8 uint16 uint32 uintptr
	// byte
	// float32 float64
	// complex64 complex 128

	//Using var
	var age int32 = 24
	var isCool = true
	isCool = false

	// Shorthand
	//name := "Matt"
	//size := 2.3

	name, email := "Matt", "matt@email.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", isCool)
}
