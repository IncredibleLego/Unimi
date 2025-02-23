package main

import "fmt"

func main() {
    /*
    Programma che richiede all'utente di inserire 5 numeri int e ne calcola la somma
    */
    const K = 5
	var n int
	somma := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma += n  // s = s + n
	}
	fmt.Println(somma)
}
