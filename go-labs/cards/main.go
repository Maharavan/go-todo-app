package main

import (
	"fmt"
)

func main() {

	cards := []string{"Jack", newCard()}
	cards = append(cards, "Diamond")
	for ind, card := range cards {
		fmt.Println(ind, card)
	}
}

func newCard() string {
	card := "Ace"
	return card
}
