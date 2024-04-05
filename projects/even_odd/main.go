//Directions
//Create a new package
//In the main function, create a slice of ints from 0 through 10
//Iterate through the slice. For each element, check to see if the number is even or odd.
//If the value is even, print out "even". If it is odd, print out "odd"

package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for number := range ints {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
