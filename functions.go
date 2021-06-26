package main

import "fmt"

// if few parameters are of the same type, list them and then tell the type
func say(greeting, name string) {
	fmt.Printf("%v %v - multiple parameters, same type\n", greeting, name)
}

func pntr(number *int) {
	*number = 30
}

func variadic_func(values ...int) {
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Printf("%v - variadic parameters", result)
}

func functions() {
	say("Hello", "Audrius")
	a := 20
	pntr(&a)
	fmt.Printf("%v - change value in a function\n", a)

	variadic_func(1, 2, 3, 4, 5)
}
