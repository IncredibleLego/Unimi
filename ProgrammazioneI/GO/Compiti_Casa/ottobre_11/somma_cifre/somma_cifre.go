//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Programma che dato in imput un numero di tre cifre calcola se la somma delle tre è > 10 oppure no")
	fmt.Print("Inserisci il numero (di tre cifre): ")
	fmt.Scan(&n)
	somma := n/100 + (n%100)/10 + (n%100)%10
	if somma > 10 {
		fmt.Println("Si, la somma delle cifre di ", n, "è uguale a ", somma, ", pertanto è >10")
	} else {
		fmt.Println("No, la somma delle cifre di ", n, "è uguale a ", somma, ", pertanto è <10")
	}
}
