package main

import "fmt"

// getUserInputs reads revenue, expense, and taxRate from stdin and returns them.
// Inputs are expected as numbers (float64).
func getIUserInputs() (float64, float64, float64) {
	var revenue, expense, taxRate float64

	fmt.Print("Enter the Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the Expense: ")
	fmt.Scan(&expense)

	fmt.Print("Enter the Taxes: ")
	fmt.Scan(&taxRate)

	return revenue, expense, taxRate
}

// earnBeforeTax computes earnings before tax (EBT) = revenue - expense.
func earnBeforeTax(revenue float64, expense float64) float64 {
	ebt := revenue - expense
	return ebt
}

// earnAfterTax computes earnings after tax (EAT) = ebt * (1 - taxRate/100).
func earnAfterTax(ebt float64, taxRate float64) float64 {
	eat := ebt * (1 - taxRate/100)
	return eat
}

// main is the entry point for the profit calculator program.
func main() {
	fmt.Println("Welcome to profit calculator")

	var revenue, expense, taxes float64
	revenue, expense, taxes = getIUserInputs()

	ebt := earnBeforeTax(revenue, expense)
	fmt.Println("the ebt is ", ebt)

	eat := earnAfterTax(ebt, taxes)
	fmt.Println("the eat is ", eat)
}
