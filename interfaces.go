package main

import "fmt"

func interf() {
	var w Writer = ConsolWriter{}
	w.Write([]byte("Hello Go!"))

	// typical way to implement interfaces is using struct, but it can be other types too
	myInt := IntCounter(10)
	var inc Incrementer = &myInt
	for i := 0; i < 2; i++ {
		fmt.Printf("%v - infterface with int\n", inc.Increment())
	}
}

// ex1 - interface with struct
type Writer interface {
	Write([]byte) (int, error)
}

type ConsolWriter struct{}

func (cw ConsolWriter) Write(data []byte) (int, error) {
	n, err := fmt.Printf("%v - interfaces #1\n", string(data))
	return n, err
}

// ex2 - interface with int
type Incrementer interface {
	Increment() int
}

// we add method to this type
// we can't add method directly to int because we do not have control of it
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
