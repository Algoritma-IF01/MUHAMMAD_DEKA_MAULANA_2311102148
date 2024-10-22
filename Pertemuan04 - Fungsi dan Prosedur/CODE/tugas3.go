package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk menghitung jumlah soal yang terjawab (soal) dan total skor (skor)
func hitungSkor(soal *int, skor *int, waktu []int) {
	*soal = 0
	*skor = 0
	
	// Iterasi melalui slice waktu untuk menghitung jumlah soal yang terjawab dalam waktu kurang dari 301 detik
	for _, w := range waktu {
		if w < 301 {
			(*soal)++
			*skor += w
		}
	}
}

func main() {
	// Membuat scanner untuk membaca input dari standar input (biasanya dari keyboard)
	scanner := bufio.NewScanner(os.Stdin)
	
	// Variabel untuk menyimpan nama pemenang, jumlah soal maksimum yang terjawab, dan waktu minimum
	var pemenang string
	var maxSoal, minWaktu int
	
	// Inisialisasi nilai maxSoal dan minWaktu
	maxSoal = 0
	minWaktu = 999999

	for {
		// Meminta pengguna untuk memasukkan nama dan waktu pengerjaan soal
		fmt.Println("Masukkan nama dan waktu (Ketik 'Selesai' untuk berhenti) : ")
		scanner.Scan()
		input := scanner.Text()

		// Jika pengguna mengetik "Selesai", maka program akan berhenti menerima input
		if strings.ToLower(input) == "selesai" {
			break
		}

		// Memisahkan input menjadi potongan-potongan data
		data := strings.Fields(input)
		if len(data) != 9 {
			fmt.Println("Input tidak valid, masukkan nama dan 8 nilai waktu pengerjaan soal.")
			continue
		}

		// Menyimpan nama peserta
		nama := data[0]
		
		// Membuat slice waktu untuk menyimpan delapan waktu pengerjaan soal
		waktu := make([]int, 8)
		for i := 0; i < 8; i++ {
			fmt.Sscanf(data[i+1], "%d", &waktu[i])
		}

		// Variabel untuk menyimpan jumlah soal yang terjawab dan skor
		var soal, skor int
		hitungSkor(&soal, &skor, waktu)

		// Logika untuk menentukan pemenang
		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}

	// Menampilkan hasil akhir
	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
