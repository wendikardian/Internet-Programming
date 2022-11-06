package main

import "fmt"

func main() {
	// Pass 1 as a new argument
	ask(1, "dog")
	
	// Pass 2 as a new argument
	ask(2, "cat")
	
	// Pass 3 as a new argument
	ask(3, "fish")
}

// Edit the method to accept number as an argument
func ask(number int,question string) {
	var input string
	// Remove the 1 line below and paste the given code
	fmt.Printf("[Question %d] Please input the following word: %s\n", number, question)
	
	fmt.Scan(&input)
	fmt.Printf("%s was input\n", input)
}
