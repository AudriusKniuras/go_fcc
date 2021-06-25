package main

import (
	"fmt"
	"strconv"
)

// i := 42 will not work - show var declaration can only be used in functions
// variable at the package level must be declared with var
var name = "Audrius"
var name2 string = "Audrius"

var (
	name3 string = "Audrius"
	age   int    = 26
)

// variable exposed outside of "main" (this) package
var Name4 string = "Audrius"

func variables() {
	// variable initilisation
	var i float32 = 42
	j := 52.0 // defaults to float64, no way to initialize float32 without declaring it
	fmt.Printf("%v, %T", i, i)
	fmt.Println()
	fmt.Printf("%v, %T", j, j)
	fmt.Println()

	// variable conversion
	z := int(i)
	fmt.Println(z)

	var x string
	// for conversion to string a strconv package is needed. Below will not work
	// x = string(z)
	x = strconv.Itoa(z) // Itoa - integer to ascii; atoi - ascii to integer
	fmt.Println(x)
}
