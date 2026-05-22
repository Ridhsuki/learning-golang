package main

import "fmt"

func main() {
	number := [5]int{10, 20, 30, 40, 50}

	//menampilkan jumlah elemen array
	fmt.Println("Jumlah elemen array : ", len(number))

	//mengakses array index ke-1
	fmt.Println("Index ke-1 :", number[1])

	//mengubah nilai array index ke-1
	number[1] = 200
	fmt.Println("Update index ke-1 :", number[1])

	//menampilkan semua array menggunakan perulangan for
	for index, value := range number {
		fmt.Println("Isi index ke :", index, " = ", value)
	}

	//membandingkan array
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	//membandingkan array apakah sama panjang dan isinya sama
	fmt.Println("arr1 == arr2 :", arr1 == arr2)

	//membandingkan array apakah sama panjang dan isinya berbeda
	fmt.Println("arr1 != arr2 :", arr1 != arr2)
}
