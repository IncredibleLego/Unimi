//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
)

const LEN_ALFABETO = 26

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int){
	for i:= 0; i < len(s); i++{
		if s[i] > 96 && s[i] < 123{
			contaMinu[s[i]-97]++
		}
	}
}

func main(){
	var alfabeto [LEN_ALFABETO]int
	for i:=0; i < LEN_ALFABETO; i++{
		alfabeto[i] = 0
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		riga := scanner.Text()
		aggiornaConteggi(riga, &alfabeto)
	}
	for i:=0; i < LEN_ALFABETO; i++{
		fmt.Println(string(rune(i+97)), alfabeto[i])
	}
}