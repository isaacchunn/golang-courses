package main

import "fmt"

func main() {
	fact := factorial(5)
	fact2 := factorialR(5)
	fmt.Println(fact)
	fmt.Println(fact2)
}

// Finding factorials
func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

//Recursive
func factorialR(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
