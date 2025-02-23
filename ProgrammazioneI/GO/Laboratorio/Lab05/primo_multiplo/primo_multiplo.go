//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, m int
	for i:=0; i<5; i++{
		fmt.Print("Inserisci un numero intero >=0: ")
		fmt.Scan(&n)
		if n < 0{
			fmt.Println("Il numero inserito è < 0, riprovare")
			i--
		}else{
			if m == 0{
				if n % 3 == 0{
					m = n
				}
			}
		}
	}
	if m == 0{
		fmt.Println("no")
	}else{
		fmt.Println(m)
	}
}