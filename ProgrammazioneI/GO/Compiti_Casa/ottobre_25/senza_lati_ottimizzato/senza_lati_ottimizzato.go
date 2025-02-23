//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Println("Programma che stampa un quadrato senza lati di n*n (versione ottimizzata)")
	fmt.Print("Inserire n: ")
	fmt.Scan(&n)
	rigaVuota(n)
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
	rigaVuota(n)
}

func rigaVuota(n int) int{
	for x:=0; x<n; x++ {
		fmt.Print(" ")
	}
	fmt.Println()
	return 0
}