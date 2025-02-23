//Autore: Francesco Corrado

package main

import (
	"fmt"
)

func main(){
	var litri, prezzo float64
	var tipo int
	fmt.Println("Programma che calcola il prezzo di un rifornimento")
	for{
		fmt.Println("Inserire il tipo di rifornimento da fare secondo la tabella: \n 0=benzina (1,78 al litro) \n 1=diesel (1,98 al litro) \n 2=alcol (1.2 al litro) \n 3=cherosene (1.1 al litro)")
		fmt.Scan(&tipo)
		if tipo>-1 && tipo<4{
			break
		}else{
			fmt.Println("Il valore inserito non è valido. Inserire un valore tra 0 e 4")
		}
	}
	for{
		fmt.Print("Inserire la quantità di litri da rifornire: ")
		fmt.Scan(&litri)
		if litri >0{
			break
		}else {
			fmt.Println("Inserire un valore positivo")
		}
	}
	if tipo == 0{
		prezzo=1.78
	}else if tipo == 1{
		prezzo=1.98
	}else if tipo == 2{
		prezzo=1.2
	}else if tipo == 3{
		prezzo=1.1
	}
	finale:= prezzo * litri
	fmt.Println("Il prezzo finale è uguale a ", finale)
}