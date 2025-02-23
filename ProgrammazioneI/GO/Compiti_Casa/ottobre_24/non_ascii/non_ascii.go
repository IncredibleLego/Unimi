//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s string
	var count int
	fmt.Println("Conta quanti caratteri non-ascii sono in una stringa")
	fmt.Print("Inserisci una stringa: ")
	fmt.Scan(&s)
	for i:=0; i < len(s); i++{
		c := rune(s[i])
		if c < 0 || c > 128 {
			count++
		}
	}
	fmt.Println("Nella stringa sono presenti", count/2, "caratteri non-ascii")
}