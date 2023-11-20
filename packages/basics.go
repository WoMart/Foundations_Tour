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
}

// Here's a function with parameters
// Note the ordering of param name and type
// The return type is declared after the function signature
func add(x int, y int) int {
	return x + y
}

// And here is a non-returning function, equivalent to void
// Void functions are simply defined without a type
func text(str string) {
	fmt.Println(str)
}
