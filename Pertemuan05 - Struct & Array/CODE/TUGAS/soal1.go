package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N) : ")
	fmt.Scanln(&n)

	array := make([]int, n) // membuat array dengan panjang n
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d : ", i) // minta input utk setiap elemen array
		fmt.Scanln(&array[i]) // membaca elemen array
	}

	var choice, x, index int
	for {
		// menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Menampilkan keseluruhan isi dari array.")
		fmt.Println("2. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
		fmt.Println("3. Menampilkan elemen-elemen array dengan indeks genap saja.")
		fmt.Println("4. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x.")
		fmt.Println("5. Menghapus elemen array pada indeks tertentu.")
		fmt.Println("6. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
		fmt.Println("7. Menampilkan standar deviasi dari bilangan yang ada di dalam array.")
		fmt.Println("8. Menampilkan frekuensi dari setiap bilangan tertentu di dalam array.")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// menampilkan semua elemen array
			fmt.Println("Isi array :", array)

		case 2:
			// menampilkan elemen dengan indeks ganjil
			fmt.Print("Elemen dengan indeks ganjil : ")
			for i := 1; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 3:
			// menampilkan elemen dengan indeks genap
			fmt.Print("Elemen dengan indeks genap : ")
			for i := 0; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 4:
			fmt.Print("Masukkan nilai x : ")
			fmt.Scanln(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d : ", x)
			for i := x; i < len(array); i += x { // iterasi elemen dengan indeks kelipatan x
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 5:
			// menghapus elemen pada indeks tertentu
			fmt.Print("Masukkan indeks yang ingin dihapus : ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(array) { // validasi indeks
				array = append(array[:index], array[index+1:]...) // menghapus elemen dengan menggunakan slicing
				fmt.Println("Isi array setelah penghapusan :", array)
			} else {
				fmt.Println("Indeks tidak valid!") // pesan jika indeks diluar jangkauan
			}

		case 6:
			// menghitung rata-rata elemen array
			total := 0
			for _, value := range array {
				total += value
			}
			avg := float64(total) / float64(len(array)) // rata-rata dihitung sebagai total dibagi jumlah elemen
			fmt.Printf("Rata-rata dari array : %.2f\n", avg)

		case 7:
			// menghitung standar deviasi
			total := 0
			for _, value := range array {
				total += value
			}
			mean := float64(total) / float64(len(array))
			variance := 0.0
			for _, value := range array {
				variance += math.Pow(float64(value)-mean, 2)
			}
			stdDev := math.Sqrt(variance / float64(len(array)))
			fmt.Printf("Standar deviasi dari array : %.2f\n", stdDev)

		case 8:
			// menampilkan frekuensi setiap bilangan
			frequency := make(map[int]int) // map untuk menyimpan frekuensi elemen
			for _, value := range array {
				frequency[value]++ // increment frekuensi untuk setiap elemen
			}
			fmt.Println("Frekuensi setiap bilangan dalam array : ")
			for num, freq := range frequency {
				fmt.Printf("%d : %d kali\n", num, freq)
			}

		case 0:
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
