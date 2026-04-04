package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue, err := getInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err = getInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err = getInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

	os.WriteFile("profit.txt", []byte(fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)), 0644)
}

func getInput(inputText string) (float64, error) {
	var userInput float64

	fmt.Print(inputText)

	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid value provided. Make sure it is a positive value.")
	}

	return userInput, nil
}

func calculate(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (100 - taxRate) / 100
	ratio = ebt / profit

	return ebt, profit, ratio
}
