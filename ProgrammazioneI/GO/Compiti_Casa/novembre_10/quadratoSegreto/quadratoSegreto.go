//Autore: Francesco Corrado

package main

import(
    "fmt"
	"math"
	"unicode"
	"strconv"
)

func numeroSegreto(str string) int{
	var s, m int
	m = 1
	for x:=len(str)-1; x>=0; x--{
		j := rune(str[x])
		if unicode.IsDigit(j){
			d, _ := strconv.Atoi(string(j))
			s += d * m
			m *= 10
		}
	}
	return s
}

func main(){
	var n int
    var s string
	fmt.Print("Inserire il numero segreto: ")
	fmt.Scan(&s)
	n = numeroSegreto(s)
	sqr := math.Pow(float64(n), 2.0)
	fmt.Println("Il numero segreto è", n, "e la sua radice quadrata", sqr)
}