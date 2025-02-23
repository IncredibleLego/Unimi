//Autore: Francesco Corrado

package main

import "fmt"

func main(){
	var num1, den1, num2, den2 int
	fmt.Println("Programma che date in input due frazioni stabilisce se esse sono equivalenti")
	fmt.Print("Inserire numeratore e denominatore della prima frazione (separati da spazio): ")
	fmt.Scan(&num1, &den1)
	fmt.Print("Inserire numeratore e denominatore della seconda frazione (separati da spazio): ")
	fmt.Scan(&num2, &den2)
	if den1*num2 == num1*den2 {
		fmt.Println("equivalenti")
	}else {
		fmt.Println("non equivalenti")
	}
}