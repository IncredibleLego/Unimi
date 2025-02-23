//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	fmt.Println("Programma che stampa l'ultimo carattere di una stringa")
	s:="Garibaldi fu ferito"
	c:=rune(s[len(s)-1])
	fmt.Println("L'ultimo carattere della stringa '", s, "' e': ", string(c))
}