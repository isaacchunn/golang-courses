package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is a slice of strings
// Extends or borrows behaviour from slice of strings
type deck []string

// Creates a new deck of cards and returns a list of strings
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

// Create a new receiver function for printing
// Any variable of type deck now gets access to the print method
func (d deck) print() {
	//d is the instance variable and reference to the exact copy
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Returns two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts our deck of cards to string
func (d deck) toString() string {
	//Deck is essentially a slice of strings
	return strings.Join([]string(d), ",")
}

// Function to save to file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Function to read from file and instantiate our deck objects
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		//Code of 1
		os.Exit(1)
	}
	//Slice of strings
	s := strings.Split(string(bs), ",") //Ace of Spades, Two of Spades, Three of Spades,...
	return deck(s)
}

// Function to shuffle the deck
func (d deck) shuffle() {

	//Make the source to make it random
	source := rand.NewSource(time.Now().UnixNano())
	//New rand object with our source
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		//Swap code
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
