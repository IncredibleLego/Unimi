//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var a, b, somma int
	var s string
	b = -1
	for{
		fmt.Scan(&a)
		if a == -1{
			break
		}
		somma+= a
		if b != -1{
			if a >= b{
				s = s + "+"
			}else{
				s = s + "-"
			}
		}
		b = a
	}
	fmt.Print(s, "\n")
	fmt.Println("somma:", somma)
}