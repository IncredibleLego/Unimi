//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("Inserire il numero di righe e colonne ")
	fmt.Scan(&n)
	for r:=0; r<n; r++{
		for c:=0; c<n; c++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
