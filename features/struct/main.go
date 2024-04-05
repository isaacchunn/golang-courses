package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Define a new struct and properties it should ahve
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//The order depends on the order of fields in the struct
	//alex := person{"Alex", "Anderson"}
	//A much better approach
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//A third way to declare alex
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	//Print out all values and keys
	alex.print()

	//Indented struct area
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//Root type works as shortcut in Go
	jim.updateName("jimmy")
	jim.print()
}

//Receiver function with structs
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//Pointer to pass by reference
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
