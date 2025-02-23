//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
)

func main(){
	const MAX = 10
	var N, n int
	n = rand.Intn(10)+1
	fmt.Println("Indovina un numero casuale generato tra [1,10]. Hai", MAX/2, "tentativi")
	for t:=1; t<=MAX/2; t++{
		fmt.Print("Indovina il numero (tentaivo numero ", t, "): ")
		fmt.Scan(&N)
		if N<1 || N>MAX{
			fmt.Println("fuori intervallo")
			t--
		}
		if N == n{
			fmt.Println(t, "tentativi")
			break
		}
		if t == MAX/2 {
			fmt.Println("hai perso, il numero era", n)
		}
	}
}