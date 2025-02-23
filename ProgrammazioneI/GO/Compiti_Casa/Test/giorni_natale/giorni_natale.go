//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var mese, giorno int
	fmt.Println("Programma che data in imput una data (giorno, mese) calcola i giorni mancanti a Natale")
	fmt.Print("Inserire il giorno (1-30) ed il mese (1-12) separati da uno spazio: ")
	fmt.Scan(&giorno, &mese)
	natale := (11-mese)*30 + 30 - giorno + 25
	fmt.Println("A Natale mancano", natale, "giorni dal", giorno, "/", mese)
}
