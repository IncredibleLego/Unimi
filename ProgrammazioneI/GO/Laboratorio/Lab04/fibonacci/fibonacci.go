//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, a, b, s int
	fmt.Print("Inserire un numero n: ")
	fmt.Scan(&n)
	a=-1
	b=-1
	for r:=1; r<=n; r++{
		s = a + b
		if s <= 0{
			s=1
		}
		for i:=1; i<=s; i++{
			fmt.Print("*")
		}
		a = b
		b = s
		fmt.Println()
	}
}