package main

import "fmt"

/*
NOTES:
Interfaces are not generic types
Interfaces are implicit -> do not have to say custom type satisfies interface
Interfaces are a contract to help us manage types
Interfaces are tough. Step #1 is understanding hwo to read them
*/

type bot interface {
	getGreeting() string
}

//Promoted as bot as they have the same functions and satisfies the interface
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//No function overloading in Go
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// VERY CUSTOM LOGIC FOR GENERATING an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
