//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func valCarta (n int) string{
	var s string
	v := n - (n/13)*13
	switch{
		case v == 0:
			s = "A"
		case v == 10:
			s = "J"
		case v == 11:
			s = "Q"
		case v == 12:
			s = "K"
		default:
			s = fmt.Sprint(v+1)
	}
	switch{
		case n<=12:
			s += " ♥"
		case n>12 && n<=25:
			s += " ♦"
		case n>25 && n<=38:
			s += " ♣"
		case n>38 && n<=51:
			s += " ♠" 
	}
	return s

}

func main(){
	var n int
	fmt.Println("Programma che dato un numero stampa la carta corrispondente in un mazzo di carte da poker")
	for{
		fmt.Print("Inserire un numero (0-51): ")
		fmt.Scan(&n)
		if n>=0 && n<=51{
			break
		}else{
			fmt.Println("Il numero inserito non e' compreso nel range (0-51), riprova")
		}
	}
	carta := valCarta(n)
	fmt.Println("La carta corrispondente al numero", n, "è il", carta)
}