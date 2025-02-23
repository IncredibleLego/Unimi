//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Programma che inserito un numero procede a calcolare la somma di tutti i numeri da 1 a esso tramite la formula di Gauss")
	fmt.Print("Inserire il numero (intero): ")
	fmt.Scan(&n)
	somma := (n * (n + 1)) / 2
	fmt.Println("La somma di tutti i numeri da 1 a", n, "è uguale a", somma)
}
