package main

import (
	"fmt"
)

func genap(n int) int {
	return n / 2
}

func ganjil(n int) int {
	return 3*n + 1
}

func skiena(n int)  {
	fmt.Println("Urutan Skiena : ", n)
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = genap(n)
		} else {
			n = ganjil(n)
		}
	}
	fmt.Println(1)
}

func main() {
	const nmax = 1000000
	var n int
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&n)

	if n < nmax {
		skiena(n)
	} else {
		fmt.Println("N terlalu besar")
	}
}