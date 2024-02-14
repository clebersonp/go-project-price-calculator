package main

import (
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	jobs := make([]*prices.TaxIncludedPriceJob, len(taxRates))

	for i, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
		jobs[i] = priceJob
	}

	for _, job := range jobs {
		fmt.Println(*job)
	}
}
