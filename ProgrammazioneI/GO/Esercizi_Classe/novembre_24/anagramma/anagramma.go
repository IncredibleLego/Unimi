//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func anagramma(s string) []string{ //Crea Tutti gli anagrammi di una stringa
	if s=="" {
		fmt.Println("Fine2")
		return []string{""}
	}else{
		var ris []string
		for i:=0; i<len(s); i++{
			primo := rune(s[i])
			resto := s[:i]+s[i+1:]
			anagResto:=anagramma(resto)
			for _, x:= range anagResto{
				ris = append(ris, string(primo)+x)
				fmt.Println("Lungh", len(ris))
			}
		}
		fmt.Println("Fine?")
		return ris
	}
}

func main(){
	var str string
	fmt.Scan(&str)
	fmt.Println("Stringa:", str)
	fmt.Println("Funzione:", anagramma(str))
}