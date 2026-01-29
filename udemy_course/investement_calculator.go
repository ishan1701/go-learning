package main

import (
	"fmt"
	"math"
)

// func main() {
// 	fmt.Print("Investment Calculator")

// 	var investmentAmount float64 = 100
// 	var rateOfInterest float64 = 5.5
// 	var numberOfYears float64 = 1

// 	var futureValue float64 = investmentAmount * math.Pow((1+rateOfInterest/100), numberOfYears)

// 	fmt.Println("the future  amount is ", futureValue)

// }

// /////////////// learning constants ////////////////////////
func learningContants() {
	const pi float64 = 3.14
	const e = 2.71

	fmt.Println("The value of pi is ", pi)
	fmt.Println("The value of e is ", e)
}

func getInputs() (float64, float64, float64) {
	var investmentAmount float64
	var rateOfInterest float64
	var numberOfYears float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the rate of interest: ")
	fmt.Scan(&rateOfInterest)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&numberOfYears)

	return investmentAmount, rateOfInterest, numberOfYears

}

///////////////// without declarative variables ////////////////////////

func main() {
	var investmentAmount, rateOfInterest, numberOfYears float64

	investmentAmount, rateOfInterest, numberOfYears = getInputs()

	futureValue := investmentAmount * math.Pow((1+rateOfInterest/100), float64(numberOfYears))
	fmt.Println("the future  amount is ", futureValue)

	fmt.Printf("the type of futureValue is %T", futureValue)

	learningContants()

}

///////////////// get  inputs from the terminal ////////////////////////
