// Main needs to exist in `package main`
// Try changing the package name and/or file name and see the error message
package main

/*
	Packages can with separate statements, like:
		import "fmt"
		import "rsc.io/quote"

	Or using a single import statement and encapsulating the packages in (parenthesis), like below
*/
import (
	"fountations/tour/mascot"
)

func main() {
	Mascot()
}

// Project Setup tutorial
func Mascot() {
	mascot.Run()
}
