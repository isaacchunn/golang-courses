package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	fmt.Println(prices)
// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	//Using subsets of arrays
// 	//Including:NotIncluding
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)
// }
