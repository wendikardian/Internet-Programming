package main

import "fmt"

func main() {
    // Declare the input variable of type string
    var input string
    
    // Print "Please input the following word: dog"
    fmt.Println("Please input the following word: dog");
    
    // Use Scan to read the input value
    fmt.Scan(&input)
    
    // Print the input value
    fmt.Printf("%s was input", input)
}