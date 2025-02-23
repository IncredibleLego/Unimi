//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func operazioni(n1, n2 int) (int, int, int){
	return n1+n2, n1*n2, n1-n2
}

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(operazioni(a, b))
}