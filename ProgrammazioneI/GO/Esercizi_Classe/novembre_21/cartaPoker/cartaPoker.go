//Autore: Francesco Corrado

package main

import(
	"fmt"
)

type Carta struct{
	seme int //0=Cuori, 1=Quadri...
	valore int //0=A, 1=2...9=10, 10=J, 11=Q, 12=K
}

func mazzo() []Carte{
	var m []Carte
	for s:=0; s<4;s++{
		for v:=0, v<13; v++{
			var c Carta
			c.Seme = seme
			c.Valore = v
			m = append(m, c)
		}
	}
	return m
}