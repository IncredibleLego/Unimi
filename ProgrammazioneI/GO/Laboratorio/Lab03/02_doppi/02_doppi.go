//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	const K = 5
	var n int
	for i:=1; i<=K; i++{
		fmt.Print("Inserire un numero: ")
		fmt.Scan(&n)
		fmt.Println("Doppio: ", n*2)
	}
}