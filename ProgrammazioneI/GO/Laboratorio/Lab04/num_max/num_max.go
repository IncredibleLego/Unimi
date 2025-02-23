//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, c int
	max:=-1
	for i:=0; i<10; i++{
		fmt.Print("Inserisci un intero positivo: ")
		fmt.Scan(&n)
		if n < 0{
			fmt.Println("Il numero inserito non è valido, riprovare")
			i--
		}else if n > max {
			max = n
			c = 0
		}else if n == max{
			c++
		}
	}
	fmt.Println("Il numero massimo inserito è", max, "e si ripete", c+1, "volte")
}