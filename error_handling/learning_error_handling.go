package main

import (
	"errors"
	"fmt"
	"os"
)

func getNumber() (int, error) {
	fmt.Println("Enter the number gt 0")
	var number int
	var err error
	fmt.Scan(&number)

	if number < 0 {
		err = errors.New("the number is less than 0")
	}
	return number, err
}
func main() {
	fmt.Println("Learning Error Handling in Go")
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error occurred while opening the file:", err)
	} else {
		fmt.Println("File opened successfully:", file.Name())
		defer file.Close()

	}

	number, e := getNumber()
	if e != nil {
		// fmt.Println("Error:", e)
		panic(e)
	} else {
		fmt.Println("The number entered is:", number)
	}

	fmt.Println("Ending the code. GoodBye")

}
