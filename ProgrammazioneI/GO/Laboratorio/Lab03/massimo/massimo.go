//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
)

func main(){
	const NUM=10
	const MAX=21
	var n, max int
	max = 0
	fmt.Println("Programma che genera 10 numeri interi casuali tra 10 e 30 e stampa il massimo")
	for i:=1; i<=NUM; i++{
		n = rand.Intn(MAX)+10 //Cosi parto da 10 e arrivo a 30 (31 e escluso), anziche 0-20
		fmt.Print(n , " ")
		if n > max{
			max = n
		}
	}
	fmt.Println("\nNunmero massimo generato: ", max)

}