package main

import "fmt"

func main() {
	// Call the ask function 3 times with an argument
	ask("dog")
	ask("cat")
	ask("fish")
}

// Edit the method to accept an argument
func ask(question string) {
	var input string
	// Change the following code to use fmt.Printf
	fmt.Printf("Please input the following word: %s\n", question)
	fmt.Scan(&input)
	fmt.Printf("%s was input\n", input)
}