package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"taxRate"`
	InputPrices       []float64         `json:"inputPrices"`
	TaxIncludedPrices map[string]string `json:"taxIncludedPrices"`
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

	job.TaxIncludedPrices = taxedPrices

	// write data as JSON format
	err := filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
