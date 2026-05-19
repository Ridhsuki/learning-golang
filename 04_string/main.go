package main

import (
	"fmt"
	"strings"
)

func main() {
	// Mendeklarasikan variabel string
	name := "Elliot"
	message := "Welcome to our program!"

	// Deklarasi string dengan backtick
	longText := `
	Ini adalah contoh multi-line string di Go.
	Line 2
	Line 3
	`

	// Menampilkan nilai variabel
	fmt.Println("Name:", name)
	fmt.Println("Message:", message)
	fmt.Println("Long Text:", longText)

	// Menggabungkan string
	greeting := message + " " + name
	fmt.Printf("Greeting: %s\n\n", greeting)

	// Operation String
	text := "Hello, World!"

	// Mengubah menjadi huruf kecil
	fmt.Println("Lowercase: ", strings.ToLower(text)) // hello, world!

	// Mengubah menjadi huruf besar
	fmt.Println("Uppercase: ", strings.ToUpper(text)) // HELLO, WORLD!

	// Mengecek apakah string dimulai dengan "Hello"
	fmt.Println(`Starts With "Hello"? `, strings.HasPrefix(text, "Hello")) // true 

	// Mengecek apakah mengandung kata "World"
	fmt.Println(`Contains "World"? `, strings.Contains(text, "World")) // true 

	// Memisahkan string berdasarkan delimiter
	parts := strings.Split("apple,banana,cerry", ",")
	fmt.Println("Split: ", parts) // [apple banana cerry]
	// fmt.Printf("Print asli: %#v\n", parts) //Print asli: []string{"apple", "banana", "cerry"}

	// Mengganti bagian string
	newText := strings.ReplaceAll(text, "World", "Sekai")
	fmt.Println("Replace: ", newText) // Hello, Sekai
}
