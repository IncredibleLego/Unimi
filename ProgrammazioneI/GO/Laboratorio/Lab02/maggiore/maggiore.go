//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var min, max int
	fmt.Println("Programma che inseriti due valori in input stampa il valore max")
	fmt.Print("Inserisci due valori int (separati da uno spazio): ")
	fmt.Scan(&min, &max)
	if min > max {
		min, max = max, min
	}
	fmt.Println("Valore max:", max)
}