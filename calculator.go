package main

import "fmt"

func main() {
	var a, b float64
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
			fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
		case "-":
			fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
		case "*":
			fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
		case "/":
			if b != 0 {
				fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
			} else {
				fmt.Println("Error: TIDAK BISA PEMBAGIAN DENGAN 0")
			}

		}

		fmt.Printf("do you want calculate again ? (yes/no)\n")
		fmt.Scan(&choice)
		if choice != "yes" {
			break
		}

	}

}
