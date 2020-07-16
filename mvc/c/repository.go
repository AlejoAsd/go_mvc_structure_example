package c

import (
	"github.com/alejoasd/go_mvc_structure_example/mvc"
)

// NewRC returns a repository for C.
// Note that since B has no specific implementation, it simply returns the
// default repository. Having to create this function can potentially be
// skipped.
func NewRC(tx *mvc.DB) mvc.RI {
	return mvc.NewR(nil)
}
