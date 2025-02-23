//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var n, c int
	fmt.Println("Programma che stampa tutti i numeri primi minori di n")
	fmt.Print("Inserisci n: ")
	fmt.Scan(&n)
	fmt.Println("I numeri primi da 2 fino a", n, "sono:")
	for d := 2; d <= n; d++ {
		for c = 2; c < d; c++ {
			if d%c == 0 {
				break
			}
		}
		if c == d {
			fmt.Print(d, " ")
		}
	}
	fmt.Println(" ")
}
