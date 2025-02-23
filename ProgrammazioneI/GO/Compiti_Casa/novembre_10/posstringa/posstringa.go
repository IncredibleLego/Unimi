//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
)

func stampaOccorrenze(s, substr string) {
	var i, f, j int
	f = -1
	for {
		if i == 0{
			j = strings.Index(s, substr)
		}else{
			j = strings.Index(s[i:], substr)
		}
		if j != -1{
			if f == -1 || j+i-1 != f{
				fmt.Print(i+j, " ")
				f = j+i-1
			}
		}
		if i == len(s)-1{
			return
		}
		i++
	}
}

func main(){
	var s, t string
	fmt.Print("Inserisci una stringa: ")
	fmt.Scan(&s)
	fmt.Print("Inserisci la substringa da trovare: ")
	fmt.Scan(&t)
	stampaOccorrenze(s, t)
	fmt.Println()
}