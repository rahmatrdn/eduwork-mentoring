package main

import (
	"fmt"
	"log"
)

var currencies = []string{"USD", "EUR", "GBP", "JPY", "IDR"}

type ExchangeRate struct {
	From string
	To   string
	Rate float64
}

var rates = []ExchangeRate{
	//{"USD", "IDR", 15000},
	//{"EUR", "IDR", 17000},
	//{"GBP", "IDR", 20000},
	//{"JPY", "IDR", 104},
	{"IDR", "USD", 1 / 15000.0},
	{"IDR", "EUR", 1 / 17000.0},
	{"IDR", "GBP", 1 / 20000.0},
	{"IDR", "JPY", 1 / 104.0},
}

var rateMap = make(map[string]map[string]float64)

func init() {
	for _, rate := range rates {
		if rateMap[rate.From] == nil {
			rateMap[rate.From] = make(map[string]float64)
		}
		rateMap[rate.From][rate.To] = rate.Rate
	}
}
func convertCurrency(amount float64, from, to string) (float64, error) {
	if fromRate, ok := rateMap[from]; ok {
		if rate, ok := fromRate[to]; ok {
			return amount * rate, nil
		}
		return 0, fmt.Errorf("no rate available from %s to %s", from, to)
	}
	return 0, fmt.Errorf("currency %s not supported", from)
}

func main() {
	var amount float64
	var from, to string

	fmt.Println("Enter the amount to convert:")
	fmt.Scan(&amount)
	fmt.Println("Enter the source currency (e.g., USD):")
	fmt.Scan(&from)
	fmt.Println("Enter the target currency (e.g., IDR):")
	fmt.Scan(&to)

	result, err := convertCurrency(amount, from, to)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, from, result, to)
}
