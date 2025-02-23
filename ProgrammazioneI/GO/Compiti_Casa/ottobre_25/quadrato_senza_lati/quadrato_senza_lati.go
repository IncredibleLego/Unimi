//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Println("Programma che stampa un quadrato senza lati di n*n")
	fmt.Print("Inserire n: ")
	fmt.Scan(&n)
	for r1:=0; r1<n; r1++ {
		fmt.Print(" ")
	}
	fmt.Println()
	for r:=0; r<n-2; r++ {
		for c:=0; c<n; c++{
			if c == 0 || c == (n-1){
				fmt.Print(" ")
			}else{
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	for r2:=0; r2<n; r2++ {
		fmt.Print(" ")
	}
	fmt.Println()
}

//Nella stessa cartella e' presente una verisone ottimizzata di questo programma