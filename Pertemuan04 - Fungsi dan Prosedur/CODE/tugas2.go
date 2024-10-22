package main

import (
	"fmt"
	"math"
)

// Menghitung jarak antara dua titik
func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x2-x1, y2-y1)
}

// Mengecek apakah titik ada di dalam lingkaran
func titikDiDalamLingkaran(x, y, cx, cy, r float64) bool {
	return hitungJarak(x, y, cx, cy) <= r
}

func main() {
	var x1, y1, r1, x2, y2, r2, x, y float64

	fmt.Scan(&x1, &y1, &r1) // Lingkaran 1
	fmt.Scan(&x2, &y2, &r2) // Lingkaran 2
	fmt.Scan(&x, &y)        // Titik yang akan dicek

	// Cek apakah titik berada dalam lingkaran 1 atau 2
	diDalam1 := titikDiDalamLingkaran(x, y, x1, y1, r1)
	diDalam2 := titikDiDalamLingkaran(x, y, x2, y2, r2)

	// Tentukan output berdasarkan posisi titik
	switch {
	case diDalam1 && diDalam2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case diDalam1:
		fmt.Println("Titik di dalam lingkaran 1")
	case diDalam2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
