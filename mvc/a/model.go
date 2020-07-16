package a

// A is a database model
import (
	"fmt"
)

type A struct {
	I int
	P bool
}

// Custom is a custom operation specific to A
func (a *A) Custom() error {
	fmt.Println("\tA's Custom method called")

	return nil
}
