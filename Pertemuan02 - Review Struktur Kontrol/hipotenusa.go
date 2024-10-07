package main

import "fmt"

func main (){
	var a, b, c float64
	var hipotenusa bool
	fmt.Print("Masukkan nilai A ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai B ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai C ")
	fmt.Scanln(&c)
	
	hipotenusa = (c * c) == (a*a + b*b)
	fmt.Println("Sisi c adalah hipotenusa segitiga a, b, dan c : ", hipotenusa)
}