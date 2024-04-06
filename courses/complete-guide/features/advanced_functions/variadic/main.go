package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	//Converts all parametsr and makes it into a slice of ints
	sum := sumup(1, 10, 15)
	//Passing in argument as slice and pull all elements out
	anotherSum := sumup(1, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// Typical way to sum up numbers in the list
// func sumup(numbers []int) int {
// 	sum := 0
// 	for _, val := range numbers {
// 		sum += val //sum = sum + val
// 	}
// 	return sum
// }

// Typical way to sum up numbers in the list
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, val := range numbers {
		sum += val //sum = sum + val
	}
	return sum
}
