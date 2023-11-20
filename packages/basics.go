package packages

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// Uppercase Main is exported
func Main() {
	// https://go.dev/tour/basics/1
	fmt.Println("Random number: ", rand.Intn(10))

	// https://go.dev/tour/basics/2
	fmt.Printf("Its square root of 7 is %g.\n", math.Sqrt(7))

	main()
}

// Lowercase main is not exported (it's internal)
func main() {
	fmt.Println(
		"\nExported names need to begin with uppercase letter.")
	fmt.Println(
		"Package's Main function is accessible to packages importing it, but main isn't.")
	fmt.Println(
		"Try it yourself!")
	fmt.Println()

	fmt.Println("Check out the code for samples of different ways to define functions.")
}

// Here's a function with parameters
// Note the ordering of param name and type
// The return type is declared after the function signature
func add(x int, y int) int {
	return x + y
}

// Consecutive parameters of the same type can share one type definition
// This should be at the end of the list of parameters
func subtract(x, y int) int {
	return x - y
}

func sampleWithMultipleParamLists(x, y int, s1, s2 string) {
	fmt.Println(x + y)
	fmt.Println(s1 + s2)
}

// This is a non-returning function, equivalent to void
// Void functions are simply defined without a type
func text(str string) {
	fmt.Println(str)
}

// Return several values using the (tuple) syntax
func swap(x, y string) (string, string) {
	return y, x
}

// Return values can be given names in the function signature
// and referenced without var or type definition inside the function body
func swap2(x, y int) (a, b int) {
	a = y
	b = x

	// This is a "Naked return", but `return a, b` would also work
	// Naked returns should only be used for short and simple functions (readability)
	return
}

// VARIABLES
// I've put some print statements so no variable is unused to satisfy the compiler

// Package level variables
var packageLevelVar string = "sample"
var isAccessibleInFunctions, isGlobalWithinPackage bool

// I swear, the golang code cleanup aligns these in columns, I don't like it...
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func variables() {
	fmt.Println(
		packageLevelVar,
		isAccessibleInFunctions,
		isGlobalWithinPackage)

	// Storngly typed

	var stronglyTypedString string
	stronglyTypedString = "Try to replace me with a non-string type value"
	fmt.Println(stronglyTypedString)

	// Implicitly typed using `var`
	// This can be done at the package level as well (see above)

	var implicitBool, implicitInt, implicitString = true, 2137, "fun"
	fmt.Println(implicitBool, implicitInt, implicitString)

	// Implicitly typed using `:=` ("short assignment")
	// This cannot be done at the package level

	implicitBool2, implicitInt2 := true, 2005
	fmt.Println(implicitBool2, implicitInt2)

	// Types (this I just copied from Golang Tour, no reason to rewrite)

	/*
		Some primitive types (copied from golang tour):
			bool

			string

			int  int8  int16  int32  int64
			uint uint8 uint16 uint32 uint64 uintptr

			byte // alias for uint8

			rune // alias for int32
				// represents a Unicode code point

			float32 float64

			complex64 complex128
	*/

	// Formats name of type and set value in string
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
