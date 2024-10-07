package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var char1, char2, char3 rune

	fmt.Print("Masukkan 5 angka : ")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	fmt.Scanln()
	fmt.Print("Masukkan 3 karakter : ")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	fmt.Println()
	fmt.Print("Output dari 5 angka = ")
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Print("Output dari 3 karakter = ")
	fmt.Printf("%c%c%c\n", char1+1, char2+1, char3+1)
}