//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, x int
	fmt.Print("Inserisci un intero positivo: ")
	fmt.Scan(&n)
	for c:=1; c<=n; c++{
		for r:=0; r<x;r++{
			fmt.Print(" ")
		}
		fmt.Println("*")
		x++
	}
}