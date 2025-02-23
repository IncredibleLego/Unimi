//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main() {
	const ali1=10.0/100.0
	const ali2=25.0/100.0
	const con1=32000
	const con2=64000
	var r int
	var c bool
	var tasse float64
	fmt.Println("Programma che dato reddito e coniugato calcola le tasse da pagare")
	fmt.Print("Inserisci il reddito: ")
	fmt.Scan(&r)
	fmt.Print("Coniugato? (t/f): ")
	fmt.Scan(&c)
	if c{ //Per gli operatori bool basta mettere il numero come condizione
		if r<=con2{
			tasse = float64(r)*ali1
		}else{
			tasse = float64(r)*ali2
		}
	}else{
		if r<=con1{
			tasse = float64(r)*ali1
		}else{
			tasse = float64(r)*ali2
		}
	}

	fmt.Println("Le tasse da pagare sono: ", int(tasse))

}