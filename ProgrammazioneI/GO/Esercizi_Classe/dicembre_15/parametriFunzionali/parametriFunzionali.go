//Autore: Francesco Corrado

package main

import(
	"fmt"
)

type Persona{
	nome, cognome string
	matricola string
	nascita Data //Bisogna definire il tipo data
}

var p []Persona //Va popolata es leggendo da un file


func lessCognome(i, j int) bool{
	return p[i].cognome < p[j].cognome
}

func lessMatricola(i, j int) bool{
	return p[i].matricola < p[j].matricola
}

sortSlice(p, lessCognome)
sort.Slice(p, lessCognome)

sort.Slice(p, func(i, j) bool{ //Funzione senza nome (O lambda), senza nome definita al momento
	return p[i].matricola<p[j].matricola
})

sort.Slice(p, func(i, j) bool{
	di:= p[i].nascita
	dj:= p[j].nascita
	return precede (di, dj)
})






func main(){
	
}

