package main  // Mendefinisikan package utama

import "fmt"  // Mengimpor package fmt untuk input-output

func main() {  // Fungsi utama program
    var celsius float64  // Mendeklarasikan variabel celsius bertipe float64 untuk menyimpan input suhu dalam derajat Celsius

    fmt.Print("Masukkan suhu dalam derajat Celsius : ")  // Mencetak pesan untuk meminta pengguna memasukkan suhu dalam derajat Celsius
    fmt.Scanln(&celsius)  // Membaca input suhu dari pengguna dan menyimpannya di variabel celsius

    // konversi ke Fahrenheit
    fahrenheit := (celsius * 9 / 5) + 32  // Menghitung suhu dalam derajat Fahrenheit menggunakan rumus konversi dari Celsius

    // konversi ke Reamur
    reamur := celsius * 4 / 5  // Menghitung suhu dalam derajat Reamur menggunakan rumus konversi dari Celsius

    // konversi ke Kelvin
    kelvin := (fahrenheit + 459.67) * 5 / 9  // Menghitung suhu dalam derajat Kelvin menggunakan rumus konversi dari Fahrenheit

    // Mencetak hasil konversi dalam Fahrenheit, Reamur, dan Kelvin
    fmt.Printf("Suhu dalam derajat Fahrenheit : %.2f\n", fahrenheit)  // Menampilkan suhu dalam Fahrenheit dengan dua angka di belakang koma
    fmt.Printf("Suhu dalam derajat Reamur : %.2f\n", reamur)  // Menampilkan suhu dalam Reamur dengan dua angka di belakang koma
    fmt.Printf("Suhu dalam derajat Kelvin : %.2f\n", kelvin)  // Menampilkan suhu dalam Kelvin dengan dua angka di belakang koma
}
