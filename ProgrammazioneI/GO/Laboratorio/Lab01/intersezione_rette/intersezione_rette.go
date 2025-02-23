//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var m1, q1, m2, q2 float64
	fmt.Println("Programma che date due rette calcola il punto di intersezione tra esse")
	fmt.Print("Inserire m e q della prima retta (separati da spazio): ")
	fmt.Scan(&m1)
	fmt.Scan(&q1)
	fmt.Print("Inserire m e q della seconda retta (separati da spazio): ")
	fmt.Scan(&m2)
	fmt.Scan(&q2)
	m3 := m1 - m2
	q3 := q2 - q1
	x3 := q3 / m3
	y3 := m1*x3 + q1
	fmt.Println("L'intersezione delle rette si trova nel punto (", x3, ",", y3, ")")
}
