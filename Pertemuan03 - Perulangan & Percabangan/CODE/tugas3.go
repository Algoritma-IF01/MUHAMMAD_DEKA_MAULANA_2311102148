package main

import (
	"fmt"
)

// fungsi hitungBiaya menghitung biaya pengiriman berdasarkan berat
func hitungBiaya(berat int) (kg, sisaGram, biayaKg, biayaSisaGram int, biayaDitambahkan bool) {
	const biayaPerKg = 10000
	kg = berat / 1000
	sisaGram = berat % 1000
	biayaKg = kg * biayaPerKg
	biayaDitambahkan = true

	// menggunakan switch utk menentukan apakah biaya sisa gram akan dihitung
	switch {
	case kg >= 10:
		biayaSisaGram = sisaGram * 5
		biayaDitambahkan = false
	case sisaGram >= 500:
		biayaSisaGram = sisaGram * 5
	default:
		biayaSisaGram = sisaGram * 15
	}

	return
}

// fungsi tampilkanBiaya mencetak hasil perhitungan biaya
func tampilkanBiaya(berat int) {
	kg, sisaGram, biayaKg, biayaSisaGram, biayaDitambahkan := hitungBiaya(berat)
	totalBiaya := biayaKg
	if biayaDitambahkan {
		totalBiaya += biayaSisaGram
	}

	// mencetak hasil detail perhitungan
	fmt.Printf("Detail berat : %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaSisaGram)
	fmt.Printf("Total biaya : Rp. %d\n\n", totalBiaya)
}

func main() {
	// loop untuk menerima input dari user sebanyak 3 kali
	for i := 1; i <= 3; i++ {
		var berat int
		fmt.Printf("Contoh #%d :\n", i)
		fmt.Print("Berat parsel (gram) : ")
		fmt.Scan(&berat)
		tampilkanBiaya(berat)
	}
}
