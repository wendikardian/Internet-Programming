package main

import "fmt"

func main() {
	name := "John"

	fmt.Println(name)
	fmt.Println(&name)
	// Declare the namePtr pointer of type string and assign the address of the variable name
	var namePtr*string = &name
	// Print the value of namePtr
	fmt.Println(namePtr)

}
