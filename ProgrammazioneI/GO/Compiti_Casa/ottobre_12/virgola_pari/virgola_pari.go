//Autore: Francesco Corrado
package main

import "fmt"

func main() {
	var n float64
	fmt.Println("Programma che dato un numero float verifica se il numero dopo la virgola sia pari o dispari")
	fmt.Print("Inserisci il numero (float): ")
	fmt.Scan(&n)
	i := float64(int(n))
	fmt.Println("n - i =", n-i+0.01)
	d := int((n - i + 0.01) * 10.0)
	if d == 0 {
		fmt.Println("Il numero non è ne pari ne dispari essendo 0")
	} else if d%2 == 0 {
		fmt.Println("Il numero dopo la virgola è pari")
	} else {
		fmt.Println("Il numero dopo la virgola è dispari")
	}
}
