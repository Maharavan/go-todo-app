package main

func main() {

	cards := newDeckFromFile("deckcard.txt")
	cards.shuffle()
	cards.printDeck()

}
