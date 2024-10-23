package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cekDigit(n int) (bool, bool) {
	ganjilSemua := true
	genapSemua := true
	samaSemua := true
	digitSebelumnya := n % 10

	for n > 0 {
		digit := n % 10
		if digit%2 == 0 {
			ganjilSemua = false
		} else {
			genapSemua = false
		}
		if digit != digitSebelumnya {
			samaSemua = false
		}
		digitSebelumnya = digit
		n /= 10
	}

	return samaSemua && ganjilSemua, genapSemua
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	var esTebak, esCendol, lamang int

	for i := 0; i < N; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		ganjilSemua, genapSemua := cekDigit(n)
		if ganjilSemua {
			fmt.Println("Es Tebak")
			esTebak++
		} else if genapSemua {
			fmt.Println("Es Cendol")
			esCendol++
		} else {
			fmt.Println("Lamang")
			lamang++
		}
	}

	fmt.Printf("Jumlah Es Tebak: %d\n", esTebak)
	fmt.Printf("Jumlah Es Cendol: %d\n", esCendol)
	fmt.Printf("Jumlah Lamang: %d\n", lamang)
}