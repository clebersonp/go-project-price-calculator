package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		// len(prices) is the initial empty slots of slice
		// because we are using the initial as the slices length
		// we must access with its index position to avoid empty values at the beginning of slice
		pricesTaxed := make([]float64, len(prices))
		for priceIndex, price := range prices {
			pricesTaxed[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = pricesTaxed
	}

	fmt.Println(result)
}
