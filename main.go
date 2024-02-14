package main

import (
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/iomanager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	jobs := make([]*prices.TaxIncludedPriceJob, len(taxRates))

	for i, taxRate := range taxRates {
		var iom iomanager.IOManager
		inputFilePath := "prices.txt"
		outputFilePath := fmt.Sprintf("result_%.0f.json", taxRate*100)
		iom = filemanager.New(inputFilePath, outputFilePath)
		//iom = cmdmanager.New() // can be interchangeable because of iomanager.IOManager interface contracts
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
			continue
		}
		jobs[i] = priceJob
	}

	for _, job := range jobs {
		if job != nil {
			fmt.Println(*job)
		}
	}
}
