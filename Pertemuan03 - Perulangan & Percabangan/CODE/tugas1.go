package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64
	oleng := false

	for {
		// meminta pengguna utk memasukkan berat belanjaan di kedua kantong
		fmt.Print("Masukkan berat belanjaan di kedua kantong : ")
		fmt.Scan(&kantong1, &kantong2)

		// jika total berat melebihi 150 unit, tampilkan pesan dan akhiri program
		if kantong1+kantong2 > 150 {
			fmt.Println("Berat melebihi 150")
			fmt.Println("Proses selesai.")
			break
		}

		// jika salah satu dari berat bernilai negatif, akhiri program
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// menghitung selisih absolut antara berat di kedua kantong
		selisih := math.Abs(kantong1 - kantong2)

		// menentukan apakah sepeda motor akan oleng
		// jika selisih lebih dari 9 unit, set `oleng` menjadi true
		oleng = selisih > 9

		// mencetak apakah sepeda motor Pak Andi akan oleng berdasarkan selisih berat di kedua kantong
		fmt.Printf("Sepeda motor pak Andi akan oleng : %v\n", oleng)
	}
}
