//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var x int
	var c rune
	var s string
	fmt.Println("Programma che conta il numero di 'a' in una stringa")
	fmt.Print("Inserire una stringa: ") //Funziona solo se la stringa e' senza spazi
	fmt.Scan(&s)
//	s="Garibaldi andava a cavallo"
	for i:=0; i < len(s); i++ {
		c=rune(s[i])
		if c == 'a' {
			x++
		}
	}
	fmt.Println("Nella frase '", s, "' sono presenti", x, "lettere 'a'")
}