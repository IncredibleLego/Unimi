//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func haDoppie (s string) bool{
	for i:=1; i<len(s); i++{
		if s[i] == s[i-1]{
			return true
		}
	}
	return false
}

func main() {
	var s string
	var c int
	for{
		fmt.Scan(&s)
		if s == "ciao"{ //Trovare condizione per CTRL D
			break
		}
		if haDoppie(s){
			c++
		}
	}
	fmt.Println("Nell'elenco sono presenti", c, "parole che includono lettere doppie")
}
