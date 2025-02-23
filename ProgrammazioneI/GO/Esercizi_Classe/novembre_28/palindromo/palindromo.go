//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func isPalin(s string) bool{
	if len(s) <= 1{
		return true
	}else{
		return s[0]==s[len(s)-1] && isPalin(s[1:len(s)-1])
	}
}

func main(){
	var str string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&str)
	if isPalin(str){
		fmt.Println("La stringa è palindroma")
	}else{
		fmt.Println("La stringa non è palindroma")
	}
}