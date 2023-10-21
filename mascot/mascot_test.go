package mascot_test

import (
	"fountations/tour/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
