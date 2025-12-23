package main

import "fmt"
import "reflect"

func sayBye(conName string) {
	message := "Have a nice day"
	fmt.Println(message)
	fmt.Println("We hope you enjoyed the", conName)
}

func variableTypeInference() {
	var someVar int
	// type should be defined
	someVar = 42
	fmt.Println(someVar)

	someVar2 := "dd"
	// automatic type inference
	fmt.Println(someVar2)

}

func pointersBasics(name string) {

	pointerExample := 40
	fmt.Println(pointerExample)
	fmt.Printf("the address of local var is %s\n", &pointerExample)
	fmt.Printf("the address of argument var is %s\n", &name)

}
func main() {

	// define the variable
	var conferenceName = "Go learning conference"
	fmt.Println(conferenceName)

	// define the constant
	const conferenceTickets = 50

	fmt.Println("hello. Welcome to ", conferenceName,
		". Have a nice day")
	fmt.Println("the number of tickets available is ", conferenceTickets)

	//conferenceTickets = 20

	fmt.Printf("Type of name: %T\n", conferenceName)
	fmt.Println(reflect.TypeOf(conferenceName))

	sayBye(conferenceName)

	variableTypeInference()

	// call function
	pointersBasics(conferenceName)
	fmt.Println(&conferenceName)                 // wll return the address of variable
	fmt.Println(reflect.TypeOf(&conferenceName)) // this will return string
	fmt.Printf("the address of var is %s\n", &conferenceName)

}
