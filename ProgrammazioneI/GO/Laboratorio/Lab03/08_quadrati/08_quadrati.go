//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Println("Programma che dato un numero intero stampa i numeri tra 1 e n che sono quadrati perfetti")
	fmt.Print("Inserire n: ")
	fmt.Scan(&n)
	for i:=1; i<=n; i++{
		for k:=1; k<=i; k++{
			if k*k == i{
				fmt.Print(i, " ")
			}
			if k*k > i{
				break
			}
		}
	}
	fmt.Println()
}