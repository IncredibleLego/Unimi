//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
	"unicode"
	"strconv"
	"math"
)

func contaOccorrenze(str string, x string) int{
	c := strings.Count(str, x)
	return c
}

func posVocali(str string, x string) int{
	a := strings.IndexAny(str, x)
	b := strings.LastIndexAny(str, x)
	fmt.Println("(4) la prima vocale di", str, "è in posizione", a, ", l'ultima in pozione", b)
	return 0
}

func trovaNumeri (str string) string{
	var sOut string
	for i:=0; i < len(str); i++{
		r := rune(str[i])
		if unicode.IsNumber(r){
			sOut += string(r)
		}
	}
	return sOut
}

func main(){
	var VOCALI, S1, S2, s string
	VOCALI = "aeiouAEIOU"
	S1 = "c"
	S2 = "sc"
	fmt.Scan(&s)
	if strings.Contains(s, S2){
		fmt.Println("(1)", s, "contiene", S2)
	}
	if strings.ContainsAny(s, VOCALI){
		fmt.Println("(2)", s, "ha vocali")
	}
	fmt.Println("(3)", s, "ha", contaOccorrenze(s, S1), S1)
	posVocali(s, VOCALI)
	fmt.Println("(5)", strings.Repeat(S2, 3))
	fmt.Println("(6)", strings.Repeat(S1, 5))
	str := trovaNumeri(s)
	fmt.Println("(7) stringa", str)
	n, err := strconv.Atoi(str)
	_ = err
	fmt.Printf("(8) intero %d\n", n)
	l := len(str)
	p := math.Pow(10, float64(l))
	f, err := strconv.ParseFloat(str, 64)
	f /= p
	fmt.Printf("(9) float %f\n", f)
}