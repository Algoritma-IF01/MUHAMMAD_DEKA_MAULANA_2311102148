package main

import (
	"fmt"
)

// selectionSortAsc adalah fungsi utk mengurutkan array menggunakan algoritma selection sort dalam urutan menaik.
func selectionSortAsc(arr []int) []int {
	n := len(arr) // menyimpan panjang array utk digunakan dalam perulangan
	for i := 0; i < n-1; i++ {
		minIdx := i // menandai indeks elemen terkecil (pada awal adalah elemen pertama)
		// mencari elemen terkecil dalam array yg belum terurut
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j // memperbarui indeks jika ditemukan elemen yg lebih kecil
			}
		}
		// menukar elemen yg ditemukan dengan elemen pertama yg belum terurut
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr // mengembalikan array yg sudah diurutkan
}

func main() {
	var n int
	fmt.Scan(&n) // membaca jumlah data yg akan diproses

	// membuat slice utk menyimpan hasil pengurutan dari setiap kasus
	results := make([][]int, n)

	// loop utk memproses setiap input
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m) // membaca jumlah elemen dalam kasus ini
		houses := make([]int, m) // membuat slice utk menyimpan data rumah

		// membaca setiap data rumah dan memasukkannya ke dalam slice
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j]) 
		}

		// mengurutkan data rumah dan menyimpannya dalam results
		results[i] = selectionSortAsc(houses)
	}

	// menampilkan hasil pengurutan utk setiap kasus
	for i := 0; i < len(results); i++ {
		sortedHouses := results[i]
		for j := 0; j < len(sortedHouses); j++ {
			// menambahkan spasi antar elemen
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(sortedHouses[j]) // menampilkan elemen yg sudah terurut
		}
		fmt.Println() // menambahkan baris baru setelah menampilkan hasil pengurutan
	}
}
