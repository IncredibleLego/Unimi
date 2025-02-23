//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, somma int
	const d = 10
	fmt.Println("Programma che calcola la somma delle cifre di un numero n inserito")
	fmt.Print("Inserisci n: ")
	fmt.Scan(&n)
	x := n
	for{
		somma += x % d
		x = x / d
		if x < d{
			somma += x
			break
		}
	}
	fmt.Println("La somma di tutte le cifre di", n, "è uguale a", somma)
}