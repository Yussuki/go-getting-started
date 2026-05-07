package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fileManager := filemanager.New("prices.txt", fmt.Sprintf("results/%.0f.json", taxRate*100))
		cmdManager := cmdmanager.New()
		priceJob := prices.New(taxRate, cmdManager)
		if err := priceJob.Process(); err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
