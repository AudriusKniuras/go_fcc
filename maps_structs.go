package main

import (
	"fmt"
	"reflect"
)

func maps() {
	// simple map
	users := map[string]int{
		"audrius": 26,
		"eivilte": 25,
	}
	fmt.Printf("%v - simple map\n", users)

	// if entries not available during the initilization, make can be used
	users = make(map[string]int)
	users = map[string]int{"audrius": 26, "eivilte": 25}
	fmt.Printf("%v - map created with make\n", users)

	// deleting elements in a map
	delete(users, "audrius")
	fmt.Printf("%v - deleted element\n", users)

	// accessing non-existant key returns 0, not an error !!
	fmt.Printf("%v - non existant key returns 0\n", users["audrius"])

	value, exists := users["audrius"] // convention is to use "ok" instead of "exists" in these cases
	fmt.Printf("%v, %v - value, key exists?\n", value, exists)
}

// fields in the struct need to be capitalized too to export them to other packages
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func structs() {
	aDoctor := Doctor{
		Number:    1,
		ActorName: "Jon Depp",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Printf("%v - accessing value in a struct\n", aDoctor.ActorName)

	// anonymous structs
	bDoctor := struct {
		name   string
		number int
	}{name: "John Depp", number: 2}
	fmt.Printf("%v - anonymous struct\n", bDoctor)

	// composition through embedding - similar thing to inheritance
	// Traditional OOP inheritance: Bird class with f-tions, Duck class inherits Bird
	// This is not a traditional inheritance, bird is not an animal, it just has these fields - composition inheritance
	// Generally this is done using interfaces, not structs
	type animal struct {
		name   string
		origin string
	}
	type bird struct {
		animal   // embed Animal into Bird struct
		speedKPH float32
		canFly   bool
	}
	b := bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speedKPH = 48
	b.canFly = false
	fmt.Printf("%v - struct embedding\n", b)

	// if literal syntax is used we need to know the structure of the inherited struct
	c := bird{
		animal:   animal{name: "varna", origin: "LT"},
		speedKPH: 30,
		canFly:   true,
	}
	fmt.Printf("%v - in literal syntax, need to know inherited struct's structure\n", c)

	// struct tags - a comment (a simple string) for struct's elements
	type tag_example struct {
		name string `required_max:"100"`
		age  int
	}
	// to get the tag we need to use the reflect package
	t := reflect.TypeOf(tag_example{})
	field, _ := t.FieldByName("name")
	fmt.Printf("%v - outputting the tag with the reflect package\n", field.Tag)
}
