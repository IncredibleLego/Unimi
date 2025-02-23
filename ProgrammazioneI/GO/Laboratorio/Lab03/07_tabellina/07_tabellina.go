//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Println("Programma che dato un numero n ne stampa la tabellina")
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for i:=1; i<=10; i++{
		fmt.Print(n*i, "\t")
	}
	fmt.Println()
}