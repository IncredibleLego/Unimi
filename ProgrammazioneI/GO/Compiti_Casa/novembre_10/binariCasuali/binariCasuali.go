//Autore: Francesco Corrado

package main

import(
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"math"
)

func genBinary(r *rand.Rand) string{
	var str string
	l := rand.Intn(20) + 1
	for x:=0; x<l; x++{
		n := rand.Intn(2)
		str = str + strconv.Itoa(n)
	}
	return str
}

func convertDecimal(str string) int{
	var n int
	var p float64
	for x:=len(str)-1; x>=0; x--{
		j := rune(str[x])
		d, _ := strconv.Atoi(string(j))
		n = n + d * int(math.Pow(2.0, p))
		p += 1
	}
	return n
}

func main(){
	start := time.Now()
	var s string
	var i int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Programma che genera 1000 stringhe binarie da 1 a 20 caratteri e le converte")
	fmt.Println("binary:decimal =")
	for n:=0; n<1000; n++{
		s = genBinary(r)
		i = convertDecimal(s)
//		i = strconv.FormatInt(i, 10)
		fmt.Println("\t", s, ":", i)
	}
	end := time.Now()
	t := end.Sub(start)
	fmt.Println("Tempo esecuzione: ", t)
}