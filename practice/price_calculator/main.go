package main

import (
	"fmt"

	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/prices"
)

func main() {
	taxRates := []float64 {0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		priceJobFM := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJobFM.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

		// cmdm := cmdmanager.New()
		// priceJobCM := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		// priceJobCM.Process()
	}

}
