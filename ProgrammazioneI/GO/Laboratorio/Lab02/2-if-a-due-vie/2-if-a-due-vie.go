package main

import "fmt"

func main() {
	/*
	 Programma che dato un numero int inserito dall'utente stabilisce se esso è pari o dispari
	*/
	
	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
