# learning-golang
## Install Go
- `wget https://go.dev/dl/go1.26.3.linux-amd64.tar.gz`
- `sudo rm -rf /usr/local/go`
- `sudo tar -C /usr/local -xzf go1.26.3.linux-amd64.tar.gz`
- `export PATH=$PATH:/usr/local/go/bin`
- `go version`

## Create Project
- `go mod init`
- `touch main.go` [main.go](./main.go)
- `go run main.go`
- `go build`
- `./learning-golang`
___

## Variables & Constants
### `var` & `const`
- `mkdir -p 01_variables && touch 01_variables/main.go` [01_variables/main.go](01_variables/main.go)
- `go run 01_variables/main.go`

___

## Data Types
### `Integer` (bilangan bulat)
- `mkdir -p 02_number && touch 02_number/main.go` [02_number/main.go](02_number/main.go)
- `go run 02_number/main.go`

#### `Signed Integer`
- Bisa menyimpan nilai _positif, nol, & negatif_

| Tipe Data | Ukuran | Rentang Nilai |
|-----|-----|-----|
| `int8` | 8 bit | -128 hingga 127 |
| `int16` | 16 bit | -32.768 hingga 32.767 |
| `int32` | 32 bit| -2.147.483.648 hingga 2.147.483.647 |
| `int64` | 64 bit |  ±9,2 x 10¹⁸ (sekitar) |

#### `Unsigned Integer`
- Hanya bisa menyimpan nilai _nol & positif_

| Tipe Data | Ukuran | Rentang Nilai |
|-----|-----|-----|
| `uint8` | 8 bit | 0 hingga 255 |
| `uint16` | 16 bit | 0 hingga 65.535 |
| `uint32` | 32 bit | 0 hingga 4.294.967.295 |
| `uint64` | 64 bit | 0 hingga 18.446.744.073.709.551.615 |

### `float point` (bilangan desimal)
| Tipe Data | Ukuran | Rentang Nilai |
|-----|-----|-----|
| `float32` | 32 bit | 1.18 × 10⁻³⁸ sampai 3.4 × 10³⁸ |
| `float64` | 64 bit | 2.23 × 10⁻³⁰⁸ sampai 1.80 × 10³⁰⁸ |

### `complex numbers` (bilangan kompleks)

| Tipe Data | Ukuran | Rentang Nilai |
|-----|-----|-----|
| `complex64` | 64 bit | (Real 32-bit + Imaginary 32-bit) |
| `complex128` | 128 bit | (Real 64-bit + Imaginary 64-bit) |

### `Boolean` (true/false)
- `mkdir -p 03_boolean && touch 03_boolean/main.go` [03_boolean/main.go](03_boolean/main.go)
- `go run 03_boolean/main.go`

### `String`
- `mkdir -p 04_string && touch 04_string/main.go` [04_string/main.go](04_string/main.go)
- `go run 04_string/main.go`

#### `string operation`
| Operation | Cara | Keterangan |
|-----|-----|-----|
| Menggabungkan string | `s := "Halo" + ", " + "Dunia"` | Menggabungkan string |
| Panjang string | `len("Go") → 2` | Panjang string |
| Akses karakter berdasarkan indeks | `"Go"[0] → 'G'` | Akses karakter berdasarkan indeks |
| Immutable | `"Go"[0] = 'g' → ❌ error` | Immutable |

___

## Type Conversion (Cast)
- `mkdir -p 05_type_conversion && touch 05_type_conversion/main.go` [05_type_conversion/main.go](05_type_conversion/main.go)
- `go run 05_type_conversion/main.go`

___

## Arrays
- `mkdir -p 06_arrays && touch 06_arrays/main.go` [06_arrays/main.go](06_arrays/main.go)
- `go run 06_arrays/main.go`

### Function Arrays
- `mkdir -p 07_function_arrays && touch 07_function_arrays/main.go` [07_function_arrays/main.go](07_function_arrays/main.go)
- `go run 07_function_arrays/main.go`

| Keyword | Deskripsi | Contoh Singkat |
|-----|-----|-----|
| `len(array)` | Menghitung jumlah elemen dalam array | `len(arr) → 5` |
| `range` | Digunakan dalam loop untuk mengambil index dan value | `for i, v := range arr {}` |
| `Index (arr[i])` | Mengambil atau mengubah nilai pada index tertentu | `arr[0] = 10` |
| `==` | Membandingkan dua array, apakah sama panjang dan isinya sama | `arr1 == arr2 → true/false` |
| `!=` | Membandingkan dua array, apakah berbeda | `arr1 != arr2 → true/false` |

___