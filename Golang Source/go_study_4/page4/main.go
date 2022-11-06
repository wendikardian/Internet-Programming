package main

import "fmt"

func main() {
	name := "John"

	fmt.Println(name)
	// Declare the namePtr pointer and assign the address of variable name
	var namePtr*string = &name
	
	// Update the value of variable name with "Kate" through dereferencing of namePtr
	*namePtr = "Kate"
	
	// Print the updated value of name variable
    fmt.Println(name)	

}
