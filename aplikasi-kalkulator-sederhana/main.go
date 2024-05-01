package main

import "fmt"

func main() {
	fmt.Println("Selamat datang...\npada program\naplikasi kalkulator serhana")
	fmt.Println("###########################")

	fmt.Print("Silahkan pilih Operator\n(+ , - , * , /) : ")
	var operator string
	fmt.Scanln(&operator)

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("input tidak valid !")
		main()
	}

	var angka1 float64
	var angka2 float64
	fmt.Print("Masukkan angka pertama : ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scanln(&angka2)

	switch operator {
	case "+":
		fmt.Print(angka1, " + ", angka2, " = ")
		fmt.Printf("%.2f\n", angka1+angka2)

	case "-":
		fmt.Print(angka1, " - ", angka2, " = ")
		fmt.Printf("%.2f\n", angka1-angka2)

	case "*":
		fmt.Print(angka1, " * ", angka2, " = ")
		fmt.Printf("%.2f\n", angka1*angka2)

	case "/":
		fmt.Print(angka1, " / ", angka2, " = ")
		fmt.Printf("%.2f\n", angka1/angka2)

	}

	var selesai string
	fmt.Print("Apakah anda ingin mengulang ? (y/n) : ")
	fmt.Scanln(&selesai)

	if selesai == "y" || selesai == "Y" {
		main()
	} else if selesai == "n" || selesai == "N" {
		fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator serhana")
	} else {
		fmt.Println("input tidak valid !\nProgram berakhir :)")
	}

}
