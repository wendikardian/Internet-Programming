package main

import "fmt"

func main() {
	totalScore := 0
	// Pass the address of the totalScore variable as the third argument
	ask(1, "dog", &totalScore)
	ask(2, "cat", &totalScore)
	ask(3, "fish", &totalScore)

	fmt.Println("Score", totalScore)
}

// Add a pointer parameter scorePtr of type *int as the third parameter
func ask(number int, question string, scorePtr*int) {
	var input string
	fmt.Printf("[Question %d] Please input the following word: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("Correct!")
		// Calculate through dereferencing the pointer
		*scorePtr += 10
		
	} else {
		fmt.Println("Incorrect!")
	}
}
