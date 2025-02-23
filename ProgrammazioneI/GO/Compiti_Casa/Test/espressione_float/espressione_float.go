//Autore: Francesco Corrado

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Programma che calcola l'espressione ax^2+bx+c dati a, b e c in input")
	fmt.Print("Inserire a, b e c (separati da uno spazio): ")
	fmt.Scan(&a, &b, &c)
	delta := b*b - 4*a*c
	s := math.Sqrt(delta)
	if delta > 0 {
		x1 := (b*-1 + s) / (2 * a)
		x2 := (b*-1 - s) / (2 * a)
		fmt.Println("Le due soluzioni sono reali e distinte e sono", x1, x2)
	} else if delta == 0 {
		x := (b*-1 + s) / (2 * a)
		fmt.Println("Le due soluzioni sono coincidenti e sono uguali a ", x)
	} else {
		fmt.Println("Non esistono soluzioni reali in quanto delta < 0")
	}
}
