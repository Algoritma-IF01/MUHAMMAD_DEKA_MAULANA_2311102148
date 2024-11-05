package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Fungsi cekDigit mengecek apakah semua digit dalam angka adalah ganjil, genap, atau semua digitnya sama
func cekDigit(n int) (bool, bool) {
	ganjilSemua := true  // Menandakan apakah semua digit adalah ganjil
	genapSemua := true   // Menandakan apakah semua digit adalah genap
	samaSemua := true    // Menandakan apakah semua digit sama
	digitSebelumnya := n % 10 // Menyimpan digit terakhir sebagai pembanding awal

	for n > 0 {
		digit := n % 10   // Mengambil digit terakhir
		if digit%2 == 0 {
			ganjilSemua = false // Jika digit genap, ganjilSemua jadi false
		} else {
			genapSemua = false  // Jika digit ganjil, genapSemua jadi false
		}
		if digit != digitSebelumnya {
			samaSemua = false   // Jika ada digit berbeda, samaSemua jadi false
		}
		digitSebelumnya = digit // Update digit sebelumnya
		n /= 10                 // Menghapus digit terakhir
	}

	return samaSemua && ganjilSemua, genapSemua // Mengembalikan hasil pengecekan
}

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari stdin
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())  // Membaca jumlah angka yang akan diinputkan

	var esTebak, esCendol, lamang int     // Variabel untuk menghitung hasil kategori

	for i := 0; i < N; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text()) // Membaca angka input

		ganjilSemua, genapSemua := cekDigit(n) // Memeriksa kategori angka
		if ganjilSemua {
			fmt.Println("Es Tebak") // Menampilkan "Es Tebak" jika semua digit ganjil
			esTebak++
		} else if genapSemua {
			fmt.Println("Es Cendol") // Menampilkan "Es Cendol" jika semua digit genap
			esCendol++
		} else {
			fmt.Println("Lamang")    // Menampilkan "Lamang" jika digit campuran
			lamang++
		}
	}

	// Menampilkan jumlah dari setiap kategori
	fmt.Printf("Jumlah Es Tebak: %d\n", esTebak)
	fmt.Printf("Jumlah Es Cendol: %d\n", esCendol)
	fmt.Printf("Jumlah Lamang: %d\n", lamang)
}