//Autore: Francesco Corrado

package main

import (
	"fmt"
	"math"
)

func main(){
	var x, y float64
	const EPSILON = 1e-5
//	fmt.Println("Programma che date due coordinate stabilisce se esse sono più o meno distanti di 10^-5 da (0,0)")
	//NON METTERE MAI PRINT QUANDO CI SONO I TEST
	fmt.Scan(&x, &y)
//Guardato da soluzioni, serve a fare x*x + y*y senza errori
//	distanza:=math.Sqrt((x-0)*(x-0)+(y-0)*(y-0)) //Distanza tra due punti (0,0) e (x,y)
	if math.Hypot(x, y) < EPSILON {
		fmt.Println("molto vicino all'origine")
	}else {
		fmt.Println("non vicino all'origine")
	}
}