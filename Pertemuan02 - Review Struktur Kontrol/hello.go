package main

import "fmt"

func main() {
	var greetings = "Selamat datang di dunia DAP"
	var a, b int
	
	fmt.Println(greetings)
	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}