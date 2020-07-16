package c

import (
	"fmt"
	"github.com/alejoasd/go_mvc_structure_example/mvc"
)

// SC is a specific service implementation for C
type SC struct {
	mvc.SI
}

// NewSC creates a new Service for C
// This function can be made even more generic if we wanted to.
// That would allow us to not have to write the
func NewSC() *SC {
	// Using the specific repository
	RA := NewRC(nil)
	// By using the SI default implementation we get all CRUD operations
	// out of the box without writing a line of code. We can override and
	// extend any methods that we need to change.
	S := mvc.NewS(RA)
	return &SC{
		SI: S,
	}
}

func (s *SC) Delete(in interface{}) interface{} {
	fmt.Println("SC Delete")

	return s.SI.GetR().Create(in)
}
