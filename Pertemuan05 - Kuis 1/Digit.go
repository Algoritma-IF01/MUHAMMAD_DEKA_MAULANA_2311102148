package main

import (
	"fmt"
	"strconv"
)

func potongBilangan(n int) (int, int) {
	// konversi bilangan ke string utk mendapatkan panjang digit
	str := strconv.Itoa(n)
	panjang := len(str)

	// hitung posisi tengah
	tengah := panjang / 2

	// jika panjang digit ganjil, bilangan pertama lebih panjang 1 digit
	if panjang%2 != 0 {
		tengah++
	}

	// potong bilangan menjadi dua bagian
	bil1, _ := strconv.Atoi(str[:tengah])
	bil2, _ := strconv.Atoi(str[tengah:])

	return bil1, bil2
}

func main() {
	var n int
	fmt.Print("Masukan bilangan bulat positif : ")
	fmt.Scanln(&n)

	bil1, bil2 := potongBilangan(n)
	fmt.Println("Bilangan pertama : ", bil1)
	fmt.Println("Bilangan kedua : ", bil2)
	fmt.Println("Hasil penjumlahan : ", bil1+bil2)
}