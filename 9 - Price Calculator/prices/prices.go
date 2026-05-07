package prices

import (
	"fmt"

	"example.com/price-calculator/iomanager"
	"example.com/price-calculator/utils"
)

type TaxedPriceJob struct {
	IOManager   iomanager.IOManager `json:"-"` // json "-" says we should ignore this field
	TaxRate     float64             `json:"tax_rate"`
	Prices      []float64           `json:"prices"`
	TaxedPrices map[string]string   `json:"taxed_prices"`
}

func New(taxRate float64, ioManager iomanager.IOManager) *TaxedPriceJob {
	return &TaxedPriceJob{
		IOManager: ioManager,
		TaxRate:   taxRate,
		Prices:    []float64{10, 20, 30},
	}
}

func (job *TaxedPriceJob) LoadData() error {
	stringPrices, err := job.IOManager.ReadFile()
	if err != nil {
		return err
	}

	floatPrices, err := utils.StringsToFloat(stringPrices)
	if err != nil {
		return err
	}

	job.Prices = floatPrices
	return nil
}

func (job TaxedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)

	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxedPrices = result

	return job.IOManager.WriteResult(job)
}
