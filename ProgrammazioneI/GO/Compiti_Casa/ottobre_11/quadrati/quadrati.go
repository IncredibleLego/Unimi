//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var n, somma int
	fmt.Println("Programma che inserito un valore n calcola la somma di tutti i quadrati da 1 a n")
	fmt.Print("Inserisci il numero n: ")
	fmt.Scan(&n)
	for x := 0; x <= n; x++ {
		somma += x * x
	}
	fmt.Println("La somma dei quadrati da 1 a", n, "è uguale a ", somma)
}
