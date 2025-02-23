//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
)

func main(){
	var c rune
	var s string
	fmt.Print("Inserire il carattere da trovare nella stringa: ")
	fmt.Scanf("%c", &c)
	fmt.Print("String c: ", string(c))
	fmt.Print("Inserire la stringa: ")
	fmt.Scan(&s)
	pos:= strings.Index(s, string(c))
	fmt.Println("Il carattere ", string(c), " si trova per la prima volta nella posizione ", pos+1)

}