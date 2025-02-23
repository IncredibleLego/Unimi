//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var distanza, carburante float64
	fmt.Println("Programma che data la distanza totale percorsa calcola il consumo medio e la resa di un motore")
	fmt.Print("Inserisci la distanza percorsa (in km): ")
	fmt.Scan(&distanza)
	fmt.Print("Inserisci la quantità di carburante utilizzata (in l): ")
	fmt.Scan(&carburante)
	consumo := carburante / distanza
	resa := distanza / carburante
	fmt.Println("Il consumo medio è stato di", consumo, "l/km")
	fmt.Println("La resa media è stata di", resa, "km/l")
}
