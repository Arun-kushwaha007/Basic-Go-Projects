package main
// import "fmt"

// func main() {
// 	fmt.Println("Welcome to the Go programming tutorial!")
// 	fmt.Println("In this tutorial, we will cover the basics of Go.")
// 	fmt.Println("Let's get started with variables and data types.")
	
// 	// Declare a variable
// 	var message string = "Hello, Go!"
// 	fmt.Println(message)

// 	// Declare an integer variable
// 	var number int = 42
// 	fmt.Println("The number is:", number)

// 	// Declare a float variable
// 	var pi float64 = 3.14
// 	fmt.Println("The value of pi is:", pi)
// 	// Declare a boolean variable
// 	var isGoFun bool = true
// 	fmt.Println("Is Go fun?", isGoFun)
// 	// Declare a constant
// 	const piConstant float64 = 3.14159
// 	fmt.Println("The constant value of pi is:", piConstant)
// 	// Declare a slice
// 	var fruits []string = []string{"Apple", "Banana", "Cherry"}
// 	fmt.Println("Fruits:", fruits)
// 	// Declare a map
	
// 	// := 
// 	name := "John Doe"  // here we use the short variable declaration syntax. It // is equivalent to var name string = "John Doe", // but it is more concise.
// 	// used when we dont now the type of variable . //used for declaring and initializing variables in a single line. It is // only allowed in function bodies.
// 	// Print the name variable
// 	fmt.Println(name); 

// 	// use of Printf :
// 	fmt.Printf("The name is: %v\n", name)  // %v is a verb that formats the value in a default format. It is used to print // the value of the variable in a human-readable format.

// 	// .Scan 
// 	fmt.Print("Enter your name: ")
// 	var inputName string
// 	fmt.Scan(&inputName)  // &inputName is used to pass the address of the variable inputName to the Scan function. This allows the function to modify the value of inputName. // it just use to take input from the user.
// 	fmt.Println(inputName)

// }