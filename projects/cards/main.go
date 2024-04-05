package main

func main() {
	//only a value of type string will be assigned to this variable
	//var card string = "Ace of Spades"
	//Alternative way of writing is this and let the compiler figure it out
	//:= only applies to NEW variables (initialization stage)
	//card := "Ace of Spades"
	cards := newDeck()
	cards.print()
}
