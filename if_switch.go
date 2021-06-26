package main

import "fmt"

func ifs() {
	if true {
		fmt.Println("simple if")
	}
	dict := map[string]int{
		"a": 1,
	}

	// evaluate and execute if in one line
	if pop, ok := dict["a"]; ok {
		fmt.Printf("value: %v - get value, and evaluate in if statement\n", pop)
	}

	// simple switch
	a := 2
	switch a {
	case 1, 2, 3:
		fmt.Println("1,2,3 - simple switch with multiple cases")
	case 4, 5, 6:
		fmt.Println("4,5,6 - simple switch with multiple cases")
	default:
		fmt.Println("Other value - simple switch")
	}
	// switch with evaluation
	switch i := 2 + 3; i {
	case 1, 2, 3:
		fmt.Println("1,2,3 - simple switch with evaluation")
	case 4, 5, 6:
		fmt.Println("4,5,6 - simple switch with evaluation")
	default:
		fmt.Println("Other value - simple switch with evaluation")
	}

	// tagless switch
	i := 10
	switch {
	case i <= 10:
		fmt.Println("i <= 10 - tagless switch")
	case i > 10:
		fmt.Println("i > 10 - tagless switch")
	default:
		fmt.Println("error")
	}

	// fallthrough switch
	// switch cases have implicit break, to have another case evaluated if one already did, use fallthrough
	i = 10
	switch {
	case i <= 10:
		fmt.Println("i <= 10 - fallthrough switch")
		fallthrough
	case i > 10:
		fmt.Println("i > 10 - fallthrough switch")
	default:
		fmt.Println("error")
	}

	// type switch
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("i is int - type switch")
	case string:
		fmt.Println("i is string - type switch")
	default:
		fmt.Println("unrecognized type - type switch")
	}

}
