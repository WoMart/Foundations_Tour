// Main needs to exist in `package main`
// Try changing the package name and/or file name and see the error message
package main

/*
	Packages can with separate statements, like:
		import "fmt"
		import "rsc.io/quote"

	Or using a "factored" import statement with packages encapsulated in (parenthesis), like below
*/
import (
	"fountations/tour/mascot"
	"fountations/tour/packages"
)

func main() {
	Basics()
}

// Project Setup tutorial
func Mascot() {
	mascot.Run()
}

func Basics() {
	packages.Main()
}
