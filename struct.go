package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	Nim     string
	Jurusan string
}

func tambahMahasiswa(mahasiswa Mahasiswa, listMahasiswa *[]Mahasiswa) {
	*listMahasiswa = append(*listMahasiswa, mahasiswa)
}

func hapusMahasiswa(nim string, listMahasiswa *[]Mahasiswa) {
	for i, mhs := range *listMahasiswa {
		if mhs.Nim == nim {
			*listMahasiswa = append((*listMahasiswa)[:i], (*listMahasiswa)[i+1:]...)
			fmt.Println("Mahasiswa dengan NIM", nim, "telah dihapus.")
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan")
}

func tampilkanData(listMahasiswa []Mahasiswa) {
	fmt.Println("Data Mahasiswa:")
	for _, mhs := range listMahasiswa {
		fmt.Println("Nama:", mhs.Nama)
		fmt.Println("NIM", mhs.Nim)
		fmt.Println("Jurusan:", mhs.Jurusan)

	}
}

func main() {
	var listMahasiswa []Mahasiswa

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Hapus Mahasisa")
		fmt.Println("3. Tampilkan Data")
		fmt.Println("4. Keluar")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var name, nim, jurusan string
			fmt.Println("Tambah Mahasiswa")
			fmt.Print("Name: ")
			fmt.Scanln(&name)
			fmt.Print("NIM: ")
			fmt.Scanln(&nim)
			fmt.Print("Jurusan: ")
			fmt.Scanln(&jurusan)
			tambahMahasiswa(Mahasiswa{Nama: name, Nim: nim, Jurusan: jurusan}, &listMahasiswa)
		case 2:
			var nim string
			fmt.Println("Hapus Mahasiswa")
			fmt.Print("NIM Mahasiswa yang akan dihapus: ")
			fmt.Scanln(&nim)
			hapusMahasiswa(nim, &listMahasiswa)
		case 3:
			tampilkanData(listMahasiswa)
		case 4:
			fmt.Println("Thank You !")
			return
		default:
			fmt.Println("PILIH YANG ADA-ADA AJA !!!")
		}
	}
}
