package main

import "fmt"

func main() {
	age := 23

	if age > 18 {
		fmt.Println("Anda Sudah Dewasa")
	} else {
		fmt.Println("anda belum dewasa")
	}

	fmt.Println("============")

	days := 3

	switch days {
	case 1:
		fmt.Println("senin")
	case 2:
		fmt.Println("selasa")
	case 3:
		fmt.Println("rabu")
	case 4:
		fmt.Println("kamis")
	case 5:
		fmt.Println("jumat")
	case 6:
		fmt.Println("sabtu")
	case 7:
		fmt.Println("migggu")
	}

}
