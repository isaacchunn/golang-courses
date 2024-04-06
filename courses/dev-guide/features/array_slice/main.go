package main

import "fmt"

//Arrays -> fixed length list of things
//Slice -> an array that can grow or shrink but must contain the same types

func main() {
	//Declare a slice in go
	cards := []string{"Ace of Diamonds", newCard()}
	//Appending a new item into our slice by returning a new slice that is reassigned
	cards = append(cards, "Six of Spades")

	//Using a for loop
	//index - is the index of element in array
	//card  - is current card iterating
	//range - takes the slice and loops over it
	for index, card := range cards {
		fmt.Println(index, card)
	}

	//prints all cards in a list format [...]
	fmt.Println(cards)
}

//A function that returns string
func newCard() string {
	return "Five of Diamonds"
}
