package main

// import "fmt"

func main() {
	// cards := newDeckFromFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
