package main

import (
	"fmt"
)

func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	var numMatches int
	var winners []string

	// meminta pengguna memasukkan nama Klub A
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&clubA)
	
	// meminta pengguna memasukkan nama Klub B
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&clubB)
	
	// meminta jumlah pertandingan yang akan dimainkan
	fmt.Print("Masukkan jumlah pertandingan: ")
	fmt.Scanln(&numMatches)

	// perulangan utk setiap pertandingan
	for i := 1; i <= numMatches; i++ {
		// memasukkan skor utk Klub A
		fmt.Printf("Pertandingan %d (skor %s): ", i, clubA)
		fmt.Scanln(&scoreA)
		
		// memasukkan skor utk Klub B
		fmt.Printf("Pertandingan %d (skor %s): ", i, clubB)
		fmt.Scanln(&scoreB)

		// jika skor negatif dimasukkan, menghentikan input lebih lanjut
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// menentukan pemenang atau hasil imbang
		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
		} else {
			winners = append(winners, "Draw")
		}
	}

	// menampilkan hasil setiap pertandingan
	fmt.Println("\nHasil Pertandingan:")
	for i, winner := range winners {
		if winner == "Draw" {
			fmt.Printf("Hasil %d : Draw\n", i+1)
		} else {
			fmt.Printf("Hasil %d : %s\n", i+1, winner)
		}
	}
	fmt.Println("Pertandingan selesai")
}
