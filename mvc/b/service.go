package b

import (
	"fmt"
	"github.com/alejoasd/go_mvc_structure_example/mvc"
)

// SB is a specific service implementation for B
type SB struct {
	mvc.SI
}

// NewSB creates a new Service for B
// This function can be made even more generic if we wanted to.
// That would allow us to not have to write the
func NewSB() *SB {
	// Using the specific repository
	RA := NewRB(nil)
	// By using the SI default implementation we get all CRUD operations
	// out of the box without writing a line of code. We can override and
	// extend any methods that we need to change.
	S := mvc.NewS(RA)
	return &SB{
		SI: S,
	}
}

func (s *SB) Get(in interface{}) interface{} {
	fmt.Println("SB Get")

	return s.SI.GetR().Create(in)
}
