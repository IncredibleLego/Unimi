//Autore: Francesco Corrado

package main

import(
	"fmt"
)
/*
prende una slice di interi e ne restituisce una che contiene solo i numeri di quella 
in ingresso che sono numeri pari
*/

func estraiPari(in []int) (out []int){
	for i:=0; i < len(in); i++{
		if in[i] % 2 == 0{
			out = append(out, in[i])
		}
	}
	return out
} 

/*
che prende un intero e una slice di interi e ne restituisce una senza i multipli del numero 
passato di quella in ingresso. 
Es.: rimuoviMultipli(5, in) restituisce, a partire da in, una slice senza i multipli di 5.
*/

func rimuoviMultipli(m int, in []int) (out []int){
	for i:=0; i < len(in); i++{
		if in[i] % m != 0{
			out = append(out, in[i])
		}
	}
	return out
} 

func main(){
	var s1, s2 []int
	s1 = []int{1, 2, 3, 4, 5, 6, 7}
	o1 := estraiPari(s1)
	fmt.Println("Con i pari: ", s1)
	fmt.Println("Solo i pari: ", o1)
	s2 = []int{1, 5, 6, 8, 10, 9, 25, 26, 50, 56, 57}
	o2 := rimuoviMultipli(5, s2)
	fmt.Println("Con i multipli: ", s2)
	fmt.Println("Senza i multipli: ", o2)
}