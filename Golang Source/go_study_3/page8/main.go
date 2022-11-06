package main

import "fmt"

func main() {
	totalScore := ask(1, "dog")
	// Add the return value from the ask function to the totalScore
	totalScore += ask(2, "cat")
	totalScore += ask(3, "fish")
	fmt.Println("Score", totalScore)
	
}

func ask(number int, question string) int {
	var input string
	fmt.Printf("[Question %d] Please input the following word: %s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%s was input\n", input)
    
    return 10
	
}
