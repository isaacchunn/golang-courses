package main

import "fmt"

// Create a new type of 'deck'
// Which is a slice of strings
// Extends or borrows behaviour from slice of strings
type deck []string

//Creates a new deck of cards and returns a list of strings
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//Create a new receiver function for printing
//Any variable of type deck now gets access to the print method
func (d deck) print() {
	//d is the instance variable and reference to the exact copy
	for i, card := range d {
		fmt.Println(i, card)
	}
}
