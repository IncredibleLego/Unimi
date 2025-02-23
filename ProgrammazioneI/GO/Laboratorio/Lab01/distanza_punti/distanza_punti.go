//Autore: Francesco Corrado
package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Println("Programma che date le coordinate di due punti calcola la distanza tra di essi")
	fmt.Print("Inserire la x e la y del primo punto (separate da spazio): ")
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Print("Inserire la x e la y del secondo punto (separate da spazio): ")
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	distanza := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	fmt.Println("La distanza tra i due punti dati è uguale a", distanza)
}
