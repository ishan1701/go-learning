package main

import "fmt"
import "strings"

func sayBye(conName string) {
	message := "Have a nice day"
	fmt.Println(message)
	fmt.Println("We hope you enjoyed the", conName)
}

func simpleIfElse() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Printf("Hello %s. The type is %T\n", name)

	if strings.ToLower(name) == "ishan" {
		fmt.Println("Welcome back Ishan!")
	} else {
		fmt.Println("Hello there, nice to meet you!")
	}

}
func forLoopBasics() {
	list1 := []string{"a", "b", "c"}

	for index, value := range list1 {
		fmt.Println("the index is ", index, " the value is ", value)
	}
	// this is another way of doing the same
	for i := 0; i < len(list1); i++ {
		fmt.Println("the index is ", i, " the value is ", list1[i])
	}

}
func slicesBasics() {
	var list = []string{"a", "b", "c"}
	fmt.Println(list)

	fmt.Printf("The format of list is %T\n", list)

	list1 := []int{}
	list1 = append(list1, 3)
	fmt.Println(list1)

}

func goList() {
	var list = []string{"a", "b", "c"}
	fmt.Println(list)

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
	//var conferenceName = "Go learning conference"
	//fmt.Println(conferenceName)
	//
	//// define the constant
	//const conferenceTickets = 50
	//
	//fmt.Println("hello. Welcome to ", conferenceName,
	//	". Have a nice day")
	//fmt.Println("the number of tickets available is ", conferenceTickets)
	//
	////conferenceTickets = 20
	//
	//fmt.Printf("Type of name: %T\n", conferenceName)
	//fmt.Println(reflect.TypeOf(conferenceName))
	//
	//sayBye(conferenceName)
	//
	//variableTypeInference()
	//
	//// call function
	//pointersBasics(conferenceName)
	//fmt.Println(&conferenceName)                 // wll return the address of variable
	//fmt.Println(reflect.TypeOf(&conferenceName)) // this will return string
	//fmt.Printf("the address of var is %s\n", &conferenceName)

	//goList()
	//slicesBasics()
	//forLoopBasics()

	simpleIfElse()

}
