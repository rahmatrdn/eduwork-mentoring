package main

import "fmt"

const (
	dollarToRupiahRate = 15000 // Rate konversi : 1 Dollar = 15000 Rupiah
	maxExchangeDollar  = 100   // Batas saldo maksimum untuk konversi
)

func main() {

	var saldoDollar float64

	fmt.Println("<==== Selamat datang di program konversi mata uang (USD To Rp) ====>")
	fmt.Print("Masukkan saldo USD yang mau di konversi : ")
	fmt.Scanln(&saldoDollar) // Input user ( saldo yang ingin di konversi )

	// Memanggil fungsi untuk melakukan transaksi konversi Dollar ke rupiah
	saldoRupiah := exchangeDollarToRupiah(saldoDollar)
	if saldoRupiah == saldoDollar {
		fmt.Println("Saldo di balikkan ke awal ->", saldoDollar)
	} else {
		fmt.Printf("Saldo setelah di konversi: %.f\n", saldoRupiah)
	}

	// Memberikan opsi kepada pengguna untuk memilih kembali melakukan konversi atau keluar dari program.
	var selesai string
	fmt.Println("Apakah anda ingin konversi lagi? (y/n)")
	fmt.Scan(&selesai)

	if selesai == "y" || selesai == "Y" {
		main()
	} else if selesai == "n" || selesai == "N" {
		fmt.Println("Terima kasih telah menggunakan program ini !")
	} else {
		fmt.Println("Keywoard tidak valid, program berakhir !")
	}
}

// Membuat fungsi untuk melakukan konversi Dollar ke Rupiah
func exchangeDollarToRupiah(dollar float64) float64 {
	if dollar > maxExchangeDollar {
		fmt.Println("Maksimal saldo untuk di konversi : 100")
		return dollar // Saldo tidak berubah jika melebihi batas konversi
	}

	// Melakukan konversi Dollar ke Rupiah
	rupiah := dollar * dollarToRupiahRate
	return rupiah
}
