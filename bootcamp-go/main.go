package main

import (
	"fmt"
)

func main () {
	const pi = 3.14159

	var radius float64 = 5.0
	var area float64

	area = pi * radius * radius

	fmt.Println("Area of the circle:", area)
}