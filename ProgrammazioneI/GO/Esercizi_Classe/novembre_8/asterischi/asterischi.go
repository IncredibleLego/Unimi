//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

// Crea una stringa isdentica a s ma sostituisce le sequenze di * con il numero di * presenti
func converti(s string) string{
	var t string
	var j int
	for i:=0; i<len(s); i++ {
		if s[i] != '*' {
			t += string(s[i])
		} else {
			for j = i; j < len(s) && s[j] == '*'; j++ {} //Bisogna mettere prima il controllo lunghezza per evitare il panico dato il controllo in ordine di go
			// s[i:j] contiene solo asterischi (j escluso)
			// la sequenza è lunga j-1
			t += strconv.Itoa(j-i)
			i = j - 1
		}
	}
	return t
}


func main(){
	scanner:= bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		riga := scanner.Text()
		rigaConvertita := converti(riga)
		fmt.Println(rigaConvertita)
	}
}