package main

import "fmt"

func main() {
	age := 32 //regular variable

	//Declare a pointer
	agePointer := &age

	fmt.Println("Age:", *agePointer)

	adultYears := getAdultYears(&age)
	fmt.Println(adultYears)
}

//Values are copied and not pass by reference
func getAdultYears(age *int) int {
	return *age - 18
}
