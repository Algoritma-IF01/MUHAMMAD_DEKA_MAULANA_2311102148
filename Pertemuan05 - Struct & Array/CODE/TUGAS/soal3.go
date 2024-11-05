package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(tab *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan karakter dalam satu baris (akhiri dengan titik '.') : ")
	input, _ := reader.ReadString('\n')

	for i, char := range input {
		if char == '.' || i >= NMAX {
			break
		}
		if char != ' ' && char != '\n' {
			tab[*n] = char
			*n++
		}
	}
}

func cetakArray(tab tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()
}

func balikkanArray(tab *tabel, n int) {
	for i := 0; i < n/2; i++ {
		tab[i], tab[n-i-1] = tab[n-i-1], tab[i]
	}
}

func palindrom(tab tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if tab[i] != tab[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

	balikkanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)
}