# go-learning
gonlang and kubernetes  learning


# how to apply 
1. SOLID principles in Go
2. Design Patterns in Go
3. Kubernetes Operators with Go
4. Build Kubernetes Controllers in Go



# GO is a static type language. 
1. I need to declare the type upfront
2. in case ogf :=, the go will assigned the type by itself
3. The types must be known before in case of fmt.Scan(&anyvar)

# how to create module via cmd line
```
➜  go-learning git:(udemy_learning) ✗ mkdir control_learning
➜  go-learning git:(udemy_learning) ✗ cd control_learning 
➜  control_learning git:(udemy_learning) ✗ go mod init control_learning/bank
go: creating new go.mod: module control_learning/bank
➜  control_learning git:(udemy_learning) ✗ ls -lrt
total 8
-rw-r--r--  1 ishan.kumar  staff  40 Jan 29 05:50 go.mod
➜  control_learning git:(udemy_learning) ✗ cat go.mod 
module control_learning/bank

go 1.25.5
➜  control_learning git:(udemy_learning) ✗ 
```


# why package is manatory
1. every file exists inside the package. That is something mandatory
2. go module can nave multiple packages
3. the main package is the plcase where the program starts execution. Because unlike python where a file runs via python comamnd, in go the code can be build as executable to run that.
4. I can build the module via command ` go build`. This will create an executable. The example is `udemy_course`
5. ./file_name is used to run the executable
6. The execubale can  be run anywhere  eventhough the Go is not installed


# GO key concepts
1. variables are defined as camel case
2. file name as snake case


# define the variables
1. var a float64=10
2. a:= 10

In the first case the type iis float as this declared
In the 2ne case, a will be of int.
If I make a:=10.0, thin the type of a becomes float
3. %T can be used to know the type if the variable


# import from another file
1. need to import
```
import "
```

# scan
1. It is used to get the inputs from the user
2. Scan doesnt support multi word input.


# functions
1. func getInputs() (float64, float64, float64) --> these 3 are the return types
2. when I call a function which and want to use the return values, I need to decalre the variables in that local function scope
```

func main() {
	var investmentAmount, rateOfInterest, numberOfYears float64. --> this needs to be defined unlike python

	investmentAmount, rateOfInterest, numberOfYears = getInputs()
}

```
3. I should use := inside the function to create a new variable rather to use var
```
func earnAfterTax(ebt float64, taxRate float64) float64 {
	eat := ebt * (1 - taxRate/100)       -> This one
	return eat
}
```


# formating text( strings)
## Sprintf

# Control statement and loops
1. In Go, else must be on the same line where the if is ending. I mean after the closing } 
```
func withdrawl(amount float64) (float64){
	if amount > accountBalance {
		fmt.Printf("The account balance is low. Please check your balance")
		return accountBalance
	} else {                 --------> this has to be in same line
		accountBalance -= amount
		return accountBalance
	}
}
```
2. go has only one loop
3. Conditional For Loops
```
Besides the for variations introduced in the last lectures, there also is another common variation (which will also be shown again later in the course):

for someCondition {
  // do something ...
}
someCondition is an expression that yields a boolean value or a variable that contains a boolean value (i.e., true or false).

In that case, the loop will continue to execute the code inside the loop body until the condition / variable yields false.
```

# error handing
1. Go doesnt have any try except block like in python
2. Errors are values returned from functions.
3. 


## what is defer in closing a file
1. Deferred functions run after the surrounding function returns, not after the current block. 
2. so defer will close the file after the main function existed
```
func example() {
    defer fmt.Println("world")
    fmt.Println("hello")
}

```
output
hello
world


```
defer.close()
```
The below is equivalent  to below
```
f = open("a.txt")
try:
    data = f.read()
finally:
    f.close()
```
