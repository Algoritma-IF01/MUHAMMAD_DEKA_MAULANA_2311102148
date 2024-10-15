package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(k)
func f(k float64) float64 {
	pembilang := math.Pow((4*k + 2), 2)
	penyebut := (4*k + 1) * (4*k + 3)
	return pembilang / penyebut
}

func akar2(k int) float64 {
	hasil := 1.0
	for i := 0; i <= k; i++ {
		hasil *= f(float64(i))
	}
	return hasil
}

func main() {
	var K int

	for i := 1; i <= 3; i++ {
		fmt.Print("Nilai K = ")
		fmt.Scan(&K)

		hampiranAkar2 := akar2(K)
		fmt.Printf("Nilai akar 2 = %.10f\n\n", hampiranAkar2)
	}

	fmt.Println("Program selesai!")
}
