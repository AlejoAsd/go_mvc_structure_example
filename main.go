package main

import (
	"fmt"
	"github.com/alejoasd/go_mvc_structure_example/mvc"
	"github.com/alejoasd/go_mvc_structure_example/mvc/a"
	"github.com/alejoasd/go_mvc_structure_example/mvc/b"
	"github.com/alejoasd/go_mvc_structure_example/mvc/c"
)

// process calls all methods of a Service implementation
func process(s mvc.SI, value interface{}) {
	s.Create(value)
	s.Get(value)
	s.Update(value)
	s.Delete(value)
}

func main() {
	fmt.Println("===== A =====\n")
	SA := a.NewSA()
	process(SA, &a.A{})
	SA.Custom(&a.A{})

	fmt.Println("\n===== B =====\n")
	SB := b.NewSB()
	process(SB, &b.B{})

	fmt.Println("\n===== C =====\n")
	SC := c.NewSC()
	process(SC, &c.C{})
}
