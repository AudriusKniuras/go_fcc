package main

import "fmt"

func looping() {
	// simple loop
	for i := 0; i < 2; i++ {
		fmt.Printf("%v - simple loop\n", i)
	}
	// range loop
	s := []int{10, 20, 30}
	for key, value := range s {
		fmt.Printf("key: %v, value: %v - range loop\n", key, value)
	}

	// looping over string
	str := "Hello"
	for key, value := range str {
		fmt.Printf("key: %v, value: %v\n", key, string(value)) // without string(value) it prints ascii numbers
	}
}
