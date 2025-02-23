//Autore: Francesco Corrado

package main

import(
	"fmt"
)
func main(){
	var s, maxPari string
	var nPari, max int
	const d=10
	for{
		fmt.Scan(&s)
		if s == "000"{
			break
		}
		for i:=0; i<len(s); i++{
			if (s[i] % d) % 2 == 0{
				nPari++
			}
		}
		if nPari > max{
			max = nPari
			maxPari = s
		}
		nPari = 0
	}
	fmt.Println("Il numero che contiene più cifre pari è", maxPari, "che ne contiene", max)
}