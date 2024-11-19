package main

import "fmt"

func main() {
	// deklarasi variabel untuk jumlah ikan (x) dan kapasitas wadah (y)
	var jumlahIkan, kapasitasWadah int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas per wadah (y) : ")
	fmt.Scan(&jumlahIkan, &kapasitasWadah)

	// membaca input berat ikan menggunakan array kapasitas 1000
	beratIkan := make([]float64, 1000) // kapasitas array 1000
	fmt.Println("Masukkan berat ikan (dipisah dengan spasi) : ")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Scan(&beratIkan[i])
	}

	// membagi ikan ke dalam wadah
	var wadah []float64
	var totalBeratPerWadah []float64
	totalBerat := 0.0
	// membagi ikan ke dalam wadah yang masing-masing berisi hingga kapasitasWadah ikan
	for i := 0; i < jumlahIkan; i++ {
		totalBerat += beratIkan[i]
		// menambahkan berat ikan ke wadah
		wadah = append(wadah, beratIkan[i])

		// jika jumlah ikan dalam wadah mencapai kapasitas, simpan total berat wadah
		if len(wadah) == kapasitasWadah || i == jumlahIkan-1 {
			var sum float64
			for _, berat := range wadah {
				sum += berat
			}
			totalBeratPerWadah = append(totalBeratPerWadah, sum) // menyimpan total berat untuk wadah tersebut
			wadah = []float64{} // reset wadah utk wadah berikutnya
		}
	}

	// menampilkan total berat ikan di setiap wadah
	for _, total := range totalBeratPerWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	// menghitung rata-rata berat ikan di setiap wadah
	rataRata := totalBerat / float64(len(totalBeratPerWadah))
	fmt.Printf("%.2f\n", rataRata)
}
