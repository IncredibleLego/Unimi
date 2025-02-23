//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
	"time"
)

func cartaPoker (n int) string{
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
		s += "♥"
	case n>12 && n<=25:
		s += "♦"
	case n>25 && n<=38:
		s += "♣"
	case n>38 && n<=51:
		s += "♠" 
	}
	return s
}

func main(){
	fmt.Println("Programma che genera casualmente una mano di poker")
	r:= rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<4; i++{
		c := r.Intn(51)
		fmt.Println(cartaPoker(c))
//		if i == 0{
	//		a := r.Int
	//	}
	}
}