package b

import (
	"github.com/alejoasd/go_mvc_structure_example/mvc"
)

// NewRB returns a repository for B.
// Note that since B has no specific implementation, it simply returns the
// default repository. Having to create this function can potentially be
// skipped.
func NewRB(tx *mvc.DB) mvc.RI {
	return mvc.NewR(nil)
}
