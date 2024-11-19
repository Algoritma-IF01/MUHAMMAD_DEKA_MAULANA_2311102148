package main

import (
	"fmt"
)

// definisi tipe array utk data berat balita
type arrBalita [100]float64

// fungsi utk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// fungsi utk menghitung rata-rata berat balita
func hitungRerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	// input jumlah data berat balita
	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scan(&n)

	// input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d : ", i+1)
		fmt.Scan(&berat[i])
	}

	// menghitung minimum, maksimum, dan rata-rata
	hitungMinMax(berat, n, &bMin, &bMax)
	rerata := hitungRerata(berat, n)

	// menampilkan hasil
	fmt.Printf("\nBerat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita : %.2f kg\n", rerata)
}