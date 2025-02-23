package main

import "fmt"

func main() {
	/*
	Programma che chiede all'utente di inserire un numero intero e stampa se è positivo, negativo o nullo
	*/

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}
