//Autore: Francesco Corrado

package main

import "fmt"

func main(){
	var n int
	fmt.Println("Programma che chiede un orario intero e stabilisce se esso si trova al mattino")
	fmt.Print("Inserire un orario (intero): ")
	fmt.Scan(&n)
	if n >= 6 && n < 13 {
		fmt.Println("mattino")
	} 
}