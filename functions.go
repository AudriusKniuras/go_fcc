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
	fmt.Printf("%v - variadic parameters\n", result)
}

func returning_func(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func multiple_returns(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("division by 0")
	}
	return a / b, nil
}

func functions() {
	say("Hello", "Audrius")
	a := 20
	pntr(&a)
	fmt.Printf("%v - change value in a function\n", a)

	variadic_func(1, 2, 3, 4, 5)
	b := []int{1, 2, 3, 4, 5}
	variadic_func(b...)

	c := returning_func(b...)
	fmt.Printf("%v - returning func\n", c)

	result, err := multiple_returns(8.4, 1)
	if err != nil {
		fmt.Printf("%v - f-tion with multiple return vars\n", err)
		return
	}
	fmt.Printf("%v - f-tion with multiple return vars\n", result)

	// anonymous function
	var f func() = func() { fmt.Printf("anonymous func\n") }
	g := func() { fmt.Printf("anonymous func2\n") }
	f()
	g()

	// methods
	greeter := greeter{greeting: "Hello", name: "Go"}
	greeter.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Printf("%v %v - method function", g.greeting, g.name)
}
