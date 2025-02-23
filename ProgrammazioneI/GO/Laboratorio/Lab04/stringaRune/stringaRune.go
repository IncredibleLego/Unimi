//Autore: Francesco Corrado

package main

import(
	"fmt"
	"unicode/utf8"
)

func main(){
	var s string
	var l int
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	l = utf8.RuneCountInString(s)
	fmt.Print("Lunghezza: ", l)
	for i:=1; i<l; i++{
		if s[i] > s[i-1]{
			fmt.Print("+")
		}else if s[i] < s[i-1]{
			fmt.Print("-")
		}else{
			fmt.Print("=")
		}
	}
	fmt.Println()
}