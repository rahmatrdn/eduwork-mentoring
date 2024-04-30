package main

import (
	"fmt"
	"strings"
)

func convertCurrency(amount float64, fromCurrency string, toCurrency string, percentageChange float64) float64 {
	exchangeRates := map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"GBP": 0.73,
		"JPY": 110.67,
	}

	fromRate := exchangeRates[fromCurrency]
	toRate := exchangeRates[toCurrency]

	newToRate := toRate + (toRate * percentageChange / 100)

	result := amount * (fromRate / newToRate)
	return result
}

func main() {
	var amount float64
	var fromCurrency, toCurrency string
	var percentageChange float64

	fmt.Print("Masukkan mata uang asal (misalnya USD, EUR, GBP, JPY, CNY) : ")
	fmt.Scanln(&fromCurrency)
	fromCurrency = strings.ToUpper(fromCurrency)

	if fromCurrency != "USD" && fromCurrency != "EUR" && fromCurrency != "GBP" && fromCurrency != "JPY" {
		fmt.Println("Input tidak valid !")
		main()
	}

	fmt.Print("Masukkan jumlah uang : ")
	fmt.Scanln(&amount)

	fmt.Print("Masukkan mata uang tujuan (misalnya USD, EUR, GBP, JPY, CNY): ")
	fmt.Scanln(&toCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	if toCurrency != "USD" && toCurrency != "EUR" && toCurrency != "GBP" && toCurrency != "JPY" {
		fmt.Println("Input tidak valid.")
		main()
	}

	fmt.Print("Masukkan persentase perubahan nilai tukar : ")
	fmt.Scanln(&percentageChange)

	result := convertCurrency(amount, fromCurrency, toCurrency, percentageChange)

	fmt.Printf("Konversi berhasil dari %.2f %s ke %s adalah: %.2f %s\n", amount, fromCurrency, toCurrency, result, toCurrency)

	var selesai string
	fmt.Print("Apakah Anda ingin melakukan konversi lagi? (y/n) : ")
	fmt.Scanln(&selesai)

	if selesai == "y" || selesai == "Y" {
		main()
	} else if selesai == "n" || selesai == "N" {
		fmt.Println("Terima kasih telah menggunakan Program ini !")
	} else {
		fmt.Println("Input tidak valid.")
	}
}
