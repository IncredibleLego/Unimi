//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("Numero: ")
	fmt.Scan(&n)
	x:=0
	y:=n
	for i:=0; i<n/2; i++{
		for j:=0; j<x; j++{
			fmt.Print("* ")
		}
		for j:=0; i<y; j++{
			if i%2 == 0{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for j:=0; j<x; j++{
			fmt.Print(" *")
		}
		fmt.Println()
		if i%2 == 0{
			if i<n/2{
				x++
				y -= 4
			}else{
				x--
				y += 4
			}
		}
	}
}