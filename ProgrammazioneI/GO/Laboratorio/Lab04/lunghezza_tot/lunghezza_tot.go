//Autore: Francesco Corrado

/* Scrivere un programma lunghezza_tot.go che legge da standard input un int totLen e una sequenza di 
stringhe, sommandone le lunghezze fino a raggiungere (o superare) totLen. Raggiunto totLen, il programma 
stampa la somma delle lunghezze e la concatenazione delle stringhe lette e termina. */

package main

import(
	"fmt"
)

func main(){
	var totLen, i int
	var s, sommaS string
	fmt.Print("Inserire totLen: ")
	fmt.Scan(&totLen)
	for i<totLen{
		fmt.Print("Inserire una stringa: ")
		fmt.Scan(&s)
		sommaS+= s
		i+=len(s)
	}
	fmt.Println("Somma lunghezze:", i, "\nConcatenzione stringhe:", sommaS)
}