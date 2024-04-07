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
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Could not proces job")
		// 	fmt.Println(err)
		// }
	}

	//Wait for each chan to be done so instead of 12s it is now 3s
	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
	//But how do we handle error channels? Use the select statement
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}

	}
}
