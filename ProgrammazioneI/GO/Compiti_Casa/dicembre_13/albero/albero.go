//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, nSpazi, nAster int
	fmt.Scan(&n)
	nSpazi = n-1
	nAster = 1
	if n == 1{
		fmt.Print(" ")
	}
	for r:=0; r<n; r++{
		for s:=0; s < nSpazi; s++{
			fmt.Print(" ")
		}
		for a:=0; a < nAster; a++{
			fmt.Print("*")
		}
		fmt.Println()
		nSpazi -= 1
		nAster = nAster + 2
	}
	pos := (nAster-2)/2 -1
	for x:=0; x < 3; x++{
		for s:=0; s < pos; s++{
			fmt.Print(" ")
		}
		for p:=0; p < 3; p++{
			fmt.Print("+")
		}
		fmt.Println()
	}
}