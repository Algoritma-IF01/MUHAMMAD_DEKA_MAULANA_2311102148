package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// fungsi utk menghitung jumlah soal yang terjawab (soal) dan total skor (skor)
func hitungSkor(soal *int, skor *int, waktu []int) {
	*soal = 0
	*skor = 0
	
	// iterasi melalui slice waktu utk menghitung jumlah soal yang terjawab dlm waktu kurang dari 301 detik
	for _, w := range waktu {
		if w < 301 {
			(*soal)++
			*skor += w
		}
	}
}

func main() {
	// membuat scanner utk membaca input dari standar input
	scanner := bufio.NewScanner(os.Stdin)
	
	// variabel utk menyimpan nama pemenang, jumlah soal maksimum yg terjawab, dan waktu minimum
	var pemenang string
	var maxSoal, minWaktu int
	
	// inisialisasi nilai maxSoal dan minWaktu
	maxSoal = 0
	minWaktu = 999999

	for {
		// meminta pengguna utk memasukkan nama dan waktu pengerjaan soal
		fmt.Println("Masukkan nama dan waktu (Ketik 'Selesai' untuk berhenti) : ")
		scanner.Scan()
		input := scanner.Text()

		// jika pengguna mengetik "Selesai", maka program akan berhenti menerima input
		if strings.ToLower(input) == "selesai" {
			break
		}

		// memisahkan input menjadi potongan-potongan data
		data := strings.Fields(input)
		if len(data) != 9 {
			fmt.Println("Input tidak valid, masukkan nama dan 8 nilai waktu pengerjaan soal.")
			continue
		}

		// menyimpan nama peserta
		nama := data[0]
		
		// membuat slice waktu utk menyimpan 8 waktu pengerjaan soal
		waktu := make([]int, 8)
		for i := 0; i < 8; i++ {
			fmt.Sscanf(data[i+1], "%d", &waktu[i])
		}

		// variabel utk menyimpan jumlah soal yg terjawab dan skor
		var soal, skor int
		hitungSkor(&soal, &skor, waktu)

		// logika utk menentukan pemenang
		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}

	// menampilkan hasil akhir
	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
