//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var anno int
	fmt.Println("Programma che dato un anno calcola se è bisestile")
	fmt.Print("Inserire l'anno: ")
	fmt.Scan(&anno)
	if (anno%100 != 0 && anno%4 == 0) || anno%400 == 0 {
		fmt.Println("L'anno ", anno, "è bisestile")
	} else {
		fmt.Println("L'anno ", anno, "non è bisestile")
	}
}
