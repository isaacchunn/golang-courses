package main

import (
	"fmt"
	"isaac/price-calc/filemanager"
	"isaac/price-calc/prices"
)

func main() {
	//Initialize some variables
	//prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not proces job")
			fmt.Println(err)
		}
	}
}
