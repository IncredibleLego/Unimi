//Autore: Francesco Corrado

package main

import(
    "fmt"
	"math"
)

func main(){
    var n, g int
	var ok bool
	var mappa map[int]string
	mappa = make(map[int]string)
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	count := 0
	a := n
	for a > 0 {
	   a = a/10
	   count++
	}
	c := int(math.Pow(10, float64(count-1)))
	for x:=n; x>0;x = x - c*10*g {
		g=x/c
		_, ok = mappa[g]
		if ok == false{
			var nome string
			fmt.Print("parola per ", g, " ? ")
			fmt.Scan(&nome)
			mappa[g]=nome
		}
		if x < 10{ //Così facendo se il secondo numero è uno zero non funziona es 203 rimane 03
			break
		}
		c = c/10
	}
	c = int(math.Pow(10, float64(count-1)))
	for x:=n; x>0;x = x - c*10*g {
		var numero string
		g=x/c
		numero = mappa[g]
		fmt.Print(numero)
		if x > 10{
			fmt.Print(" - ")
		}
		if x < 10{
			break
		}
		c = c/10
	}
	fmt.Println()
}