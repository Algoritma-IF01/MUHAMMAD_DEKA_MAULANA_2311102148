package main

import "fmt"

func main (){
	var (
		satu, dua, tiga string
		temp 			string
	)
	fmt.Print("Masukkan Input String : ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan Input String : ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan Input String : ")
	fmt.Scanln(&tiga)
	fmt.Print("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Print(" Output akhir = " + satu + " " + dua + " " + tiga)
}