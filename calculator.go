package main

import "fmt"

func main() {
	var a, b int
	var operator, choice string
	fmt.Println("Aplikasi Kalkulator")
	fmt.Println("###############################")

	for {
		fmt.Println("Input angka pertama: ")
		_, err := fmt.Scan(&a)
		if err != nil {
			continue
		}

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

		fmt.Printf("do you want calculate again ? (yes/no)\n")
		fmt.Scan(&choice)
		if choice != "yes" {
			break
		}

	}

}
