//Autore: Francesco Corrado

/* Scrivere un programma max_char.go che legge da standard input una sequenza di 5 caratteri ASCII 
(byte) e stampa il maggiore in ordine lessicografico (cioè con il codice ASCII più alto). */

package main

import(
	"fmt"
)

func main(){
	var s string
	var MAX int
	fmt.Print("Inserire una stringa di soli 5 caratteri ASCII: ")
	fmt.Scan(&s)
	MAX = int(s[0])
	for i:=1; i<len(s); i++{
		if s[i] > s[i-1]{
			MAX = int(s[i])
		}
	}
	fmt.Println("Il carattere con maggiore ordine lessicografico è", string(MAX), "con indice ASCII", MAX)
}