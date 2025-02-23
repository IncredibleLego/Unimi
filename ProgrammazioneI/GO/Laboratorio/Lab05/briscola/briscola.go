//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func puntiCarta (s string) int{
	var punti int
	for i:=0; i<len(s); i++{
		switch s[i]{
			case 'A':
				punti+=11
			case '3':
				punti+=10
			case 'K':
				punti+=4
			case 'Q':
				punti+=3
			case 'J':
				punti+=2
			case '7', '6', '5', '4', '2', '0':
				punti+=0
			default:
				punti = -1
				break
		}
	}
	return punti
}

func main(){
	var s string
	for{
		fmt.Print("Inserisci una mano di briscola (deve avere tre caratteri): ")
		fmt.Scan(&s)
		if len(s) > 3 || len(s) < 3{
			fmt.Println("La stringa non ha tre caratteri, riprovare")
		}else{
			break
		}
	}
	fmt.Println("mano", s, ":", puntiCarta(s), "punti")
}