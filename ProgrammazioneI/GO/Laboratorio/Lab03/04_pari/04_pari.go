//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("Numeri casuali generati tra [0,10]:")
	const K = 10
	var n, c int
	for i:=1; i<=K; i++{
		n = rand.Intn(11)
		fmt.Print(n, " ")
		if n % 2 == 0{
			c++
		}
	}
	fmt.Println("\nI valori pari che sono stati generati sono stati", c)
}