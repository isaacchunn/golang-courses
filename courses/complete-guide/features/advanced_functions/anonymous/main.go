package main

import "fmt"

type transformFn func(int) int

//File to describe anonymous functions and closures
func main() {
	numbers := []int{1, 2, 3}

	//Usage of anonymous functions
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	//Usage of closures as a factory function
	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

// Using functions as values as other functions (?)
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

//Closures
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
