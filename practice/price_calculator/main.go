package main

import (
	"fmt"

	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/prices"
)

func main() {
	taxRates := []float64 {0, 0.07, 0.1, 0.15}
	doneChans := make([]chan error, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		priceJobFM := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJobFM.Process(doneChans[idx])
	}

	for _, doneChan := range doneChans {
		if <- doneChan != nil {
			fmt.Println("Could not process job")
			fmt.Println(doneChan)
		}
	}

}
