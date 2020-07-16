package a

import (
	"github.com/alejoasd/go_mvc_structure_example/mvc"
)

// RA is a specific repository implementation for A
type RA struct {
	mvc.RI
}

// NewRA creates a new RA
// Note that because RA uses mvc.R, all the RI methods are implemented
// by default. RI methods can be overriden to tailor them for the model, and
// new methods can be added.
func NewRA(tx *mvc.DB) mvc.RI {
	return &RA{
		RI: mvc.NewR(nil),
	}
}

// RA replaces one of the base methods and adds a new method
func (r *RA) Create(value interface{}) error {
	// Perform the basic operation
	if err := r.RI.Create(value); err != nil {
		return err
	}
	// Do some postprocessing on the input value
	// Since this is a repository specifically tailored to A
	// We should be able to case the value to A
	valueA := value.(*A)
	valueA.I += 1000

	return nil
}

func (r *RA) Custom(in interface{}) error {
	// Since these are a repository and method specifically tailored to A
	// We should be able to case the value to A
	valueA := in.(*A)
	return valueA.Custom()
}
