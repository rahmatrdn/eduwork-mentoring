package main

import (
	"fmt"
)

const (
	dollarToRupiah = 15000
	euroToRupiah   = 17000
	gbpToRupiah    = 20000
	jpyToRupiah    = 104
)

func main() {
	var amount float64
	var sourceCurrency, targetCurrency string
	var percentageChange float64
	var decision string

	for {
		fmt.Print("Enter amount in dollars: ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Error reading amount:", err)
			continue
		}

		fmt.Print("Enter source currency (USD): ")
		_, err = fmt.Scan(&sourceCurrency)
		if err != nil {
			fmt.Println("Error reading source currency:", err)
			continue
		}

		fmt.Print("Enter target currency type (IDR, EUR, GBP, JPY): ")
		_, err = fmt.Scan(&targetCurrency)
		if err != nil {
			fmt.Println("Error reading target currency type:", err)
			continue
		}

		fmt.Print("Enter percentage change in exchange rate: ")
		_, err = fmt.Scan(&percentageChange)
		if err != nil {
			fmt.Println("Error reading percentage change:", err)
			continue
		}

		convertedAmount, err := convertCurrency(amount, sourceCurrency, targetCurrency, percentageChange)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Printf("Converted amount in %s: %.2f\n", targetCurrency, convertedAmount)

		fmt.Println("Do you want to perform another conversion? (yes/no)")
		fmt.Scan(&decision)
		if decision != "yes" {
			break
		}
	}
}

func convertCurrency(amount float64, sourceCurrency, targetCurrency string, percentageChange float64) (convertedAmount float64, err error) {
	baseRates := map[string]float64{
		"USDIDR": dollarToRupiah,
		"USDEUR": euroToRupiah,
		"USDGBP": gbpToRupiah,
		"USDJPY": jpyToRupiah,
	}

	key := sourceCurrency + targetCurrency
	baseRate, ok := baseRates[key]
	if !ok {
		return 0, fmt.Errorf("unsupported currency type %s to %s", sourceCurrency, targetCurrency)
	}

	adjustedRate := baseRate + (baseRate * percentageChange / 100)
	convertedAmount = amount * adjustedRate
	return convertedAmount, nil
}
