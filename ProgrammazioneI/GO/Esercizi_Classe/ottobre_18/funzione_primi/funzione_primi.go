//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var n, stampati int
	fmt.Println("Programma che dato un numero stampa i primi n numeri primi utilizzando una funzione")
	fmt.Print("Inserire il numero: ")
	fmt.Scan(&n)
	for d := 2; stampati < n; d++ {
		if isPrime(d) {
			fmt.Print(d, " ")
			stampati++
		}
	}
	fmt.Println("")
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
