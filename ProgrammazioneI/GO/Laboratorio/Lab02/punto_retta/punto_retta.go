//Autore: Francesco Corrado

package main

import (
	"fmt"
)

func main(){
	const distanza=1e-06 //Non utilizzata alla fine
	var x1, y1, m, q float64
	fmt.Println("Programma che dati in input i valori di una retta e di un punto stabisce la posizione del secondo rispetto alla prima")
	fmt.Print("Inserire x e y del punto P: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Inserire m e q della retta y=mx+q: ")
	fmt.Scan(&m, &q)
	if y1 == m*x1+q{
		fmt.Println("sulla retta")
	}else if m*x1+q < y1{
		fmt.Println("sopra")
	}else {
		fmt.Println("sotto")
	}
}