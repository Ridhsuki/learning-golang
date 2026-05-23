package main

import "fmt"

func main() {
	// Membuat Slice
	arr := [6]int{10, 20, 30, 40, 50, 60}

	slice1 := arr[:] // ini untuk mengambil seluruh element
	fmt.Println("ini adalah Slice1 :", slice1)

	slice2 := arr[:3] // ini untuk mengambil element dari index 0 sampai 3
	fmt.Println("ini adalah Slice2 :", slice2)

	slice3 := arr[3:] // ini untuk mengambil element dari index 3 sampai akhir
	fmt.Println("ini adalah Slice3 :", slice3)

	slice4 := arr[1:4] // ini untuk mengambil element dari index 1 sampai 4
	fmt.Println("ini adalah Slice4 :", slice4)

	// Function Slice
	// 'make' digunakan untuk membuat slice
	s := make([]int, 3, 5)
	fmt.Printf("Awal: %v | Panjang: %d | Kapasitas: %d\n", s, len(s), cap(s))

	// `append()` Menambahkan element baru ke slice
	s = append(s, 10, 20, 30)
	fmt.Printf("Setelah di Append : %v | Panjang: %d | Kapasitas: %d\n", s, len(s), cap(s))

	// `copy()` Menyalin element dari slice `src` ke `dst`
	s2 := make([]int, 3)
	copy(s2, s)

	fmt.Println("Setelah di copy :", s2)
	fmt.Println("Jumlah element yang di copy", len(s2), "karena kapasitas Slice s2 adalah 3")
}
