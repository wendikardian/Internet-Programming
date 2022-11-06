package main

import "fmt"

func main() {
	x := 10
	y := 10

	// Call the calculate function, and pass x and the address of y variable as the argument
	calculate(x, &y)

	fmt.Println("When setting an integer as an argument:", x)
	fmt.Println("When setting a pointer as an argument:", y)
}

func calculate(x int, yPtr *int) {
	// Describe the process of adding 1 to the x variable and yPtr
	x += 1
	*yPtr += 1
}
