package main
import (
	"fmt"
	"math/rand"
)
/*Scrivere un programma che generi casualmente un valore fino a 10, ci aggiunga 1 e stampi true nel caso il valore stampato sia 10, altrimenti false */
/*CORREZIONE 3/01: Scrivere un programma che generi casualmente un valore tra 1 e 10, lo stampa e poi stampa true se il numero è uguale a 1 o 10, altrimenti false */
func main() {
	n := rand.Intn(10) + 1
	fmt.Println(n)
	fmt.Println(n == 1 || n == 10)
}
