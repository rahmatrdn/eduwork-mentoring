package main

import "fmt"
func main() {

	//Tipe data int
	age := 25

	// Tipe data string
	name := "Fani"

	// Tipe data float64
	temperature := 98.6

	// Tipe data byte (uint8)
	var binaryData byte = 0b11010101

	// Tipe data bool
	isStudent := true

	// Tipe data nil
	var noValue interface{} // interface{} dapat berisi nil

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Temperature: %1f\n", temperature)
	fmt.Printf("Binary Data: %08b\n", binaryData)
	fmt.Printf("Is Student: %v\n", isStudent)
	fmt.Printf("No Value: %v\n", noValue)
}