package main

import (
	"fmt"
)

func pointersBasics() {
	some_value := 10.23
	address := &some_value
	fmt.Printf("The some_value is %f \n address is %p\n value at address is %f\n", some_value, address, *address)
	fmt.Println("sime print is ", *address)
}

func someOperationOnPointer(p1 *int, p2 *int) (int, int) {

	fmt.Println("values at pointer are ", p1, p2)
	fmt.Println("values at pointer are ", *p1, *p2)

	// lets modify the values at the pointer
	fmt.Print("Enter new value for num2")
	var newValue int
	fmt.Scan(&newValue)
	*p1 = 30
	*p2 = newValue

	fmt.Println("values at pointer after modification are ", *p1, *p2)

	sum := *p1 + *p2
	fmt.Println("sum of values at pointer is ", sum)
	if *p2 == 0 {
		panic("division by zero")
	}
	div := *p2 / *p1
	fmt.Println("division of values at pointer is ", div)
	return sum, div
}
func main() {
	fmt.Println("Welcome to the poiinters")
	// pointersBasics()

	var num1, num2 int = 10, 20

	p1 := &num1
	p2 := &num2

	sum, div := someOperationOnPointer(p1, p2)

	fmt.Println("values after modification in main are ", num1, num2)
	fmt.Println("sum returned is ", sum)
	fmt.Println("division returned is ", div)

}
