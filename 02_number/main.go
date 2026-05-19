package main

import "fmt"

func main() {
	// Signed integers
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = -100 // Ukuran tergantung arsitektur (32-bit atau 64-bit)

	// Unsigned integers
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 100 // Ukuran tergantung arsitektur

	fmt.Println("Signed Integers:")
	fmt.Printf("int8   : %v\n", i8)
	fmt.Printf("int16  : %v\n", i16)
	fmt.Printf("int32  : %v\n", i32)
	fmt.Printf("int64  : %v\n", i64)
	fmt.Printf("int    : %v\n", i)

	fmt.Println("\nUnsigned Integers:")
	fmt.Printf("uint8  : %v\n", u8)
	fmt.Printf("uint16 : %v\n", u16)
	fmt.Printf("uint32 : %v\n", u32)
	fmt.Printf("uint64 : %v\n", u64)
	fmt.Printf("uint   : %v\n", u)

	// Float point
	var f32 float32 = 3.14
	var f64 float64 = 3.1415926535

	fmt.Println("\nFloat Point:")
	fmt.Printf("float32 : %v\n", f32)
	fmt.Printf("float64 : %v\n", f64)

	// Bilangan kompleks
	var c64 complex64 = 3 + 4i
	var c128 complex128 = 3 + 4i

	fmt.Println("\nBilangan Kompleks:")
	fmt.Printf("complex64 : %v\n", c64)
	fmt.Printf("complex128 : %v\n", c128)

	// Alias
	fmt.Println("Satu =", 1)
}
