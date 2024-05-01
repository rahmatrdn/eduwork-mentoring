package main

import "fmt"

func cobaFunc(a int) int {
	// cars := []string{"bmw", "volvo", "honda", "daihatsu"}
	// cars2 := []string{"hello", "world"}
	// slice3 := append(cars, cars2...)

	// cars[1] = "toyota"
	// cars = append(cars, "audi", "suzuki")

	// fmt.Println(slice3)
	// fmt.Println(len(cars))
	// fmt.Println(cap(cars))
	// fmt.Println(cars)

	// carsCopy := cars[:len(cars)]
	// copy(carsCopy, cars)
	// fmt.Println(carsCopy)

	// x := 10
	// x += 3
	// fmt.Println(x)

	// a := 5
	// b := 10
	// fmt.Println(a ^ b)

	// for i := 0; i < 50; i += 10 {
	// 	fmt.Println(i)
	// }

	// adj := [2]string{"big", "tasty"}
	// fruits := [3]string{"apple", "banana", "orange"}
	// for i := 0; i < len(adj); i++ {
	// 	for j := 0; j < len(fruits); j++ {
	// 		fmt.Println(adj[i], fruits[j])
	// 	}
	// }

	// fruits := [3]string{"apple", "banana", "orange"}
	// for idx, val := range fruits {
	// 	fmt.Println(idx, val)
	// }

	// for _, val := range fruits {
	// 	fmt.Println(val)
	// }

	// for idx, _ := range fruits {
	// 	fmt.Println(idx)
	// }

	if a == 11 {
		return 12
	}
	fmt.Println(a)
	return cobaFunc(a + 1)
}

func main() {
	cobaFunc(1)
}
