//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
)

func main(){
	s := os.Args[1] //Prendo la stringa, os.Args[0] è il nome del programma
	n := len(s)
	for y:=0; y<n; y++{
		for x:=0; x < n; x++{
			if x == y || x == (x-1-y) {
				fmt.Printf("%c", s[y])
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//export GOPROXY = direct