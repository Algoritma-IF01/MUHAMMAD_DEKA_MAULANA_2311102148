package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fungsi selection sort utk mengurutkan slice integer
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// menukar elemen
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var jumlahBaris int
	fmt.Print("Masukkan jumlah baris : ")
	fmt.Scanln(&jumlahBaris)

	scanner := bufio.NewScanner(os.Stdin)
	result := make([][]int, jumlahBaris)

	for baris := 0; baris < jumlahBaris; baris++ {
		fmt.Printf("Masukkan angka untuk baris %d : ", baris+1)
		scanner.Scan()
		input := scanner.Text()

		angkaStr := strings.Fields(input)
		var ganjil []int
		var genap []int

		// memisahkan bilangan ganjil dan genap
		for i := 0; i < len(angkaStr); i++ {
			angka, _ := strconv.Atoi(angkaStr[i])
			if angka%2 == 0 {
				genap = append(genap, angka)
			} else {
				ganjil = append(ganjil, angka)
			}
		}

		// mengurutkan ganjil dan genap menggunakan selectionSort
		selectionSort(ganjil)
		selectionSort(genap)

		// menyimpan hasil (menggabungkan secara manual)
		result[baris] = append(ganjil, genap...)
	}

	// menampilkan hasil
	fmt.Println("\nOutput : ")
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}