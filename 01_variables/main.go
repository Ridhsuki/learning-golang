package main

import "fmt"

func main() {
	// 1. Deklarasi dengan kata kunci var
	var name string = "Ahmad"
	var age int = 35

	// 2. Deklarasi dengan :=
	city := "New York"
	year := 2027

	// 3. Konstanta
	const pi = 3.14

	fmt.Println("Name : ", name)
	fmt.Println("Age  : ", age)
	fmt.Println("City : ", city)
	fmt.Println("Year : ", year)
	fmt.Println("PI   : ", pi)
	// pi := 3.1 // Ini akan Error
}
