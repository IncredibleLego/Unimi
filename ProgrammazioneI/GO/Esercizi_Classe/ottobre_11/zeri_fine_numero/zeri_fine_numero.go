//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Programma che dato un numero intero positivo stabilisce se finisce con zero uno o più zeri")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)
	if n%1000 == 0 {
		fmt.Println("Il numero termina con più di due zeri")
	} else if n%100 == 0 {
		fmt.Println("Il numweo termina con due zeri")
	} else if n%10 == 0 {
		fmt.Println("Il numero termina con uno zero")
	} else {
		fmt.Println("Il numero termina senza zeri")
	}
}
