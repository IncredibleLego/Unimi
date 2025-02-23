//Autore: Francesco Corrado

package main

import(
    "fmt"
	"math/rand"
	"strconv"
)

type Carta struct{
	valore, seme string
}

func carta(n int) (Carta, bool){
	var card Carta
	if n < 0 || n > 51{
		return card, false
	}
	v := n - (n/13)*13
	switch{
		case v == 0:
			card.valore = "A"
		case v == 10:
			card.valore = "J"
		case v == 11:
			card.valore = "Q"
		case v == 12:
			card.valore = "K"
		default:
			card.valore = fmt.Sprint(v+1)
	}
	switch{
		case n <= 12:
			card.seme = "\u2665"
		case n >= 13 && n <= 25:
			card.seme = "\u2666"
		case n >= 26 && n <= 38:
			card.seme = "\u2663"
		case n >= 39:
			card.seme = "\u2660"
	}
	return card, true
}

func estraiCarta() Carta{
	n := rand.Intn(52)
	card,_ := carta(n)
	return card
}

func dai4Carte() []Carta{
	var array []Carta
	for i:=0; i < 4; i++{
		n := rand.Intn(52)
		card,_ := carta(n)
		array = append(array, card)
	}
	return array
}

func getValoreBJ(card Carta) int{
	var output int
	if card.valore == "J" || card.valore == "Q" || card.valore == "K"{
		output = 10
	}else if card.valore == "A"{
		output = 11
	}else{
		output, _ = strconv.Atoi(card.valore)
	}
	return output
}

func mazzoPoker() []Carta{
	var mazzo []Carta
	for i:=0; i < 52; i++{
		card, _ := carta(i)
		mazzo = append(mazzo, card)
	}
	return mazzo
}

func mischia(mazzo []Carta) []Carta{
	for i:=0; i < 52; i++{
		r := rand.Intn(52)
		mazzo[i], mazzo[r] = mazzo[r], mazzo[i]
	}
	return mazzo
}

func main(){
	const SEMI = 4
	const VALORI = 13
	const CMAZZO = 52
	cart := estraiCarta()
	fmt.Println("Carta estratta: ", cart)
	fmt.Println("Valore carta nel black jack: ", getValoreBJ(cart))
	arr := dai4Carte()
	fmt.Println("Quattro carte estratte: ", arr)
	mazzo := mazzoPoker()
	fmt.Println("Creo un mazzo di carte: \n", mazzo)
	mazzo = mischia(mazzo)
	fmt.Println("Rimescolo il mazzo: \n", mazzo)
}