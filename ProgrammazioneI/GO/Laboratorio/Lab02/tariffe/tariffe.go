//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var e int
	var s bool
	var tariffa float64
	fmt.Println("Programma che data l'età e occupazione dell'utente calcola il costo di un biglietto di ingresso")
	fmt.Print("Inserire l'età: ")
	fmt.Scan(&e)
	fmt.Print("L'utente è uno/a studente/ssa? (t/f): ")
	fmt.Scan(&s)
	if (e>17 && s == true) || (e>=9 && e<18){
		tariffa= 5
	}else if e>=0 && e<10{
		tariffa=0
	}else if (e>=18 && e<=26) || (e>=65){
		tariffa=7.5
	}else{
		tariffa=10
	}
	fmt.Println("Il biglietto di ingresso costa ", tariffa, "euro")
}