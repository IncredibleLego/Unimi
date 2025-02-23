//Autore: Francesco Corrado

package main

import "fmt"

/* //SOLUZIONE 1
func main() {
	var n, s, h int
	fmt.Println("Programma che cacola la media di un insieme di valori")
	fmt.Print("Inserire il valore: ")
	fmt.Scan(&h)
	for h > 0 {
		n++
		s += h
		fmt.Print("Inserire il valore: ")
		fmt.Scan(&h)
	}
	media := float64(s) / float64(n)
	fmt.Println("La media è uguale a ", media)
} */

/* //SOLUZIONE 2

func main() {
	var n, s, h int
	fmt.Println("Programma che cacola la media di un insieme di valori")
	h = 1
	for h > 0 {
		fmt.Print("Inserire il valore: ")
		fmt.Scan(&h)
		if h > 0 {
			n++
			s += h
		}
	}
	media := float64(s) / float64(n)
	fmt.Println("La media è uguale a ", media)
} */

//SOLUZIONE 3

func main() {
	var n, s, h int
	fmt.Println("Programma che cacola la media di un insieme di valori")
	for {
		fmt.Print("Inserire il valore: ")
		fmt.Scan(&h)
		if h <= 0 {
			break
		}
		n++
		s += h
	}
	media := float64(s) / float64(n)
	fmt.Println("La media è uguale a ", media)
}
