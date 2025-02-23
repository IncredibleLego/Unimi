package main

import "fmt"

func main() {
    /*
    Programma che chiede un valore n in input e procede a stampare tutti i valori da 0 fino a n interi 
	utilizzando 8 come intero (es 9 = 1 1, 1 intero che è 8 e 1 resto) e per ciascuno di essi una barra
	verticale |
    */

	const BASE = 8
	var n int

	fmt.Print("un int: ")
	fmt.Scan(&n)

	length := float64(n)
	for i := 0; float64(i)/BASE < length; i++ { // i < n*BASE
		fmt.Print("|\t")
	}
	fmt.Println()

	for i := 0; float64(i)/BASE < length; i++ {
		fmt.Print(i/BASE, i%BASE, "\t")
	}
	fmt.Println()
}