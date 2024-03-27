package main

import "fmt"

func main() {
	// Declaring constants
	const pi = 3.14

	// Declaring variables with explicit types
	var radius float64 = 5.0
	var area float64

	// Calculating the area of a circle
	area = pi * radius * radius

	// Printing the result
	fmt.Println("Area of the circle", area)
}
