package main

import "fmt"

func main() {
	fmt.Println("Benar : ", true)
	fmt.Println("Salah : ", false)

	// Deklarasi variabel boolean
	var isRaining bool = true
	var isSunny bool = false
	isLoggedIn := true
	hasPermission := false

	// Menampilkan nilai boolean
	fmt.Println("Apakah hujan?", isRaining)
	fmt.Println("Apakah cerah?", isSunny)
	fmt.Println("Apakah sudah login?", isLoggedIn)
	fmt.Println("Apakah punya akses?", hasPermission)
}
