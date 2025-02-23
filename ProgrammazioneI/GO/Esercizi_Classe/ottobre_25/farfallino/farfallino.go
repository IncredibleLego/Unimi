//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	n:= farfallino(s)
	fmt.Println(n)
}

func farfallino (s string) string{
	var t string
	for _, c:= range s{
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'{
			t+=string(c)+"f"+string(c)
		}else{
			t+=string(c)
		}
	}
	return t
}