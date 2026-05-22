package main

import "fmt"

func main() {
	var angka [5]int = [5]int{10,20,30,40,50}

	fmt.Println("array:", angka)
	fmt.Println("element index 0:", angka[0])


	angka[0] = 100
	fmt.Println("element index 0 sebelumnya 10 menjadi", angka[0])
}
