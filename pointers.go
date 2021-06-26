package main

import "fmt"

func pointer() {
	// longer syntax, but a:=42; b:= &a is the same
	var a int = 42
	var b *int = &a
	fmt.Printf("%v and %v (prints memory location) - simple pointer\n", a, b)
	fmt.Printf("memory location: %v, dereferenced: %v - simple pointer (2)\n", &a, *b)

	// uninitialized pointer is pointing to "nil"
	var ms *myStruct
	fmt.Printf("%v - unitialized pointer\n", ms)
	// new allocates memory for that type
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Printf("%v - changing pointer field\n", (*ms).foo)

	// compiler understand accessing pointer fields
	// ms is a pointer and doesnt have .foo field, but compiler understand that we are changing
	// the data that ms points to
	ms.foo = 50
	fmt.Printf("%v - accessing pointers without *\n", ms.foo)
}

type myStruct struct {
	foo int
}
