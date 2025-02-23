//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, n2 int
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&n)
	for{
		if n2 == 1{
			break
		}
		if n % 2 == 0{
			n2 = n/2
		}else{
			n2 = 3*n+1
		}
		fmt.Print(n2, " ")
		n = n2
	}
	fmt.Println()
}