package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.TaxIncludedPrices = make(map[string]float64)
	for _, price := range job.InputPrices {
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
