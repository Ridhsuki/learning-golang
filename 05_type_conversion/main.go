package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Number to Number
	var a int = 10
	var b float64 = float64(a) // int to float64

	fmt.Println("Nilai a:", a)
	fmt.Println("Nilai b:", b)

	// Number to String
	var score int = 95
	var text string = strconv.Itoa(score) // int to string

	fmt.Println("Nilai ujian:", text)

	// String to Number
	var price string = "100"
	number, err := strconv.Atoi(price) // string to int
	if err != nil {
		fmt.Println("Pesan Error: ", err.Error())
	} else {
		fmt.Println("Angka: ", number)
	}

	// Bool to String
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean to String:", str)

	// String to Bool
	val, _ := strconv.ParseBool("true")
	fmt.Println("String to Boolean:", val)
}
