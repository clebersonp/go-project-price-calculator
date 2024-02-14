package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
	"fmt"
)

// TaxIncludedPriceJob - to ignore serialization of a field just put -(hyphen) after json: tag
type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	// convert lines values into float values
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	// load data before process
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		// return after passed error to errorChan to enforce the break
		return
	}

	taxedPrices := make(map[string]string)
	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		taxedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	job.TaxIncludedPrices = taxedPrices

	// write data as JSON format
	err = job.IOManager.WriteResult(job)
	if err != nil {
		errorChan <- err
		return
	}

	// at the end we need to communicate the channel
	// we must omitted return statement because goroutine won't be able to receive the return data anymore
	// now the communication only works with channel
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
