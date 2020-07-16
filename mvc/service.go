package mvc

import "fmt"

// SI is the base Service interface
type SI interface {
	// getR returns the service's repository
	GetR() RI
	Create(value interface{}) interface{}
	Get(out interface{}) interface{}
	Update(value interface{}) interface{}
	Delete(value interface{}) interface{}
}

// S is the base Service implementation
type S struct {
	RI
}

// NewS creates a new service
func NewS(r RI) *S {
	if r == nil {
		r = &R{}
	}
	return &S{RI: r}
}

// getR returns the service's repository
func (s *S) GetR() RI { return s.RI }

// Create creates a resource
func (s *S) Create(in interface{}) interface{} {
	fmt.Println("S Create")

	return s.RI.Create(in)
}

// Get gets a resource
func (s *S) Get(in interface{}) interface{} {
	fmt.Println("S Get")

	return s.RI.Get(in)
}

// Update updates a resource
func (s *S) Update(in interface{}) interface{} {
	fmt.Println("S Update")

	return s.RI.Update(in)
}

// Delete deletes a resource
func (s *S) Delete(in interface{}) interface{} {
	fmt.Println("S Delete")

	return s.RI.Delete(in)
}
