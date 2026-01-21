# arrays and slices in go
1. Arrays in go must have fixed size.
2. The logic as 
```go
var bookings = [10]string{} // this will create an array of size 10 which can store string values.
```

3. the list contains only of one data type. Unlike python list which can contain multiple data types.
4. I can also defined the array like this
```go
var bookings = [5]string{"a", "b", "c", "d", "e"}
or
var bookiings = [5]string
bookings[0] = "a"
````
5. Its better to use slices in go than arrays.
6. Slices are dynamic in size.
7. Slices are built on top of arrays.

# loops in go
1. Go has only one loop which is for loop.
2. There is no while do dowhile or foreach loop in go.
3. there is something called finite and infinite loops in go.
4. Finite loop
```go
list1 := []string{"a", "b", "c"}

for index, value := range list1 {
fmt.Println("the index is ", index, " the value is ", value)
}
// this is another way of doing the same
for i := 0; i < len(list1); i++ {
fmt.Println("the index is ", i, " the value is ", list1[i])
}
````

# conditionals in go
1. Go has if else and switch case conditionals.

