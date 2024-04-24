package main

import "fmt"

func main() {
	var a, b int
	var operator string
	fmt.Println("Aplikasi Kalkulator")
	fmt.Println("###############################")

	fmt.Println("Input angka pertama: ")
	fmt.Scan(&a)

	fmt.Println("Input angka kedua: ")
	fmt.Scan(&b)

	fmt.Println("Pilih Operator(+,_,*,/)")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("%d + %d = %d\n", a, b, a+b)
	case "-":
		fmt.Printf("%d - %d = %d\n", a, b, a-b)
	case "*":
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
	case "/":
		fmt.Printf("%d / %d = %d\n", a, b, a/b)

	}
}
