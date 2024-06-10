package main

import "fmt"

	const (
		dollarToRupiah = 15000
		maxDollarExchange = 100
	)
	func main()  {
		saldoDolar := 50.0

		fmt.Println("saldo Dollar Awal : ", saldoDolar)

		saldoRupiah := exchangeDollarToRupiah(saldoDolar)

		fmt.Println("saldo Rupiah : ", saldoRupiah)
	}

func exchangeDollarToRupiah(dollar float64) float64 {
	if dollar > maxDollarExchange {
		fmt.Println("maaf tidak bisa menukar lebih dari 100")
		return dollar
	}
	rupiah := dollar * dollarToRupiah

	return rupiah
}