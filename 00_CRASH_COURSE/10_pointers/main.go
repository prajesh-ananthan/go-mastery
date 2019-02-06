package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println(b) // 0xc0000140c0

	// Use * to read val from address
	fmt.Println(*b)  // 5
	fmt.Println(*&a) // 5

	// Change value with pointer
	*b = 10
	fmt.Println(a) // 10
}
