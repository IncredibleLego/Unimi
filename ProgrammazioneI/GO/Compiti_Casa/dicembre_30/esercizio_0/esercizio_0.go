//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"strings"
)

func Occorrenze(str string) [26]int{
	var occorrenze [26]int
	for i:=0; i<26; i++{ //Riempie l'array occorrenze di 26 spazi, tutti = a 0
		occorrenze[i] = 0
	}
	s:= strings.ToLower(str) //Prende la stringa e la rende tutta minuscola
	for i:=0; i<len(s); i++{
		r:=rune(s[i]) //Prende uno per uno i caratteri della stringa sotto forma di rune
		if r >= 97 && r <= 122{ //Deve essere tra 97 e 122 per essere una lettera minuscola,
			occorrenze[r-97]++  //altrimenti è un altro carattere, es il punto di domanda da non
		}                       //considerare. Aumenta la pos r-97 ovvero della lettera 
	}                           //corrispondente di 1
	return occorrenze
}

func main(){
	var s string
	text:=os.Args[1:] //Crea un array con tutti i valori da riga di comando
	for i:=0; i<len(text); i++{ //Prende uno per uno i valori dell'array e li mette in una stringa
		s = s + text[i]
	}
	occ:=Occorrenze(s)
	for i:=0; i<26; i++{
		fmt.Print(string(rune(i+97)), "  ", occ[i], " ") //Stampa finale di lettera + occorrenza
	}
	fmt.Println()
}