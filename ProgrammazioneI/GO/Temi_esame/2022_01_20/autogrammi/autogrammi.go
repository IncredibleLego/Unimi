package main

import(
	"fmt"
	"strings"
	"unicode"
)
/* 	- conta le stringhe (separate da white space) presenti nella frase, escludendo i numeri (sequenze di sole cifre e eventuale punteggiatura attaccata alla fine)
	- calcola la lunghezza della parola più corta nella frase, **non considerando** l'eventuale punteggiatura come parte delle parole da misurare ("minima:" sarà considerata di lunghezza 6 e non 7), ed escludendo i numeri nella frase
	- calcola la lunghezza della parola più lunga nella frase, **non considerando** la punteggiatura come parte delle parole da misurare, ed escludendo i numeri nella frase
	- nota bene: per 'lunghezza' si intende il numero di byte
	- nota bene: in caso di stringa vuota, la funzione deve restituire (0,0,0)
*/
func CalcQuanteMinMax(frase string) (quante, min, max int){
	min = 100
	var num, check int
	parole := strings.Split(frase, " ")
	fmt.Println("lunghezza parole: ", len(parole))
	for i:=0; i < len(parole); i++{
		s := string(parole[i])
		var conta int
		for j:=0; j < len(s); j++{
			if unicode.IsDigit(rune(s[j])) == true && check == 0{
				num++
				check++
			}else if unicode.IsPunct(rune(s[j])) == false{
				conta++
			}
		}
		check = 0
		if conta < min{
			min = conta
		}
		if conta > max{
			max = conta
		}
	}
	quante = len(parole) - num
	return quante, min, max
}
/*
func TrovaNumDopoKeyword(frase, keyWord string) (num int){

}

func Autogramma(frase string) bool{

}

func StampaAnalisiAutogramma(frase string){

}
*/
func main(){
	q, mi, ma := CalcQuanteMinMax("questa frase contiene: 13 parole, lunghezza massima: 8, lunghezza minima: 5, esclusi numeri")
	fmt.Println("q, mi, ma: ", q, mi, ma)
}