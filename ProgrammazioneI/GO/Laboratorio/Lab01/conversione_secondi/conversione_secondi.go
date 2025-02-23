//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var secondi int
	fmt.Println("Programma che converte un numero dato di secondi in giorni, ore, minuti e secondi")
	fmt.Print("Inserisci il numero di secondi da convertire: ")
	fmt.Scan(&secondi)
	gg := secondi / 86400
	hh := (secondi % 86400) / 3600
	min := ((secondi % 86400) % 3600) / 60
	sec := ((secondi % 86400) % 3600) % 60
	fmt.Print(gg, ":", hh, ":", min, ":", sec, "\n")
}
