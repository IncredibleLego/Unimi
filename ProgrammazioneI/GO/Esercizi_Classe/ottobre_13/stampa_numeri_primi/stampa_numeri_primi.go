//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var n, stampati, c int
	fmt.Println("Programma che dato un numero stampa i primi n numeri primi")
	fmt.Print("Inserire il numero: ")
	fmt.Scan(&n)
	for d := 2; stampati < n; d++ {
		for c = 2; c < d; c++ {
			if d%c == 0 {
				break
			}
		}
		if c == d {
			fmt.Print(d, " ")
			stampati++
		}
	}
	fmt.Println(" ")
}
