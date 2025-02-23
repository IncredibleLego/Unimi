//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math"
)

func main(){
	fmt.Println("Programma che dati 5 input stabilisce quale è più vicino a un valore target")
	const TARGET = 12
	var n, v, N float64
	v = 25.0
	for i:=1; i<=5; i++{
		fmt.Print("Inserire il numero numero ", i, ": ")
		fmt.Scan(&n)
		if n>=0 && n<=20{
			if math.Abs(TARGET - n) < v{
				v = math.Abs(TARGET -n)
				N = n
			}
		}else{
			fmt.Println("Valore fuori dal range consentito [0,20], riprovare")
			i--
		}
	}
	fmt.Println("Il valore più vicino a TARGET è", N)
}