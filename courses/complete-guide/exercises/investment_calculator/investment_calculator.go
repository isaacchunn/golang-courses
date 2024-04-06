package main

// This file works with fmt scans and constants
import (
	"fmt"
	"math"
)

func main() {
	//Using constants
	const inflationRate = 2.5
	//Wait for scan later
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//Use the fmt scan feature
	//NOTE: Scan limitation lies in that we cant fetch multi word input values. Fetching
	//Text that consists of more than one single word is tricky with this function

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
