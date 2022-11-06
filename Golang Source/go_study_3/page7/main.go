package main

import "fmt"

func main() {
	// Assign the return value to totalScore
	totalScore := ask(1, "dog")

	// Paste the given code specified in the instruction
	fmt.Println("Score", totalScore)
	
}

// Set the type of the return value
func ask(number int, question string) int {
	var input string
	fmt.Printf("[Question %d] Please input the following word: %s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%s was input\n", input)
	// Return 10 as the return value
	return 10
	
}
