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
	"fmt"
	"fountations/tour/mascot"

	// To download a package, use go get <Name> in terminal
	// Go package library: https://pkg.go.dev/
	"rsc.io/quote"
)

func main() {
	ProjectSetup()
}

func ProjectSetup() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Glass())
}
