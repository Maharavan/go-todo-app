package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

// slice of string,creating my own type
type deck []string

func newDeck() deck {
	cards := deck{}

	cardsuits := []string{"Spades", "diamond", "hearts", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardsuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards

}

func (data deck) printDeck() {
	for ind, out := range data {
		fmt.Println(ind, out)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveTofile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0777)
}

func newDeckFromFile(filename string) deck {
	retrieve_data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(retrieve_data), ","))
}

func (d deck) shuffle() deck {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	// for i := range d {
	// 	newPos := rand.IntN(len(d) - 1)
	// 	d[i], d[newPos] = d[newPos], d[i]
	// }
	return d
}
