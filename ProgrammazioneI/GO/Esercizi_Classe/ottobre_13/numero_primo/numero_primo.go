//Autore: Francesco Corrado

package main

import "fmt"

/* //SOLUZIONE 1

func main() {
	var n int
	var trovatoDivisore bool
	fmt.Println("Programma che dato un numero di imput calcola se esso è primo")
	fmt.Print("Inserire il numero: ")
	fmt.Scan(&n)
	for d := 2; d < n; d++ {
		if n%d == 0 {
			trovatoDivisore = true
			break
		}
	}
	if trovatoDivisore {for {
	break
}
		fmt.Println("Il numero ", n, "non è primo")
	} else {
		fmt.Println("Il numero ", n, "è primo")
	}
} */

//SOLUZIONE 2

func main() {
	var n, d int
	fmt.Println("Programma che dato un numero di imput calcola se esso è primo")
	fmt.Print("Inserire il numero: ")
	fmt.Scan(&n)
	for d = 2; d < n; d++ {
		if n%d == 0 {
			break //Tendenzialmente conviene non cambiare le variabili di un ciclo dentro al ciclo, altrimenti non si può controllare
		}
	}
	if d < n { //Permette di controllare se il si è usciti in modo naturale o con il break, per quello d va dichiarata fuori
		fmt.Println("Il numero ", n, "non è primo")
	} else {
		fmt.Println("Il numero ", n, "è primo")
	}
}
