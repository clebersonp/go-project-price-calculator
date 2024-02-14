package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
	"fmt"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert lines values into float values
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	// load data before process
	job.LoadData()

	taxedPrices := make(map[string]string)
	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		taxedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	job.TaxIncludedPrices = make(map[string]float64)
	for key, value := range taxedPrices {
		floatValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			fmt.Println("Failed to convert string to float!")
			fmt.Println(err)
			return
		}
		job.TaxIncludedPrices[key] = floatValue
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
