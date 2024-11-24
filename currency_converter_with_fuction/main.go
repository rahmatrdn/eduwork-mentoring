package main

import (
	"fmt"
	"strings"
)

func konversiMataUang(jumlah float64, mataUangAsal string, mataUangTujuan string, nilaiTukarLama float64, persenPerubahan float64) float64 {

	nilaiTukarBaru := nilaiTukarLama + (nilaiTukarLama * persenPerubahan / 100)

	hasilKonversi := jumlah * nilaiTukarBaru
	return hasilKonversi
}

func main() {
	
	var jumlah float64
	var mataUangAsal, mataUangTujuan string
	var nilaiTukarLama, persenPerubahan float64

	fmt.Println("=== Program Konversi Mata Uang ===")
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan mata uang asal (contoh: USD): ")
	fmt.Scan(&mataUangAsal)

	fmt.Print("Masukkan mata uang tujuan (contoh: IDR): ")
	fmt.Scan(&mataUangTujuan)

	fmt.Print("Masukkan nilai tukar lama (contoh: 14.000 untuk USD ke IDR): ")
	fmt.Scan(&nilaiTukarLama)

	fmt.Print("Masukkan persentase perubahan nilai tukar (contoh: 10 untuk 10%): ")
	fmt.Scan(&persenPerubahan)

	mataUangAsal = strings.ToUpper(mataUangAsal)
	mataUangTujuan = strings.ToUpper(mataUangTujuan)

	hasilKonversi := konversiMataUang(jumlah, mataUangAsal, mataUangTujuan, nilaiTukarLama, persenPerubahan)

	fmt.Printf("\nHasil Konversi:\n")
	fmt.Printf("%.2f %s dengan nilai tukar baru (%.2f%% perubahan) menjadi %.2f %s\n",
		jumlah, mataUangAsal, persenPerubahan, hasilKonversi, mataUangTujuan)
}
