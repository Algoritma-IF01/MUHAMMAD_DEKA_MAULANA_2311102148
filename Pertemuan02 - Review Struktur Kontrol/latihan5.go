package main  // Mendefinisikan package utama

import "fmt"  // Mengimpor package fmt untuk input-output

func main() {  // Fungsi utama program
    var a, b, c, d, e int         // Mendeklarasikan lima variabel integer untuk menyimpan lima angka
    var char1, char2, char3 rune  // Mendeklarasikan tiga variabel rune untuk menyimpan tiga karakter

    fmt.Print("Masukkan 5 angka : ")  // Mencetak pesan untuk meminta pengguna memasukkan lima angka
    fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)  // Membaca lima angka dari input pengguna dan menyimpannya di variabel a, b, c, d, e

    fmt.Scanln()  // Membaca input baru untuk menangani sisa karakter seperti newline dari input angka sebelumnya
    fmt.Print("Masukkan 3 karakter : ")  // Mencetak pesan untuk meminta pengguna memasukkan tiga karakter
    fmt.Scanf("%c%c%c", &char1, &char2, &char3)  // Membaca tiga karakter dari input pengguna dan menyimpannya di variabel char1, char2, char3

    fmt.Println()  // Menambahkan baris kosong untuk estetika
    fmt.Print("Output dari 5 angka = ")  // Mencetak pesan "Output dari 5 angka = "
    fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)  // Mengonversi lima angka menjadi karakter ASCII dan mencetaknya

    fmt.Print("Output dari 3 karakter = ")  // Mencetak pesan "Output dari 3 karakter = "
    fmt.Printf("%c%c%c\n", char1+1, char2+1, char3+1)  // Mencetak tiga karakter yang di-input oleh pengguna, tetapi setiap karakter ditambahkan 1 (karakter berikutnya dalam tabel ASCII)
}
