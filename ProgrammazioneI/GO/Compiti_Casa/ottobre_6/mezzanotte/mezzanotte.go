//Autore: Francesco Corrado

package main

import "fmt"

func main(){
	var h, m int
	fmt.Print("Ore?: ")
	fmt.Scan(&h)
	fmt.Print("Minuti?: ")
	fmt.Scan(&m)
	oreRimaste := (23 - h) * 60
	tot := 60 - m + oreRimaste
	fmt.Println("Minuti rimasti: ", tot)
}