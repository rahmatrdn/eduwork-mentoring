package main

import (
	"fmt"
)

const (
	dollarToRupiah    = 15000
	euroToRupiah      = 17000
	gbpToRupiah       = 20000
	jpyToRupiah       = 104
	maxDollarExchange = 1000
)

func main() {
	var saldoDollar float64
	var currencyType string
	var decision string

	for {
		fmt.Print("Enter amount in dollars: ")
		_, err := fmt.Scan(&saldoDollar)
		if err != nil {
			fmt.Println("Error reading amount:", err)
			continue
		}

		if saldoDollar > maxDollarExchange {
			fmt.Printf("Cannot exchange more than %d dollars\n", maxDollarExchange)
			continue
		}

		fmt.Print("Enter currency type (IDR, EUR, GBP, JPY): ")
		_, err = fmt.Scan(&currencyType)
		if err != nil {
			fmt.Println("Error reading currency type:", err)
			continue
		}

		saldoRupiah, err := exchangeDollarToRupiah(saldoDollar, currencyType)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Printf("Balance in Rupiah: %.2f\n", saldoRupiah)

		fmt.Println("Do you want to perform another conversion? (yes/no)")
		fmt.Scan(&decision)
		if decision != "yes" {
			break
		}
	}
}

func exchangeDollarToRupiah(dollar float64, currencyType string) (rupiah float64, err error) {
	switch currencyType {
	case "IDR":
		rupiah = dollar * dollarToRupiah
	case "EUR":
		rupiah = dollar * euroToRupiah
	case "GBP":
		rupiah = dollar * gbpToRupiah
	case "JPY":
		rupiah = dollar * jpyToRupiah
	default:
		return 0, fmt.Errorf("unsupported currency type %s", currencyType)
	}
	return rupiah, nil
}
