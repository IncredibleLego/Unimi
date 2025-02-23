//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	n:= contaNonAscii(s)
	fmt.Println("I caratteri non ascii sono ", n)
}

func contaNonAscii (s string) int{
	count:=0
	for _, c:= range s{
		if c > 127{
			count++
		}
	}
	return count
}