package main

import "fmt"

func main() {
	totalScore := 0
	ask(1, "dog")
	ask(2, "cat")
	ask(3, "fish")

	fmt.Println("Score", totalScore)
}

func ask(number int, question string) {
	var input string
	fmt.Printf("[Question %d] Please input the following word: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("Correct!")
		// Adds 10 to the totalScore
		totalScore += 10
		
		
	} else {
		fmt.Println("Incorrect!")
	}
}
