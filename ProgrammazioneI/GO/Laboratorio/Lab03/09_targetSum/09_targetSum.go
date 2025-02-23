//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("Programma che genera numeri casuali fra 1 e 10 fino a che la loro somma non supera 50 e ne stampa la somma")
	const TARGET = 50
	var n, somma int
	for somma < TARGET{
		n = rand.Intn(10)+1
		fmt.Print(n, " ")
		somma += n
	}
	fmt.Println("\nSomma:", somma)
}