package prices

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrice []float64
	TaxIncludedPrices map[string]float64
}
