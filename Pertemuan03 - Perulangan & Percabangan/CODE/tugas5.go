package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		var b int
		fmt.Printf("Bilangan ke-%d : ", i)
		fmt.Scan(&b)

		// Cetak faktor
		fmt.Print("Faktor : ")
		jumlahFaktor := 0
		for j := 1; j <= b; j++ {
			if b%j == 0 {
				fmt.Printf("%d ", j)
				jumlahFaktor++
			}
		}
		fmt.Println()

		// Cek apakah prima
		fmt.Printf("Prima : %t\n\n", jumlahFaktor == 2)
	}

	fmt.Println("Program selesai!")
}
