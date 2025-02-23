//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s string
	fmt.Println("Programma stringhe")
	s = "Pippo"
	for i:=0; i<len(s); i++{
		c:=rune(s[i])
		fmt.Println(string(c))
	}
}