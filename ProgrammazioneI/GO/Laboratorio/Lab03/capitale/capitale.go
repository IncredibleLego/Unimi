//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var c, i, o, x float64
	var anni int
	fmt.Println("Programma che dato un capitale un interesse e un obbiettivo calcola gli anni necessari per ottenerlo")
	fmt.Print("Inserire il capitale iniziale: ")
	fmt.Scan(&c)
	fmt.Print("Inserire il tasso di interesse: ")
	fmt.Scan(&i)
	fmt.Print("Inserire l'obbiettivo da raggiungere: ")
	fmt.Scan(&o)
	x = c
	for{
		anni+=1
		x = x + (x * (i/100))
		if x>=o{
			break
		}
	}
	fmt.Println("Per raggiungere l'obbiettivo sono necessari", anni, "anni")
}