//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	for i:=1; i<len(s); i++{
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