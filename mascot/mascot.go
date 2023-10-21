package mascot

// Following a youtube video on go project setup in VS Code
// https://youtu.be/1MXIGYrMk80?si=hCSslaK3xmt-C4HQ

import (
	"fmt"

	"rsc.io/quote"
)

// Change the string to fail the associated text
func BestMascot() string {
	return "Go Gopher"
}

func Run() {
	fmt.Println(BestMascot())
	fmt.Println(quote.Glass())
}
