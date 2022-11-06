package main

import "fmt"

func main() {
	// Call the ask function again
	ask()
	ask()
	
}

func ask() {
    var input string
	fmt.Println("Please input the following word: dog")
    fmt.Scan(&input)
    fmt.Printf("%s was input\n", input)
}
