package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "My name is " + p.firstName + " " + p.lastName + ". I am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	// Init person struct
	person1 := Person{
		firstName: "Vicky",
		lastName:  "Chai",
		city:      "KL",
		gender:    "f",
		age:       29}

	fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("Ananthan")
	fmt.Println(person1.greet())
}
