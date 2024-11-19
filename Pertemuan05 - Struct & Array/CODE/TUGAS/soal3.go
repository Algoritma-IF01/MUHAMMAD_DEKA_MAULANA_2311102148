package main

import (
	"bufio"  // Untuk membaca input dari pengguna
	"fmt"    // Untuk mencetak output ke layar
	"os"     // Untuk mengakses sistem file dan input/output standar
)

const NMAX int = 127 // Menentukan ukuran maksimum array yang akan digunakan

type tabel [NMAX]rune // Mendefinisikan tipe data tabel yang berupa array dengan elemen bertipe rune

// Fungsi untuk mengisi array dengan karakter yang dimasukkan pengguna
func isiArray(tab *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin) // Membaca input dari pengguna
	fmt.Print("Masukkan karakter dalam satu baris (akhiri dengan titik '.') : ")
	input, _ := reader.ReadString('\n') // Membaca input pengguna hingga ditemukan baris baru

	// Loop untuk memproses setiap karakter dalam input
	for i, char := range input {
		if char == '.' || i >= NMAX { // Jika karakter adalah titik atau indeks melebihi batas
			break // Menghentikan input
		}
		if char != ' ' && char != '\n' { // Mengabaikan spasi dan karakter baris baru
			tab[*n] = char // Menyimpan karakter dalam array
			*n++           // Menambah indeks n
		}
	}
}

// Fungsi untuk mencetak elemen array
func cetakArray(tab tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", tab[i]) // Mencetak karakter per karakter
	}
	fmt.Println() // Mencetak baris baru setelah mencetak array
}

// Fungsi untuk membalikkan urutan elemen dalam array
func balikkanArray(tab *tabel, n int) {
	for i := 0; i < n/2; i++ { // Mengulang setengah panjang array
		tab[i], tab[n-i-1] = tab[n-i-1], tab[i] // Menukar elemen di posisi i dengan posisi n-i-1
	}
}

// Fungsi untuk mengecek apakah array membentuk palindrom
func palindrom(tab tabel, n int) bool {
	for i := 0; i < n/2; i++ { // Membandingkan elemen dari kedua ujung
		if tab[i] != tab[n-i-1] { // Jika elemen tidak sama, bukan palindrom
			return false
		}
	}
	return true // Jika semua elemen cocok, maka palindrom
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n) // Memanggil fungsi untuk mengisi array dengan input pengguna

	fmt.Print("Teks: ")
	cetakArray(tab, n) // Mencetak teks yang dimasukkan

	if palindrom(tab, n) { // Mengecek apakah teks adalah palindrom
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

	balikkanArray(&tab, n) // Membalikkan urutan array
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n) // Mencetak teks yang dibalik
}
