package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//Declare the variables for revenue expenses and tax rates
	var revenue, expenses, tax_rate float64
	var ebt, profit, ratio float64
	//We have to calcualte the earnings before tax, earnings after tax and the ratio
	//Scan the individual values for all types
	revenue, err := scanData("Revenue:")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err = scanData("Expenses:")
	if err != nil {
		fmt.Println(err)
		return
	}
	tax_rate, err = scanData("Tax Rate:")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Calculate profits
	ebt, profit, ratio = calculateProfits(revenue, expenses, tax_rate)

	//Then print out the stuff
	formattedString := fmt.Sprintf("EBT: %2f\nProfit: %2f\nRatio: %2f\n", ebt, profit, ratio)
	fmt.Print(formattedString)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func scanData(text string) (float64, error) {
	var data float64
	fmt.Print(text)
	fmt.Scan(&data)

	//Check if data is invalid
	if data <= 0.0 {
		return 0, errors.New("Value must be positive number.")
	}
	//no errors
	return data, nil
}

func calculateProfits(revenue float64, expenses float64, tax_rate float64) (ebt float64, profit float64, ratio float64) {
	//Then we can just calculate the stuff
	//Before tax is revenue - expenses?
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
