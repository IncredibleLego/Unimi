//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Println("Programma che dato un numero determina se esso è primo o no")
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for i:=2; i<n; i++{
		if i == n -1{
			fmt.Println("primo")
		}else if n % i == 0{
			fmt.Println("non primo")
			break
		}
	}
	if n == 2{
		fmt.Println("primo")
	}
}