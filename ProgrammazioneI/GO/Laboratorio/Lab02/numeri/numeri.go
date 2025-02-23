//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var x int
	fmt.Println("Programma che dato un numero stabilisce se un intero e' divisibile per 10 o positivo o entrambi")
	fmt.Print("Inserire il numero: ")
	fmt.Scan(&x)
	if x>= 0 {
		fmt.Print("positivo ")
	}
	if x % 10 == 0{
		fmt.Print("divisibile per 10")
	}
	fmt.Println()
}