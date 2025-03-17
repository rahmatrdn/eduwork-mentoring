package main

import (
	"fmt"
)

// Struct Mahasiswa
type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
}

// Slice untuk menyimpan data mahasiswa
var daftarMahasiswa []Mahasiswa

// Fungsi untuk menambahkan mahasiswa
func tambahMahasiswa(nama, nim, jurusan string) {
	mahasiswa := Mahasiswa{Nama: nama, NIM: nim, Jurusan: jurusan}
	daftarMahasiswa = append(daftarMahasiswa, mahasiswa)
	fmt.Println("Mahasiswa berhasil ditambahkan!\n")
}

// Fungsi untuk menghapus mahasiswa berdasarkan NIM
func hapusMahasiswa(nim string) {
	index := -1
	for i, m := range daftarMahasiswa {
		if m.NIM == nim {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan.\n")
		return
	}

	// Menghapus mahasiswa dari slice
	daftarMahasiswa = append(daftarMahasiswa[:index], daftarMahasiswa[index+1:]...)
	fmt.Println("Mahasiswa dengan NIM", nim, "berhasil dihapus.\n")
}

// Fungsi untuk menampilkan semua data mahasiswa
func tampilkanData() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Belum ada data mahasiswa.\n")
		return
	}

	fmt.Println("Daftar Mahasiswa:")
	for _, m := range daftarMahasiswa {
		fmt.Printf("Nama: %s, NIM: %s, Jurusan: %s\n", m.Nama, m.NIM, m.Jurusan)
	}
	fmt.Println()
}

// Fungsi untuk menampilkan menu
func menu() {
	for {
		fmt.Println("=== Menu Manajemen Mahasiswa ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Hapus Mahasiswa")
		fmt.Println("3. Tampilkan Data Mahasiswa")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var nama, nim, jurusan string
			fmt.Print("Masukkan Nama: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan NIM: ")
			fmt.Scanln(&nim)
			fmt.Print("Masukkan Jurusan: ")
			fmt.Scanln(&jurusan)
			tambahMahasiswa(nama, nim, jurusan)
		case 2:
			var nim string
			fmt.Print("Masukkan NIM Mahasiswa yang akan dihapus: ")
			fmt.Scanln(&nim)
			hapusMahasiswa(nim)
		case 3:
			tampilkanData()
		case 4:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.\n")
		}
	}
}

// Fungsi utama (main)
func main() {
	menu()
}
