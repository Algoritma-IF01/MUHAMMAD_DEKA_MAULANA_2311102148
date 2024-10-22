package main

import (
	"fmt"
)

// Fungsi rekursif unutk menhitung deret Fibonacci
func fibonacci(n int) int {
	if n == 0{
		return 0
	} else if n == 1{
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	// Menampilkan deret Fibonacci hingga suku ke-10
	fmt.Println("Deret Fibonacci hingga suku ke-10 : ")
	for i := 0; i <= 10; i++{
		fmt.Printf("Fibonacci (%d) = %d\n", i, fibonacci (i))
	}
}