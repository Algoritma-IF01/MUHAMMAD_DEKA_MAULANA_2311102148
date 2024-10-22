package main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b{
		fmt.Println(permutasi (a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++{
		hasil = hasil * i
	}
	return hasil
}

func permutasi (n, r int) int {
	return faktorial(n) / faktorial(n-r)
}