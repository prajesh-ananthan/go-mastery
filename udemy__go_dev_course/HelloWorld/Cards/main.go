package main

import "fmt"

func main() {
	// var card string = "Ace of Strings"
	card := "Ace of Strings"
	card = "Five of Diamonds"

	fmt.Println(card)
}

// Function return
func newCard() string {
	return "Five of Diamonds"
}
