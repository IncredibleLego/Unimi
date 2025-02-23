package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
/*
che aggiorna una dispensa prendendo i dati da un file CSV il cui nome è passato come parametro.
	In particolare il file conterrà gli acquisti nel formato "approv,<prodotto>,<peso>" e i consumi nel formato 
	"uso,<prodotto>,<peso>", dove le quantità (pesi) sono sempre espresse in Kg. (Non vanno fatti controlli sul 
	formato dei dati, si può assumere che siano sempre esattamente come nel file di esempio).
	Se il file non esiste o non è leggibile, la funzione deve terminare subito restituendo 'false'.
	Nel caso in cui un prodotto dovesse "andare in negativo"  durante l'aggiornamento della dispensa, la 
	funzione dovrà terminare subito e restituire 'false', mentre restituirà 'true' in caso contrario.
*/
func AggiornaDispensa(dispensa map[string]float64, filename string) bool{
	file, err := os.Open(filename)
	if err != nil{
		return false
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan(){
		riga := scanner.Text()
		parametri := strings.Split(riga, ",")
		q, _ := strconv.ParseFloat(string(parametri[2]), 64)
		if parametri[0] == "approv"{
			dispensa[parametri[1]] = dispensa[parametri[1]] + q
		}else{
			x := dispensa[parametri[1]]
			if x - q < 0{
				return false
			}else{
				dispensa[parametri[1]] = dispensa[parametri[1]] - q
			}
		}
	}
	return true
}
/*
che, data una dispensa opportunamente popolata di provviste, restituisce la 
rimanenza in Kg dell'alimento passato come parametro; restituisce 0 se il prodotto non è presente in lista. */
func Rimanenza(dispensa map[string]float64, alimento string) float64{
	val, ok := dispensa[alimento]
	if ok{
		return val
	}else{
		return 0
	}
}

func main(){
	disp := make(map[string]float64)
	nFile := os.Args[1]
	risultato := AggiornaDispensa(disp, nFile)
	if len(os.Args) > 2{
		alimenti := os.Args[2:]
		for i:=0; i < len(alimenti); i++{
			cibo := string(alimenti[i])
			fmt.Println(cibo, Rimanenza(disp, cibo))
		}
	}else{
		if risultato == true{
			fmt.Println("file corretto")
		}else{
			fmt.Println("file non corretto")
		}
	}
}