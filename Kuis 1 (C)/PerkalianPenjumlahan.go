package main

import (
    "fmt"
)

// fungsi rekursif untuk menghitung perkalian
func multiply(n int, m int) int {
    // kasus dasar : jika m adalah 0, hasilnya adalah 0
    if m == 0 {
        return 0
    }
    // kasus rekursif : n ditambahkan dengan hasil perkalian n dan m-1
    return n + multiply(n, m-1)
}

func main() {
    var n, m int
    fmt.Print("Masukkan bilangan bulat n: ")
    fmt.Scan(&n)
    fmt.Print("Masukkan bilangan bulat m: ")
    fmt.Scan(&m)

    hasil := multiply(n, m)

    fmt.Printf("Hasil dari %d x %d adalah: %d\n", n, m, hasil)
}