package main

import (
	"fmt"
	"strings"
)

func main() {
	// Nama mata uang dengan array
	currencies := [...]string{"USD", "EUR", "GBP", "JPY"}

	// Nilai tukar mata uang dengan slice
	exchangeRates := []float64{1, 0.85, 0.75, 110.0}

	// Mengaitkan antara array dan slice dengan Map
	currencyMap := make(map[string]float64)
	for i, currency := range currencies {
		currencyMap[currency] = exchangeRates[i]
	}

	// Input user: jumlah uang, mata uang asal, dan mata uang tujuan
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&amount)

	fmt.Print("Masukkan mata uang asal: ")
	fmt.Scan(&fromCurrency)
	fromCurrency = strings.ToUpper(fromCurrency)

	fmt.Print("Masukkan mata uang tujuan: ")
	fmt.Scan(&toCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	// Memanggil fungsi konversi mata uang
	convertedAmount := convertCurrency(amount, fromCurrency, toCurrency, currencyMap)

	// Output hasil konversi
	fmt.Printf("konversi berhasil dari %.2f %s menjadi %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)

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

// Membuat fungsi untuk melakukan konversi mata uang
func convertCurrency(amount float64, fromCurrency, toCurrency string, currencyMap map[string]float64) float64 {
	fromRate := currencyMap[fromCurrency]
	toRate := currencyMap[toCurrency]

	convertedAmount := (amount / fromRate) * toRate
	return convertedAmount
}
