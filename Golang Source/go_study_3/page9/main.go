package main

import "fmt"

func main() {
	totalScore := ask(1, "dog")
	totalScore += ask(2, "cat")
	totalScore += ask(3, "fish")

	fmt.Println("Score", totalScore)
}

func ask(number int, question string) int {
	var input string
	fmt.Printf("[Question %d] Please input the following word: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
	    // Add the process where the word in the question matches the input value
	    fmt.Println("Correct!")
	    return 10
	    
	} else {
	    // Add the process where the word in the question doesn't match the input value
	    fmt.Println("Incorrect!")
	    return 0
	    
	}
}
