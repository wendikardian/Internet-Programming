package main

import "fmt"

func main() {
	totalScore := 0
	fn(totalScore)

	// Print the value of the totalScore variable
	fmt.Println(totalScore)
	// Print the address of the totalScore variable
	fmt.Println(&totalScore)

}

func fn(totalScore int) {
	totalScore += 10
	// Print the value of the fn function's totalScore variable
	fmt.Println(totalScore)
	// Print the address of the fn function's totalScore variable
	fmt.Println(&totalScore)
}
