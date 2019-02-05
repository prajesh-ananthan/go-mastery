package main

import "fmt"

func main() {

	// Using String
	// var name string = "Prajesh"
	var age int = 32

	// Using const
	const isCool bool = true

	// Shorthand operator
	name := "Prajesh"
	lastName, email := "ananthan", "prajesh.ananthan@outlook.com"

	fmt.Println(name, age, isCool, lastName, email)

	// %T	a Go-syntax representation of the type of the value
	// Print the datatype of
	fmt.Printf("%T\n", isCool)
}
