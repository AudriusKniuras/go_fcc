package main

import "fmt"

// enumerated constants
const (
	_ = iota // if 0 enum is not used, throw it away
	a
	b
	c
)

// can't use powers in const, because it's a function, and cannot be used in const declaration
// but we can use bitshifting
const (
	_  = iota
	KB = 1 << (10 * iota) // 2^10
	MB                    // 2^100 (1 << 20)
	GB
	TB
	PB
)

// each const is bitshifted by 1, so each occupies 1 bit in a number
const (
	isAdmin          = 1 << iota // 1 << 0 = 00001 = 1
	isHeadquarters               // 1 << 1 = 00010 = 1
	CanSeeFinancials             // 1 << 2 = 00100 = 4

	canSeeAfrica // 1 << 3 = 01000 = 8
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func contants() {
	// const declaration
	const myConst int = 42

	fmt.Println("simple enum:")
	fmt.Printf("%v\n", a) // 1
	fmt.Printf("%v\n", b) // 2
	fmt.Printf("%v\n", c) // 3

	fmt.Println("enum with powers of number")
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)

	fmt.Println("enum with bitshifting for access rights")
	var roles byte = isAdmin | CanSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)                              // 100101
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin) // 000001 & 100101 = 000001
	fmt.Printf("Is HQ? %v", isHeadquarters&roles == isHeadquarters)
}
