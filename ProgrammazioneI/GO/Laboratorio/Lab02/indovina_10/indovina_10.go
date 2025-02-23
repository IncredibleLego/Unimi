//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, x int
	fmt.Println("Programma che richiede di indovinare un numero tra 1 e 10")
	n=7 //Sarebbe meglio dichiararlo come CONST
	fmt.Print("Inserire un numero tra 1 e 10: ")
	fmt.Scan(&x)
	if x<1 || x>10{
		fmt.Println("Valore non valido")
	}else if n == x {
		fmt.Println("Hai indovinato!")
	}else {
		fmt.Println("Non hai indovinato")
	}
}

