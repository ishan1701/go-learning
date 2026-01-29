# go-learning
gonlang and kubernetes  learning


# how to apply 
1. SOLID principles in Go
2. Design Patterns in Go
3. Kubernetes Operators with Go
4. Build Kubernetes Controllers in Go





## why package is manatory
1. every file exists inside the package. That is something mandatory
2. go module can nave multiple packages
3. the main package is the plcase where the program starts execution. Because unlike python where a file runs via python comamnd, in go the code can be build as executable to run that.
4. I can build the module via command ` go build`. This will create an executable. The example is `udemy_course`
5. ./file_name is used to run the executable
6. The execubale can  be run anywhere  eventhough the Go is not installed


## GO key concepts
1. variables are defined as camel case
2. file name as snake case


## define the variables
1. var a float64=10
2. a:= 10

In the first case the type iis float as this declared
In the 2ne case, a will be of int.
If I make a:=10.0, thin the type of a becomes float
3. %T can be used to know the type if the variable


## import from another file
1. need to import
```
import "
```