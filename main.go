package main

import (
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/iomanager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// slice channels of bool type
	doneChannels := make([]chan bool, len(taxRates))

	jobs := make([]*prices.TaxIncludedPriceJob, len(taxRates))

	for i, taxRate := range taxRates {

		// create a new chan for each taxRate
		doneChannels[i] = make(chan bool)

		var iom iomanager.IOManager
		inputFilePath := "prices.txt"
		outputFilePath := fmt.Sprintf("result_%.0f.json", taxRate*100)
		iom = filemanager.New(inputFilePath, outputFilePath)
		//iom = cmdmanager.New() // can be interchangeable because of iomanager.IOManager interface contracts
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRate)

		// starting a new goroutine to perform the process in parallel
		// pass channel to process function
		// one this process will become a goroutine,
		// it does not return any normal value anymore, instead we need use channel to get value
		// in this case the return was an error, we just ignore for now
		go priceJob.Process(doneChannels[i])

		//if err != nil {
		//	fmt.Println("Could not process job")
		//	fmt.Println(err)
		//	continue
		//}
		jobs[i] = priceJob
	}

	// go through all channels in slice to waiting for it communication to end the application when all process are done!
	for _, done := range doneChannels {
		<-done
	}

	for _, job := range jobs {
		if job != nil {
			fmt.Println(*job)
		}
	}
}
