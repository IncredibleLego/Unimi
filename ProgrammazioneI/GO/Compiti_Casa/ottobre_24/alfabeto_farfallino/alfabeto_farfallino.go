//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s, str string
	fmt.Println("Programma che data una stringa la trasforma in alfabeto farfallino")
	fmt.Print("Inserisci una stringa (minuscola): ")
	fmt.Scan(&s)
	for i:=0; i<len(s); i++{
		c:=rune(s[i])
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			str += string(c) + "f" + string(c)
		}else{
			str += string(c)
		}
	}
	fmt.Println(str)
}