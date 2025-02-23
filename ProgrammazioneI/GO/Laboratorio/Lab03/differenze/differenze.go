//Autore: Francesco Corrado

package main

import(
	"fmt"
//	"math"
)
/* Soluzione originale
func main(){
	var n1, n2, diff float64
	var c int
	for{
		fmt.Scan(&n1)
		if n1 == 0{
			break
		}
		if c > 0{
			diff = math.Round(n1 - n2)
			fmt.Println(diff)
		}
		n2 = n1
		c++
	}
} */


//Soluzione presa da soluzioni caricate
func main() {
	var prev, next float64

	fmt.Print("scrivi un numero ")
	fmt.Scan(&next)

	for next != 0 {
		prev = next
		fmt.Print("scrivi un numero ")
		fmt.Scan(&next)
		fmt.Println("diff = ", next-prev)
	}
}
