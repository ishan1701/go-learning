package main

import "fmt"

const maxAttempts int = 2

var accountBalance float64 = 0.0

func getUserInput() int {
	var input int
	fmt.Scan(&input)

	switch input {
	case 1, 2, 3, 4:
		return input
	default:
		fmt.Printf("The user input %d is wrong\n", input)
		return 0
	}
}

func returnBalance() float64 {
	return accountBalance
}

func deposit(amount float64) float64 {
	accountBalance += amount
	return accountBalance
}

func withdrawl(amount float64) {
	if amount > accountBalance {
		fmt.Printf("The account balance is low. Please check your balance\n")
	}

	accountBalance -= amount
}

func main() {
	counter := 0
	var userInput int

	for {
		for {
			fmt.Println("Welcome to the Bank")
			fmt.Print("Enter 1- acc balance\n 2 - deposit \n 3- withdrawl\n 4 - exit\n")

			input := getUserInput()
			if input != 0 {
				userInput = input
				break
			}

			if counter >= maxAttempts {
				fmt.Println("Max attempts reached. Exiting.")
				break
			}

			fmt.Printf("the counter = %d", counter)
			counter += 1
		}

		switch userInput {
		case 1:
			balance := returnBalance()
			fmt.Printf("The account balance is %f\n", balance)

		case 2:
			var amount float64
			fmt.Println("Enter the amount to be deposited:")
			fmt.Scan(&amount)

			newBalance := deposit(amount)
			fmt.Printf("The new account balance after deposit is %f\n", newBalance)

		case 3:
			var amount float64
			fmt.Println("Enter the amount to be withdrawn:")
			fmt.Scan(&amount)

			withdrawl(amount)
			fmt.Printf("Collect your cash.\n")

		case 4:
			fmt.Println("Exiting the bank. Have a nice day!")
		}

		fmt.Println("Do you want to perform more?. Enter 1 for yes and 0 for no")
		var more int
		fmt.Scan(&more)
		if more == 0 {
			break
		}
	}
}
