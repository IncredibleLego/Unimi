//Autore: Francesco Corrado

package main

import (
	"fmt"
	"strings"
)

func main() {
	const K = 5
	var parola string
	for i:=0; i<K; i++ {
		var primaVocale rune = -1
		fmt.Scan(&parola)
   		for _, r := range parola {
			if strings.ContainsRune("aeiou", r) {
				primaVocale = r
				break
			}
   		}
		if primaVocale == -1 {
			fmt.Println("la parola", parola, "non contiene vocali")
		} else {
			fmt.Println(string(primaVocale))
		}
	}
}