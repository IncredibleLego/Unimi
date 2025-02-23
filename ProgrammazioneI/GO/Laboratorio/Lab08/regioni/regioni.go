//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	mappa := make(map[string][]string)
	nomeRegioni := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() //Salto la prima riga di intestazione
	for scanner.Scan(){
		riga := scanner.Text()
		pezzi := strings.Split(riga, ",")
		slice := mappa[pezzi[2]]
		slice = append(slice, pezzi[0])
		mappa[pezzi[2]] = slice
	}
	for i:=0; i < len(nomeRegioni); i++{
		for key, value := range mappa {
			if key == nomeRegioni[i]{
				fmt.Println(key,":", value)
			}
		}
	}
}