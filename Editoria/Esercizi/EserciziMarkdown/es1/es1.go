//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var x string
	fmt.Print("Input: ")
	fmt.Scan(&x)
	fmt.Print("Output: ")
	for i:=0; i<len(x); i++{
		fmt.Print(x[i], " ")
	}
	fmt.Println()
}