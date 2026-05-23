## Slice
- `mkdir -p 08_slice && touch 08_slice/main.go` [08_slice/main.go](08_slice/main.go)
- `go run 08_slice/main.go`

| Cara Membuat Slice | Penjelasan | Contoh Kode |
|-----|-----|-----|
| `Slice seluruh array` | Mengambil semua elemen array | `s := arr[:]` |
| `Slice dari indeks awal ke tertentu` | Mengambil elemen dari indeks `0` sampai `2` (tidak termasuk `3`) | `s := arr[:3]` |
| `Slice dari indeks tertentu ke akhir` | Mengambil elemen dari indeks `2` sampai akhir array | `s := arr[2:]` |
| `Slice dari indeks i ke j` | Mengambil elemen dari indeks `1` sampai `3` (tidak termasuk `4`) | `s := arr[1:4]` |


```swift
Hubungan Array dan Slice
────────────────────────────────────────────

Array (panjang tetap, elemen tersimpan langsung di memori)

+----+----+----+----+----+
| 10 | 20 | 30 | 40 | 50 |
+----+----+----+----+----+
   0    1    2    3    4   ← indeks array


Slice (punya pointer, length, capacity)

s := arr[1:4]

Slice: [20 30 40]


+----------------------------------+
| Pointer  ───────────────► arr[1] |
|                     (nilai 20)   |
|                                  |
| Length   = 3  → (20,30,40)       |
| Capacity = 4  → (20,30,40,50)    |
+----------------------------------+
```

### Function Slice

| Keyword | Deskripsi | Contoh Singkat |
|-----|-----|-----|
| `len(slice)` | Mengembalikan panjang slice (jumlah elemen yang ada) | `len([]int{1,2,3}) → 3` |
| `cap(slice)` | Mengembalikan kapasitas slice (jumlah elemen maksimum sebelum realokasi) | `cap(make([]int, 3, 5)) → 5` |
| `append(slice, elem...)` | Menambahkan elemen baru ke slice, jika penuh akan membuat array baru | `append([]int{1,2}, 3, 4) → [1 2 3 4]` |
| `copy(dst, src)` | Menyalin elemen dari slice `src` ke `dst`, lalu mengembalikan jumlah elemen yang tersalin | `copy([]int{0,0,0}, []int{1,2}) → 2` |
| `make([]T, len, cap)` | Membuat slice baru dengan panjang `len` dan kapasitas `cap` | `make([]int, 3, 5) → [0 0 0]` |
| `s[i:j]` | Membuat slice baru dari indeks `i` hingga `j-1` dari slice/array asal | `s := []int{1,2,3,4}; s[1:3] → [2 3]` |
