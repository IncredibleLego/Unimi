//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var g, m, a int
	fmt.Println("Programma che data una data calcola se è corretta")
	fmt.Print("Inserisci la data (g m a) separata da una virgola: ")
	fmt.Scan(&g, &m, &a)
	/*	if m > 12 {
			fmt.Println("Il mese inserito non corretto: troppo grande")
		}
		if (m == 11) || (m == 4) || (m == 6) || (m == 9) && g > 31 */ //Esercizio mio
	if g <= 0 || g > 31 || m <= 0 || m > 12 {
		fmt.Println("La data inserita è sbagliata")
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		if g == 31 {
			fmt.Println("La fata inserita è sbagliata")
		} else {
			fmt.Println("La data è corretta")
		}
	} else if m == 2 {
		bisest := (a%100 != 0 && a%4 == 0) || a%400 == 0
		var gf int
		if bisest {
			gf = 29
		} else {
			gf = 28
		}
		if g <= gf {
			fmt.Println("La data è corretta")
		} else {
			fmt.Println("La data è sbagliata")
		}
	}
}
