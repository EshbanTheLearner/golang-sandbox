package main

func main() {

	// cards := newDeck()
	cards := newDeckFromFile("myCards")
	cards.print()
}
