package main

import (
	"fmt"
)

const (
	RupiahtoDollar          = 0.000065
	EurotoDollar            = 1.12
	PoundsterlingtoDollar   = 1.32
	YentoDollar             = 0.0069
	maxDollarExchange		= 1000
)

func main() {
	for {
		fmt.Println("Selamat datang di program konversi mata uang")

		var jumlah float64
		var pilihan int

		fmt.Print("Masukkan jumlah Dollar yang ingin dikonversi (max=$1000): ")
		fmt.Scanln(&jumlah)


		if jumlah > 1000 {
			fmt.Println("Maaf, jumlah konversi melebihi batas maksimum 1000 Dollar.")
			fmt.Println("Silakan coba lagi dengan jumlah yang lebih kecil.")
			continue
		}

		fmt.Println("\nPilih mata uang tujuan:")
		fmt.Println("1. Rupiah")
		fmt.Println("2. Euro")
		fmt.Println("3. Poundsterling")
		fmt.Println("4. Yen Jepang")
		fmt.Print("Masukkan pilihan (1-4): ")
		fmt.Scanln(&pilihan)

		var hasil float64
		var matauang string

		switch pilihan {
		case 1:
			hasil = jumlah / RupiahtoDollar
			matauang = "Rupiah"
		case 2:
			hasil = jumlah / EurotoDollar
			matauang = "Euro"
		case 3:
			hasil = jumlah / PoundsterlingtoDollar
			matauang = "Poundsterling"
		case 4:
			hasil = jumlah / YentoDollar
			matauang = "Yen Jepang"
		default:
			fmt.Println("Pilihan tidak valid")
			continue
		}

		fmt.Printf("%.2f Dollar = %.2f %s\n", jumlah, hasil, matauang)

		var lanjut string
		fmt.Print("\nApakah Anda ingin melakukan konversi lagi? (y/n): ")
		fmt.Scanln(&lanjut)

		if lanjut != "y" && lanjut != "Y" {
			fmt.Println("Terima kasih telah menggunakan program konversi mata uang.")
			break
		}

		fmt.Println()
	}
}