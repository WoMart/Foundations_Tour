package packages

import (
	"fmt"
	"math"
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
