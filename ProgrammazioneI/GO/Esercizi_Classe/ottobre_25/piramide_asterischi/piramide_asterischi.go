//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Println("Programma che stampa una piramide progressiva di asterischi in base al valore inserito n")
	fmt.Print("Inserire n: ")
	fmt.Scan(&n)
	for r:=0; r<n; r++ {
		for c:=0; c<n-r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}