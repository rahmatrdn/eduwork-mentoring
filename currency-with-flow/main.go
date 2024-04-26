package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	USDToEUR = 0.93
	USDToGBP = 0.80
	USDToJPY = 156.76
	EURToUSD = 1.07
	EURToGBP = 0.86
	EURToJPY = 168.18
	GBPToUSD = 1.25
	GBPToEUR = 1.17
	GBPToJPY = 196.09
	JPYToUSD = 0.0064
	JPYToEUR = 0.0059
	JPYToGBP = 0.0051
)

func main() {
	fmt.Println("Selamat datang :)\ndi program sederhana\nKonversi Mata Uang")
	fmt.Println("==================")
	fmt.Println("USD\nEUR\nGBP\nJPY")
	fmt.Print("Pilih mata uang asal: ")

	var fromCurrency string
	fmt.Scanln(&fromCurrency)
	fromCurrency = strings.ToUpper(fromCurrency)

	if fromCurrency != "USD" && fromCurrency != "EUR" && fromCurrency != "GBP" && fromCurrency != "JPY" {
		fmt.Println("Error: Masukkan mata uang yang benar !")
		main()
	}
	fmt.Print("Masukkan jumlah uang: ")
	var amount string
	fmt.Scanln(&amount)

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a valid number.")
		main()
	}

	if amountFloat > 1000 {
		fmt.Println("Maaf, jumlah yang di konversi cukup besar (max : 1000)")
		main()
	} else if amountFloat <= 0 {
		fmt.Println("Masukkan uang yang benar !")
		main()
	}

	fmt.Print("Pilih mata uang yang diinginkan: ")
	var toCurrency string
	fmt.Scanln(&toCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	if toCurrency != "USD" && toCurrency != "EUR" && toCurrency != "GBP" && toCurrency != "JPY" {
		fmt.Println("Error: Masukkan mata uang yang benar !")
		main()
	}
	switch fromCurrency {
	case "USD":
		switch toCurrency {
		case "EUR":
			result := amountFloat * USDToEUR
			fmt.Printf("Konversi berhasil... dari %.2f USD menjadi %.2f EUR\n", amountFloat, result)

		case "GBP":
			result := amountFloat * USDToGBP
			fmt.Printf("Konversi berhasil... dari %.2f USD menjadi %.2f GBP\n", amountFloat, result)

		case "JPY":
			result := amountFloat * USDToJPY
			fmt.Printf("Konversi berhasil... dari %.2f USD menjadi %.2f JPY\n", amountFloat, result)

		default:
			fmt.Println("konversi tidak support !")
		}
	case "EUR":
		switch toCurrency {
		case "USD":
			result := amountFloat * EURToUSD
			fmt.Printf("Konversi berhasil... dari %.2f EUR menjadi %.2f USD\n", amountFloat, result)

		case "GBP":
			result := amountFloat * EURToGBP
			fmt.Printf("Konversi berhasil... dari %.2f EUR menjadi %.2f GBP\n", amountFloat, result)

		case "JPY":
			result := amountFloat * EURToJPY
			fmt.Printf("Konversi berhasil... dari %.2f EUR menjadi %.2f JPY\n", amountFloat, result)

		default:
			fmt.Println("konversi tidak support !")
		}

	case "GBP":
		switch toCurrency {
		case "USD":
			result := amountFloat * GBPToUSD
			fmt.Printf("Konversi berhasil... dari %.2f GBP menjadi %.2f USD\n", amountFloat, result)

		case "EUR":
			result := amountFloat * GBPToEUR
			fmt.Printf("Konversi berhasil... dari %.2f GBP menjadi %.2f EUR\n", amountFloat, result)

		case "JPY":
			result := amountFloat * GBPToJPY
			fmt.Printf("Konversi berhasil... dari %.2f GBP menjadi %.2f JPY\n", amountFloat, result)

		default:
			fmt.Println("konversi tidak support !")
		}

	case "JPY":
		switch toCurrency {
		case "USD":
			result := amountFloat * JPYToUSD
			fmt.Printf("Konversi berhasil... dari %.2f JPY menjadi %.2f USD\n", amountFloat, result)

		case "EUR":
			result := amountFloat * JPYToEUR
			fmt.Printf("Konversi berhasil... dari %.2f JPY menjadi %.2f EUR\n", amountFloat, result)

		case "GBP":
			result := amountFloat * JPYToGBP
			fmt.Printf("Konversi berhasil... dari %.2f JPY menjadi %.2f GBP\n", amountFloat, result)

		default:
			fmt.Println("konversi tidak support !")
		}

	}

	var selesai string
	fmt.Print("Apakah anda ingin melakukan konversi lagi? (y/n) : ")
	fmt.Scanln(&selesai)

	if selesai == "y" || selesai == "Y" {
		main()
	} else if selesai == "n" || selesai == "N" {
		fmt.Println("Terima kasih telah menggunakan program ini")
	} else {
		fmt.Println("Error: input tidak valid !\nProgram berakhir")
	}
}
