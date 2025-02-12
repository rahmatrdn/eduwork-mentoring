package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Struct untuk menyimpan nilai tukar mata uang
type ExchangeRate struct {
	BaseCurrency   string
	TargetCurrency string
	Rate           float64
}

// Map untuk menyimpan nilai tukar mata uang
var exchangeRates = map[string]map[string]float64{
	"USD": {"EUR": 0.92, "GBP": 0.78, "JPY": 130.0},
	"EUR": {"USD": 1.09, "GBP": 0.85, "JPY": 142.0},
	"GBP": {"USD": 1.28, "EUR": 1.18, "JPY": 167.0},
	"JPY": {"USD": 0.0077, "EUR": 0.007, "GBP": 0.006},
}

// Mutex untuk mengamankan akses paralel ke data nilai tukar
var mutex = &sync.Mutex{}

// Fungsi untuk konversi mata uang
func convertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if rates, ok := exchangeRates[fromCurrency]; ok {
		if rate, ok := rates[toCurrency]; ok {
			return amount * rate, nil
		}
		return 0, errors.New("nilai tukar untuk mata uang tujuan tidak tersedia")
	}
	return 0, errors.New("mata uang asal tidak ditemukan")
}

// Fungsi untuk memperbarui nilai tukar (simulasi layanan eksternal)
func updateExchangeRates() {
	mutex.Lock()
	defer mutex.Unlock()

	// Simulasi pembaruan nilai tukar
	exchangeRates["USD"]["EUR"] = 0.93
	exchangeRates["USD"]["GBP"] = 0.79
	fmt.Println("Nilai tukar diperbarui!")
}

// Fungsi untuk menjalankan pembaruan otomatis
func startAutoUpdate(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			updateExchangeRates()
		}
	}()
}

// Fungsi utama dengan antarmuka pengguna
func main() {
	// Mulai pembaruan otomatis setiap 1 jam
	startAutoUpdate(1 * time.Hour)

	fmt.Println("Aplikasi Konversi Mata Uang")
	fmt.Println("==========================")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Konversi Mata Uang")
		fmt.Println("2. Keluar")
		fmt.Print("Pilihan Anda: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Print("Masukkan jumlah uang: ")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Print("Masukkan mata uang asal (contoh: USD): ")
			var fromCurrency string
			fmt.Scanln(&fromCurrency)

			fmt.Print("Masukkan mata uang tujuan (contoh: EUR): ")
			var toCurrency string
			fmt.Scanln(&toCurrency)

			result, err := convertCurrency(amount, fromCurrency, toCurrency)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Hasil konversi: %.2f %s\n", result, toCurrency)
			}
		} else if choice == 2 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
