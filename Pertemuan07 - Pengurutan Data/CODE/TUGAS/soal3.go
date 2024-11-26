package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fungsi utk mengurutkan array menggunakan insertion sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// fungsi utk memeriksa apakah jarak antar elemen tetap
func cekJarakTetap(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}
	jarak := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != jarak {
			return false, 0
		}
	}
	return true, jarak
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan data (diakhiri dengan bilangan negatif) : ")
	scanner.Scan()
	input := scanner.Text()

	angkaStr := strings.Fields(input)
	var data []int

	// menggunakan for i utk iterasi
	for i := 0; i < len(angkaStr); i++ {
		num, err := strconv.Atoi(angkaStr[i])
		if err != nil {
			fmt.Println("Input tidak valid!")
			return
		}
		if num < 0 {
			break
		}
		data = append(data, num)
	}

	// mengurutkan array menggunakan insertion sort
	insertionSort(data)

	// memeriksa apakah jarak tetap
	isTetap, jarak := cekJarakTetap(data)

	// menampilkan hasilnya
	fmt.Println("\nKeluaran : ")
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()
	if isTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}