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
	fmt.Printf("%v - reference type\n", a)
	fmt.Printf("%v - reference type\n", b)

	d := make([]int, 2, 100) // creating slice with len and capacity. By default elements are 0's [0 0] with capacity of 100
	fmt.Printf("%v, len: %v, capacity: %v - capacity, len showcase\n", d, len(d), cap(d))

	c := append(a, 4, 5, 6, 7, 8, 9) // variadic function - accept any number of elements
	fmt.Printf("%v, len: %v, capacity: %v - variadic function (append)\n", c, len(c), cap(c))

	// e := append(a, []int{4, 5, 6}) will not work - appended elements must be of the same type (int), not int[]
	e := append(a, []int{4, 5, 6}...) // ... expands the slice into individual elements
	fmt.Printf("%v - expanding slice with ... \n", e)

	// pushing/popping elements in a slice
	a = []int{1, 2, 3, 4, 5}
	b = a[1:]
	c = a[:len(a)-1]
	fmt.Printf("top removed: %v, bottom removed: %v", b, c)
}
