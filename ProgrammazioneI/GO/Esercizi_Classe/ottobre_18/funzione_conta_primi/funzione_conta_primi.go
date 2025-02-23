//Autore: Francesco Corrado

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	n = 2
	fmt.Println("Programma che stampa i primi n numeri primi utilizzando una funzione")
	for {
		x1 := float64(countPrime(n))
		x2 := asintotico(n)
		fmt.Println(n, x1/x2)
		n++
	}
	//	fmt.Println("")
}
func isPrime(x int) bool {
	//Dato un numero x la funzione restituisce true o false a seconda che x sia primo o no
	var d int
	for d = 2; d < x; d++ {
		if x%d == 0 {
			return false
		}
	}
	return true
}

func countPrime(n int) int {
	//Conta il numero di numeri primi
	var d, c int
	for d = 2; d < n; d++ {
		if isPrime(d) {
			c++
		}
	}
	return c
}

func asintotico(n int) float64 {
	return float64(n) / math.Log(float64(n))
}
