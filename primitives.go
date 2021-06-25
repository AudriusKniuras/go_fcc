package main

import "fmt"

func primitives() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 2
	fmt.Printf("%v, %T\n", m, m)

	fmt.Println("Bit operations:")
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010 = 2
	fmt.Println(a | b)  // 1011 = 11
	fmt.Println(a ^ b)  // XOR 1001 = 9 (1 of them has 1 but not both)
	fmt.Println(a &^ b) // ANDNOT 0100 = 8

	fmt.Println("Bit shifting:")
	c := 8              // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = 1

	fmt.Println("Complex numbers:")
	var d complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", real(d), real(d))
	fmt.Printf("%v, %T\n", imag(d), imag(d))

	fmt.Println("Strings:")
	e := "this is a string" // in go a string is a slice of bytes
	f := []byte(e)          // encodes string into bytes
	fmt.Printf("%v, %T\n", f, f)
}
