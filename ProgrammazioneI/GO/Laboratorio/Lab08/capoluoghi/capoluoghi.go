//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Capoluogo struct {
	Nome, Sigla, Regione string
	Superficie int
}

func main(){
	mappa := make(map[string]Capoluogo)
	var sigleDaStampare []string
	sigleDaStampare = os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() //Salta la prima riga di intestazione
	for scanner.Scan(){
		var cap Capoluogo
		riga := scanner.Text()
		pezzi := strings.Split(riga, ",")
		cap.Nome = pezzi[0]
		cap.Sigla = pezzi[1]
		cap.Regione = pezzi[2]
		cap.Superficie, _ = strconv.Atoi(pezzi[4])
		mappa[cap.Sigla] = cap
	}
	for i:= 0; i < len(sigleDaStampare); i++{
		fmt.Println(mappa[sigleDaStampare[i]])
	}
}