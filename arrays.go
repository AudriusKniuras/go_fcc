package main

import "fmt"

func arrays() {
	// simple array
	grades := [3]int{50, 63, 42}
	fmt.Printf("Grades: %v\n", grades)

	// simple array, ... means creating array with the len of elements
	grades2 := [...]int{50, 63, 42}
	fmt.Printf("Grades: %v\n", grades2)

	// array of arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// different than in other languages, array is a value, not a pointer to memory
	a := [...]int{1, 2, 3}
	b := a // copies entire array into b
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	// use pointers to have a copy of the array
	c := &a
	fmt.Println(c)
}

func slices() {
	a := []int{1, 2, 3}
	// unlike arrays, slices are reference types
	b := a // b points to a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := make([]int, 3, 100) // int slice with len of 3, and capacity of 100
	fmt.Printf("%v, len: %v, capacity: %v\n", c, len(c), cap(c))
}
