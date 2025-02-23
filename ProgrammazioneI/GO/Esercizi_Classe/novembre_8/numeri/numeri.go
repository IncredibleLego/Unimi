//Autore: Francesco Corrado

package main

/*
  Converte le righe come in questo esempio:

          ci sono 12 cani e 345 gatti

          ci sono dodici cani e trecentoquarantacinque gatti
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func f(x int) string {
	return "<<" + strconv.Itoa(x) + ">>"
}

// Crea una stringa identica a s a parte il fatto che le sequenze di
// cifre vengono passate alla funzione f prima di essere concatenate
func converti(s string) string {
	var t string
	var j int
	for i := 0; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			t += string(s[i])
		} else {
			for j = i; j < len(s) && unicode.IsDigit(rune(s[j])); j++ {
			}
			// s[i:j] contiene solo cifre (j escluso!)
			n, _ := strconv.Atoi(s[i:j])
			t += f(n)
			i = j - 1
		}
	}
	return t
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		rigaConvertita := converti(riga)
		fmt.Println(rigaConvertita)
	}
}
