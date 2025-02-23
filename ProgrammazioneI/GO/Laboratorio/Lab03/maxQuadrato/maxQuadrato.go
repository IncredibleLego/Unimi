//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, k int
	fmt.Println("Programma che dato un numero stampa il massimo quadrato minore di n")
	fmt.Print("Inserire un valore n: ")
	fmt.Scan(&n)
	for{
		if k*k >= n{
			if k*k > n{
				k--
			}
			break
		}
		k++
	}
	fmt.Println("Il massimo quadrato possibile minore di", n, "è", k, "^ 2", "=", k*k,)
}