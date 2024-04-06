package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	//Using a make function to make an array of 2 with max size of 5
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	//Append does not overwrite the first 2 slots if we add smth to user names
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	//Making maps with max size of 3
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["d"] = 4.8
	courseRatings["t"] = 4.8
	//courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}
}
