//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("fornire due numeri (int): ")
	_, err := fmt.Scan(&n1, &n2)
	fmt.Println(err) //stampo l'errore , se è nil, vuol dire che non si sono verificati errori
	fmt.Println(n1, n2)
	n1, n2 = n2, n1
	fmt.Println(n1, n2)
}
