//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("Numeri casuali generati tra [1,10]:")
	const K = 10
	var n, somma int
	for i:=1; i<=K; i++{
		n = rand.Intn(10)+1
		fmt.Print(n, " ")
		somma = somma + n
	}
	fmt.Println("\nLa somma di tutti i valori generati è uguale a: ", somma)
}