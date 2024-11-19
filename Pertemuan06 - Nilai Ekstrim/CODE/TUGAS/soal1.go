package main

import "fmt"

func main() {
    var n int
    var weights [1000]float64 // array dengan kapasitas 1000 untuk menampung berat kelinci

    fmt.Print("Masukkan jumlah anak kelinci : ")
    fmt.Scan(&n)

    // validasi input utk memastikan jumlah anak kelinci berada dalam rentang 1-1000
    if n <= 0 || n > 1000 {
        fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
        return
    }

    fmt.Println("Masukkan berat anak kelinci : ")
    // membaca input berat kelinci ke dalam array weights
    for i := 0; i < n; i++ {
        fmt.Scan(&weights[i])
    }

    // inisialisasi berat terkecil dan terbesar dengan berat kelinci pertama
    minWeight, maxWeight := weights[0], weights[0]

    // iterasi melalui array utk menemukan berat terkecil dan terbesar
    for i := 1; i < n; i++ {
        if weights[i] < minWeight { // cek apakah berat saat ini lebih kecil dari berat terkecil
            minWeight = weights[i]
        }
        if weights[i] > maxWeight { // cek apakah berat saat ini lebih besar dari berat terbesar
            maxWeight = weights[i]
        }
    }

    // menampilkan hasilny
    fmt.Printf("Berat kelinci terkecil : %.2f\n", minWeight)
    fmt.Printf("Berat kelinci terbesar : %.2f\n", maxWeight)
}
