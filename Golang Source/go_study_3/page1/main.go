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

	score := 0

	if question == input {
		fmt.Println("Correct!")
		score = 10
	} else {
		fmt.Println("Incorrect!")
	}

	return score
}
