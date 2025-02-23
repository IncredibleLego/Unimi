//Autore: Francesco Corrado

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Programma con lo scopo di dimostrare il teorema di format")
	var n, Fn, c int
	fmt.Print("Inserire il numero: ")
	fmt.Scan(&n)
	for d := 2; d < n; d++ {
		for c = 2; c < d; c++ {
			if d%c == 0 {
				break
			}
		}
		if c == d {
			Fn++
		}
	}
	fmt.Println("Fn = ", Fn, "N =", n)
	soluzione := float64(Fn) / (float64(n) / math.Log(float64(n)))
	fmt.Println("Utilizzando il teorema di Fermat su", n, "otteniamo:", soluzione)
}

/* //VERSIONE CON 3 RIPETIZIONI, NON FUNZIONA
func main() {
	fmt.Println("Programma con lo scopo di dimostrare il teorema di format")
	var n, Fn, c int
	var sol1, sol2, sol3 float64
	_ = sol1
	_ = sol2
	_ = sol3
	num := 1
	for x := 0; x < 3; x++ {
		fmt.Println("Ripetizione numero", num, "del ciclo")
		fmt.Print("Inserire il numero: ")
		fmt.Scan(&n)
		for d := 2; Fn < n; d++ {
			for c = 2; c < d; c++ {
				if d%c == 0 {
					break
				}
			}
			if c == d {
				fmt.Print(d, " ")
				Fn++
			}
		}
		fmt.Println("")
		soluzione := float64(Fn) / (float64(n) / math.Log(float64(n)))
		if num == 1 {
			sol1 = soluzione
		} else if num == 2 {
			sol2 = soluzione
		} else {
			sol3 = soluzione
		}
		soluzione = 0.0
		Fn = 0
		num++
	}
	fmt.Println("Dopo aver eseguito il controllo 3 volte, e avendo ottenuto", sol1, sol2, sol3, "possiamo confermare il teorema di Fermat")
} */
