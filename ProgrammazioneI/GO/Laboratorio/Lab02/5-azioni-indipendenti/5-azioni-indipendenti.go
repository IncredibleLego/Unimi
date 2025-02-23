package main

import "fmt"

func main() {
	/* 
	Programma che chiede un numero intero in input e stampa "Fizz" se è divisibile per 3 e
	"Buzz" se è divisibile per 5
	*/

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
/* DOMANDE */
/*
- D1. Se al posto del secondo if ci fosse un else if (legato al primo if), il programma si comporterebbe nello stesso modo? Dareste le stesse specifiche?
- R1. In quel caso nel caso in cui il numero fosse divisibile sia per 3 sia per 5 verrebbe stampato
	  solo il Fizz per il 3 e il programma non arriverebbe al print del Buzz

- D2. Perché? Spiegate la vostra risposta alla domanda D1
- R2. Spiegato sopra

*/
