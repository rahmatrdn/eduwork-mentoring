package main

import (
	"fmt"
)

const (
	DollartoRupiah          = 15573.00
	EurotoRupiah            = 17191.58
	PoundsterlingtoRupiah   = 20169.64
	YentoRupiah             = 106.14
	maxDollarExchange      = 1000
)

func main() {
	for {
		fmt.Println("Selamat datang di program konversi mata uang")

		var jumlah float64
		var pilihan int

		fmt.Print("Masukkan jumlah Rupiah yang ingin dikonversi: ")
		fmt.Scanln(&jumlah)

		fmt.Println("\nPilih mata uang tujuan:")
		fmt.Println("1. Dollar AS")
		fmt.Println("2. Euro")
		fmt.Println("3. Poundsterling")
		fmt.Println("4. Yen Jepang")
		fmt.Print("Masukkan pilihan (1-4): ")
		fmt.Scanln(&pilihan)

		var hasil float64
		var matauang string

		switch pilihan {
		case 1:
			hasil = jumlah / DollartoRupiah
			matauang = "Dollar AS"
		case 2:
			hasil = jumlah / EurotoRupiah
			matauang = "Euro"
		case 3:
			hasil = jumlah / PoundsterlingtoRupiah
			matauang = "Poundsterling"
		case 4:
			hasil = jumlah / YentoRupiah
			matauang = "Yen Jepang"
		default:
			fmt.Println("Pilihan tidak valid")
			continue
		}

		fmt.Printf("%.2f Rupiah = %.2f %s\n", jumlah, hasil, matauang)

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