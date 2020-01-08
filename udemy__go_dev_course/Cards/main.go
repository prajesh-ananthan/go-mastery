package main

func main() {

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// This is possible because of the creation of the receiver
	// cards.print()

	cards := newDeck()
	cards.print()
}

func newDeck() deck {
	// Declare cards to type deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Iterate over cardValues slice
	// Use _ to remove index declaration
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}

	}
	return cards
}

func newCard() string {
	return "This is another Ace of Strings"
}
