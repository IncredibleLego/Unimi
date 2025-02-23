//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s string
	fmt.Print("Inserire una stringa di soli caratteri ASCII: ")
	fmt.Scan(&s)
	fmt.Print(string(s[0]))
	for i:=1; i<len(s); i++{
		if s[i] < s[i-1]{
			fmt.Print("-", string(s[i]))
		}else{
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}