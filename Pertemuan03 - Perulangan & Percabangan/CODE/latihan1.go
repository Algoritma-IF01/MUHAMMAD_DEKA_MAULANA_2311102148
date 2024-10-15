package main

import "fmt"

func main() {
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d\n ", i)
		fmt.Print("Masukkan warna pertama : ")
		fmt.Scan(&warna1)
		fmt.Print("Masukkan warna kedua : ")
		fmt.Scan(&warna2)
		fmt.Print("Masukkan warna ketiga : ")
		fmt.Scan(&warna3)
		fmt.Print("Masukkan warna keempat : ")
		fmt.Scan(&warna4)

		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
			hasil = false
		}
	}
	println("BERHASIL", hasil)
}