package a

import (
	"fmt"
	"github.com/alejoasd/go_mvc_structure_example/mvc"
)

// SA is a specific service implementation for A
type SA struct {
	mvc.SI
}

// NewSA creates a new Service for A
// This function can be made even more generic if we wanted to.
// That would allow us to not have to write the
func NewSA() *SA {
	// Using the specific repository
	RA := NewRA(nil)
	// By using the SI default implementation we get all CRUD operations
	// out of the box without writing a line of code. We can override and
	// extend any methods that we need to change.
	S := mvc.NewS(RA)
	return &SA{
		SI: S,
	}
}

// getR returns this service's specific repository implementation.
func (s *SA) getR() *RA {
	return s.SI.GetR().(*RA)
}

// Create replaces the standard service's Create operation for an implementation
// specific for A.
func (s *SA) Create(value interface{}) interface{} {
	fmt.Println("SA Create")

	// Since this is a Service specifically tailored to A
	// We should be able to case the value to A
	valueA := value.(*A)

	// Perform some pre-processing on A's specific values
	valueA.I *= 2
	valueA.P = true

	// Call the repository implementation
	return s.SI.GetR().Create(value)
}

// Create replaces the standard service's Create operation for an implementation
// specific for A.
func (s *SA) Custom(value interface{}) interface{} {
	fmt.Println("SA Custom")

	// Since this is a Service specifically tailored to A
	// We should be able to case the value to A
	valueA := value.(*A)

	// Perform some pre-processing on A's specific values
	valueA.I *= 2
	valueA.P = true

	// Call the repository implementation
	return s.getR().Custom(value)
}
