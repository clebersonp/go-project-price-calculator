package main

import (
	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	jobs := make([]*prices.TaxIncludedPriceJob, len(taxRates))

	for i, taxRate := range taxRates {
		inputFilePath := "prices.txt"
		outputFilePath := fmt.Sprintf("result_%.0f.json", taxRate*100)
		ioManager := filemanager.New(inputFilePath, outputFilePath)
		priceJob := prices.NewTaxIncludedPriceJob(ioManager, taxRate)
		priceJob.Process()
		jobs[i] = priceJob
	}

	for _, job := range jobs {
		fmt.Println(*job)
	}

	fmt.Println("------------------")

	// reset jobs slice
	jobs = make([]*prices.TaxIncludedPriceJob, len(taxRates))
	for i, taxRate := range taxRates {
		ioManger := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(ioManger, taxRate)
		priceJob.Process()
		jobs[i] = priceJob
	}

	fmt.Println("-------------------")
	for _, job := range jobs {
		fmt.Println(*job)
	}
}
