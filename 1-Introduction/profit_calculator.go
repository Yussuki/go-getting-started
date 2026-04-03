package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getInput("Revenue: ")
	expenses = getInput("Expenses: ")
	taxRate = getInput("Tax Rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)

}

func getInput(inputText string) (userInput float64) {

	fmt.Print(inputText)
	fmt.Scan(&userInput)

	return userInput
}

func calculate(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (100 - taxRate) / 100
	ratio = ebt / profit

	return ebt, profit, ratio
}
