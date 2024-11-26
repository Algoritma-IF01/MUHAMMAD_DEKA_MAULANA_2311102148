package main

import "fmt"

const nMax = 7919

// struct Buku
type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

// struct DaftarBuku
type DaftarBuku struct {
	Pustaka  []Buku
	nPustaka int
}

// fungsi utk DaftarkanBuku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d (ID Judul Penulis Penerbit Eksemplar Tahun Rating):\n", i+1)
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		pustaka.Pustaka = append(pustaka.Pustaka, buku)
	}
	pustaka.nPustaka = n
}

// fungsi utk CetakTerfavorit
func CetakTerfavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}
	terfavorit := pustaka.Pustaka[0]
	for i := 1; i < len(pustaka.Pustaka); i++ {
		buku := pustaka.Pustaka[i]
		if buku.Rating > terfavorit.Rating {
			terfavorit = buku
		}
	}
	fmt.Printf("Buku terfavorit: %s oleh %s (%s, %d) - Rating: %d\n", terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun, terfavorit.Rating)
}

// fungsi utk UrutBuku
func UrutBuku(pustaka *DaftarBuku) {
	// menggunakan Insertion Sort utk mengurutkan berdasarkan rating
	for i := 1; i < len(pustaka.Pustaka); i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.Pustaka[j].Rating < key.Rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		pustaka.Pustaka[j+1] = key
	}
}

// fungsi utk Cetak5Terbaru
func Cetak5Terbaru(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("%s oleh %s (%s, %d) - Rating: %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

// fungsi utk mencari buku berdasarkan rating
func CariBuku(pustaka DaftarBuku, rating int) {
	// menggunakan Binary Search utk mencari rating
	low, high := 0, len(pustaka.Pustaka)-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka.Pustaka[mid].Rating == rating {
			buku := pustaka.Pustaka[mid]
			fmt.Printf("Buku ditemukan : %s oleh %s (%s, %d) - Rating : %d, Eksemplar : %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating, buku.Eksemplar)
			return
		}
		if pustaka.Pustaka[mid].Rating < rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku

	var n int
	fmt.Print("Masukkan jumlah buku : ")
	fmt.Scan(&n)

	// daftarkan buku
	DaftarkanBuku(&pustaka, n)

	// cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// cari buku berdasarkan rating
	var rating int
	fmt.Print("Masukkan rating yang ingin dicari : ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}