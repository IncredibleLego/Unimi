//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var carattere byte
	var somma int
	fmt.Print("Una seqeunza di caratteri terminata da '.': ")
	for{
		fmt.Scanf("%c", &carattere)
		if carattere=='.'{
			fmt.Println("somma:", somma, "\nbye")
			break
		}
		somma++
		if carattere >= 'a' && carattere <= 'z'{
			fmt.Println(string(carattere), "è la", int(carattere)-96, "^a")
		}else if carattere >= '0' && carattere <= '9'{
			fmt.Println(string(carattere), "-", int(carattere)-48)
		}else{
			fmt.Println(string(carattere), "- altro")
		}
	}
}