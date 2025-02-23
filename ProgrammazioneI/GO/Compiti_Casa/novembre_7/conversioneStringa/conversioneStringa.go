//Autore: Francesco Corrado
//Versione sviluppata da me, su 8_novembre es classe soluzione prof

package main

import(
	"fmt"
)

func main(){
	var s, o string
	fmt.Print("Inserisci una stringa: ")
	fmt.Scan(&s)
	o= convertiStringa(s)
	fmt.Println(o)
}

func convertiStringa (s string) string{
	var x string
	var n int
	var c rune
	for i:=0; i < len(s); i++{
		c= rune(s[i])
		if c == '*'{
			for y:=i; c == '*' && y < len(s); y++{
				c = rune(s[y])
				n++
			}
			x = x + fmt.Sprint(n)
			i+=n
			n=0
		}else{
			x = x + string(c)
		}
	}
	return x
}