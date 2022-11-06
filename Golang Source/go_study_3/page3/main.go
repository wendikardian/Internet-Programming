package main

import "fmt"

func main() {
	// Cut and paste the code below
    ask()
	// Cut and paste the code above

	// Call the ask function

}

func ask() {
	var input string
	fmt.Println("Please input the following word: dog")
	fmt.Scan(&input)
	fmt.Printf("%s was input", input)
}

// Define the ask function
